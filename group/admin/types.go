package admin

import (
	"gitlab.com/MicahParks/wigole/group"
)

type Parameters struct {
	group.Parameters
}

type Response struct {
	Groupid string
	group.Response
}
