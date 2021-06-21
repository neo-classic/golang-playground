package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neo-classic/golang-playground/rest/02_gorilla_mux/config"
	"github.com/neo-classic/golang-playground/rest/02_gorilla_mux/domain"
	"gopkg.in/go-playground/validator.v9"
)

type TaskService interface {
	Create(ctx context.Context, task domain.Task) (*domain.Task, error)
	Fetch(ctx context.Context) ([]*domain.Task, error)
	FetchByTag(ctx context.Context, tag string) ([]*domain.Task, error)
	FetchByDueDate(ctx context.Context, y, m, d int) ([]*domain.Task, error)
	GetByGUID(ctx context.Context, guid domain.TaskGUID) (*domain.Task, error)
	Delete(ctx context.Context, guid domain.TaskGUID) error
	DeleteAll(ctx context.Context) error
}

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Error(msg string)
}

type TaskHTTP struct {
	service  TaskService
	validate *validator.Validate
	cfg      *config.Config
	log      Logger
}

func NewTaskHTTP(ctx context.Context, s TaskService, v *validator.Validate, cfg *config.Config, log Logger) {
	h := &TaskHTTP{
		service:  s,
		validate: v,
		cfg:      cfg,
		log:      log,
	}

	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/task/", h.createTask).Methods("POST")
	router.HandleFunc("/task/", h.getAllTasks).Methods("GET")
	router.HandleFunc("/task/", h.deleteAllTasks).Methods("DELETE")
	router.HandleFunc("/task/{guid}/", h.getTask).Methods("GET")
	router.HandleFunc("/task/{guid}/", h.deleteTask).Methods("DELETE")

	router.HandleFunc("/tag/{tag}/", h.tagHandler).Methods("GET")
	router.HandleFunc("/due/{year:[0-9]+}/{month:[0-9]+}/{day:[0-9]+}/", h.dueHandler).Methods("GET")

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", cfg.Server.Port), router)
	h.log.Error(err.Error())
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
