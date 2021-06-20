package http

import (
	"context"
	"net/http"

	"github.com/neo-classic/golang-playground/rest/01_http/domain"
	"gopkg.in/go-playground/validator.v9"
)

type TaskService interface {
	Create(ctx context.Context, task domain.Task) (*domain.Task, error)
	Fetch(ctx context.Context) ([]*domain.Task, error)
	GetByGUID(ctx context.Context, guid domain.TaskGUID) (*domain.Task, error)
}

type TaskHTTP struct {
	service  TaskService
	validate *validator.Validate
}

func NewTaskHTTP(s TaskService, v *validator.Validate) {
	h := &TaskHTTP{
		service:  s,
		validate: v,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/task/", h.taskHandler)
}
