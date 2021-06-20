package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/neo-classic/golang-playground/rest/01_http/config"
	"github.com/neo-classic/golang-playground/rest/01_http/domain"
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
	Debug(ctx context.Context, msg string)
	Info(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)
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

	mux := http.NewServeMux()
	mux.HandleFunc("/task/", h.taskHandler)
	mux.HandleFunc("/tag/", h.tagHandler)
	mux.HandleFunc("/due/", h.dueHandler)

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", cfg.Server.Port), mux)
	h.log.Error(ctx, err.Error())
}
