package task

import (
	"context"

	"github.com/neo-classic/golang-playground/rest/01_http/domain"
	"github.com/pkg/errors"
)

type Repository interface {
	Create(ctx context.Context, task *domain.Task) error
	Fetch(ctx context.Context) ([]*domain.Task, error)
	GetByGUID(ctx context.Context, guid domain.TaskGUID) (*domain.Task, error)
}

type Service struct {
	repo Repository
}

func (s *Service) Create(ctx context.Context, task domain.Task) (*domain.Task, error) {
	task.GUID = domain.NewTaskGUID()
	err := s.repo.Create(ctx, &task)
	if err != nil {
		return nil, errors.Wrap(err, "[task_service] create error")
	}

	return &task, nil
}

func (s *Service) Fetch(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := s.repo.Fetch(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "[task_service] fetch error")
	}
	return tasks, nil
}

func (s *Service) GetByGUID(ctx context.Context, guid domain.TaskGUID) (*domain.Task, error) {
	task, err := s.repo.GetByGUID(ctx, guid)
	if err != nil {
		return nil, errors.Wrap(err, "[task_service] getByGUID error")
	}

	return task, nil
}
