package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	sUser "gitlab.com/MicahParks/wigole/api/stats/user"
)

type creds struct {
	Password string
	Username string
}

func main() {
	// Request your own stats. Print your username and rank.
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
	sU := sUser.New()
	resp, err := sU.Do(u)
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
	println(resp.User, resp.Rank)
}
