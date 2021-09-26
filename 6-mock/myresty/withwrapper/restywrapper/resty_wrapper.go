package restywrapper

import "github.com/go-resty/resty/v2"

type RestyWrapper interface {
	Get(path string) (*resty.Response, error)
}
