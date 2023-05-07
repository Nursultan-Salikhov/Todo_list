package service

import (
	"Todo_list/internal/model"
	"Todo_list/internal/repository"
)

type TaskService struct {
	repo *repository.Repository
}

func NewTaskService(repo *repository.Repository) *TaskService {
	return &TaskService{repo: repo}
}

func (t *TaskService) Create(task model.Task) (int, error) {
	return t.repo.Create(task)
}

func (t *TaskService) GetAll() ([]model.Task, error) {
	return t.repo.GetAll()
}

func (t *TaskService) GetCompleted() ([]model.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskService) GetUncompleted() ([]model.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskService) Update(taskId int, updateTask model.UpdateTask) error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskService) Delete(taskId int) error {
	//TODO implement me
	panic("implement me")
}
