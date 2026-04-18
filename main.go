package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/JanisJuska/Go-task-cli/task"
	"github.com/JanisJuska/Go-task-cli/utils"
)

func main() {

	if len(os.Args) <= 1 {
		log.Fatalf("No argument passed.\n")
	} else {

		var tasks, err = utils.ReturnListFromFile("todos.json")
		if err != nil {
			log.Fatal(err)
		}
		var idCount uint
		idCount, err = utils.ReturnIdCount("todos.json")
		if err != nil {
			log.Fatal(err)
		}

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

			tasks, err = utils.AddNewTaskToJSON(newTask, "todos.json")
			if err != nil {
				log.Fatal(err)
			}

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
			tasks, err = utils.SearchByIDAndMarkDone(argString, tasks)
			if err != nil {
				log.Fatal(err)
			}

		case "delete":
			tasks, err = utils.SearchByIDAndDelete(argString, tasks)
			if err != nil {
				log.Fatal(err)
			}
		case "help":
			utils.ShowHelp()
		default:
			log.Fatal("Wrong argument passed. Please pass 'help' as argument to read the manual\n")
		}
	}
}
