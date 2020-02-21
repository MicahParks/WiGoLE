package detail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole/user"
)

var Url = "network/detail?"

func (p *Parameters) BuildUrl() string {
	// TODO Check to see if any values are legally zero.
	url := ""
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
	return url
}

func (p *Parameters) Do(u *user.User) (*Response, error) {
	url := p.BuildUrl()
	resp, err := u.Do("GET", Url+url, nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := &Response{}
	if err = json.Unmarshal(body, r); err != nil {
		return nil, err
	}
	return r, nil
}

func New() *Parameters {
	return &Parameters{}
}
