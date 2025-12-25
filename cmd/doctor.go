package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SriPrarabdha/pdfx/internal/pdf"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system dependencies required by pdfx",
	Run: func(cmd *cobra.Command, args []string) {
		results := pdf.CheckDependencies()

		fmt.Println("pdfx system check:")
		for _, d := range results {
			if d.Available {
				fmt.Printf("  ✔ %s (%s)\n", d.Name, d.Description)
			} else {
				fmt.Printf("  ✘ %s (%s) — NOT FOUND\n", d.Name, d.Description)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
