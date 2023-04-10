package storage

import (
	"gonban/internal/model"
)

type Storage interface {
	GetAllTasks() []model.Task
	GetTaskById(needId int) (model.Task, error)
	AddTask(newTask model.Task) int
}
