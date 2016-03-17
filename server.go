package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	content := loadPage(path)
	fmt.Fprintf(w, "%s", content)
}

func loadPage(path string) []byte {

	if val, ok := memoryFiles[path]; ok {
		return val
	}

	return memoryFiles["index.html"]
}
