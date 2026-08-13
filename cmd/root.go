package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	bucketFlag string
	regionFlag string
)

var rootCmd = &cobra.Command{
	Use:   "cosh",
	Short: "FTP-like CLI tool for Tencent Cloud COS",
	Long: fmt.Sprintf("%s╔══════════════════════════════════════════╗%s\n%s║  %s☁  COSH%s - Tencent Cloud COS CLI         %s║%s\n%s║  %sBrowse, upload, download files easily   %s║%s\n%s╚══════════════════════════════════════════╝%s",
		colorCyan, colorReset,
		colorCyan, colorWhite, colorCyan, colorCyan, colorReset,
		colorCyan, colorDim, colorCyan, colorReset,
		colorCyan, colorReset),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&bucketFlag, "bucket", "", "COS bucket name (or set in config)")
	rootCmd.PersistentFlags().StringVar(&regionFlag, "region", "", "COS region (default: ap-guangzhou)")
}
