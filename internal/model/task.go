package model

import "fmt"

const maxNameLen = 1000

type Task struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" binding:"required,max=1000" db:"name"`
	Done bool   `json:"done" db:"done"`
}

type UpdateTask struct {
	Name *string `json:"name"`
	Done *string `json:"done"`
}

func (ut UpdateTask) Validate() error {
	if ut.Done == nil && ut.Name == nil {
		return fmt.Errorf("UpdateTask struct don't have values")
	}

	return nil
}
