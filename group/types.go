package group

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

type Parameters struct {
	Groupid string
}

type GroupMemberResponse struct {
	Success bool
	Group   *Group
	Users   []*GroupMember
}

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
