package webserver

import (
	"gonban/internal/usecase"
	"net/http"
)

type Handler struct {
	taskInteractor *usecase.TaskInteractor
}

func NewHandler(taskInteractor *usecase.TaskInteractor) *Handler {
	return &Handler{
		taskInteractor: taskInteractor,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	return
}
