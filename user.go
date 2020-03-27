package wigole

import (
	"io"
	"net/http"
	"net/url"
)

var BaseUrl = "https://api.wigle.net/api/v2/"

type User struct {
	BaseUrl  string
	Client   http.Client
	ApiName  string
	ApiToken string
}

func (u *User) Do(apiPath string, body io.Reader, method string, values url.Values) (*http.Response, error) {
	println(BaseUrl + apiPath) // TODO Logging stuff.
	req, err := http.NewRequest(method, BaseUrl+apiPath, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(u.ApiName, u.ApiToken)
	req.URL.RawQuery = values.Encode()
	resp, err := u.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewUser(apiName string, apiToken string) *User {
	return &User{
		BaseUrl:  BaseUrl,
		Client:   http.Client{},
		ApiName:  apiName,
		ApiToken: apiToken,
	}
}
