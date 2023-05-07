package postgres

import (
	"Todo_list/internal/model"
	"Todo_list/internal/repository"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Task struct {
	db *sqlx.DB
}

func NewTask(db *sqlx.DB) *Task {
	return &Task{db: db}
}

func (t *Task) Create(task model.Task) (int, error) {

	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, done) VALUES ($1, $2) RETURNING id", taskTable)
	row := t.db.QueryRow(query, task.Name, task.Done)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *Task) GetAll() ([]model.Task, error) {
	var res []model.Task

	query := fmt.Sprintf("SELECT t.id, t.name, t.done FROM %s t", taskTable)
	err := t.db.Select(&res, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Task) GetCompleted() ([]model.Task, error) {
	var res []model.Task

	query := fmt.Sprintf("SELECT t.id, t.name, t.done FROM %s t WHERE t.done", taskTable)
	err := t.db.Select(&res, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Task) GetUncompleted() ([]model.Task, error) {
	var res []model.Task

	query := fmt.Sprintf("SELECT t.id, t.name, t.done FROM %s t WHERE t.done = FALSE", taskTable)
	err := t.db.Select(&res, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Task) Update(taskId int, updateTask model.UpdateTask) error {
	setValues := make([]string, 0, 2)
	args := make([]interface{}, 0, 3)
	argId := 1

	if updateTask.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *updateTask.Name)
		argId++
	}

	if updateTask.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *updateTask.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s t SET %s WHERE t.id = $%d",
		taskTable, setQuery, argId)

	args = append(args, taskId)

	res, err := t.db.Exec(query, args...)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return repository.ErrDataNotFound
	}

	return nil
}

func (t *Task) Delete(taskId int) error {
	query := fmt.Sprintf("DELETE FROM %s t WHERE t.id=$1", taskTable)

	res, err := t.db.Exec(query, taskId)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return repository.ErrDataNotFound
	}

	return nil
}
