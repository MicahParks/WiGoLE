package geocode

import (
	"github.com/MicahParks/wigole/api/network"
)

// Parameters holds all information that can be used for an API call to network/geocode.
type Parameters struct {
	Addresscode string
}

// Response is the response from an API call for network/geocode.
type Response struct {
	Success bool
	Results []*network.GeocodingResponse
}
