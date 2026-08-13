package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"cosh/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage cosh configuration",
}

var (
	cfgSecretID  string
	cfgSecretKey string
	cfgRegion    string
	cfgBucket    string
)

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config (interactive or via flags)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfgSecretID != "" && cfgSecretKey != "" {
			cfg := &config.Config{
				SecretID:  cfgSecretID,
				SecretKey: cfgSecretKey,
				Region:    cfgRegion,
				Bucket:    cfgBucket,
			}
			if cfg.Region == "" {
				cfg.Region = "ap-guangzhou"
			}
			if err := config.Save(cfg); err != nil {
				return fmt.Errorf("save config: %w", err)
			}
			fmt.Printf("Config saved to %s\n", config.ConfigPath())
			return nil
		}
		_, err := config.InteractiveSetup()
		if err != nil {
			return fmt.Errorf("config setup failed: %w", err)
		}
		return nil
	},
}

func init() {
	configInitCmd.Flags().StringVar(&cfgSecretID, "secret_id", "", "COS Secret ID (AK)")
	configInitCmd.Flags().StringVar(&cfgSecretKey, "secret_key", "", "COS Secret Key (SK)")
	configInitCmd.Flags().StringVar(&cfgRegion, "region", "", "COS region (default: ap-guangzhou)")
	configInitCmd.Flags().StringVar(&cfgBucket, "bucket", "", "COS bucket name (e.g. mybucket-1250000000)")

	configCmd.AddCommand(configInitCmd)
	rootCmd.AddCommand(configCmd)
}
