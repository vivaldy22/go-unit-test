package restywrapper

import "github.com/go-resty/resty/v2"

type myRestyWrapper struct {
	resty resty.Client
}

func NewMyRestyWrapper(resty resty.Client) RestyWrapper {
	return &myRestyWrapper{
		resty: resty,
	}
}

func (m *myRestyWrapper) Get(path string) (*resty.Response, error) {
	return m.resty.R().Get(path)
}
