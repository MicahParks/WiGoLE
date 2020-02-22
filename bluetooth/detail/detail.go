package detail

import (
	"fmt"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const Url = "bluetooth/detail?"

func (p *Parameters) BuildUrl() (url string, err error) {
	if len(p.Netid) != 0 {
		url += fmt.Sprintf("&netid=%s", p.Netid)
	}
	if len(p.ReverseAddress) != 0 {
		url += fmt.Sprintf("&reverseAddress=%s", p.ReverseAddress)
	}
	return url, nil
}

func (p *Parameters) Do(u *user.User) (*wigole.Response, error) {

}
