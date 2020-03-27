// Package wigole is the wrapper for the WiGLE API.
package wigole

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

var (
	ErrAuth          = errors.New("basic auth failure")
	ErrFail          = errors.New("failed API call")
	ErrTooMany       = fmt.Errorf("%w\nmessage: too many queries today", ErrFail)
	basicAuthFailure = []byte("Basic auth failure")
)

// Do wraps building the URL, building the request body, doing the request, and parsing and unmarshaling the response.
// Returns any errors.
func Do(apiPath string, b builder, method string, response interface{}, user *User) error {
	values, err := b.Url()
	if err != nil {
		return err
	}
	pBody, err := b.Body()
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
	failResp := &failResp{}
	if err = json.Unmarshal(rBody, failResp); err == nil {
		if !failResp.Success && len(failResp.Message) != 0 {
			if failResp.Message == ErrTooMany.Error() {
				return ErrTooMany
			}
			return fmt.Errorf("%w\nmessage: %s", ErrFail, failResp.Message)
		}
	}
	if err = json.Unmarshal(rBody, response); err != nil {
		if bytes.Equal(rBody, basicAuthFailure) {
			return ErrAuth
		}
		return err
	}
	return nil
}
