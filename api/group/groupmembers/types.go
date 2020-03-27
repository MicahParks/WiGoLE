package groupmembers

import (
	"gitlab.com/MicahParks/wigole/api/group"
)

type Parameters struct {
	group.Parameters
}

type GroupMemberResponse struct {
	group.GroupMemberResponse
}
