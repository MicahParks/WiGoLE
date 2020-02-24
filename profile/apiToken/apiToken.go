package apiToken

import (
	"io"
	"net/url"

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

func (p *Parameters) Url() (val url.Values, err error) {
	val = url.Values{}
	if len(p.Type) != 0 {
		val.Set("type", string(p.Type))
	}
	return val, nil
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
