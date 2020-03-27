package geocode

import (
	"io"
	"net/url"

	"gitlab.com/MicahParks/wigole"
)

const (
	ApiPath = "network/geocode"
	Method  = "GET"
)

// Body builds the request body reader for an API call to network/geocode.
func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

// Url builds the URL values for an API call to network/geocode.
func (p *Parameters) Url() (values url.Values, err error) {
	values = url.Values{}
	if len(p.Addresscode) != 0 {
		values.Set("addresscode", p.Addresscode)
	}
	return values, nil
}

// Do wraps the API call for network/geocode.
func (p *Parameters) Do(u *wigole.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

// New initializes and returns a pointer to parameters that can be used to make an API call to network/geocode.
func New() *Parameters {
	return &Parameters{}
}
