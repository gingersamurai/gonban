package model

import (
	"time"

	"github.com/pkg/errors"
)

type Task struct {
	Id          int       `json:"id"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Performer   string    `json:"performer"`
	Deadline    time.Time `json:"deadline"`
}

type RequestTask struct {
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Performer   string    `json:"performer"`
	Deadline    time.Time `json:"deadline"`
}

var ErrTaskNotFound = errors.New("task not found")
