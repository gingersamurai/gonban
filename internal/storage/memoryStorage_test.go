package storage_test

import (
	"gonban/internal/model"
	"gonban/internal/storage"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		ms := storage.NewMemoryStorage()
		needDeadline, _ := time.Parse(
			"02/01/2006",
			"09/05/2023",
		)
		needTask := model.Task{
			Id:          5,
			Status:      "TODO",
			Name:        "buy beer",
			Description: "I need by some beer",
			Performer:   "Ivan Khomyakov",
			Deadline:    needDeadline,
		}
		ms.AddTask(needTask)
		needTask.Id = 0
		haveTask, err := ms.GetTaskById(0)
		assert.NoError(t, err)
		assert.Equal(t, needTask, haveTask)
	})
	t.Run("task not found", func(t *testing.T) {
		ms := storage.NewMemoryStorage()

		_, err := ms.GetTaskById(0)
		assert.Error(t, err)
	})

}
