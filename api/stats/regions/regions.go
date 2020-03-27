package regions

import (
	"io"
	"net/url"

	"gitlab.com/MicahParks/wigole"
)

const (
	ApiPath = "stats/regions"
	Method  = "GET"
)

// Body builds the request body reader for an API call to stats/regions.
func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

// Url builds the URL values for an API call to stats/regions.
func (p *Parameters) Url() (values url.Values, err error) {
	values = url.Values{}
	values.Set("country", p.Country)
	return values, nil
}

// Do wraps the API call for stats/regions.
func (p *Parameters) Do(u *wigole.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

// New initializes and returns a pointer to parameters that can be used to make an API call to stats/regions.
func New() *Parameters {
	return &Parameters{
		Country: "US",
	}
}
