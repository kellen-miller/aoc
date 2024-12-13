package io

import (
	"bufio"
	"log"
	"os"
)

func OpenFile(relativePath string) (*os.File, func()) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fullPath := pwd + "/" + relativePath
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

func GetScanner(relativePath string) (*bufio.Scanner, func()) {
	file, closeFile := OpenFile(relativePath)
	return bufio.NewScanner(file), closeFile
}
