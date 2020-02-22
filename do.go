package wigole

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole/user"
)

var (
	ErrAuth        = errors.New("basic auth failure")
	errAuthResp    = []byte(`{"success":false,"message":"too many queries today"}`)
	ErrTooMany     = errors.New("too many queries today")
	errTooManyResp = []byte("Basic auth failure")
)

func Do(builder Builder, method string, response interface{}, apiUrl string, user *user.User) error {
	url, err := builder.Url()
	if err != nil {
		return err
	}
	pBody, err := builder.Body()
	if err != nil {
		return err
	}
	resp, err := user.Do(method, apiUrl+url, pBody)
	if err != nil {
		return err
	}
	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if bytes.Equal(rBody, errAuthResp) {
		return errTooMany
	}
	if err = json.Unmarshal(rBody, response); err != nil {
		// TODO Check for things like "Basic auth failure".
		if bytes.Equal(rBody, errTooManyResp) {
			return errAuth
		}
		return err
	}
	return nil
}
