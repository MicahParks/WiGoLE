package search

import (
	"fmt"
	"io"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const (
	ApiUrl = "cell/search?"
	Method = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (url string, err error) {
	url, err = p.ParentUrl()
	if err != nil {
		return "", err
	}
	if len(p.Cell_op) != 0 {
		url += fmt.Sprintf("&cell_op=%s", p.Cell_op)
	}
	if len(p.Cell_net) != 0 {
		url += fmt.Sprintf("&cell_net=%s", p.Cell_net)
	}
	if len(p.Cell_id) != 0 {
		url += fmt.Sprintf("&cell_id=%s", p.Cell_id)
	}
	url += fmt.Sprintf("&showGsm=%v&showCdma=%v&showLte=%v&showWcdma=%v", p.ShowGsm, p.ShowCdma, p.ShowLte, p.ShowWcdma)
	return url, nil
}

func (p *Parameters) Do(u *user.User) (*wigole.Response, error) {
	resp := &wigole.Response{}
	if err := wigole.Do(p, Method, resp, ApiUrl, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	params := Parameters{}
	params.ShowGsm = true
	params.ShowCdma = true
	params.ShowLte = true
	params.ShowWcdma = true
	return &params
}
