package servestatic

import (
	"os"
	"path/filepath"
	"regexp"
)

var mNameInPathRegexp = regexp.MustCompile(".*/([^/]+)")
var mGroupLength = 2

// Get the trailing name in path; get 'world' in '/hello/world'.
func GetTrailingNameInPath(path string) string {
	res := mNameInPathRegexp.FindSubmatch([]byte(path))
	if len(res) != mGroupLength {
		return ""
	}
	return string(res[1])
}

// Whether the target file exists and is a regular file.
func IsFileRegular(paths ... string) (bool, string) {
	des := filepath.Join(paths...)
	stat, err := os.Stat(des)
	return !os.IsNotExist(err) && stat.Mode().IsRegular(), des
}

func (m *FileServer) GetFilePathFromStatics(filePath string) (bool, string) {
	return IsFileRegular(m.RootDir, filePath)
}

func (m *FileServer) GetFilePathFromStaticsAndDir(dirPath, fileName string) (bool, string) {
	return IsFileRegular(m.RootDir, dirPath, fileName)
}
