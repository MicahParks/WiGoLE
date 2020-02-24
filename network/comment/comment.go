package comment

import (
	"io"
	"net/url"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const (
	ApiPath = "network/comment"
	Method  = "POST"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

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

func (p *Parameters) Do(u *user.User) (*NetCommentResponse, error) {
	resp := &NetCommentResponse{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{}
}
