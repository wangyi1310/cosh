package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"cosh/internal/config"
	"cosh/internal/cosclient"
)

var bucketsCmd = &cobra.Command{
	Use:   "buckets",
	Short: "List all COS buckets",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}
		client, err := cosclient.NewServiceClient(cfg)
		if err != nil {
			return err
		}
		resp, _, err := client.Service.Get(context.Background())
		if err != nil {
			return fmt.Errorf("list buckets: %w", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "%sNAME%s\t%sREGION%s\t%sCREATED%s\n",
			colorCyan, colorReset, colorCyan, colorReset, colorCyan, colorReset)
		for _, b := range resp.Buckets {
			fmt.Fprintf(w, "%s%s%s\t%s%s%s\t%s%s%s\n",
				colorWhite, b.Name, colorReset,
				colorYellow, b.Region, colorReset,
				colorDim, b.CreationDate, colorReset)
		}
		w.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(bucketsCmd)
}
