package user

import (
	"io"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const (
	ApiUrl = "profile/user?"
	Method = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (url string, err error) {
	return "", nil
}

func (p *Parameters) Do(u *user.User) (*Person, error) {
	resp := &Person{}
	if err := wigole.Do(p, Method, resp, ApiUrl, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{}
}
