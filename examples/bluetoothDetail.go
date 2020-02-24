package main

import (
	"encoding/json"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/bluetooth/detail"
	"gitlab.com/MicahParks/wigole/user"
)

type creds struct {
	Password string
	Username string
}

func main() {
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
	d := detail.New()
	d.Netid = "SOMETHING"
	resp, err := d.Do(u)
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
	println(len(resp.LocationData))
}
