package utils

import (
	"encoding/json"
	"io/fs"
	"os"
	"roadmap.sh/task-tracker/tasks"
)

const (
	fileName string = "task-tracker.json"
)

func WriteToFile(cnt tasks.TaskContainer) error {
	fileCnt, err := json.Marshal(cnt.Items)
	if err != nil {
		return err
	}

	if nErr := os.WriteFile(fileName, fileCnt, fs.ModePerm); nErr != nil {
		return nErr
	}

	return nil
}

func ReadFileToContainer() (content []byte, readFileErr error) {
	cnt, e := os.ReadFile(fileName)
	if e != nil {
		return nil, e
	}
	return cnt, nil
}
