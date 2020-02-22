package main

import (
	"gitlab.com/MicahParks/wigole/cell/mccMnc"
	"gitlab.com/MicahParks/wigole/user"
)

func main() {
	// Search for metadata on the MNC 110. Print the first result's key.
	password := "password"
	username := "username"
	u := user.New(password, username)
	m := mccMnc.New()
	m.Mnc = 110
	resp, err := m.Do(u)
	if err != nil {
		panic(err)
	}
	for k, v := range resp {
		if len(v) != 0 {
			println(k)
			return
		}
	}
}
