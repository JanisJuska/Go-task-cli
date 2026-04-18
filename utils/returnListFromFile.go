package utils

import (
	"encoding/json"

	"github.com/JanisJuska/Go-task-cli/task"
)

func ReturnListFromFile(filename string) ([]task.Task, error) {
	fileData, err := OpenAndReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasksList []task.Task
	err = json.Unmarshal(fileData, &tasksList)
	if err != nil {
		return nil, err
	}

	return tasksList, nil
}
