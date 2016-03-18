package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"mime"
	"path/filepath"
)

func loadInMemory(path string, f os.FileInfo, err error) error {
	filePath, err := fileName(path, f)
	if err == nil && !fileIsBinary(filePath) {
		memoryFiles[filePath], _ = ioutil.ReadFile(path)
	}
	return nil
}

func fileName(path string, f os.FileInfo) (string, error) {
	prefix := configuration.HtmlRoot + string(os.PathSeparator)

	if f.IsDir() {
		return "", errors.New("File is directory")
	}

	return strings.TrimPrefix(path, prefix), nil
}

func fileIsBinary(path string) bool {
	ctype := mime.TypeByExtension(filepath.Ext(path))
	return strings.HasPrefix(ctype, "image");
}
