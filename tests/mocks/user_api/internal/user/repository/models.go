package repository

import "github.com/neo-classic/golang-playground/tests/mocks/user_api/internal/domain"

type user struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (u *user) toDomain() *domain.User {
	return &domain.User{
		ID:    domain.UserID(u.ID),
		Email: domain.Email(u.Email),
	}
}
