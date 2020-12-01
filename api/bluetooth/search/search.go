package search

import (
	"io"
	"net/url"
	"strconv"

	"github.com/MicahParks/wigole"
)

const (
	// ApiPath is the path relative to the BaseUrl to make the API call.
	ApiPath = "bluetooth/search"
	// Method is the HTTP method to use when doing an API call.
	Method = "GET"
)

// Body builds the request body reader for an API call to bluetooth/search.
func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

// Url builds the URL values for an API call to bluetooth/search.
func (p *Parameters) Url() (values url.Values, err error) {
	values, err = p.SearchUrl()
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

// Do wraps the API call for bluetooth/search.
func (p *Parameters) Do(u *wigole.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

// New initializes and returns a pointer to parameters that can be used to make an API call to bluetooth/search.
func New() *Parameters {
	p := Parameters{
		ShowBt:  true,
		ShowBle: true,
	}
	p.SearchParameters = *wigole.NewSearch()
	return &p
}
