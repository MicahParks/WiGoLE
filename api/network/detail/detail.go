package detail

import (
	"io"
	"net/url"
	"strconv"

	"github.com/MicahParks/wigole"
)

const (
	// ApiPath is the path relative to the BaseUrl to make the API call.
	ApiPath = "network/detail"
	// Method is the HTTP method to use when doing an API call.
	Method = "GET"
)

// Body builds the request body reader for an API call to network/detail.
func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

// Url builds the URL values for an API call to network/detail.
func (p *Parameters) Url() (values url.Values, err error) {
	// TODO Check to see if any values are legally zero.
	values = url.Values{}
	if len(p.NetId) != 0 {
		values.Set("netId", p.NetId)
	}
	if p.Operator != 0 {
		values.Set("operator", strconv.FormatUint(p.Operator, 10))
	}
	if p.Lac != 0 {
		values.Set("lac", strconv.FormatUint(p.Lac, 10))
	}
	if p.Cid != 0 {
		values.Set("cid", strconv.FormatUint(p.Cid, 10))
	}
	if len(p.Type) != 0 {
		// It's possible for a user of the API to make their own Network type, but we'll allow it.
		values.Set("type", string(p.Type))
	}
	if p.System != 0 {
		values.Set("system", strconv.FormatUint(p.System, 10))
	}
	if p.Network != 0 {
		values.Set("network", strconv.FormatUint(p.Network, 10))
	}
	if p.Basestation != 0 {
		values.Set("basestation", strconv.FormatUint(p.Basestation, 10))
	}
	return values, nil
}

// Do wraps the API call for network/detail.
func (p *Parameters) Do(u *wigole.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

// New initializes and returns a pointer to parameters that can be used to make an API call to network/detail.
func New() *Parameters {
	return &Parameters{}
}
