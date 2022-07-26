package pathUtils

import (
	"log"
	"os"
	"path"
	"strings"
)

// /for/bar/ => /foo/bar
// foo => /foo
func NormalizePath(dir string) string {
	if len(dir) > 0 && dir[0] == '.' {
		p, err := os.Getwd()
		if err != nil {
			log.Println(err)
		} else {
			dir = p + "/" + dir
		}
	}
	dir = path.Clean(dir)
	if dir != "/" {
		dir = strings.TrimSuffix(dir, "/")
	}
	return dir
}

// /foo/bar => /bar
func TrimPrefix(fullPath, prefix string) string {
	return NormalizePath(strings.TrimPrefix(fullPath, NormalizePath(prefix)))
}
