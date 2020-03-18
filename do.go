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
	failResp := &FailResp{}
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
