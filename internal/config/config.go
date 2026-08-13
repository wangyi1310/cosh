package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	SecretID  string `toml:"secret_id"`
	SecretKey string `toml:"secret_key"`
	Region    string `toml:"region"`
	Bucket    string `toml:"bucket"`
}

func ConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".cosh.toml")
}

func Load() (*Config, error) {
	path := ConfigPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config not found: %s\nRun 'cosh config init' first", path)
	}
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	if cfg.SecretID == "" || cfg.SecretKey == "" {
		return nil, fmt.Errorf("secret_id and secret_key are required")
	}
	if cfg.Region == "" {
		cfg.Region = "ap-guangzhou"
	}
	return &cfg, nil
}

func Save(cfg *Config) error {
	path := ConfigPath()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(cfg)
}

func InteractiveSetup() (*Config, error) {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("=== cosh config setup ===")

	fmt.Print("Secret ID: ")
	secretID, _ := r.ReadString('\n')
	secretID = strings.TrimSpace(secretID)

	fmt.Print("Secret Key: ")
	secretKey, _ := r.ReadString('\n')
	secretKey = strings.TrimSpace(secretKey)

	fmt.Print("Region [ap-guangzhou]: ")
	region, _ := r.ReadString('\n')
	region = strings.TrimSpace(region)
	if region == "" {
		region = "ap-guangzhou"
	}

	fmt.Print("Default Bucket (optional, e.g. mybucket-1250000000): ")
	bucket, _ := r.ReadString('\n')
	bucket = strings.TrimSpace(bucket)

	cfg := &Config{
		SecretID:  secretID,
		SecretKey: secretKey,
		Region:    region,
		Bucket:    bucket,
	}

	if err := Save(cfg); err != nil {
		return nil, err
	}
	fmt.Printf("Config saved to %s\n", ConfigPath())
	return cfg, nil
}
