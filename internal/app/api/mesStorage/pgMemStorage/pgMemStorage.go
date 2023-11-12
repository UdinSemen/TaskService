package pgMemStorage

import (
	"TaskService/internal/app/api/config"
	"TaskService/internal/app/api/model"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PgMemStorage struct {
	connect *sql.DB
}

func NewPgMemStorage(cfg *config.DbConfig) (*PgMemStorage, error) {
	const op = "pgMemStorage.NewPgMemStorage"

	connect, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &PgMemStorage{
		connect: connect,
	}, nil
}

func (pg *PgMemStorage) Ping() error {
	if err := pg.connect.Ping(); err != nil {
		return err
	}
	return nil
}

func (pg *PgMemStorage) AddTask(ctx context.Context, task model.Task) error {
	return nil
}

func (pg *PgMemStorage) GetTask(ctx context.Context, id string) (model.Task, error) {
	return model.Task{}, nil
}

func (pg *PgMemStorage) UpdateTask(ctx context.Context, task model.Task) (model.Task, error) {
	return model.Task{}, nil
}

func (pg *PgMemStorage) DeleteTask(ctx context.Context, id string) error {
	return nil
}

func (pg *PgMemStorage) AddManyTasks(ctx context.Context, tasks []model.Task) error {
	return nil
}

func (pg *PgMemStorage) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	return []model.Task{}, nil
}
