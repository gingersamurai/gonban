package storage_test

import (
	"github.com/stretchr/testify/assert"
	"gonban/internal/entity"
	"gonban/internal/interfaces/storage"
	"testing"
	"time"
)

func TestMemoryTaskStorage(t *testing.T) {
	t.Run("simpleTest", func(t *testing.T) {
		deadline, _ := time.Parse("02/01/2006", "15/05/2023")
		testData := entity.Task{
			Id:          5,
			Status:      "TODO",
			Name:        "do math homework",
			Description: "demidovich p432 ex12",
			Performer:   "Nazim Malyshev",
			Deadline:    deadline,
		}

		ms := storage.NewMemoryTaskStorage()
		ms.Add(testData)
		testData.Id = 1
		resultData, err := ms.GetById(testData.Id)
		assert.NoError(t, err)
		assert.Equal(t, testData, resultData)

		err = ms.DeleteById(5)
		assert.Error(t, err)

		err = ms.DeleteById(1)
		assert.NoError(t, err)
		_, err = ms.GetById(1)
		assert.Error(t, err)
	})
}
