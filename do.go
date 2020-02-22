package wigole

import (
	"encoding/json"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole/user"
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
	if err = json.Unmarshal(rBody, response); err != nil {
		// TODO Check for things like "Basic auth failure".
		return err
	}
	return nil
}
