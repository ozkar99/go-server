package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

var memoryFiles = make(map[string][]byte)

func main() {
	buildConfiguration()                                //Load configuration
	filepath.Walk(configuration.HtmlRoot, loadInMemory) //load files on memory.

	fmt.Println("Listening on ", configuration.Listener, " with html root as ", configuration.HtmlRoot, " ...")

	http.HandleFunc("/", handler)
	http.ListenAndServe(configuration.Listener, nil)
}
