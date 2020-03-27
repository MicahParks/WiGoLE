package apitoken

import (
	"io"
	"net/url"

	"gitlab.com/MicahParks/wigole"
)

const (
	// ApiPath is the path relative to the BaseUrl to make the API call.
	ApiPath = "profile/apiToken"
	// Method is the HTTP method to use when doing an API call.
	Method = "GET"
)

// Body builds the request body reader for an API call to profile/apitoken.
func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

// Url builds the URL values for an API call to profile/apitoken.
func (p *Parameters) Url() (val url.Values, err error) {
	val = url.Values{}
	if len(p.Type) != 0 {
		val.Set("type", string(p.Type))
	}
	return val, nil
}

// Do wraps the API call for profile/apitoken.
func (p *Parameters) Do(u *wigole.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

// New initializes and returns a pointer to parameters that can be used to make an API call to profile/apitoken.
func New() *Parameters {
	return &Parameters{}
}
