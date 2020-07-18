package config

import (
	"flag"
	"log"

	"github.com/cheprasov/go-image-resizer/pkg/pathUtils"
)

type Config struct {
	SourcePath  string
	OutputPath  string
	Width       uint
	Height      uint
	Quality     uint8
	Prefix      string
	Extension   string
	IsLargeOnly bool
	IsVerbose   bool
}

var typeByExt = map[string]bool{
	"jpg": true,
	"png": true,
}

func GetConfig() *Config {
	sourcePath := flag.String("source-path", "", "Please provide --source-path")
	outputPath := flag.String("output-path", "", "Please provide --output-path")
	width := flag.Uint("width", 0, "Please provide --width")
	height := flag.Uint("height", 0, "Please provide --height")
	quality := flag.Uint("quality", 100, "Please provide --quality")
	largeOnly := flag.Bool("large-only", false, "Please provide --large-only")
	convertTo := flag.String("convert-to", "", "Please provide --convert-to")
	prefix := flag.String("prefix", "", "Please provide --prefix")
	verbose := flag.Bool("verbose", false, "Please provide --verbose")
	flag.Parse()

	config := Config{
		SourcePath:  pathUtils.NormalizePath(*sourcePath),
		OutputPath:  pathUtils.NormalizePath(*outputPath),
		Width:       *width,
		Height:      *height,
		Quality:     uint8(*quality),
		Extension:   *convertTo,
		IsLargeOnly: *largeOnly,
		Prefix:      *prefix,
		IsVerbose:   *verbose,
	}
	return &config
}

func GetValidConfig() *Config {
	c := GetConfig()

	if len(c.SourcePath) == 0 {
		log.Fatal("Please provide --source-path")
	}

	if len(c.OutputPath) == 0 {
		log.Fatal("Please provide --output-path")
	}

	if c.SourcePath == c.OutputPath && len(c.Prefix) == 0 {
		log.Fatal("--source-path and --source-path should be not equal, or please use --prefix")
	}

	if _, ok := typeByExt[c.Extension]; !ok && len(c.Extension) > 0 {
		log.Fatal("Incorrect type for --convert-to, allowed types: jpg, png")
	}

	return c
}

func (c *Config) IsCustomType() bool {
	return len(c.Extension) != 0
}
