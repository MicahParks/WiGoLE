package main

import (
	"encoding/json"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole/network/search"
	"gitlab.com/MicahParks/wigole/user"
)

type creds struct {
	Password string
	Username string
}

func main() {
	// Search for all Wifi networks that have the SSID of "Harambe", print the number of results.
	// It'll cap out at 100, but you get the idea.
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
	s := search.New()
	s.Ssid = "Harambe"
	resp, err := s.Do(u)
	if err != nil {
		panic(err)
	}
	println(resp.TotalResults)
}
