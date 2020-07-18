package img

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/cheprasov/go-image-resizer/pkg/fs"
)

type ImageLoaderType func(filename string) (*image.Image, error)

func LoadImage(filename string) (*image.Image, error) {
	var loaders = make([]ImageLoaderType, 2, 2)

	// Mechanism for loading image with wrong ext
	if fs.IsJpegFile(filename) {
		loaders[0] = LoadJpegImage
		loaders[1] = LoadPngImage
	} else if fs.IsPngFile(filename) {
		loaders[0] = LoadPngImage
		loaders[1] = LoadJpegImage
	}

	for _, loader := range loaders {
		img, err := loader(filename)
		if err != nil {
			continue
		}
		return img, nil
	}

	return nil, fmt.Errorf("%s is not image file", filename)
}

func LoadJpegImage(filename string) (*image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	return &img, err
}

func LoadPngImage(filename string) (*image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	return &img, err
}
