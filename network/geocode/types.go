package geocode

import (
	"gitlab.com/MicahParks/wigole/network"
)

type Parameters struct {
	Addresscode string
}

type Response struct {
	Success bool
	Results []*network.GeocodingResponse
}
