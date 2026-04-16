package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/JanisJuska/Go-task-cli/task"
	"github.com/JanisJuska/Go-task-cli/utils"
)

var tasks = returnListFromFile("todos.json")
var idCount uint = utils.ReturnIdCount("todos.json")

func main() {

	allArgs := os.Args[1:]
	firstArg := allArgs[0]
	afterArgs := allArgs[1:]
	argString := strings.Join(afterArgs, " ")

	switch strings.ToLower(firstArg) {
	case "add":
		idCount++
		var newTask task.Task
		newTask.ID = idCount
		newTask.Title = argString
		newTask.Done = false

		tasks = addNewTaskToJSON(newTask, "todos.json")

		fmt.Printf("'%v' succesfully added to the list\n", argString)
	case "list":

		fmt.Println()
		fmt.Printf("%-4s | %-30s | %s\n", "ID", "Title", "Done")
		fmt.Println("----------------------------------------------")

		for _, task := range tasks {
			fmt.Println(task.String())
		}

		fmt.Println()
	}

}

func addNewTaskToJSON(newTask task.Task, filename string) []task.Task {

	tasksList := returnListFromFile(filename)
	dataFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Cannot open the file due to: %v\n", err)
	}
	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Cannot read the file due to: %v\n", err)
	}

	defer dataFile.Close()

	tasksList = append(tasksList, newTask)

	fileData, err = json.MarshalIndent(tasksList, "", "  ")
	if err != nil {
		log.Fatalf("Cannot Marshal file due to: %v\n", err)
	}

	fmt.Printf("fileData: %v\n", string(fileData))

	_, err = dataFile.Write(fileData)
	if err != nil {
		log.Fatalf("Cannot write to file due to: %v\n", err)
	}

	return tasksList

}

func returnListFromFile(filename string) []task.Task {
	dataFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Cannot open the file due to: %v\n", err)
	}
	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Cannot read the file due to: %v\n", err)
	}

	defer dataFile.Close()

	var tasksList []task.Task
	json.Unmarshal(fileData, &tasksList)

	return tasksList
}
