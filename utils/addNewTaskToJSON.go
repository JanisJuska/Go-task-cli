package utils

import (
	"encoding/json"
	"os"

	"github.com/JanisJuska/Go-task-cli/task"
)

func AddNewTaskToJSON(newTask task.Task, filename string) ([]task.Task, error) {

	tasksList, err := ReturnListFromFile(filename)
	if err != nil {
		return nil, err
	}
	fileData, err := OpenAndReadFile(filename)
	if err != nil {
		return nil, err
	}

	tasksList = append(tasksList, newTask)

	fileData, err = json.MarshalIndent(tasksList, "", "  ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(filename, fileData, 0644)
	if err != nil {
		return nil, err
	}

	return tasksList, nil
}
