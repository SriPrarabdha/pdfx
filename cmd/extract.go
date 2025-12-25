package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SriPrarabdha/pdfx/internal/pdf"
)

var deletePages string
var deleteOutput string

var deleteCmd = &cobra.Command{
	Use:   "delete [pdf file]",
	Short: "Delete selected pages from a PDF",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if deletePages == "" {
			return fmt.Errorf("--pages is required")
		}

		if err := pdf.DeletePages(args[0], deleteOutput, deletePages); err != nil {
			return err
		}

		fmt.Println("Updated PDF created:", deleteOutput)
		return nil
	},
}

func init() {
	deleteCmd.Flags().StringVar(&deletePages, "pages", "", "Pages to delete (e.g. 2,4-6)")
	deleteCmd.Flags().StringVarP(&deleteOutput, "output", "o", "output.pdf", "Output PDF file")
	rootCmd.AddCommand(deleteCmd)
}
