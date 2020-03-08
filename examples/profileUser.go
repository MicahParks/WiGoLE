package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	pUser "gitlab.com/MicahParks/wigole/api/profile/user"
)

type creds struct {
	Password string
	Username string
}

func main() {
	// Get and print your user info.
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
	pU := pUser.New()
	resp, err := pU.Do(u)
	if err != nil {
		if errors.Is(err, wigole.ErrFail) {
			println(err.Error())
		}
		if err == wigole.ErrAuth {
			println("Failed to authenticate with creds.json.")
			return
		}
		panic(err)
	}
	fmt.Printf("%+v\n", *resp)
}
