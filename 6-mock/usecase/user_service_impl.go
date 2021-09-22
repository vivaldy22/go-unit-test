package usecase

import (
	"errors"

	"github.com/vivaldy22/go-unit-test/6-mock/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) GetByID(id int64) (user User, err error) {
	if id <= 0 {
		err = errors.New("id not valid")
		return
	}

	result := u.repo.GetByID(id)

	user = User{
		ID:   result.ID,
		Name: result.Name,
		Age:  result.Age,
	}

	return
}
