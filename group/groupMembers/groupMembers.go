package groupMembers

import (
	"io"
	"net/url"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const (
	ApiUrl = "group/admin?"
	Method = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (values url.Values, err error) {
	values = url.Values{}
	if len(p.Groupid) != 0 {
		values.Set("groupid", p.Groupid)
	}
	return values, nil
}

func (p *Parameters) Do(u *user.User) (*GroupMemberResponse, error) {
	resp := &GroupMemberResponse{}
	if err := wigole.Do(ApiUrl, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{}
}
