package service

import (
	"github.com/go-resty/resty/v2"
)

type service struct {
	resty resty.Client
}

func NewMyResty(resty resty.Client) *service {
	return &service{
		resty: resty,
	}
}

func (s *service) Get(path string) (code int, err error) {
	resp, err := s.resty.R().Get(path)
	if err != nil {
		return
	}

	code = resp.StatusCode()

	return
}
