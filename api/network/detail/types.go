package detail

import (
	"github.com/MicahParks/wigole"
	"github.com/MicahParks/wigole/api/network"
)

// Parameters holds all information that can be used for an API call to network/detail.
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

// Response is the response from an API call for network/detail.
type Response struct {
	Results []*wigole.WiFiNetworkWithLocation
	wigole.Detail
}
