package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"cosh/internal/config"
	"cosh/internal/cosclient"
)

var putCmd = &cobra.Command{
	Use:   "put <local-file> <remote-key>",
	Short: "Upload a local file to COS",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		localPath := args[0]
		remoteKey := args[1]

		if remoteKey == "." {
			remoteKey = filepath.Base(localPath)
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}
		client, err := cosclient.NewClient(cfg, bucketFlag)
		if err != nil {
			return err
		}

		f, err := os.Open(localPath)
		if err != nil {
			return fmt.Errorf("open file: %w", err)
		}
		defer f.Close()

		stat, err := f.Stat()
		if err != nil {
			return err
		}

		fmt.Printf("%s↑%s Uploading %s%s%s (%s%s%s) -> %s%s%s\n",
			colorCyan, colorReset,
			colorWhite, localPath, colorReset,
			colorDim, formatSize(stat.Size()), colorReset,
			colorYellow, remoteKey, colorReset)

		_, err = client.Object.Put(context.Background(), remoteKey, f, nil)
		if err != nil {
			return fmt.Errorf("upload: %w", err)
		}

		fmt.Printf("%s✓%s Upload complete.\n", colorGreen, colorReset)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
