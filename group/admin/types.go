package admin

import (
	"gitlab.com/MicahParks/wigole/group"
)

type Parameters struct {
	group.Parameters
}

type GroupResponse struct {
	Groupid string
	group.GroupMemberResponse
}
