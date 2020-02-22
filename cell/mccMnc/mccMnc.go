package mccMnc

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const ApiUrl = "cell/mccMnc?"
const Method = "GET"

func (a *AllStrings) Convert() (*Response, error) {
	resp := Response{}
	resp.Type = a.Type
	resp.CountryName = a.CountryName
	resp.CountryCode = a.CountryCode
	resp.Brand = a.Brand
	resp.Operator = a.Operator
	resp.Status = a.Status
	resp.Bands = a.Bands
	resp.Notes = a.Notes
	key, err := strconv.Atoi(a.Key)
	if err != nil {
		return nil, err
	}
	resp.Key = key
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

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (url string, err error) {
	if p.Mcc != 0 {
		url += fmt.Sprintf("&mcc=%d", p.Mcc)
	}
	if p.Mnc != 0 {
		url += fmt.Sprintf("&mnc=%d", p.Mnc)
	}
	return url, nil
}

func (p *Parameters) Do(u *user.User) ([]*Response, error) {
	m := make(map[string]map[string]map[string]string) // ;_;
	if err := wigole.Do(p, Method, &m, ApiUrl, u); err != nil {
		return nil, err
	}
	resp := make([]*Response, 0)
	if !(len(m) <= 1) {
		return nil, wigole.ErrUpdate
	}
	var master string
	for k := range m { // TODO This is only the case when MCC is included find a way to fix this for MNC.
		master = k
	}
	for k, v := range m[master] {
		a := AllStrings{Key: k}
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
		resp = append(resp, r)
	}
	return resp, nil
}

func (p *Parameters) DoRaw(u *user.User) (map[string]map[string]map[string]string, error) {
	m := make(map[string]map[string]map[string]string)
	if err := wigole.Do(p, Method, m, ApiUrl, u); err != nil {
		return nil, err
	}
	return m, nil
}

func New() *Parameters {
	return &Parameters{}
}
