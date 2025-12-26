package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SriPrarabdha/pdfx/internal/pdf"
)

var extractPages string
var extractOutput string

var extractCmd = &cobra.Command{
	Use:   "Extract [pdf file]",
	Short: "Extract selected pages from a PDF",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if extractPages == "" {
			return fmt.Errorf("--pages is required")
		}

		if err := pdf.ExtractPages(args[0], extractOutput, extractPages); err != nil {
			return err
		}

		fmt.Println("Updated PDF created:", extractOutput)
		return nil
	},
}

func init() {
	extractCmd.Flags().StringVar(&extractPages, "pages", "", "Pages to extract (e.g. 2,4-6)")
	extractCmd.Flags().StringVarP(&extractOutput, "output", "o", "output.pdf", "Output PDF file")
	rootCmd.AddCommand(extractCmd)
}
