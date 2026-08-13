package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	cos "github.com/tencentyun/cos-go-sdk-v5"

	"cosh/internal/config"
	"cosh/internal/cosclient"
)

var lsCmd = &cobra.Command{
	Use:   "ls [prefix]",
	Short: "List objects in bucket",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}
		bucket := bucketFlag
		if bucket == "" {
			bucket = cfg.Bucket
		}
		client, err := cosclient.NewClient(cfg, bucket)
		if err != nil {
			return err
		}

		prefix := ""
		if len(args) > 0 {
			prefix = args[0]
		}

		opt := &cos.BucketGetOptions{
			Prefix:    prefix,
			Delimiter: "/",
		}
		resp, _, err := client.Bucket.Get(context.Background(), opt)
		if err != nil {
			return fmt.Errorf("list objects: %w", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "%sTYPE%s\t%sSIZE%s\t%sLAST MODIFIED%s\t%sKEY%s\n",
			colorCyan, colorReset, colorCyan, colorReset,
			colorCyan, colorReset, colorCyan, colorReset)

		for _, p := range resp.CommonPrefixes {
			fmt.Fprintf(w, "%sDIR%s\t%s-%s\t%s-%s\t%s%s%s\n",
				colorBlue, colorReset, colorDim, colorReset,
				colorDim, colorReset, colorYellow, p, colorReset)
		}
		for _, obj := range resp.Contents {
			size := formatSize(obj.Size)
			fmt.Fprintf(w, "%sFILE%s\t%s%s%s\t%s%s%s\t%s\n",
				colorGreen, colorReset,
				colorWhite, size, colorReset,
				colorDim, obj.LastModified, colorReset,
				obj.Key)
		}
		w.Flush()

		total := len(resp.CommonPrefixes) + len(resp.Contents)
		fmt.Printf("\n%s%d%s items\n", colorCyan, total, colorReset)
		return nil
	},
}

func formatSize(bytes int64) string {
	const (
		KB int64 = 1024
		MB       = KB * 1024
		GB       = MB * 1024
	)
	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.1fGB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.1fMB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.1fKB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
