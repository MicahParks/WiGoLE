package search

import (
	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/network"
)

// Parameters holds all information that can be used for an API call to network/search.
type Parameters struct {
	Encryption network.Encryption // "None", "WEP", "WPA", "WPA2", "WPA3", "Unknown". Case insensitive.
	Freenet    bool               // Default to false.
	Paynet     bool               // Default to false.
	Netid      string             // BSSID exact match. Need first three octets. Format "0A:2C:EF:3D:25:1B".
	wigole.SearchSsid
}

// Response is the response from an API call for network/search.
type Response struct {
	wigole.NetSearchResponse
}
