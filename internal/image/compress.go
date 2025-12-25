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

var qualityMap = map[string]int{
	"good":   85,
	"better": 65,
	"best":   40,
}

func CompressImage(input, output, level string) error {
	q, ok := qualityMap[level]
	if !ok {
		return fmt.Errorf("invalid level: %s (use good, better, best)", level)
	}

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
		return jpeg.Encode(out, img, &jpeg.Options{Quality: q})

	case ".png":
		encoder := png.Encoder{CompressionLevel: png.BestCompression}
		if level == "good" {
			encoder.CompressionLevel = png.DefaultCompression
		}
		return encoder.Encode(out, img)

	default:
		return fmt.Errorf("unsupported output format for compression: %s", ext)
	}
}
