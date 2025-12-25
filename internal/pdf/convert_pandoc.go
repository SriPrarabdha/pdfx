package pdf

import (
	"fmt"
	"os/exec"
)

func convertWithPandoc(input, output string) error {
	if _, err := exec.LookPath("pandoc"); err != nil {
		return fmt.Errorf("pandoc not found (required for this file type)")
	}

	cmd := exec.Command(
		"pandoc",
		input,
		"-o", output,
	)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pandoc conversion failed: %w", err)
	}

	return nil
}
