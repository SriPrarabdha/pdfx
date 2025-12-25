package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SriPrarabdha/pdfx/internal/pdf"
)

var (
	deletePages string
	deleteOutput string
)

var deleteCmd = &cobra.Command{
	Use:   "delete [pdf file]",
	Short: "Delete selected pages from a PDF",
	Long: `Delete pages from a PDF using inclusive page ranges.

Examples:
  pdfx delete file.pdf --pages 3
  pdfx delete file.pdf --pages 1,4-6
  pdfx delete file.pdf --pages n
  pdfx delete file.pdf --pages 2-n
  pdfx delete file.pdf --pages 1,3,5,n
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if deletePages == "" {
			return fmt.Errorf("--pages flag is required")
		}

		input := args[0]

		if err := pdf.DeletePages(input, deleteOutput, deletePages); err != nil {
			return err
		}

		fmt.Println("Updated PDF written to:", deleteOutput)
		return nil
	},
}

func init() {
	deleteCmd.Flags().StringVar(
		&deletePages,
		"pages",
		"",
		"Pages to delete (e.g. 1,3-5,n,2-n)",
	)

	deleteCmd.Flags().StringVarP(
		&deleteOutput,
		"output",
		"o",
		"output.pdf",
		"Output PDF file",
	)

	rootCmd.AddCommand(deleteCmd)
}
