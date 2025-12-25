package imageutil

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func ConvertImage(input, output string) error {
	in, err := os.Open(input)
	if err != nil {
		return err
	}
	defer in.Close()

	img, _, err := image.Decode(in)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	ext := strings.ToLower(filepath.Ext(output))

	switch ext {
	case ".jpg", ".jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 90})

	case ".png":
		return png.Encode(out, img)

	default:
		return fmt.Errorf("unsupported output format: %s", ext)
	}
}
