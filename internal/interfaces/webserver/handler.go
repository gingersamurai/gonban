package webserver

import (
	"gonban/internal/usecase"
	"net/http"
	"strings"
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
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/tasks":
		h.getAllTasksHandler(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/tasks/"):
		h.getTaskByIdHandler(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/tasks":
		h.addTaskHandler(w, r)
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/tasks/"):
		h.deleteTaskHandler(w, r)
	default:
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	return
}

func (h *Handler) getAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) getTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	return
}
func (h *Handler) addTaskHandler(w http.ResponseWriter, r *http.Request) {
	return
}
func (h *Handler) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	return
}
