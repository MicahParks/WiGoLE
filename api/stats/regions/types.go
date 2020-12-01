package regions

import (
	"github.com/MicahParks/wigole/api/stats"
)

// Parameters holds all information that can be used for an API call to stats/regions.
type Parameters struct {
	Country string // Defaults to "US".
}

// Response is the response from an API call for stats/regions.
type Response struct {
	Country string
	Regions []*stats.Region
}
