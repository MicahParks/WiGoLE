package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/profile/apiToken"
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
	u := wigole.NewUser(cred.Password, cred.Username)
	aT := apiToken.New()
	aT.Type = apiToken.API
	resp, err := aT.Do(u)
	if err != nil {
		if errors.Is(err, wigole.ErrFail) {
			println(err.Error())
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
