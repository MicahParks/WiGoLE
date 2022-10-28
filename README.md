[![GoDoc](https://godoc.org/gitlab.com/MicahParks/wigole?status.svg)](https://godoc.org/gitlab.com/MicahParks/wigole) [![Go Report Card](https://goreportcard.com/badge/gitlab.com/MicahParks/wigole)](https://goreportcard.com/report/gitlab.com/MicahParks/wigole) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
# WiGoLE

This project is now archived. As of October 28th, 2022, it looks like a few of the examples still work with a valid
`creds.json` file, but mileage may vary. I'm not sure if v2 of the WiGLE API is still deployed or if v3 is just
partially backwards compatible. In any case, I would recommend using v3 of their API by generating a client from their
Swagger 2.0 (OpenAPI 2) specification. My favorite way to do that is with
[go-swagger](https://github.com/go-swagger/go-swagger). HOWEVER, I remember trying to do that for v2 of the
WiGLE API and it didn't work due to a server side implementation that wasn't 100% compatible with the specification
published, hence, this project.

GoLang wrapper for v2 WiGLE API. Reference the API directly [here](https://api.wigle.net/swagger).

# Learn
Check out the `examples` directory for examples of all implemented API calls. To test out the examples, create a
`creds.json` in your current directory with your credentials. See `examples/exampleCreds.json` and replace the
information appropriately. 

The example below is from `examples/networkSearch.go`. It shows the workflow that all implemented API calls follow.
```go
package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/network/search"
)

type creds struct {
	ApiName  string
	ApiToken string
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
	u := wigole.NewUser(cred.ApiName, cred.ApiToken)
	s := search.New()
	s.Ssid = "Harambe"
	resp, err := s.Do(u)
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
	println(resp.TotalResults)
}
```

Typical steps:
1. Read credentials into memory.
2. Create a `*wigole.User` with your credentials.
3. Create and initialize new `*api_name.Parameters` where `api_name` is the WiGoLE package for the API you'd like to
use.
4. Edit the `*api_name.Parameters`'s attributes according to your query.
5. Call `s.Do(u)` to preform the API call. Where `u` is your `*wigole.User` and `s` is your `*api_name.Parameters`.
6. Handle any errors.
7. Use the API call's response.

## Contributing
If you'd like to contribute to the project? Have a feature request? Want some more documentation? Feel free to create an
issue or submit a pull request!


## TO DO
* Get guidance on how to implement `api/file/*`
* Add an anonymous user so that API calls can be performed without login (most require login).
* Add examples and more documentation.

## Implementation coverage
- [x] `GET` bluetooth/detail
- [x] `GET` bluetooth/search
- [x] `GET` cell/mccMnc
- [x] `GET` cell/search
- [ ] `GET` file/kml/{transid}
- [ ] `GET` file/transactions
- [ ] `POST` file/upload
- [x] `POST` network/comment
- [x] `GET` network/detail
- [x] `GET` network/geocode
- [x] `GET` network/search
- [x] `GET` stats/countries
- [x] `GET` stats/general
- [x] `GET` stats/group
- [x] `GET` stats/regions
- [x] `GET` stats/site
- [x] `GET` stats/standings
- [x] `GET` stats/user
- [x] `GET` group/admin
- [x] `GET` group/groupMembers
- [x] `GET` profile/apiToken
- [x] `GET` profile/user
 
