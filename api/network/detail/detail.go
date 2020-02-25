package detail

import (
	"io"
	"net/url"
	"strconv"

	"gitlab.com/MicahParks/wigole"
)

const (
	ApiPath = "network/detail"
	Method  = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

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

func (p *Parameters) Do(u *wigole.User) (*WiFiNetworkDetailResponse, error) {
	resp := &WiFiNetworkDetailResponse{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{}
}
