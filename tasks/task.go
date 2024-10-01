package tasks

import "fmt"

type Task struct {
	Id     string
	Name   string
	Status string
}

func (tk *Task) UpdateStatus(newStatus string) {
	tk.Status = newStatus
}

func (tk *Task) GetStatus() string {
	return tk.Status
}

func (tk *Task) UpdateName(newName string) {
	tk.Name = newName
}

func (tk *Task) GetName() string {
	return tk.Name
}

func CreateTask(taskName string, newId string) (nt Task) {
	n := Task{
		Id:     newId,
		Status: taskInprogress,
		Name:   taskName,
	}
	return n
}

func (tk *Task) ShowInfo() {
	fmt.Println(*tk)
}
