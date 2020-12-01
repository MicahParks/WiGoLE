package mccmnc

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"

	"github.com/MicahParks/wigole"
)

const (
	// ApiPath is the path relative to the BaseUrl to make the API call.
	ApiPath = "cell/mccMnc"
	// Method is the HTTP method to use when doing an API call.
	Method = "GET"
)

// Convert changes numeric values of strings to integers as the expected data type.
func (a *allStrings) Convert() (*MccMnc, error) {
	resp := MccMnc{}
	resp.Mm = a.Mm
	mcc, err := strconv.Atoi(a.Mcc)
	if err != nil {
		return nil, err
	}
	resp.Mcc = mcc
	mnc, err := strconv.Atoi(a.Mnc)
	if err != nil {
		return nil, err
	}
	resp.Mnc = mnc
	return &resp, nil
}

// Body builds the request body reader for an API call to cell/mmcmnc.
func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

// Url builds the URL values for an API call to cell/mmcmnc.
func (p *Parameters) Url() (values url.Values, err error) {
	values = url.Values{}
	if p.Mcc != 0 {
		values.Set("mcc", strconv.Itoa(p.Mcc))
	}
	if p.Mnc != 0 {
		values.Set("mnc", strconv.Itoa(p.Mnc))
	}
	return values, nil
}

// Do wraps the API call for cell/mmcmnc.
func (p *Parameters) Do(u *wigole.User) (map[int][]*MccMnc, error) {
	m := make(map[string]map[string]map[string]string) // ;_;
	if err := wigole.Do(ApiPath, p, Method, &m, u); err != nil {
		return nil, err
	}
	resp := make(map[int][]*MccMnc, 0)
	var master string
	for k := range m {
		master = k
		mccMncs := make([]*MccMnc, 0)
		for _, v := range m[master] {
			a := allStrings{}
			dat, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			if err = json.Unmarshal(dat, &a); err != nil {
				return nil, err
			}
			r, err := a.Convert()
			if err != nil {
				return nil, err
			}
			mccMncs = append(mccMncs, r)
		}
		mcc, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		resp[mcc] = mccMncs
	}
	return resp, nil
}

// DoRaw wraps the API call for cell/mmcmnc but returns the raw response's map of strings.
func (p *Parameters) DoRaw(u *wigole.User) (map[string]map[string]map[string]string, error) {
	m := make(map[string]map[string]map[string]string)
	if err := wigole.Do(ApiPath, p, Method, m, u); err != nil {
		return nil, err
	}
	return m, nil
}

// New initializes and returns a pointer to parameters that can be used to make an API call to cell/mmcmnc.
func New() *Parameters {
	return &Parameters{}
}
