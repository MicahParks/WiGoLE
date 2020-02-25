package search

import (
	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/network"
)

type Parameters struct {
	Encryption network.Encryption // "None", "WEP", "WPA", "WPA2", "WPA3", "Unknown". Case insensitive.
	Freenet    bool               // Default to false.
	Paynet     bool               // Default to false.
	Netid      string             // BSSID exact match. Need first three octets. Format "0A:2C:EF:3D:25:1B".
	wigole.SearchSsid
}

type NetSearchResponse struct {
	wigole.NetSearchResponse
}
