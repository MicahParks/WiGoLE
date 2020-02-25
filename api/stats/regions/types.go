package regions

import (
	"gitlab.com/MicahParks/wigole/api/stats"
)

type Parameters struct {
	Country string // Defaults to "US".
}

type Response struct {
	Country string
	Regions []*stats.Region
}
