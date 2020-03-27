package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/group/admin"
)

type creds struct {
	ApiName  string
	ApiToken string
}

func main() {
	// Get and print admin info for your group.
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
	a := admin.New()
	a.Groupid = "I DON'T KNOW ANY GROUP ID"
	resp, err := a.Do(u)
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
	fmt.Printf("%+v\n", *resp)
}
