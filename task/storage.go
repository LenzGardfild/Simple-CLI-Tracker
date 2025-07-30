package task

import (
	"encoding/json"
	"os"
)

func SaveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadTasks(filename string) ([]Task, error) {
	var tasks []Task
	data, err := os.ReadFile(filename)
	if err != nil {
		return tasks, err
	}
	err = json.Unmarshal(data, &tasks)
	return tasks, nil
}