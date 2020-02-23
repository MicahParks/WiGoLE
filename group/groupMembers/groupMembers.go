package groupMembers

import (
	"fmt"
	"io"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const (
	ApiUrl = "group/admin?"
	Method = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (url string, err error) {
	if len(p.Groupid) != 0 {
		url += fmt.Sprintf("&groupid=%s", p.Groupid)
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
	return &Parameters{}
}
