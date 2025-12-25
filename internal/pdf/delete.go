package pdf

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func DeletePages(input, output, pages string) error {
	if _, err := os.Stat(input); err != nil {
		return fmt.Errorf("input file not found: %s", input)
	}

	total, err := PageCount(input)
	if err != nil {
		return err
	}

	toDelete, err := ParseDeletePages(pages, total)
	if err != nil {
		return err
	}

	keep := ComplementPages(toDelete, total)
	if len(keep) == 0 {
		return fmt.Errorf("cannot delete all pages")
	}

	if err := api.TrimFile(input, output, keep, nil); err != nil {
		return fmt.Errorf("failed to delete pages: %w", err)
	}

	return nil
}
