package tasks

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	taskDone       = "DONE"
	taskInprogress = "PROGRESS"
)

type ErrorMessage struct {
	msg string
}

type TaskContainer struct {
	Items []Task
}

func (receiver *ErrorMessage) Error() string {
	return receiver.msg
}

func (tk *TaskContainer) Add(taskName string) string {
	var lstItemId string
	if lstItem := len(tk.Items); lstItem > 0 {
		idx := lstItem - 1
		val, _ := strconv.Atoi(tk.Items[idx].Id)
		lstItemId = strconv.Itoa(val + 1)
	} else {
		lstItemId = "1"
	}
	nt := CreateTask(taskName, lstItemId)

	tk.Items = append(tk.Items, nt)
	return nt.Id
}

func (tk *TaskContainer) Update(taskId string, status string) (*Task, string) {
	searchTaskedId, err := tk.getTaskIndex(taskId)
	if err != nil {
		cE := ErrorMessage{
			msg: err.Error(),
		}
		return nil, cE.msg
	}

	// Check that the status argument is inline with the constants
	if b := strings.EqualFold(status, taskInprogress) || strings.EqualFold(status, taskDone); !b {
		err := errors.New(fmt.Sprintf("The status must either be %v or %v", taskDone, taskInprogress))
		cE := ErrorMessage{
			msg: err.Error(),
		}
		return nil, cE.msg
	}

	tk.Items[searchTaskedId].UpdateStatus(status)
	return &tk.Items[searchTaskedId], tk.Items[searchTaskedId].GetStatus()
}

func (tk *TaskContainer) getTaskIndex(taskId string) (int, error) {
	taskFound := false
	fetchedTaskIndex := -1
	for i := 0; i < len(tk.Items); i++ {
		if tk.Items[i].Id == taskId {
			fetchedTaskIndex = i
			taskFound = true
		}
	}

	if taskFound {
		return fetchedTaskIndex, nil
	}

	return fetchedTaskIndex, errors.New("no task id matched the id that was passed")
}

func (tk *TaskContainer) Delete(taskId string) (bool, string) {
	searchedTaskIndex, err := tk.getTaskIndex(taskId)
	if err != nil {
		return false, fmt.Sprintln(err)
	}

	i := searchedTaskIndex
	if i == len(tk.Items)-1 {
		newItem := tk.Items[:i]
		fmt.Println("new item", newItem)
		tk.Items = append(make([]Task, 0), newItem...)
	} else {
		tk.Items = append(tk.Items[:i], tk.Items[i+1:]...)
	}
	return true, fmt.Sprintf("updated task are %v ", tk.Items)
}

func (tk *TaskContainer) List() {
	for _, t := range tk.Items {
		t.ShowInfo()
	}
}

func (tk *TaskContainer) ListTasksDone() {
	for _, t := range tk.Items {
		if t.Status == taskDone {
			t.ShowInfo()
		}
	}
}

func (tk *TaskContainer) ListTasksInProgress() {
	for _, t := range tk.Items {
		if t.Status == taskInprogress {
			t.ShowInfo()
		}
	}
}

func (tk *TaskContainer) FetchData(data []Task) {
	tk.Items = append(tk.Items, data...)
}
