package io

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func OpenFile(inputPath string) (*os.File, func()) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fullPath := inputPath
	if !filepath.IsAbs(inputPath) {
		fullPath = filepath.Join(pwd, inputPath)
	}

	file, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}

	closeFunc := func() {
		if err := file.Close(); err != nil {
			log.Println(err.Error())
		}
	}

	return file, closeFunc
}

func GetScanner(path string) (*bufio.Scanner, func()) {
	file, closeFile := OpenFile(path)
	return bufio.NewScanner(file), closeFile
}
