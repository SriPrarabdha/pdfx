package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SriPrarabdha/pdfx/internal/pdf"
)

var output string

var mergeCmd = &cobra.Command{
	Use:   "merge [pdf files]",
	Short: "Merge multiple PDF files into one",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := pdf.Merge(args, output); err != nil {
			return err
		}

		fmt.Println("Merged PDF created:", output)
		return nil
	},
}

func init() {
	mergeCmd.Flags().StringVarP(&output, "output", "o", "merged.pdf", "Output PDF file")
	rootCmd.AddCommand(mergeCmd)
}
