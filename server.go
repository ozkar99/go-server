package main

import (
	"net/http"
	"mime"
	"path/filepath"
	"os"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	ctype := mime.TypeByExtension(filepath.Ext(path))

	if ctype != ""  {
		w.Header().Set("Content-Type", ctype)
		if fileIsBinary(path) {
			filePath := configuration.HtmlRoot + string(os.PathSeparator) + path
			http.ServeFile(w, r, filePath)
			return;
		}
	}

	content := loadPage(path)
	fmt.Fprintf(w, "%s", content)
}

func loadPage(path string) []byte {
	if val, ok := memoryFiles[path]; ok {
		return val
	}
	return memoryFiles["index.html"]
}
