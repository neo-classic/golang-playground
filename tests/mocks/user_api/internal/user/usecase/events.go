package usecase

import (
	"github.com/neo-classic/golang-playground/tests/mocks/user_api/internal/domain"
	"github.com/neo-classic/golang-playground/tests/mocks/user_api/internal/events"
)

type UserEvent interface {
	events.Event
	UserID() domain.UserID
}

type UserGot struct {
	userID domain.UserID
}

func (e UserGot) Name() string {
	return "event.user.got"
}

func (e UserGot) UserID() domain.UserID {
	return e.userID
}

type UserCreated struct {
	userID domain.UserID
}

func (e UserCreated) Name() string {
	return "event.user.created"
}

func (e UserCreated) UserID() domain.UserID {
	return e.userID
}

type UserUpdated struct {
	userID domain.UserID
}

func (e UserUpdated) Name() string {
	return "event.user.updated"
}

func (e UserUpdated) UserID() domain.UserID {
	return e.userID
}

type UserDeleted struct {
	userID domain.UserID
}

func (e UserDeleted) Name() string {
	return "event.user.deleted"
}

func (e UserDeleted) UserID() domain.UserID {
	return e.userID
}
