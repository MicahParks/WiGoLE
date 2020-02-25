package admin

import (
	"gitlab.com/MicahParks/wigole/api/group"
)

type Parameters struct {
	group.Parameters
}

type GroupResponse struct {
	Groupid string
	group.GroupMemberResponse
}
