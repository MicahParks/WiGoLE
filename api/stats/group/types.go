package group

import (
	"gitlab.com/MicahParks/wigole/api/group"
)

type Parameters struct {
}

type Response struct {
	Success bool
	Groups  []*group.Group
}
