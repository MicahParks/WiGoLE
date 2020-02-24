package detail

import (
	"fmt"
	"io"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const (
	ApiUrl = "network/detail?"
	Method = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (url string, err error) {
	// TODO Check to see if any values are legally zero.
	if len(p.NetId) != 0 {
		url += fmt.Sprintf("&netId=%s", p.NetId)
	}
	if p.Operator != 0 {
		url += fmt.Sprintf("&operator=%d", p.Operator)
	}
	if p.Lac != 0 {
		url += fmt.Sprintf("&lac=%d", p.Lac)
	}
	if p.Cid != 0 {
		url += fmt.Sprintf("&cid=%d", p.Cid)
	}
	if len(p.Type) != 0 {
		// It's possible for a user of the API to make their own Network type, but we'll allow it.
		url += fmt.Sprintf("&type=%s", p.Type)
	}
	if p.System != 0 {
		url += fmt.Sprintf("&system=%d", p.System)
	}
	if p.Network != 0 {
		url += fmt.Sprintf("&network=%d", p.Network)
	}
	if p.Basestation != 0 {
		url += fmt.Sprintf("&basestation=%d", p.Basestation)
	}
	return url, nil
}

func (p *Parameters) Do(u *user.User) (*WiFiNetworkDetailResponse, error) {
	resp := &WiFiNetworkDetailResponse{}
	if err := wigole.Do(p, Method, resp, ApiUrl, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{}
}
