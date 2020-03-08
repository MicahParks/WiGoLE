package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/cell/mccMnc"
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
	u := wigole.NewUser(cred.Password, cred.Username)
	m := mccMnc.New()
	m.Mnc = 110
	resp, err := m.Do(u)
	if err != nil {
		if errors.Is(err, wigole.ErrFail) {
			println(err.Error())
		}
		if err == wigole.ErrAuth {
			println("Failed to authenticate with creds.json.")
			return
		}
		panic(err)
	}
	for k, v := range resp {
		if len(v) != 0 {
			println(k)
			return
		}
	}
}
