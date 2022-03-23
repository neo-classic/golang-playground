package domain

import (
	"fmt"
	"net/mail"
)

type User struct {
	ID    UserID
	Email Email
}

type UserID int

func (u UserID) String() string {
	return fmt.Sprintf("%d", u)
}

type Email string

func (e Email) String() string {
	return string(e)
}

func (e Email) IsValid() bool {
	_, err := mail.ParseAddress(e.String())
	return err == nil
}
