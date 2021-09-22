// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	repository "github.com/vivaldy22/go-unit-test/6-mock/repository"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: id
func (_m *UserRepository) GetByID(id int64) *repository.UserEntity {
	ret := _m.Called(id)

	var r0 *repository.UserEntity
	if rf, ok := ret.Get(0).(func(int64) *repository.UserEntity); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.UserEntity)
		}
	}

	return r0
}
