package mccMnc

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/user"
)

const ApiUrl = "cell/mccMnc"
const Method = "GET"

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

func (p *Parameters) Do(u *user.User) ([]Response, error) {
	m := make(map[string]string)
	if err := wigole.Do(p, Method, m, ApiUrl, u); err != nil {
		return nil, err
	}
	resp := []Response{}
	for k, v := range m {
		key, err := strconv.ParseInt(k, 10, 0)
		if err != nil {
			return nil, err
		}
		r := Response{Key: key}
		dat, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(dat, &r); err != nil {
			return nil, err
		}
		resp = append(resp, r)
	}
	return resp, nil
}

func (p *Parameters) DoRaw(u *user.User) (map[string]string, error) {
	m := make(map[string]string)
	if err := wigole.Do(p, Method, m, ApiUrl, u); err != nil {
		return nil, err
	}
	return m, nil
}

func New() *Parameters {
	return &Parameters{}
}
