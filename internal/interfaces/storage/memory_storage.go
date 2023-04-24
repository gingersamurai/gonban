package storage

import (
	"fmt"
	"github.com/pkg/errors"
	"gonban/internal/entity"
	"sync"
)

type MemoryTaskStorage struct {
	sync.RWMutex
	data   map[int]entity.Task
	nextId int
}

func NewMemoryTaskStorage() *MemoryTaskStorage {
	data := make(map[int]entity.Task)
	nextId := 1
	return &MemoryTaskStorage{
		data:   data,
		nextId: nextId,
	}
}

func (ms *MemoryTaskStorage) Add(task entity.Task) (int, error) {
	ms.Lock()
	defer ms.Unlock()

	task.Id = ms.nextId
	ms.nextId++
	ms.data[task.Id] = task
	return task.Id, nil
}

func (ms *MemoryTaskStorage) GetById(id int) (entity.Task, error) {
	ms.RLock()
	defer ms.RUnlock()

	result, ok := ms.data[id]
	if !ok {
		return entity.Task{}, errors.New(
			fmt.Sprintf("MemoryStorage: task with id = %d not found", id),
		)
	}
	return result, nil
}

func (ms *MemoryTaskStorage) GetAll() ([]entity.Task, error) {
	ms.RLock()
	defer ms.RUnlock()

	result := make([]entity.Task, 0, len(ms.data))
	for _, v := range ms.data {
		result = append(result, v)
	}
	return result, nil
}

func (ms *MemoryTaskStorage) DeleteById(id int) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.data[id]; !ok {
		return fmt.Errorf("MemoryTaskStorage: task with id = %d not found", id)
	}

	delete(ms.data, id)
	return nil
}
