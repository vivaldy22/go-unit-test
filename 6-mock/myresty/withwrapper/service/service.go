package service

import (
	"github.com/vivaldy22/go-unit-test/6-mock/myresty/withwrapper/restywrapper"
)

type service struct {
	myResty restywrapper.RestyWrapper
}

func NewMyResty(myResty restywrapper.RestyWrapper) *service {
	return &service{
		myResty: myResty,
	}
}

func (s *service) Get(path string) (code int, err error) {
	resp, err := s.myResty.Get(path)
	if err != nil {
		return
	}

	code = resp.StatusCode()

	return
}
