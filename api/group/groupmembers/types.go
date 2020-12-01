package groupmembers

import (
	"github.com/MicahParks/wigole/api/group"
)

// Parameters holds all information that can be used for an API call to group/groupmembers.
type Parameters struct {
	group.Parameters
}

// Response is the response from an API call for group/groupmembers.
type Response struct {
	group.GroupMemberResponse
}
