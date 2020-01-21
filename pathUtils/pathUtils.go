package pathUtils

import (
    "path"
    "strings"
)

// /for/bar/ => /foo/bar
// foo => /foo
func NormalizePath(dir string) string {
    dir = path.Clean("/" + dir)
    if dir != "/" {
        dir = strings.TrimSuffix(dir, "/");
    }
    return dir
}

// /foo/bar, /foo => /bar
func TrimPrefix(fullPath, prefix string) string {
    return NormalizePath(strings.TrimPrefix(fullPath, NormalizePath(prefix)))
}