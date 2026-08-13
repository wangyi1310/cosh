package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"cosh/internal/config"
	"cosh/internal/cosclient"
)

var getCmd = &cobra.Command{
	Use:   "get <remote-key> [local-path]",
	Short: "Download a file from COS",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		remoteKey := args[0]
		localPath := ""
		if len(args) > 1 {
			localPath = args[1]
		}

		// If local path is directory or ".", append filename from remote key
		if localPath == "" || localPath == "." {
			localPath = filepath.Base(remoteKey)
		} else if info, err := os.Stat(localPath); err == nil && info.IsDir() {
			localPath = filepath.Join(localPath, filepath.Base(remoteKey))
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}
		client, err := cosclient.NewClient(cfg, bucketFlag)
		if err != nil {
			return err
		}

		resp, err := client.Object.Get(context.Background(), remoteKey, nil)
		if err != nil {
			return fmt.Errorf("download: %w", err)
		}
		defer resp.Body.Close()

		total := resp.ContentLength

		f, err := os.Create(localPath)
		if err != nil {
			return fmt.Errorf("create file: %w", err)
		}
		defer f.Close()

		pw := &progressWriter{total: total, label: remoteKey}
		n, err := io.Copy(f, io.TeeReader(resp.Body, pw))
		pw.finish()
		if err != nil {
			return fmt.Errorf("write file: %w", err)
		}

		fmt.Printf("%s✓%s Download complete. (%s%s%s)\n",
			colorGreen, colorReset,
			colorCyan, formatSize(n), colorReset)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
