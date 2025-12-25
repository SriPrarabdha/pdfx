package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	imageutil "github.com/SriPrarabdha/pdfx/internal/image"
)

var imgConvertOutput string

var imgConvertCmd = &cobra.Command{
	Use:   "img-convert [image]",
	Short: "Convert image format",
	Long: `Convert image formats locally.

Examples:
  pdfx img-convert photo.webp -o photo.png
  pdfx img-convert scan.jpg  -o scan.webp
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := imageutil.ConvertImage(args[0], imgConvertOutput); err != nil {
			return err
		}

		fmt.Println("Converted image written to:", imgConvertOutput)
		return nil
	},
}

func init() {
	imgConvertCmd.Flags().StringVarP(
		&imgConvertOutput,
		"output",
		"o",
		"output.png",
		"Output image file",
	)

	rootCmd.AddCommand(imgConvertCmd)
}
