package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/MicahParks/wigole"
	"github.com/MicahParks/wigole/api/stats/countries"
)

type creds struct {
	ApiName  string
	ApiToken string
}

func main() {
	// Query for the countries. Print out the first result's country code and its count.
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
	country := countries.New()
	resp, err := country.Do(u)
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
	println(resp.Countries[0].Country, resp.Countries[0].Count)
}
