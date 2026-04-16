package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/JanisJuska/Go-task-cli/task"
)

func ReturnIdCount(filename string) uint {
	dataFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Cannot open the file due to: %v\n", err)
	}
	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Cannot read the file due to: %v\n", err)
	}

	defer dataFile.Close()

	var dataSlice []task.Task
	err = json.Unmarshal(fileData, &dataSlice)
	if err != nil {
		return 0
	}

	return dataSlice[len(dataSlice)-1].ID
}
