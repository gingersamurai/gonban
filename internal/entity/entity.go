package entity

import (
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Performer   string    `json:"performer"`
	Deadline    time.Time `json:"deadline"`
}
