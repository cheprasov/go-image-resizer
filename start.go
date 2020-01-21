package main

import (
    "flag"
    "fmt"
    "log"
    "path"

    "./fs"
    IMG "./img"
    "./pathUtils"
)

func getFlags() (string, string, uint, uint, uint, bool) {
    sourcePath := flag.String("source-path", "", "Please provide --source-path")
    outputPath := flag.String("output-path", "", "Please provide --output-path")
    width := flag.Uint("width", 0, "Please provide --width")
    height := flag.Uint("height", 0, "Please provide --height")
    quality := flag.Uint("quality", 100, "Please provide --quality")
    skipSmall := flag.Bool("skip-small", false, "Please provide --skip-small")
    flag.Parse()

    if len(*sourcePath) == 0 {
        log.Fatal("Please provide --source-path")
    }

    if len(*outputPath) == 0 {
        log.Fatal("Please provide --output-path")
    }

    if *sourcePath == *outputPath {
        log.Fatal("--source-path and --source-path should be not equal")
    }

    if *width == 0 && *height == 0 {
        log.Fatal("--width or/and --height should be provided")
    }

    return *sourcePath, *outputPath, *width, *height, *quality, *skipSmall;
}

func main() {
    sourcePath, outputPath, width, height, quality, skipSmall := getFlags()

    sourcePath = pathUtils.NormalizePath(sourcePath);
    outputPath = pathUtils.NormalizePath(outputPath);

    fmt.Printf("STARTING IMG RESIZER, source: %s, output: %s \n", sourcePath, outputPath)

    filesMapPointer, _ := fs.ReadFolder(sourcePath, nil, true);

    for filename, _ := range (*filesMapPointer).InfoMap {
        if !fs.IsImageFile(filename) {
            continue;
        }

        fmt.Printf("Resize image: %s", filename)

        img, err := IMG.LoadImage(filename)
        if err != nil {
            fmt.Printf(" - error: can't load image\n")
            continue
        }

        resizedImg := IMG.ResizeFile(img, width, height, skipSmall);
        outFilename := pathUtils.NormalizePath(outputPath + pathUtils.TrimPrefix(filename, sourcePath))

        outDir := path.Dir(outFilename)
        if !fs.IsDirExists(outDir) {
            err = fs.MkDir(outDir)
            if err != nil {
                fmt.Printf(" - error: can't create dir %s\n", outDir)
                continue
            }
        }

        err = IMG.SaveImage(outFilename, resizedImg, quality);
        if err != nil {
            fmt.Printf(" - error: can't save image to %s\n", outFilename)
            continue
        }


        fmt.Printf(" - saved to %s\n", outFilename)
    }
}