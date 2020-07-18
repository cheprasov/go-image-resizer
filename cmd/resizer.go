package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/cheprasov/go-image-resizer/pkg/config"
	"github.com/cheprasov/go-image-resizer/pkg/fs"
	imgPkg "github.com/cheprasov/go-image-resizer/pkg/img"
	"github.com/cheprasov/go-image-resizer/pkg/pathUtils"
)

func main() {
	cfg := config.GetValidConfig()

	if cfg.IsVerbose {
		fmt.Printf("STARTING IMG RESIZER, source: %s, output: %s \n", cfg.SourcePath, cfg.OutputPath)
	}

	filesMapPointer, _ := fs.ReadFolder(cfg.SourcePath, nil, true)

	for filename, _ := range (*filesMapPointer).InfoMap {
		if !fs.IsImageFile(filename) {
			continue
		}

		if cfg.IsVerbose {
			fmt.Printf("Found image: %s", filename)
		}

		img, err := imgPkg.LoadImage(filename)
		if err != nil {
			if cfg.IsVerbose {
				fmt.Printf(" - error: can't load image\n")
			}
			continue
		}

		resizedImg := imgPkg.ResizeFile(img, cfg.Width, cfg.Height, cfg.IsLargeOnly)
		outFilename := pathUtils.NormalizePath(
			cfg.OutputPath + pathUtils.TrimPrefix(filename, cfg.SourcePath),
		)

		if len(cfg.Prefix) > 0 {
			outFilename = path.Dir(outFilename) + "/" + cfg.Prefix + path.Base(outFilename)
		}

		outDir := path.Dir(outFilename)
		if !fs.IsDirExists(outDir) {
			err = fs.MkDir(outDir)
			if err != nil {
				if cfg.IsVerbose {
					fmt.Printf(" - error: can't create dir %s\n", outDir)
				}
				continue
			}
		}

		if cfg.IsCustomType() {
			outFilename = strings.TrimSuffix(outFilename, filepath.Ext(outFilename)) + "." + cfg.Extension
		}

		err = imgPkg.SaveImage(outFilename, resizedImg, cfg.Quality)
		if err != nil {
			if cfg.IsVerbose {
				fmt.Printf(" - error: can't save image to %s\n", outFilename)
			}
			continue
		}

		if cfg.IsVerbose {
			fmt.Printf(" - saved to %s\n", outFilename)
		}
	}
}
