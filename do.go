package wigole

import (
	"encoding/json"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole/user"
)

func Do(builder Builder, method string, response interface{}, apiUrl string, user *user.User) error {
	url, err := builder.BuildUrl()
	if err != nil {
		return err
	}
	resp, err := user.Do(method, apiUrl+url, nil)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, response); err != nil {
		return err
	}
	return nil
}
