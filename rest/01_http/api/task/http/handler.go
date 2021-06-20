package http

import (
	"context"
	"net/http"

	"github.com/neo-classic/golang-playground/rest/01_http/domain"
)

type TaskService interface {
	Create(ctx context.Context, task domain.Task) (*domain.Task, error)
	Fetch(ctx context.Context) ([]*domain.Task, error)
	GetByGUID(ctx context.Context, guid domain.TaskGUID) (*domain.Task, error)
}

type TaskHTTP struct {
	service TaskService
}

func NewTaskHTTP(service TaskService) {
	h := &TaskHTTP{
		service: service,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/task/", h.taskHandler)
}
