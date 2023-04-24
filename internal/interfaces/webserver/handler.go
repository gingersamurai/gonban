package webserver

import (
	"encoding/json"
	"fmt"
	"gonban/internal/entity"
	"gonban/internal/usecase"
	"log"
	"mime"
	"net/http"
	"strconv"
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
	log.Println("got request:", r.Method, r.URL.Path)

	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/tasks":
		// GET /tasks
		h.getAllTasksHandler(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/tasks/"):
		// GET /tasks/$id
		h.getTaskByIdHandler(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/tasks":
		// POST /tasks
		h.addTaskHandler(w, r)
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/tasks/"):
		// DELETE /tasks/$id
		h.deleteTaskHandler(w, r)
	default:
		log.Println("ERROR: bad request")
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	return
}

func (h *Handler) getAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("started getAllTasksHandler")
	log.Println("finished getAllTasksHandler")

	tasks := h.taskInteractor.GetAll()
	fmt.Println(len(tasks))
	js, err := json.Marshal(tasks)
	if err != nil {
		log.Println("ERROR json.Marshal():", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		log.Println("ERROR w.Write():", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) getTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("started getTaskByIdHandler")
	defer log.Println("finished getTaskByIdHandler")

	args := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(args) != 2 {
		errString := fmt.Sprintf("need 1 argument, got %d\n", len(args)-1)
		log.Println("ERROR: ", errString)
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		log.Println("ERROR strconv.Atoi():", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.taskInteractor.GetById(id)
	if err != nil {
		log.Println("ERROR GetById: not found")
		http.Error(w, fmt.Sprintf("task with id = %d not found\n", id), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(task)
	if err != nil {
		log.Println("ERROR json.Marshall():", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		log.Println("ERROR w.Write():", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) addTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("started addTaskHandler")
	defer log.Println("finished addTaskHandler")

	type ResponseId struct {
		Id int `json:"id"`
	}

	ContentType := r.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(ContentType)
	if err != nil {
		log.Println("ERROR mime.ParseMediaType():", err.Error())
		http.Error(w, "converting type: "+err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
		log.Printf("ERROR unexpected mediaType: got %s, need application/json\n", mediaType)
		http.Error(w, "expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(r.Body)
	//dec.DisallowUnknownFields()
	var rt entity.Task
	if err := dec.Decode(&rt); err != nil {
		log.Println("ERROR dec.Decode():", err.Error())
		http.Error(w, "decoding: "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.taskInteractor.Add(rt)
	if err != nil {
		log.Println("ERROR: task interactor.Add():", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	js, err := json.Marshal(ResponseId{Id: id})
	if err != nil {
		log.Println("ERROR json.Marshal():", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		log.Println("ERROR w.Write():", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("started deleteTaskHandler")
	defer log.Println("finished deleteTaskHandler")

	args := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(args) != 2 {
		errString := fmt.Sprintf("need 1 argument, got %d\n", len(args))
		log.Println("ERROR: ", errString)
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		log.Println("ERROR strconv.Atoi():", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.taskInteractor.DeleteById(id)
	if err != nil {
		log.Println("ERROR GetById: not found")
		http.Error(w, fmt.Sprintf("task with id = %d not found\n", id), http.StatusNotFound)
		return
	}

	return
}
