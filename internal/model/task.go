package model

import "fmt"

const maxNameLen = 1000

type Task struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Done bool   `json:"done" db:"done"`
}

type UpdateTask struct {
	Name *string `json:"name"`
	Done *bool   `json:"done"`
}

func (t Task) Validate() error {
	if len(t.Name) > maxNameLen {
		return fmt.Errorf("invalid data")
	}

	return nil
}

func (ut UpdateTask) Validate() error {
	if ut.Done == nil && ut.Name == nil {
		return fmt.Errorf("UpdateTask struct don't have values")
	}

	if ut.Name != nil {
		if len(*ut.Name) > maxNameLen {
			return fmt.Errorf("invalid data")
		}
	}

	return nil
}
