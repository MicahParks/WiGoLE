package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/bluetooth/detail"
)

type creds struct {
	ApiName  string
	ApiToken string
}

func main() {
	// Search for a specific Netid that will result in the Ssid of "SONY".
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
	d := detail.New()
	d.Netid = "00:03:19:8c:b3:bf"
	resp, err := d.Do(u)
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
	println(resp.Results[0].Ssid)
}
