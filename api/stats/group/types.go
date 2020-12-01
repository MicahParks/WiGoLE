package group

import (
	"github.com/MicahParks/wigole/api/group"
)

// Parameters holds all information that can be used for an API call to stats/group.
type Parameters struct {
}

// Response is the response from an API call for stats/group.
type Response struct {
	Success bool
	Groups  []*group.Group
}
