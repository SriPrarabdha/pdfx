package pdf

import (
	"fmt"
	"os"
	// "os/exec"
	"path/filepath"
	"strings"
)

func ConvertToPDF(input, output string) error {
	if _, err := os.Stat(input); err != nil {
		return fmt.Errorf("input file not found: %s", input)
	}

	ext := strings.ToLower(filepath.Ext(input))

	switch ext {
	case ".txt", ".md", ".html", ".htm":
		return convertWithPandoc(input, output)

	case ".docx", ".pptx", ".xlsx", ".odt":
		return convertWithLibreOffice(input, output)

	default:
		return fmt.Errorf("unsupported file type: %s", ext)
	}
}
