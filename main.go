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

	if len(os.Args) <= 1 {
		log.Fatalf("No argument passed.\n")
	} else {
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
		case "done":
			tasks = utils.SearchByIDAndMarkDone(argString, tasks)

		case "delete":
			tasks = utils.SearchByIDAndDelete(argString, tasks)
		case "help":
			readmeData := utils.OpenAndReadFile("./README.md")

			readme := string(readmeData)

			readme = strings.ReplaceAll(readme, "task-cli", "")
			readme = strings.ReplaceAll(readme, "```", "")
			readme = strings.ReplaceAll(readme, "bash", "")

			// usageIndex := strings.Index(readme, "Usage")
			// fmt.Println(usageIndex)
			// usageIndex = strings.Index(readme, "Notes")
			// fmt.Println(usageIndex)

			fmt.Println(readme[423:814])

		default:
			log.Fatalf("Wrong argument passed. Please pass 'help' as argument to read the manual\n")
		}
	}
}

func addNewTaskToJSON(newTask task.Task, filename string) []task.Task {

	tasksList := returnListFromFile(filename)
	fileData := utils.OpenAndReadFile(filename)

	tasksList = append(tasksList, newTask)

	fileData, err := json.MarshalIndent(tasksList, "", "  ")
	if err != nil {
		log.Fatalf("Cannot Marshal file due to: %v\n", err)
	}

	err = os.WriteFile(filename, fileData, 0644)
	if err != nil {
		log.Fatalf("Cannot write to file due to: %v\n", err)
	}

	return tasksList

}

func returnListFromFile(filename string) []task.Task {
	fileData := utils.OpenAndReadFile(filename)

	var tasksList []task.Task
	json.Unmarshal(fileData, &tasksList)

	return tasksList
}
