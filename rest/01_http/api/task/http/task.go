package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"strings"

	"github.com/neo-classic/golang-playground/rest/01_http/domain"
)

func (h *TaskHTTP) taskHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/task/" {
		switch req.Method {
		case http.MethodPost:
			h.createTask(w, req)
			break
		case http.MethodGet:
			h.getAllTasks(w, req)
			break
		case http.MethodDelete:
			h.deleteAllTasks(w, req)
			break
		default:
			http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /task/, got %v", req.Method), http.StatusMethodNotAllowed)
		}
	} else {
		path := strings.Trim(req.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
			http.Error(w, "expect /task/<id> in task handler", http.StatusBadRequest)
			return
		}
		guid := pathParts[1]

		switch req.Method {
		case http.MethodDelete:
			h.deleteTask(w, req, guid)
			break
		case http.MethodGet:
			h.getTask(w, req, guid)
			break
		default:
			http.Error(w, fmt.Sprintf("expect method GET or DELETE at /task/<id>, got %v", req.Method), http.StatusMethodNotAllowed)
		}
	}
}

func (h *TaskHTTP) createTask(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling task create at %s\n", req.URL.Path)
	ctx := context.Background()
	input := createTaskRequest{}

	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := mapCreateRequestToDomain(&input)
	task, err := h.service.Create(ctx, *t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(mapDomainToReply(task))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (h *TaskHTTP) getAllTasks(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get all tasks at %s\n", req.URL.Path)
	ctx := context.Background()

	allTasks, err := h.service.Fetch(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(mapDomainsToReply(allTasks))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (h *TaskHTTP) deleteAllTasks(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling delete all tasks at %s\n", req.URL.Path)
	ctx := context.Background()
	err := h.service.DeleteAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TaskHTTP) deleteTask(w http.ResponseWriter, req *http.Request, guid string) {
	log.Printf("handling delete task at %s\n", req.URL.Path)
	ctx := context.Background()

	err := h.service.Delete(ctx, domain.TaskGUID(guid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (h *TaskHTTP) getTask(w http.ResponseWriter, req *http.Request, guid string) {
	log.Printf("handling get task at %s\n", req.URL.Path)
	ctx := context.Background()

	task, err := h.service.GetByGUID(ctx, domain.TaskGUID(guid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(mapDomainToReply(task))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
