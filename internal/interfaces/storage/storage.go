package storage

import (
	"fmt"
	"github.com/pkg/errors"
	"gonban/internal/entity"
	"log"
)

type MemoryTaskStorage struct {
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

func (ms *MemoryTaskStorage) Add(task entity.Task) int {
	task.Id = ms.nextId
	ms.nextId++
	ms.data[task.Id] = task
	log.Println("len:", len(ms.data))
	return task.Id
}

func (ms *MemoryTaskStorage) GetById(id int) (entity.Task, error) {
	result, ok := ms.data[id]
	if !ok {
		return entity.Task{}, errors.New(
			fmt.Sprintf("MemoryStorage: task with id = %d not found", id),
		)
	}
	log.Println("len:", len(ms.data))
	return result, nil
}

func (ms *MemoryTaskStorage) GetAll() []entity.Task {
	result := make([]entity.Task, 0, len(ms.data))
	for _, v := range ms.data {
		result = append(result, v)
	}
	log.Println("len:", len(ms.data))
	return result
}

func (ms *MemoryTaskStorage) DeleteById(id int) error {
	if _, ok := ms.data[id]; !ok {
		return errors.New(
			fmt.Sprintf("MemoryStorage: task with id = %d not found", id),
		)
	}
	delete(ms.data, id)
	log.Println("len:", len(ms.data))
	return nil
}
