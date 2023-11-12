package mesStorage

import (
	"TaskService/internal/app/api/model"
	"context"
)

type MemStorages interface {
	Ping() error
	AddTask(ctx context.Context, task model.Task) error
	GetTask(ctx context.Context, id string) (model.Task, error)
	UpdateTask(ctx context.Context, task model.Task) (model.Task, error)
	DeleteTask(ctx context.Context, id string) error
	AddManyTasks(ctx context.Context, tasks []model.Task) error
	GetAllTasks(ctx context.Context) ([]model.Task, error)
}
