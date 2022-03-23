package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/neo-classic/golang-playground/tests/mocks/user_api/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserAPIRepository(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/users/") {
			var user user
			user.ID = 1
			user.Email = "bel@gmail.com"

			content, err := json.Marshal(user)
			if err != nil {
				log.Fatal("marshal error")
			}

			io.WriteString(w, string(content))
			return
		}
		http.NotFound(w, r)
	}))

	repo := NewUserAPIRepository(server.URL)

	user, err := repo.GetByID(context.Background(), 1)
	assert.Nil(t, err)
	assert.Equal(t, domain.UserID(1), user.ID)

	fmt.Printf("%+v", user)
}
