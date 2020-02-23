package search

import (
	"fmt"
	"io"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const (
	ApiUrl = "network/search?"
	Method = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (url string, err error) {
	url, err = p.ParentUrl()
	if err != nil {
		return "", err
	}
	if len(p.Encryption) != 0 {
		// It's possible for a user of the API to make their own Encryption type, but we'll allow it.
		url += fmt.Sprintf("&encryption=%s", p.Encryption)
	}
	url += fmt.Sprintf("&freenet=%v", p.Freenet)
	url += fmt.Sprintf("&paynet=%v", p.Paynet)
	if len(p.Netid) != 0 {
		url += fmt.Sprintf("&netid=%s", p.Netid)
	}
	return url, nil
}

func (p *Parameters) Do(u *user.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(p, Method, resp, ApiUrl, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	params := Parameters{}
	params.MinQoS = 8 // Max of 7. This let's you know it's uninitialized.
	return &params
}
