package http

import (
	"context"
	"fmt"
	"github.com/neo-classic/golang-playground/rest/01_http/internal/middleware"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/task/", h.taskHandler)
	mux.HandleFunc("/tag/", h.tagHandler)
	mux.HandleFunc("/due/", h.dueHandler)

	handler := middleware.Logging(mux)
	handler = middleware.PanicRecovery(handler)

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", cfg.Server.Port), handler)
	h.log.Error(err.Error())
}
