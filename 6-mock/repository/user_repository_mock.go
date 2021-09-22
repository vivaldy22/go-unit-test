package repository

import "github.com/stretchr/testify/mock"

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (u *UserRepositoryMock) GetByID(id int64) (user *UserEntity) {
	args := u.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}

	user = args.Get(0).(*UserEntity)
	return
}
