package http

import (
	"context"
	"fmt"
	"log"
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

type TaskHTTP struct {
	service  TaskService
	validate *validator.Validate
	cfg      *config.Config
}

func NewTaskHTTP(s TaskService, v *validator.Validate, cfg *config.Config) {
	h := &TaskHTTP{
		service:  s,
		validate: v,
		cfg:      cfg,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/task/", h.taskHandler)
	mux.HandleFunc("/tag/", h.tagHandler)
	mux.HandleFunc("/due/", h.dueHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", cfg.Server.Port), mux))
}
