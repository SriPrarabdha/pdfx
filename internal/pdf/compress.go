package pdf

import (
	"fmt"
	"os"
	"os/exec"
)

var compressionLevels = map[string]string{
	"good":   "/printer",
	"better": "/ebook",
	"best":   "/screen",
}

// CompressPDF compresses a PDF using Ghostscript.
func CompressPDF(input, output, level string) error {
	if _, err := os.Stat(input); err != nil {
		return fmt.Errorf("input file not found: %s", input)
	}

	gsLevel, ok := compressionLevels[level]
	if !ok {
		return fmt.Errorf("invalid compression level: %s (use good, better, best)", level)
	}

	// Check if ghostscript is installed
	if _, err := exec.LookPath("gs"); err != nil {
		return fmt.Errorf("ghostscript (gs) not found in PATH")
	}

	args := []string{
		"-sDEVICE=pdfwrite",
		"-dCompatibilityLevel=1.4",
		"-dPDFSETTINGS=" + gsLevel,
		"-dNOPAUSE",
		"-dQUIET",
		"-dBATCH",
		"-sOutputFile=" + output,
		input,
	}

	cmd := exec.Command("gs", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pdf compression failed: %w", err)
	}

	return nil
}
