package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"cosh/internal/config"
	"cosh/internal/cosclient"
)

var mkdirCmd = &cobra.Command{
	Use:   "mkdir <remote-path>",
	Short: "Create a folder in COS",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dirPath := args[0]
		if !strings.HasSuffix(dirPath, "/") {
			dirPath += "/"
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}
		client, err := cosclient.NewClient(cfg, bucketFlag)
		if err != nil {
			return err
		}

		_, err = client.Object.Put(context.Background(), dirPath, strings.NewReader(""), nil)
		if err != nil {
			return fmt.Errorf("mkdir: %w", err)
		}

		fmt.Printf("%s📁 Created folder: %s%s%s\n",
			colorGreen, colorYellow, dirPath, colorReset)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(mkdirCmd)
}
