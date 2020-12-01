package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/MicahParks/wigole"
	sUser "github.com/MicahParks/wigole/api/stats/user"
)

type creds struct {
	ApiName  string
	ApiToken string
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
	u := wigole.NewUser(cred.ApiName, cred.ApiToken)
	sU := sUser.New()
	resp, err := sU.Do(u)
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
	println(resp.User, resp.Rank)
}
