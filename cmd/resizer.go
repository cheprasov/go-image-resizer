package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/cheprasov/go-image-resizer/pkg/config"
	"github.com/cheprasov/go-image-resizer/pkg/fs"
	imgPkg "github.com/cheprasov/go-image-resizer/pkg/img"
	"github.com/cheprasov/go-image-resizer/pkg/pathUtils"
)

var isVerboseEnabled = false

func verbose(message string, params ...interface{}) {
	if isVerboseEnabled {
		fmt.Printf(message, params...)
	}
}

func resizeImage(cfg config.Config, filename, outFilename string) (bool, error) {
	verbose("Processing image: %s\n", filename)
	img, err := imgPkg.LoadImage(filename)
	if err != nil {
		verbose("Error: can't load image: %s\n", filename)
		return false, err
	}

	resizedImg := imgPkg.ResizeFile(img, cfg.Width, cfg.Height, cfg.IsLargeOnly)

	if len(outFilename) == 0 {
		outFilename = path.Dir(filename) + "/"
	}
	if !cfg.IsSingleFile || outFilename[len(outFilename)-1] == '/' {
		// DIR
		outFilename = pathUtils.NormalizePath(
			outFilename + "/" + path.Base(filename),
		)
	}
	if len(cfg.Prefix) > 0 {
		dir, file := path.Split(outFilename)
		outFilename = path.Clean(dir + "/" + cfg.Prefix + file)
	}

	outDir := path.Dir(outFilename)
	if !fs.IsDirExists(outDir) {
		err = fs.MkDir(outDir)
		if err != nil {
			verbose("Error: can't create dir %s\n", outDir)
			return false, err
		}
	}

	if cfg.IsCustomType() {
		outFilename = strings.TrimSuffix(outFilename, filepath.Ext(outFilename)) + "." + cfg.Extension
	}

	err = imgPkg.SaveImage(outFilename, resizedImg, cfg.Quality)
	if err != nil {
		verbose("Error: can't save image to %s\n", outFilename)
		return false, err
	}

	verbose("Saved as %s\n", outFilename)

	return true, nil
}

func main() {
	cfg := config.GetValidConfig()
	isVerboseEnabled = cfg.IsVerbose

	verbose("STARTING IMG RESIZER, source: %s, output: %s \n", cfg.SourcePath, cfg.OutputPath)

	if fs.IsFileExists(cfg.SourcePath) {
		cfg.IsSingleFile = true
		if !fs.IsImageFile(cfg.SourcePath) {
			fmt.Printf("Not an Image file: %s\n", cfg.SourcePath)
			os.Exit(1)
		}
		_, err := resizeImage(*cfg, cfg.SourcePath, cfg.OutputPath)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	filesMapPointer, _ := fs.ReadFolder(cfg.SourcePath, nil, true)
	cfg.IsSingleFile = false

	var count uint = 0
	for filename := range (*filesMapPointer).InfoMap {
		if !fs.IsImageFile(filename) {
			continue
		}

		done, err := resizeImage(*cfg, filename, cfg.OutputPath)
		if err != nil {
			verbose("%s\n", err.Error())
			continue
		}
		if !done {
			verbose("Fail\n")
		}

		if cfg.Limit > 0 {
			count += 1
			if count >= cfg.Limit {
				verbose("Limit %v is reached. Stopped.\n", cfg.Limit)
				break
			}
		}
	}
}
