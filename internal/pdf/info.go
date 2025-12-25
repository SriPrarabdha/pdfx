package pdf

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func PageCount(file string) (int, error) {
	ctx, err := api.ReadContextFile(file)
	if err != nil {
		return 0, fmt.Errorf("failed to read pdf: %w", err)
	}
	return ctx.PageCount, nil
}
