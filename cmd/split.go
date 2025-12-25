package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SriPrarabdha/pdfx/internal/pdf"
)

var splitPages string

var splitCmd = &cobra.Command{
	Use:   "split [pdf file]",
	Short: "Extract selected pages into separate PDFs",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if splitPages == "" {
			return fmt.Errorf("--pages is required")
		}

		if err := pdf.Split(args[0], splitPages); err != nil {
			return err
		}

		fmt.Println("Pages extracted successfully")
		return nil
	},
}

func init() {
	splitCmd.Flags().StringVar(&splitPages, "pages", "", "Pages to extract (e.g. 1-3,5)")
	rootCmd.AddCommand(splitCmd)
}
