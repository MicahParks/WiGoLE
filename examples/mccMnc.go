package main

import (
	"gitlab.com/MicahParks/wigole/cell/mccMnc"
	"gitlab.com/MicahParks/wigole/user"
)

func main() {
	// Search for all Wifi networks that have the SSID of "Harambe", print the number of results.
	// It'll cap out at 100, but you get the idea.
	password := "password"
	username := "username"
	u := user.New(password, username)
	m := mccMnc.New()
	m.Mcc = 310
	resp, err := m.Do(u)
	if err != nil {
		panic(err)
	}
	resp
}
