package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/profile/apiToken"
	"gitlab.com/MicahParks/wigole/user"
)

type creds struct {
	Password string
	Username string
}

func main() {
	// Get and print your API info.
	cred := creds{}
	c, err := ioutil.ReadFile("creds.json")
	if err != nil {
		println("Create a properly formatted 'creds.json' file in the working directory.")
		return
	}
	if err = json.Unmarshal(c, &cred); err != nil {
		println("JSON failure for 'creds.json' file in the working directory.")
		return
	}
	u := user.New(cred.Password, cred.Username)
	aT := apiToken.New()
	aT.Type = apiToken.API
	resp, err := aT.Do(u)
	if err != nil {
		if err == wigole.ErrTooMany {
			println("Too many queries of that type for today.")
			return
		}
		if err == wigole.ErrAuth {
			println("Failed to authenticate with creds.json.")
			return
		}
		panic(err)
	}
	for _, v := range resp.Result {
		fmt.Printf("%+v\n", *v)
	}
}
