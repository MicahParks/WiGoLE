package main

import (
	"encoding/json"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/network/search"
)

type creds struct {
	Password string
	Username string
}

func main() {
	// Search for all Wifi networks that have the SSID of "Harambe", print the number of results.
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
	s := search.New()
	s.Ssid = "Harambe"
	resp, err := s.Do(u)
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
	println(resp.TotalResults)
}
