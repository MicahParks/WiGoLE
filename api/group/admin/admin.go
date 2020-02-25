package admin

import (
	"io"
	"net/url"

	"gitlab.com/MicahParks/wigole"
)

const (
	ApiPath = "group/admin"
	Method  = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (values url.Values, err error) {
	values = url.Values{}
	if len(p.Groupid) != 0 {
		values.Set("groupid", p.Groupid)
	}
	return values, nil
}

func (p *Parameters) Do(u *wigole.User) (*GroupResponse, error) {
	resp := &GroupResponse{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{}
}
