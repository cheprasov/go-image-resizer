package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cheprasov/go-image-resizer/pkg/pathUtils"
)

func MkDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) || err != nil {
		return false
	}
	return !info.IsDir()
}

func IsDirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) || err != nil {
		return false
	}
	return info.IsDir()
}

func IsJpegFile(filename string) bool {
	ext := GetFileExtension(filename)
	return ext == ".jpg" || ext == ".jpeg"
}

func IsPngFile(filename string) bool {
	ext := GetFileExtension(filename)
	return ext == ".png"
}

func IsImageFile(filename string) bool {
	ext := GetFileExtension(filename)
	if !IsAllowedFile(filename) {
		return false
	}
	return ext == ".png" || ext == ".jpg" || ext == ".jpeg"
}

type FolderContent struct {
	Folders []FileInfo
	Jpegs   []FileInfo
	Zips    []FileInfo
}

func IsAllowedFile(filename string) bool {
	if strings.HasPrefix(filename, "._") {
		return false
	}

	return true
}

func WriteFile(filename, data string) error {
	return ioutil.WriteFile(filename, []byte(data), os.ModePerm)
}

func ReadFile(filename string) (string, error) {
	if !IsFileExists(filename) {
		return "", nil
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func GetFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}

type FileInfo struct {
	FullName string
	Info     os.FileInfo
}
type FileInfoMap map[string]FileInfo

type DeepFolderInfo struct {
	BaseDir string
	InfoMap FileInfoMap
}

func IsIgnoredFile(info os.FileInfo) bool {
	if strings.HasPrefix(info.Name(), "._") {
		return true
	}
	return false
}

func ReadFolder(baseDir string, filesMapPointer *DeepFolderInfo, isDeep bool) (*DeepFolderInfo, error) {
	baseDir = pathUtils.NormalizePath(baseDir)
	if filesMapPointer == nil {
		filesMapPointer = &DeepFolderInfo{
			BaseDir: baseDir,
			InfoMap: FileInfoMap{},
		}
	}

	if baseDir == (*filesMapPointer).BaseDir {
		baseDirInfo, err := os.Stat(baseDir)
		if err != nil {
			return filesMapPointer, err
		}
		(*filesMapPointer).InfoMap[baseDir] = FileInfo{
			FullName: baseDir,
			Info:     baseDirInfo,
		}
	}

	fileInfos, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}

	var filename string
	for _, fileInfo := range fileInfos {
		filename = baseDir + "/" + fileInfo.Name()

		if IsIgnoredFile(fileInfo) {
			continue
		}

		(*filesMapPointer).InfoMap[filename] = FileInfo{
			FullName: filename,
			Info:     fileInfo,
		}

		if fileInfo.IsDir() && isDeep {
			_, err = ReadFolder(filename, filesMapPointer, isDeep)
			if err != nil {
				return nil, err
			}
		}
	}

	return filesMapPointer, nil
}
