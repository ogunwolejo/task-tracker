package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"roadmap.sh/task-tracker/tasks"
	"roadmap.sh/task-tracker/utils"
)

func main() {
	container := tasks.TaskContainer{
		Items: make([]tasks.Task, 0),
	}

	// We need to check if file exist, then we will need to pass all the data from the file to the task
	content, readFileErr := utils.ReadFileToContainer()
	if readFileErr != nil {
		log.Fatal(readFileErr)
	}
	// we decode the bytes and then append it to the container
	var fileContent []tasks.Task
	if decodeErr := json.Unmarshal(content, &fileContent); decodeErr != nil {
		log.Fatal(decodeErr)
	}
	container.FetchData(fileContent)

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Error no arg was passed")
		os.Exit(1)
	}

	switch args[0] {
	case "add":
		addItem := container.Add(args[1])
		fmt.Println(addItem, container.Items)
	case "update":
		tk, val := container.Update(args[1], args[2])
		fmt.Println(*tk, val)
	case "delete":
		isDone, val := container.Delete(args[1])
		if isDone {
			fmt.Println(val)
		} else {
			fmt.Println(val)
		}
	case "list":
		container.List()
	case "list done":
		container.ListTasksDone()
	case "list inProgress":
		container.ListTasksInProgress()
	default:
		fmt.Println("Invalid command")
	}

	utils.WriteToFile(container)
}
