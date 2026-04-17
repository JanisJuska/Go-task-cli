package utils

import (
	"log"
	"os"
)

func OpenAndReadFile(filename string) []byte {
	dataFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Cannot open the file due to: %v\n", err)
	}
	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Cannot read the file due to: %v\n", err)
	}

	defer dataFile.Close()

	return fileData
}
