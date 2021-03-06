package servestatic

import (
	"net/http"
	"strings"
)

const Slash = "/"
const DotHtml = ".html"
const IndexDotHtml = "index" + DotHtml

func (m *FileServer) serveFile(w http.ResponseWriter, r *http.Request, reqPath string, next func(resolvedLocation string)) {
	if m.UsingHostFolder {
		// Using the r.Host as a folder in m.RootDir.
		exists, location := m.GetFilePathFromStaticsAndDir(r.Host, reqPath)
		if exists {
			http.ServeFile(w, r, location)
		} else {
			next(location)
		}
	} else {
		exists, location := m.GetFilePathFromStatics(reqPath)
		if exists {
			http.ServeFile(w, r, location)
		} else {
			next(location)
		}
	}
}

// Serve potential html file; looking for `.../${test-page}/${test-page}.html` and `.../${test-page}/index.html`.
func (m *FileServer) serveHtmlFile(w http.ResponseWriter, r *http.Request, reqPath, targetHtml string, next func(resolvedLocation string)) {
	upgradedPath := reqPath + Slash + targetHtml + DotHtml
	// Using the r.Host as a folder in m.RootDir.
	exists, location := m.IsFileFromStaticsRegular(r.Host, upgradedPath)
	if !exists {
		upgradedPath = reqPath + Slash + IndexDotHtml
		exists, location = m.IsFileFromStaticsRegular(r.Host, upgradedPath)
		if !exists {
			next(location)
			return
			//upgradedPath := reqPath + DotHtml
			//exists, location = m.IsFileFromStaticsRegular(r.Host, upgradedPath)
		}
	}
	http.ServeFile(w, r, location)
}

// Checking request host and responding with their corresponding files.
func (m *FileServer) ServeFiles(w http.ResponseWriter, r *http.Request, next func(resolvedLocation string)) {
	reqPath := r.URL.Path
	if !strings.HasPrefix(reqPath, Slash) {
		reqPath = "/" + reqPath
		r.URL.Path = reqPath
	}
	if reqPath == Slash {
		// 0. Serve the index.html for the home page '/'.
		m.serveFile(w, r, reqPath+Slash+IndexDotHtml, next)
		return
	}
	if strings.HasSuffix(reqPath, Slash) {
		// 1. Redirect all trailing-slash paths to no-trailing-slash paths.
		http.Redirect(w, r, reqPath[:len(reqPath)-1], http.StatusMovedPermanently)
		return
	} else if strings.HasSuffix(reqPath, DotHtml) {
		// 2. Ignore the *.html files(, which are not served).
		next("")
		return
	} else {
		targetHtml := GetTrailingNameInPath(reqPath)
		if strings.Contains(targetHtml, ".") {
			// 3. Serve regular files(, but not *.html).
			m.serveFile(w, r, reqPath, next)
			return
		}
		// 4. Serve the target html file if no dot is found in the request.path.
		m.serveHtmlFile(w, r, reqPath, targetHtml, next)
	}
}
