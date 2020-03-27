package comment

import (
	"io"
	"net/url"

	"gitlab.com/MicahParks/wigole"
)

const (
	// ApiPath is the path relative to the BaseUrl to make the API call.
	ApiPath = "network/comment"
	// Method is the HTTP method to use when doing an API call.
	Method = "POST"
)

// Body builds the request body reader for an API call to network/comment.
func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

// Url builds the URL values for an API call to network/comment.
func (p *Parameters) Url() (values url.Values, err error) {
	values = url.Values{}
	if len(p.Netid) != 0 {
		values.Set("netid", p.Netid)
	}
	if len(p.Comment) != 0 {
		values.Set("comment", p.Comment)
	}
	return values, nil
}

// Do wraps the API call for network/comment.
func (p *Parameters) Do(u *wigole.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

// New initializes and returns a pointer to parameters that can be used to make an API call to network/comment.
func New() *Parameters {
	return &Parameters{}
}
