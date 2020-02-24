package apiToken

import (
	"fmt"
	"io"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const (
	ApiUrl = "profile/apiToken?"
	Method = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (url string, err error) {
	if len(p.Type) != 0 {
		url = fmt.Sprintf("&type=%s", p.Type)
	}
	return url, nil
}

func (p *Parameters) Do(u *user.User) (*AuthTokenResponse, error) {
	resp := &AuthTokenResponse{}
	if err := wigole.Do(p, Method, resp, ApiUrl, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{}
}
