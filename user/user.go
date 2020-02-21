package user

import (
	"io"
	"net/http"
)

var BaseUrl = "https://api.wigle.net/api/v2/"

type User struct {
	BaseUrl  string
	Client   http.Client
	Password string
	Username string
}

func (u *User) Do(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, BaseUrl+url, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(u.Username, u.Password)
	resp, err := u.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func New(password string, username string) *User {
	return &User{
		BaseUrl:  BaseUrl,
		Client:   http.Client{},
		Password: password,
		Username: username,
	}
}
