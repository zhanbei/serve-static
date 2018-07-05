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

// Whether the expected file exists?
func (m *FileServer) GetFilePathFromStatics(filePath string) (bool, string) {
	return m.GetFilePathFromStaticsAndDir("", filePath)
}

// Whether the expected file exists?
func (m *FileServer) GetFilePathFromStaticsAndDir(dirPath, fileName string) (bool, string) {
	des := filepath.Join(m.RootDir, dirPath, fileName)
	stat, err := os.Stat(des)
	return !os.IsNotExist(err) && !stat.IsDir(), des
}
