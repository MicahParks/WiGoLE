package wigole

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
)

var (
	ErrAuth          = errors.New("basic auth failure")
	errAuthResp      = []byte(`{"success":false,"message":"too many queries today"}`)
	ErrTooMany       = errors.New("too many queries today")
	basicAuthFailure = []byte("Basic auth failure")
)

func Do(apiPath string, builder Builder, method string, response interface{}, user *User) error {
	values, err := builder.Url()
	if err != nil {
		return err
	}
	pBody, err := builder.Body()
	if err != nil {
		return err
	}
	resp, err := user.Do(apiPath, pBody, method, values)
	if err != nil {
		return err
	}
	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if bytes.Equal(rBody, errAuthResp) {
		return ErrTooMany
	}
	if err = json.Unmarshal(rBody, response); err != nil {
		if bytes.Equal(rBody, basicAuthFailure) {
			return ErrAuth
		}
		return err
	}
	return nil
}
