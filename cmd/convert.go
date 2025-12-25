package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SriPrarabdha/pdfx/internal/pdf"
)

var convertOutput string

var convertCmd = &cobra.Command{
	Use:   "convert [file]",
	Short: "Convert a file to PDF",
	Long: `Convert common file formats to PDF locally.

Supported formats:
  Text:   .txt .md .html
  Office: .docx .pptx .xlsx .odt

Examples:
  pdfx convert resume.docx -o resume.pdf
  pdfx convert notes.md -o notes.pdf
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := pdf.ConvertToPDF(args[0], convertOutput); err != nil {
			return err
		}

		fmt.Println("Converted to PDF:", convertOutput)
		return nil
	},
}

func init() {
	convertCmd.Flags().StringVarP(
		&convertOutput,
		"output",
		"o",
		"output.pdf",
		"Output PDF file",
	)

	rootCmd.AddCommand(convertCmd)
}
