package utils

import (
	"encoding/json"

	"github.com/JanisJuska/Go-task-cli/task"
)

func ReturnIdCount(filename string) (uint, error) {
	fileData, err := OpenAndReadFile(filename)
	if err != nil {
		return 0, err
	}

	var dataSlice []task.Task
	err = json.Unmarshal(fileData, &dataSlice)
	if err != nil {
		return 0, err
	}

	return dataSlice[len(dataSlice)-1].ID, nil
}
