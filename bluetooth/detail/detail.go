package detail

import (
	"fmt"
	"io"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const (
	Method = "GET"
	ApiUrl = "bluetooth/detail?"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (url string, err error) {
	if len(p.Netid) != 0 {
		url += fmt.Sprintf("&netid=%s", p.Netid)
	}
	if len(p.ReverseAddress) != 0 {
		url += fmt.Sprintf("&reverseAddress=%s", p.ReverseAddress)
	}
	return url, nil
}

func (p *Parameters) Do(u *user.User) (*wigole.Response, error) {
	resp := &wigole.Response{}
	if err := wigole.Do(p, Method, resp, ApiUrl, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{}
}
