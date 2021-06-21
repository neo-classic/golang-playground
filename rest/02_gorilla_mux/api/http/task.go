package http

import (
	"context"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neo-classic/golang-playground/rest/02_gorilla_mux/domain"
)

func (h *TaskHTTP) createTask(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	h.log.Info(fmt.Sprintf("handling task create at %s\n", req.URL.Path))
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

	renderJSON(w, mapDomainToReply(task))
}

func (h *TaskHTTP) getAllTasks(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	h.log.Info(fmt.Sprintf("handling get all tasks at %s\n", req.URL.Path))

	tasks, err := h.service.Fetch(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(w, mapDomainsToReply(tasks))
}

func (h *TaskHTTP) deleteAllTasks(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	h.log.Info(fmt.Sprintf("handling delete all tasks at %s\n", req.URL.Path))
	err := h.service.DeleteAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TaskHTTP) deleteTask(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	h.log.Info(fmt.Sprintf("handling delete task at %s\n", req.URL.Path))

	guid := mux.Vars(req)["guid"]
	err := h.service.Delete(ctx, domain.TaskGUID(guid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (h *TaskHTTP) getTask(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	h.log.Info(fmt.Sprintf("handling get task at %s\n", req.URL.Path))

	guid := mux.Vars(req)["guid"]
	task, err := h.service.GetByGUID(ctx, domain.TaskGUID(guid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(w, mapDomainToReply(task))
}
