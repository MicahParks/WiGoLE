package search

import (
	"io"
	"net/url"
	"strconv"

	"gitlab.com/MicahParks/wigole"
)

const (
	// ApiPath is the path relative to the BaseUrl to make the API call.
	ApiPath = "cell/search"
	// Method is the HTTP method to use when doing an API call.
	Method = "GET"
)

// Body builds the request body reader for an API call to cell/search.
func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

// Url builds the URL values for an API call to cell/search.
func (p *Parameters) Url() (values url.Values, err error) {
	values, err = p.SsidUrl()
	if err != nil {
		return url.Values{}, err
	}
	if len(p.Cell_op) != 0 {
		values.Set("cell_op", string(p.Cell_op))
	}
	if len(p.Cell_net) != 0 {
		values.Set("cell_net", string(p.Cell_net))
	}
	if len(p.Cell_id) != 0 {
		values.Set("cell_id", string(p.Cell_id))
	}
	values.Set("showGsm", strconv.FormatBool(p.ShowGsm))
	values.Set("showCdma", strconv.FormatBool(p.ShowCdma))
	values.Set("showLte", strconv.FormatBool(p.ShowLte))
	values.Set("showWcdma", strconv.FormatBool(p.ShowWcdma))
	return values, nil
}

// Do wraps the API call for cell/search.
func (p *Parameters) Do(u *wigole.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

// New initializes and returns a pointer to parameters that can be used to make an API call to cell/search.
func New() *Parameters {
	p := Parameters{
		ShowGsm:   true,
		ShowCdma:  true,
		ShowLte:   true,
		ShowWcdma: true,
	}
	p.SearchSsid = *wigole.NewSsid()
	return &p
}
