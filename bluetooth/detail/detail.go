package detail

import (
	"io"
	"net/url"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const (
	Method  = "GET"
	ApiPath = "bluetooth/detail"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

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

func (p *Parameters) Do(u *user.User) (*WiFiNetworkWithLocation, error) {
	resp := &WiFiNetworkWithLocation{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{}
}
