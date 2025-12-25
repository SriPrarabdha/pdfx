package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	imageutil "github.com/SriPrarabdha/pdfx/internal/image"
)

var (
	imgCompressLevel  string
	imgCompressOutput string
)

var imgCompressCmd = &cobra.Command{
	Use:   "img-compress [image]",
	Short: "Compress an image file",
	Long: `Compress an image locally.

Levels:
  good   - high quality
  better - balanced
  best   - smallest size

Supported formats: jpg, png, webp
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := imageutil.CompressImage(args[0], imgCompressOutput, imgCompressLevel); err != nil {
			return err
		}

		fmt.Println("Compressed image written to:", imgCompressOutput)
		return nil
	},
}

func init() {
	imgCompressCmd.Flags().StringVar(
		&imgCompressLevel,
		"level",
		"better",
		"Compression level: good | better | best",
	)

	imgCompressCmd.Flags().StringVarP(
		&imgCompressOutput,
		"output",
		"o",
		"compressed.jpg",
		"Output image file",
	)

	rootCmd.AddCommand(imgCompressCmd)
}
