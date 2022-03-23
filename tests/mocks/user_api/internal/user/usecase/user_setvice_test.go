package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/neo-classic/golang-playground/tests/mocks/user_api/internal/domain"
	"github.com/neo-classic/golang-playground/tests/mocks/user_api/internal/user/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite

	service *UserService
	repo    *mocks.UserRepository
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, &UserServiceTestSuite{})
}

func (s *UserServiceTestSuite) SetupTest() {
	s.repo = &mocks.UserRepository{}
	s.service = NewUserService(s.repo)
}

func (s *UserServiceTestSuite) TearDownTest() {
	s.repo.AssertExpectations(s.T())
}

func (s *UserServiceTestSuite) TestUserService_GetById() {
	user := domain.User{
		ID:    1,
		Email: "qweqwe@qweqwe.qwe",
	}
	ctx := context.Background()
	s.repo.On("GetByID", ctx, domain.UserID(1)).Return(&user, nil).Once()

	u, err := s.service.GetById(ctx, domain.UserID(1))

	s.Nil(err)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	s.NoError(err)

	s.T().Logf("%+v", u)
}

type TestUserGotHandler struct {
}

func (h TestUserGotHandler) Notify(e UserEvent) {
	fmt.Printf("[new event] GetByID(event name=%s, userID=%d)\n", e.Name(), e.UserID())
}
