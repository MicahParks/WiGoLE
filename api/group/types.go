package group

// Group is used to deserialize information returned from the WiGLE API. Used for inheritance.
type Group struct {
	GroupId    string
	GroupName  string
	Owner      string
	Discovered int
	Total      int
	GenDisc    int
	AuthType   string
	GroupOwner bool
}

// Parameters is a generalized parameter holder for API calls in group/*.
type Parameters struct {
	Groupid string
}

// GroupMemberResponse is used to deserialize information returned from the WiGLE API. Used for inheritance.
type GroupMemberResponse struct {
	Success bool
	Group   *Group
	Users   []*GroupMember
}

// GroupMember is used to deserialize information returned from the WiGLE API. Used for inheritance.
type GroupMember struct {
	GroupId        string
	Username       string
	Status         string
	Discovered     int
	Total          int
	GenDisc        int
	FirstTransid   string
	LastTransid    string
	MonthCount     int
	PrevMonthCount int
}
