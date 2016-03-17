package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func loadInMemory(path string, f os.FileInfo, err error) error {
	filePath, err := fileName(path, f)
	if err == nil {
		memoryFiles[filePath], _ = ioutil.ReadFile(path)
	}
	return nil
}

func fileName(path string, f os.FileInfo) (string, error) {
	prefix := configuration.HtmlRoot + "/"

	if f.IsDir() {
		return "", errors.New("File is directory")
	}

	return strings.TrimPrefix(path, prefix), nil
}
