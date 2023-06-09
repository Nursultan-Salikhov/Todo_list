package repository

import (
	"Todo_list/internal/model"
	"errors"
)

var (
	ErrDataNotFound = errors.New("data not found")
)

type Task interface {
	Create(task model.Task) (int, error)
	GetAll() ([]model.Task, error)
	GetCompleted() ([]model.Task, error)
	GetUncompleted() ([]model.Task, error)
	Update(taskId int, updateTask model.UpdateTask) error
	Delete(taskId int) error
}

type Repository struct {
	Task
}

func NewRepository(task Task) *Repository {
	return &Repository{Task: task}
}
