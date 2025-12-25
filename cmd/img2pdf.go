package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SriPrarabdha/pdfx/internal/pdf"
)

var imgOutput string

var img2pdfCmd = &cobra.Command{
	Use:   "img2pdf [image files]",
	Short: "Convert images to a single PDF",
	Long: `Convert one or more images into a single PDF.

Supported formats:
  .jpg .jpeg .png .webp .bmp .tif .tiff

Examples:
  pdfx img2pdf img1.jpg img2.png -o images.pdf
  pdfx img2pdf *.jpg -o album.pdf
`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := pdf.ImagesToPDF(args, imgOutput); err != nil {
			return err
		}

		fmt.Println("PDF created from images:", imgOutput)
		return nil
	},
}

func init() {
	img2pdfCmd.Flags().StringVarP(
		&imgOutput,
		"output",
		"o",
		"images.pdf",
		"Output PDF file",
	)

	rootCmd.AddCommand(img2pdfCmd)
}
