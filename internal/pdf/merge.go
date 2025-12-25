package pdf

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// Merge merges multiple PDF files into a single output file.
func Merge(inputs []string, output string) error {
	// Validate input files
	for _, file := range inputs {
		if _, err := os.Stat(file); err != nil {
			return fmt.Errorf("input file not found: %s", file)
		}
	}

	// Use default pdfcpu configuration
	conf := model.NewDefaultConfiguration()

	// Merge PDFs
	if err := api.MergeCreateFile(inputs, output, false, conf); err != nil {
		return fmt.Errorf("failed to merge PDFs: %w", err)
	}

	return nil
}
