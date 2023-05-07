package model

import "fmt"

const maxNameLen = 1000

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type UpdateTask struct {
	Name *string `json:"name"`
	Done *string `json:"done"`
}

func (t Task) Validate() error {
	if len(t.Name) > maxNameLen {
		return fmt.Errorf("task name len exceeds the max value(%d): len = %d", maxNameLen, len(t.Name))
	}

	return nil
}

func (ut UpdateTask) Validate() error {
	if ut.Done == nil && ut.Name == nil {
		return fmt.Errorf("UpdateTask struct don't have values")
	}

	if ut.Name != nil {
		if len(*ut.Name) > maxNameLen {
			return fmt.Errorf("task name len exceeds the max value(%d): len = %d", maxNameLen, len(*ut.Name))
		}
	}

	return nil
}
