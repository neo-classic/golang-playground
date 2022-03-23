package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail_IsValid(t *testing.T) {
	var e Email = "test@gmail.com"
	isValid := e.IsValid()
	assert.True(t, isValid)

	e = "asd@asd."
	isValid = e.IsValid()
	assert.False(t, isValid)
}

func TestUserID_String(t *testing.T) {
	var u User
	u.ID = 1

	assert.Equal(t, "1", u.ID.String())
}
