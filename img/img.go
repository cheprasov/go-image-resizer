package img

import (
    "fmt"
    "image"
    "image/jpeg"
    "image/png"
    "os"

    "../fs"

    "github.com/nfnt/resize"
)

func LoadImage(filename string) (*image.Image, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    if fs.IsJpegFile(filename) {
        img, err := jpeg.Decode(file)
        return &img, err;
    }

    if fs.IsPngFile(filename) {
        img, err := png.Decode(file)
        return &img, err;
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
    return &img, err;
}

func LoadPngImage(filename string) (*image.Image, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    img, err := png.Decode(file)
    return &img, err;
}

func ResizeFile(img *image.Image, width, height uint, skipSmall bool) *image.Image {
    if skipSmall {
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

func SaveImage(filename string, img *image.Image, quality uint) error {
    outFile, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer outFile.Close()

    if fs.IsJpegFile(filename) {
        options := &jpeg.Options{Quality: int(quality)};
        return jpeg.Encode(outFile, *img, options)
    }

    if fs.IsPngFile(filename) {
        return png.Encode(outFile, *img)
    }

    return fmt.Errorf("%s is not image file", filename)
}
