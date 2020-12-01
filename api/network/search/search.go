package search

import (
	"io"
	"net/url"
	"strconv"

	"github.com/MicahParks/wigole"
)

const (
	// ApiPath is the path relative to the BaseUrl to make the API call.
	ApiPath = "network/search"
	// Method is the HTTP method to use when doing an API call.
	Method = "GET"
)

// Body builds the request body reader for an API call to network/search.
func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

// Url builds the URL values for an API call to network/search.
func (p *Parameters) Url() (values url.Values, err error) {
	values, err = p.SsidUrl()
	if err != nil {
		return url.Values{}, err
	}
	if len(p.Encryption) != 0 {
		// It's possible for a user of the API to make their own Encryption type, but we'll allow it.
		values.Set("encryption", string(p.Encryption))
	}
	values.Set("freenet", strconv.FormatBool(p.Freenet))
	values.Set("paynet", strconv.FormatBool(p.Paynet))
	if len(p.Netid) != 0 {
		values.Set("netid", p.Netid)
	}
	return values, nil
}

// Do wraps the API call for network/search.
func (p *Parameters) Do(u *wigole.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

// New initializes and returns a pointer to parameters that can be used to make an API call to network/search.
func New() *Parameters {
	p := Parameters{}
	p.SearchSsid = *wigole.NewSsid()
	return &p
}
