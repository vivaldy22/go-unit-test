package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vivaldy22/go-unit-test/6-mock/repository"
	. "github.com/vivaldy22/go-unit-test/6-mock/usecase"
)

func TestNewUserService(t *testing.T) {

}

func TestUserService_GetByID(t *testing.T) {
	repo := &repository.UserRepositoryMock{}
	repo.Mock.On("GetByID", int64(1)).Return(&repository.UserEntity{
		ID:   1,
		Name: "Test 1",
		Age:  21,
	}, nil)

	service := NewUserService(repo)
	user, err := service.GetByID(1)
	assert.Nil(t, err)
	assert.Equal(t, User{
		ID:   1,
		Name: "Test 1",
		Age:  21,
	}, user)
}
