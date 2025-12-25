package pdf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func Split(input, pages string) error {
	if _, err := os.Stat(input); err != nil {
		return fmt.Errorf("input file not found: %s", input)
	}

	pageList, err := ParsePageRange(pages)
	if err != nil {
		return err
	}

	outDir := filepath.Dir(input)

	if err := api.ExtractPagesFile(input, outDir, pageList, nil); err != nil {
		return fmt.Errorf("failed to split PDF: %w", err)
	}

	return nil
}
