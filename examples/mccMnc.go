package main

import (
	"gitlab.com/MicahParks/wigole/cell/mccMnc"
	"gitlab.com/MicahParks/wigole/user"
)

func main() {
	// Search for metadata on the MCC 310. Print the first result's key.
	password := "password"
	username := "username"
	u := user.New(password, username)
	m := mccMnc.New()
	m.Mcc = 310
	resp, err := m.Do(u)
	if err != nil {
		panic(err)
	}
	println(resp[0].Key)
}
