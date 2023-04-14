package usecase

import "gonban/internal/entity"

type TaskInteractor struct {
	taskRepo entity.TaskRepo
}

func NewTaskInteractor(taskRepo entity.TaskRepo) *TaskInteractor {
	return &TaskInteractor{taskRepo: taskRepo}
}

func (t *TaskInteractor) GetAllTasks() []entity.Task {
	return t.taskRepo.FindAll()
}

func (t *TaskInteractor) GetTaskById(id int) (entity.Task, error) {
	return t.taskRepo.FindById(id)
}
