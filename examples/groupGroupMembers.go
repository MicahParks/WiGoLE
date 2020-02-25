package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/group/groupMembers"
)

type creds struct {
	Password string
	Username string
}

func main() {
	// Get and print group members.
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
	u := wigole.New(cred.Password, cred.Username)
	gM := groupMembers.New()
	gM.Groupid = "I DON'T KNOW ANY GROUP ID"
	resp, err := gM.Do(u)
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
	fmt.Printf("%+v\n", *resp)
}
