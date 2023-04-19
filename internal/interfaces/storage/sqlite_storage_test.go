package storage_test

import (
	"fmt"
	"github.com/gingersamurai/gonban/internal/entity"
	"github.com/gingersamurai/gonban/internal/interfaces/storage"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestSqliteTaskStorage(t *testing.T) {
	t.Run("Add test", func(t *testing.T) {
		connInfo := "/home/gingersamurai/DataGripProjects/dg_intro/identifier.sqlite"
		ss, err := storage.NewSqliteTaskStorage(connInfo)
		assert.NoError(t, err)
		newTask := entity.Task{
			Id:          5,
			Status:      "IN PROGRESS",
			Name:        "buy strings",
			Description: "for 6-stringed guitar",
			Performer:   "Dima",
			Deadline:    time.Now(),
		}
		id := ss.Add(newTask)
		log.Println("ID ==", id)
	})

	t.Run("GetAll test", func(t *testing.T) {
		connInfo := "/home/gingersamurai/DataGripProjects/dg_intro/identifier.sqlite"
		ss, err := storage.NewSqliteTaskStorage(connInfo)
		assert.NoError(t, err)
		tasks := ss.GetAll()
		for _, v := range tasks {
			fmt.Printf("i: %-5v\ts: %-20v\tn: %-20v\td: %-30v\tp: %-20v\tdd: %-20v\n",
				v.Id, v.Status, v.Name, v.Description, v.Performer, v.Deadline)
		}
	})

	t.Run("GetById test", func(t *testing.T) {
		connInfo := "/home/gingersamurai/DataGripProjects/dg_intro/identifier.sqlite"
		ss, err := storage.NewSqliteTaskStorage(connInfo)
		assert.NoError(t, err)
		v, err := ss.GetById(6)
		assert.NoError(t, err)
		fmt.Printf("i: %-5v\ts: %-20v\tn: %-20v\td: %-30v\tp: %-20v\tdd: %-20v\n",
			v.Id, v.Status, v.Name, v.Description, v.Performer, v.Deadline)
	})
}
