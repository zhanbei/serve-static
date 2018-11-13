package servestatic

import (
	"errors"
	"net/http"
	"os"
)

var (
	ErrNotExist = os.ErrNotExist

	ErrNotDir = errors.New("file is not a directory")
)

type FileServer struct {
	// Path of the root directory.
	RootDir string `json:"path"`
	// Whether to using the req.Host as folder when to serve files.
	UsingHostFolder bool `json:"host"`
}

//func NewFileServer(rootDir string) (*FileServer, error) {
//	return NewFileServerWithHostFolder(rootDir, false)
//}

func NewFileServer(rootDir string, usingHost bool) (*FileServer, error) {
	stat, err := os.Stat(rootDir)
	if os.IsNotExist(err) {
		return nil, ErrNotExist
	}
	if !stat.IsDir() {
		return nil, ErrNotDir
	}
	return &FileServer{
		RootDir: rootDir,

		UsingHostFolder: usingHost,
	}, nil
}

func (m *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.ServeFiles(w, r, func(resolvedLocation string) {
		http.NotFound(w, r)
	})
}
