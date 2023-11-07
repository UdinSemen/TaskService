package repository

import (
	cfg "TaskService/internal/app/api/config"
	"TaskService/internal/app/api/mesStorage"
	"TaskService/internal/app/api/model"
	"database/sql"
	"fmt"
)

type Repository struct {
	Db *sql.DB
}

func InitRepository(cfg *cfg.Config) (*Repository, error) {
	const op = "repository.InitRepository"

	db, err := mesStorage.InitDb(cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Repository{Db: db}, nil
}

func (rep *Repository) AddTask(task model.Task) error {
	return nil
}

func (rep *Repository) GetTask(id string) (model.Task, error) {
	return model.Task{
		Uri:                 "1414",
		Id:                  "1414",
		Type:                "14141",
		CollectionFrequency: 0,
		Description:         "41414",
		Ti:                  nil}, nil
}

func (rep *Repository) DeleteTask(id string) error {
	return nil
}

func (rep *Repository) AddManyTasks(tasks []model.Task) error {
	return nil
}

func (rep *Repository) GetAllTasks() ([]model.Task, error) {
	return nil, nil
}
