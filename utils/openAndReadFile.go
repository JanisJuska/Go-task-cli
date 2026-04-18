package utils

import (
	"os"
)

func OpenAndReadFile(filename string) ([]byte, error) {
	dataFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	defer dataFile.Close()

	return fileData, nil
}
