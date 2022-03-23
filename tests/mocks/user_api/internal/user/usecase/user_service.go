package usecase

import (
	"context"

	"github.com/neo-classic/golang-playground/tests/mocks/user_api/internal/domain"
	"github.com/neo-classic/golang-playground/tests/mocks/user_api/internal/events"
)

type UserRepository interface {
	GetByID(ctx context.Context, id domain.UserID) (*domain.User, error)
	Fetch(ctx context.Context) ([]*domain.User, error)
	Create(ctx context.Context, user domain.User) (*domain.User, error)
}

type EventPublisher interface {
	Notify(event events.Event)
}

type UserService struct {
	repo      UserRepository
	publisher EventPublisher
}

func (u *UserService) GetById(ctx context.Context, id domain.UserID) (*domain.User, error) {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	u.publisher.Notify(UserGot{
		userID: id,
	})

	return user, nil
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo:      repo,
		publisher: events.NewEventPublisher(),
	}
}
