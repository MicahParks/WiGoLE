package main

import (
	"gitlab.com/MicahParks/wigole/network/detail"
	"gitlab.com/MicahParks/wigole/user"
)

func main() {
	// Search for the details of all Wifi networks with the BSSID. Print the number of networks.
	password := "password"
	username := "username"
	u := user.New(password, username)
	d := detail.New()
	d.NetId = "not working yet"
	resp, err := d.Do(u)
	if err != nil {
		panic(err)
	}
	println(len(resp.Results))
}
