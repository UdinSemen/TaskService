package memStorage

import (
	"TaskService/internal/app/api/model"
	"context"
	"fmt"
	"sync"
)

type MemStorage struct {
	sync.RWMutex
	tasks map[string]model.Task
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		tasks: make(map[string]model.Task),
	}
}

func (mem *MemStorage) addTask(_ context.Context, task model.Task) {
	mem.tasks[task.Id] = task
}

func (mem *MemStorage) getTask(_ context.Context, id string) (model.Task, error) {
	if len(mem.tasks) > 0 {
		task := mem.tasks[id]
		return task, nil
	}
	return model.Task{}, fmt.Errorf("task with id %s not found", id)
}

func (mem *MemStorage) AddTask(ctx context.Context, task model.Task) {
	mem.Lock()
	defer mem.Unlock()
	mem.addTask(ctx, task)
}

func (mem *MemStorage) GetTask(ctx context.Context, id string) (model.Task, error) {
	mem.RLock()
	defer mem.RUnlock()
	task, err := mem.GetTask(ctx, id)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}
