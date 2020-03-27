package apitoken

import (
	"io"
	"net/url"

	"gitlab.com/MicahParks/wigole"
)

const (
	ApiPath = "profile/apiToken"
	Method  = "GET"
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

func (p *Parameters) Do(u *wigole.User) (*AuthTokenResponse, error) {
	resp := &AuthTokenResponse{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{}
}
