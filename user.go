package wigole

import (
	"io"
	"net/http"
	"net/url"
)

// BaseUrl is the base URL for the WiGLE API. All requests will be built on this string.
var BaseUrl = "https://api.wigle.net/api/v2/"

// User stores API credentials. It is then used to preform requests.
type User struct {
	BaseUrl  string
	Client   http.Client
	ApiName  string
	ApiToken string
}

// Do will preform the API request with the given information. It preforms HTTP basic auth with the User's credentials.
func (u *User) Do(apiPath string, body io.Reader, method string, values url.Values) (*http.Response, error) {
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

// NewUser returns a pointer to a User struct.
func NewUser(apiName string, apiToken string) *User {
	return &User{
		BaseUrl:  BaseUrl,
		Client:   http.Client{},
		ApiName:  apiName,
		ApiToken: apiToken,
	}
}
