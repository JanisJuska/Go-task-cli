package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/JanisJuska/Go-task-cli/task"
)

func SearchByIDAndMarkDone(argString string, tasks []task.Task) ([]task.Task, error) {

	id, err := strconv.Atoi(argString)
	if err != nil {
		return nil, err
	}

	var taskTitle string

	for i, t := range tasks {
		if t.ID == uint(id) {
			t.Done = true

			taskTitle = t.Title

			tasks[i] = t
		}
	}

	fileData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile("todos.json", fileData, 0644)
	if err != nil {
		return nil, err
	}

	fmt.Printf("'%v' task marked as Done ✔️\n", taskTitle)

	return tasks, nil
}

func SearchByIDAndDelete(argString string, tasks []task.Task) ([]task.Task, error) {

	id, err := strconv.Atoi(argString)
	if err != nil {
		return nil, err
	}

	var taskTitle string

	for i, t := range tasks {
		if t.ID == uint(id) {

			taskTitle = t.Title

			tasks = slices.Delete(tasks, i, i+1)
		}
	}

	fileData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile("todos.json", fileData, 0644)
	if err != nil {
		return nil, err
	}

	fmt.Printf("'%v' task removed\n", taskTitle)

	return tasks, nil
}
