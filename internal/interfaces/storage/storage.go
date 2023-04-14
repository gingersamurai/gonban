package storage

import (
	"fmt"
	"github.com/pkg/errors"
	"gonban/internal/entity"
)

type MemoryStorage struct {
	data map[int]entity.Task
}

func NewMemoryStorage() *MemoryStorage {
	data := make(map[int]entity.Task)
	return &MemoryStorage{data: data}
}

func (ms *MemoryStorage) Store(task entity.Task) {
	ms.data[task.Id] = task
}

func (ms *MemoryStorage) FindById(id int) (entity.Task, error) {
	result, ok := ms.data[id]
	if !ok {
		return entity.Task{}, errors.New(
			fmt.Sprintf("MemoryStorage: task with id = %d not found", id),
		)
	}
	return result, nil
}

func (ms *MemoryStorage) FindAll() []entity.Task {
	result := make([]entity.Task, len(ms.data))
	for _, v := range ms.data {
		result = append(result, v)
	}
	return result
}
