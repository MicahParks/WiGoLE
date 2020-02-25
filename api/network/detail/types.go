package detail

import (
	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/network"
)

type Parameters struct {
	NetId       string
	Operator    uint64
	Lac         uint64
	Cid         uint64
	Type        network.Network
	System      uint64
	Network     uint64
	Basestation uint64
}

type WiFiNetworkDetailResponse struct {
	Success   bool
	Cdma      bool
	Gsm       bool
	Lte       bool
	Wcdma     bool
	Wifi      bool
	Addresses []*network.GeocodingResponse
	Results   []*wigole.WiFiNetworkWithLocation
}
