package controller

import (
	"encoding/json"
	"fmt"
	"gonban/internal/model"
	"log"
	"mime"
	"net/http"
	"strings"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling create task [%v]\n", r.URL)

	contentType := r.Header.Get("Content-type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
		http.Error(w, "expect application/json Content-Type",
			http.StatusUnsupportedMediaType)
		return
	}
	dec := json.NewDecoder(r.Body)
	var rt model.RequestTask
	err = dec.Decode(&rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// storage.
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/tasks" && r.Method == http.MethodPost:
		// POST /tasks
		return
	case r.URL.Path == "/tasks" && r.Method == http.MethodGet:
		// GET /tasks
		return
	case strings.HasPrefix(r.URL.Path, "/tasks/") && r.Method == http.MethodGet:
		// GET /tasks/$id
		return

	default:
		http.Error(w,
			fmt.Sprint("wrong URL:", r.URL),
			http.StatusMethodNotAllowed,
		)
	}
}
