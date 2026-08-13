package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"cosh/internal/config"
	"cosh/internal/cosclient"
)

var rmCmd = &cobra.Command{
	Use:   "rm <remote-key>",
	Short: "Delete an object from COS",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		remoteKey := args[0]

		cfg, err := config.Load()
		if err != nil {
			return err
		}
		client, err := cosclient.NewClient(cfg, bucketFlag)
		if err != nil {
			return err
		}

		_, err = client.Object.Delete(context.Background(), remoteKey)
		if err != nil {
			return fmt.Errorf("delete: %w", err)
		}

		fmt.Printf("%s🗑%s Deleted: %s%s%s\n",
			colorRed, colorReset, colorWhite, remoteKey, colorReset)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
