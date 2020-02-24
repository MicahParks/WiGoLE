package detail

import (
	"gitlab.com/MicahParks/wigole"
)

type Parameters struct {
	Netid          string
	ReverseAddress string
}

type Response struct {
	wigole.WiFiNetworkWithLocation
}
