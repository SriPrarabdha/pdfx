package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdfx",
	Short: "pdfx is a local-first, privacy-friendly PDF toolkit",
	Long: `pdfx is a CLI tool for working with PDFs locally.

No cloud.
No tracking.
All operations happen on your machine.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
