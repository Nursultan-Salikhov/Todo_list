package postgres

import (
	"Todo_list/internal/model"
	"github.com/jmoiron/sqlx"
)

type Task struct {
	db *sqlx.DB
}

func NewTask(db *sqlx.DB) *Task {
	return &Task{db: db}
}

func (t *Task) Create(task model.Task) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Task) GetAll() ([]*model.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Task) GetCompleted() ([]*model.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Task) GetUncompleted() ([]*model.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Task) Update(taskId int, updateTask model.UpdateTask) error {
	//TODO implement me
	panic("implement me")
}

func (t *Task) Delete(taskId int) error {
	//TODO implement me
	panic("implement me")
}
