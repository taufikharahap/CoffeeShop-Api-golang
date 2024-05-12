package repository

import (
	"coffeeshop-api-golang/config"
	"coffeeshop-api-golang/internal/models"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) GetByEmail(email string) (*config.Result, error) {
	args := r.Mock.Called(email)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) GetAllUser() (*config.Result, error) {
	args := r.Mock.Called()
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) GetAuthData(email string) (*models.User, error) {
	args := r.Mock.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (r *RepoMock) CreateUser(data *models.User) (*config.Result, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) Update(data *models.User, user_id string) (*config.Result, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) Delete(data *models.User) (*config.Result, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}
