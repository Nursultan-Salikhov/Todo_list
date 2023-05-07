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
	err := task.Validate()
	if err != nil {
		return 0, ErrUnvalidatedData
	}

	return t.repo.Create(task)
}

func (t *TaskService) GetAll() ([]model.Task, error) {
	return t.repo.GetAll()
}

func (t *TaskService) GetCompleted() ([]model.Task, error) {
	return t.repo.GetCompleted()
}

func (t *TaskService) GetUncompleted() ([]model.Task, error) {
	return t.repo.GetUncompleted()
}

func (t *TaskService) Update(taskId int, updateTask model.UpdateTask) error {
	err := updateTask.Validate()
	if err != nil {
		return ErrUnvalidatedData
	}

	err = t.repo.Update(taskId, updateTask)
	if err == repository.ErrDataNotFound {
		return ErrUnvalidatedData
	}

	return err
}

func (t *TaskService) Delete(taskId int) error {
	err := t.repo.Delete(taskId)
	if err == repository.ErrDataNotFound {
		return ErrUnvalidatedData
	}

	return err
}
