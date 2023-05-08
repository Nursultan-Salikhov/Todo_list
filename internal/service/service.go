package service

import (
	"Todo_list/internal/model"
	"errors"
)

var (
	ErrUnvalidatedData = errors.New("non-validated input data")
	ErrNotFound        = errors.New("task not found")
)

type Task interface {
	Create(task model.Task) (int, error)
	GetAll() ([]model.Task, error)
	GetCompleted() ([]model.Task, error)
	GetUncompleted() ([]model.Task, error)
	Update(taskId int, updateTask model.UpdateTask) error
	Delete(taskId int) error
}

type Service struct {
	Task
}

func NewService(task Task) *Service {
	return &Service{Task: task}
}
