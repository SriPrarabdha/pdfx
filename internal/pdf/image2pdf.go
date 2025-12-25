package pdf

import (
	"fmt"
	"io"       
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)


var supportedImageExt = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".bmp":  true,
	".tif":  true,
	".tiff": true,
}

// ImagesToPDF converts image files into a single PDF.
func ImagesToPDF(images []string, output string) error {
	if len(images) == 0 {
		return fmt.Errorf("no image files provided")
	}

	var readers []os.File

	for _, img := range images {
		if _, err := os.Stat(img); err != nil {
			return fmt.Errorf("image not found: %s", img)
		}

		ext := strings.ToLower(filepath.Ext(img))
		if !supportedImageExt[ext] {
			return fmt.Errorf("unsupported image format: %s", img)
		}

		f, err := os.Open(img)
		if err != nil {
			return fmt.Errorf("failed to open image %s: %w", img, err)
		}
		readers = append(readers, *f)
	}

	defer func() {
		for _, f := range readers {
			_ = f.Close()
		}
	}()

	outFile, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("failed to create output PDF: %w", err)
	}
	defer outFile.Close()

	var imgReaders []io.Reader
	for i := range readers {
		imgReaders = append(imgReaders, &readers[i])
	}

	importConf := pdfcpu.DefaultImportConfig()

	if err := api.ImportImages(
		nil,                 
		outFile,             
		imgReaders,          
		importConf,          
		model.NewDefaultConfiguration(),
	); err != nil {
		return fmt.Errorf("failed to create PDF from images: %w", err)
	}

	return nil
}
