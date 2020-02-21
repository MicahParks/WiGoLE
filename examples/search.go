package main

import (
	"gitlab.com/MicahParks/wigole/network/search"
	"gitlab.com/MicahParks/wigole/user"
)

func main() {
	// Search for all Wifi networks that have the SSID of "Harambe", print the number of results.
	password := "password"
	username := "username"
	u := user.New(password, username)
	s := search.New()
	s.Ssid = "Harambe"
	resp, err := s.Do(u)
	if err != nil {
		panic(err)
	}
	println(resp.TotalResults)
}
