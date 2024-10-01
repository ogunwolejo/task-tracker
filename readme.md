# Task Tracker
### Build a CLI app to track your tasks and manage your to-do list.

## Description
#### Task tracker is a project used to track and manage your tasks. In this task, you will build a simple command line interface (CLI) to track what you need to do, what you have done, and what you are currently working on. This project will help you practice your programming skills, including working with the filesystem, handling user inputs, and building a simple CLI application.


## Requirements
#### The application should run from the command line, accept user actions and inputs as arguments, and store the tasks in a JSON file. The user should be able to:
* Add, Update, and Delete tasks
* Mark a task as in progress or done
* List all tasks
* List all tasks that are done
* List all tasks that are in progress


#### Here are some constraints to guide the implementation:
* You can use any programming language to build this project.
* Use positional arguments in command line to accept user inputs.
* Use a JSON file to store the tasks in the current directory.
* The JSON file should be created if it does not exist.
* Use the native file system module of your programming language to interact with the JSON file
* Do not use any external libraries or frameworks to build this project.
* Ensure to handle errors and edge cases gracefully.


## Instructions to use the cli 
### ADD a new task
    go run main.go add <Task name>
### UPDATE Status of a task
    go run main.go update <Task ID> <TASK STATUS>
    note: <TASK STATUS> must be either DONE or PROGRESS
### DELETE a task 
    go run main.go delete <Task ID>
### SHOW all tasks
    go run main.go list
### SHOW all tasks in progress
    go run main.go "list inProgress"
### SHOW all tasks done
    go run main.go "list done"