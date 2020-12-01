package admin

import (
	"github.com/MicahParks/wigole/api/group"
)

// Parameters holds all information that can be used for an API call to group/admin.
type Parameters struct {
	group.Parameters
}

// Response is the response from an API call for group/admin.
type Response struct {
	Groupid string
	group.GroupMemberResponse
}
