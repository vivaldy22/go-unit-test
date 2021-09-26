package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vivaldy22/go-unit-test/6-mock/repository"
)

func TestUserService_GetByID(t *testing.T) {
	// panggil struct repo nya
	repoMock := &repository.UserRepositoryMock{}

	// hasil yang diinginkan
	user := &repository.UserEntity{
		ID:   10,
		Name: "Bambang",
		Age:  100,
	}

	// "program" mock supaya hasilnya keluar yang diinginkan
	repoMock.Mock.On("GetByID", mock.Anything).Return(user)

	expected := User{
		ID:   10,
		Name: "Bambang",
		Age:  100,
	}

	s := NewUserService(repoMock)
	actual, err := s.GetByID(1)
	assert.Nil(t, err)
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Age, actual.Age)

}
