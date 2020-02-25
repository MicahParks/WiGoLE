package search

import (
	"io"
	"net/url"
	"strconv"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const (
	ApiPath = "bluetooth/search"
	Method  = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (values url.Values, err error) {
	values, err = p.ParentSearch()
	if err != nil {
		return nil, err
	}
	if len(p.Netid) != 0 {
		values.Set("netid", p.Netid)
	}
	if len(p.Name) != 0 {
		values.Set("name", p.Name)
	}
	if len(p.Namelike) != 0 {
		values.Set("namelike", p.Namelike)
	}
	values.Set("showBt", strconv.FormatBool(p.ShowBt))
	values.Set("showBle", strconv.FormatBool(p.ShowBle))
	return values, nil
}

func (p *Parameters) Do(u *user.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	p := Parameters{
		ShowBt:  true,
		ShowBle: true,
	}
	return &p
}
