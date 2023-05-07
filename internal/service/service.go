package service

import "Todo_list/internal/model"

type Task interface {
	Create(task model.Task) (int, error)
	GetAll() ([]*model.Task, error)
	GetCompleted() ([]*model.Task, error)
	GetUncompleted() ([]*model.Task, error)
	Update(taskId int, updateTask model.UpdateTask) error
	Delete(taskId int) error
}

type Service struct {
	Task
}
