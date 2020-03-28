[![GoDoc](https://godoc.org/gitlab.com/MicahParks/wigole?status.svg)](https://godoc.org/gitlab.com/MicahParks/wigole) [![Go Report Card](https://goreportcard.com/badge/gitlab.com/MicahParks/wigole)](https://goreportcard.com/report/gitlab.com/MicahParks/wigole) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
# WiGoLE

GoLang wrapper for v2 [WiGLE API](https://api.wigle.net/swagger).

## TO DO
* Get guidance on how to implement file*
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
 
