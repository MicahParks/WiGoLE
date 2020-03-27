package detail

import (
	"io"
	"net/url"

	"gitlab.com/MicahParks/wigole"
)

const (
	// ApiPath is the path relative to the BaseUrl to make the API call.
	ApiPath = "bluetooth/detail"
	// Method is the HTTP method to use when doing an API call.
	Method = "GET"
)

// Body builds the request body reader for an API call to bluetooth/detail.
func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

// Url builds the URL values for an API call to bluetooth/detail.
func (p *Parameters) Url() (values url.Values, err error) {
	values = url.Values{}
	if len(p.Netid) != 0 {
		values.Set("netid", p.Netid)
	}
	if len(p.ReverseAddress) != 0 {
		values.Set("reverseAddress", p.ReverseAddress)
	}
	return values, nil
}

// Do wraps the API call for bluetooth/detail.
func (p *Parameters) Do(u *wigole.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

// New initializes and returns a pointer to parameters that can be used to make an API call to bluetooth/detail.
func New() *Parameters {
	return &Parameters{}
}
