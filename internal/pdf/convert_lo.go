package pdf

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func convertWithLibreOffice(input, output string) error {
	if _, err := exec.LookPath(libreOfficeCmd()); err != nil {
		return fmt.Errorf("libreoffice not found (required for office formats)")
	}

	outDir := filepath.Dir(output)

	cmd := exec.Command(
		libreOfficeCmd(),
		"--headless",
		"--convert-to", "pdf",
		"--outdir", outDir,
		input,
	)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("libreoffice conversion failed: %w", err)
	}

	generated := filepath.Join(
		outDir,
		strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))+".pdf",
	)

	return os.Rename(generated, output)
}
