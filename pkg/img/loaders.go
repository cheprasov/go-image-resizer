package img

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/cheprasov/go-image-resizer/pkg/fs"
	"golang.org/x/image/webp"
)

type ImageLoaderType func(filename string) (*image.Image, error)

func LoadImage(filename string) (*image.Image, error) {
	var loaders = make([]ImageLoaderType, 4)

	// Mechanism for loading image with wrong ext
	loaders[0] = LoadJpegImage
	loaders[1] = LoadPngImage
	loaders[2] = LoadGifImage
	loaders[3] = LoadWebpImage

	if fs.IsJpegFile(filename) {
	} else if fs.IsPngFile(filename) {
		loaders[0] = LoadPngImage
		loaders[1] = LoadJpegImage
	} else if fs.IsGifFile(filename) {
		loaders[0] = LoadGifImage
		loaders[2] = LoadJpegImage
	} else if fs.IsWebpFile(filename) {
		loaders[0] = LoadWebpImage
		loaders[3] = LoadJpegImage
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

func LoadGifImage(filename string) (*image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := gif.Decode(file)
	return &img, err
}

func LoadWebpImage(filename string) (*image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := webp.Decode(file)
	return &img, err
}
