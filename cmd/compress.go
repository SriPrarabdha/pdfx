package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SriPrarabdha/pdfx/internal/pdf"
)

var (
	compressLevel  string
	compressOutput string
)

var compressCmd = &cobra.Command{
	Use:   "compress [pdf file]",
	Short: "Compress a PDF file",
	Long: `Compress a PDF locally using Ghostscript.

Compression levels:
  good   - high quality, moderate compression
  better - balanced (default)
  best   - smallest file size (most aggressive)

Examples:
  pdfx compress file.pdf --level good
  pdfx compress file.pdf --level best -o small.pdf
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		input := args[0]

		if err := pdf.CompressPDF(input, compressOutput, compressLevel); err != nil {
			return err
		}

		fmt.Println("Compressed PDF written to:", compressOutput)
		return nil
	},
}

func init() {
	compressCmd.Flags().StringVar(
		&compressLevel,
		"level",
		"better",
		"Compression level: good | better | best",
	)

	compressCmd.Flags().StringVarP(
		&compressOutput,
		"output",
		"o",
		"compressed.pdf",
		"Output PDF file",
	)

	rootCmd.AddCommand(compressCmd)
}
