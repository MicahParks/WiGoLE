package detail

import (
	"github.com/MicahParks/wigole"
	"github.com/MicahParks/wigole/api/bluetooth"
)

// Parameters holds all information that can be used for an API call to bluetooth/detail.
type Parameters struct {
	Netid          string
	ReverseAddress string
}

// Response is the response from an API call for bluetooth/detail.
type Response struct {
	wigole.Detail
	Results []*bluetooth.Bluetooth
}
