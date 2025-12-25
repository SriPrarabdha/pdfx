package pdf

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func ExtractPages(input, output, pages string) error {
	if _, err := os.Stat(input); err != nil {
		return fmt.Errorf("input file not found: %s", input)
	}

	pageList, err := ParsePageRange(pages)
	if err != nil {
		return err
	}

	if err := api.TrimFile(input, output, pageList, nil); err != nil {
		return fmt.Errorf("failed to delete pages: %w", err)
	}

	return nil
}
