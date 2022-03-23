package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/neo-classic/golang-playground/tests/mocks/user_api/internal/domain"
)

type UserAPIRepository struct {
	host string
}

func NewUserAPIRepository(host string) *UserAPIRepository {
	return &UserAPIRepository{
		host: host,
	}
}

func (r *UserAPIRepository) GetByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	var user user

	//url := path.Join(r.host, "/users/", id.String())
	url := fmt.Sprintf("%s/%s/%s", r.host, "users", id.String())
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	// or
	//if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
	//}

	return user.toDomain(), nil
}
