package task

import (
	"context"
	"fmt"

	"github.com/neo-classic/golang-playground/rest/02_gorilla_mux/adapters/repository/task"
	"github.com/neo-classic/golang-playground/rest/02_gorilla_mux/domain"
	"github.com/pkg/errors"
)

type Repository interface {
	Create(ctx context.Context, task *domain.Task) error
	Fetch(ctx context.Context, filter task.Filter) ([]*domain.Task, error)
	GetByGUID(ctx context.Context, guid domain.TaskGUID) (*domain.Task, error)
	Delete(ctx context.Context, guid domain.TaskGUID) error
	DeleteAll(ctx context.Context) error
}

type Service struct {
	repo Repository
}

func NewTaskService(r Repository) *Service {
	return &Service{
		repo: r,
	}
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
	var filter task.Filter
	tasks, err := s.repo.Fetch(ctx, filter)
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

func (s *Service) Delete(ctx context.Context, guid domain.TaskGUID) error {
	err := s.repo.Delete(ctx, guid)
	if err != nil {
		return errors.Wrap(err, "[task_service] delete error")
	}

	return nil
}

func (s *Service) DeleteAll(ctx context.Context) error {
	err := s.repo.DeleteAll(ctx)
	if err != nil {
		return errors.Wrap(err, "[task_service] delete all error")
	}

	return nil
}

func (s *Service) FetchByTag(ctx context.Context, tag string) ([]*domain.Task, error) {
	var filter task.Filter
	filter.Tag = tag
	tasks, err := s.repo.Fetch(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "[task_service] fetch error")
	}
	return tasks, nil
}

func (s *Service) FetchByDueDate(ctx context.Context, y, m, d int) ([]*domain.Task, error) {
	var filter task.Filter
	filter.DueDate = fmt.Sprintf("%d-%d-%d", y, m, d)
	tasks, err := s.repo.Fetch(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "[task_service] fetch error")
	}
	return tasks, nil
}
