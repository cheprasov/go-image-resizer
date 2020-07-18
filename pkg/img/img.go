package img

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/cheprasov/go-image-resizer/pkg/fs"
	"github.com/cheprasov/resize"
)

func ResizeFile(img *image.Image, width, height uint, isLargeOnly bool) *image.Image {
	if width == 0 && height == 0 {
		return img
	}
	if isLargeOnly {
		if width > 0 && (*img).Bounds().Size().X < int(width) {
			return img
		}
		if height > 0 && (*img).Bounds().Size().Y < int(height) {
			return img
		}
	}
	resizedImg := resize.Resize(width, height, *img, resize.NearestNeighbor)
	return &resizedImg
}

func SaveImage(filename string, img *image.Image, quality uint8) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if fs.IsJpegFile(filename) {
		options := &jpeg.Options{Quality: int(quality)}
		return jpeg.Encode(outFile, *img, options)
	}

	if fs.IsPngFile(filename) {
		return png.Encode(outFile, *img)
	}

	return fmt.Errorf("%s is not image file", filename)
}
