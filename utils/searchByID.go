package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/JanisJuska/Go-task-cli/task"
)

func SearchByIDAndMarkDone(argString string, tasks []task.Task) []task.Task {

	id, err := strconv.Atoi(argString)
	if err != nil {
		log.Fatalf("Cannot convert string to number due to: %v\n", err)
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

	err = os.WriteFile("todos.json", fileData, 0644)
	if err != nil {
		log.Fatalf("Cannot write to file due to: %v\n", err)
	}

	fmt.Printf("'%v' task marked as Done ✔️\n", taskTitle)

	return tasks
}

func SearchByIDAndDelete(argString string, tasks []task.Task) []task.Task {

	id, err := strconv.Atoi(argString)
	if err != nil {
		log.Fatalf("Cannot convert string to number due to: %v\n", err)
	}

	var taskTitle string

	for i, t := range tasks {
		if t.ID == uint(id) {

			taskTitle = t.Title

			tasks = slices.Delete(tasks, i, i+1)
		}
	}

	fileData, err := json.MarshalIndent(tasks, "", "  ")

	err = os.WriteFile("todos.json", fileData, 0644)
	if err != nil {
		log.Fatalf("Cannot write to file due to: %v\n", err)
	}

	fmt.Printf("'%v' task removed\n", taskTitle)

	return tasks
}
