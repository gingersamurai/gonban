package storage

import (
	"gonban/internal/model"
	"sync"
)

type MemoryStorage struct {
	sync.RWMutex
	tasks  map[int]model.Task
	nextId int
}

func NewMemoryStorage() *MemoryStorage {
	ms := &MemoryStorage{}
	ms.tasks = make(map[int]model.Task)
	ms.nextId = 0
	return ms
}

func (ms *MemoryStorage) AddTask(newTask model.Task) int {
	ms.Lock()
	defer ms.Unlock()

	newTask.Id = ms.nextId
	ms.tasks[ms.nextId] = newTask
	ms.nextId++

	return ms.nextId
}

func (ms *MemoryStorage) GetTaskById(needId int) (model.Task, error) {
	ms.RLock()
	defer ms.RUnlock()

	if _, ok := ms.tasks[needId]; !ok {
		return model.Task{}, model.ErrTaskNotFound
	}

	return ms.tasks[needId], nil
}

func (ms *MemoryStorage) GetAllTasks() []model.Task {
	ms.RLock()
	defer ms.RUnlock()

	result := make([]model.Task, len(ms.tasks))
	for _, v := range ms.tasks {
		result = append(result, v)
	}
	return result
}
