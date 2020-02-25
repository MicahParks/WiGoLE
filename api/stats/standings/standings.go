package standings

import (
	"io"
	"net/url"
	"strconv"

	"gitlab.com/MicahParks/wigole"
)

const (
	ApiPath = "stats/standings"
	Method  = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (values url.Values, err error) {
	values = url.Values{}
	if len(p.Sort) != 0 {
		values.Set("sort", string(p.Sort))
	}
	values.Set("pagestart", strconv.Itoa(p.Pagestart))
	if p.Pageend >= p.Pagestart && p.Pageend != 0 {
		values.Set("pageend", strconv.Itoa(p.Pageend))
	}
	return values, nil
}

func (p *Parameters) Do(u *wigole.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{
		Sort:      Discovered,
		Pagestart: 0,
	}
}
