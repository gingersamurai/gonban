package usecase

import "github.com/gingersamurai/gonban/internal/entity"

type TaskStorage interface {
	Add(task entity.Task) int
	GetById(id int) (entity.Task, error)
	GetAll() []entity.Task
	DeleteById(id int) error
}

type TaskInteractor struct {
	taskStorage TaskStorage
}

func NewTaskInteractor(taskStorage TaskStorage) *TaskInteractor {
	return &TaskInteractor{taskStorage: taskStorage}
}

func (t *TaskInteractor) Add(task entity.Task) int {
	return t.taskStorage.Add(task)
}

func (t *TaskInteractor) GetById(id int) (entity.Task, error) {
	return t.taskStorage.GetById(id)
}

func (t *TaskInteractor) GetAll() []entity.Task {
	return t.taskStorage.GetAll()
}

func (t *TaskInteractor) DeleteById(id int) error {
	return t.taskStorage.DeleteById(id)
}
