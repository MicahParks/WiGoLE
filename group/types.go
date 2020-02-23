package group

type Group struct {
	GroupId    string
	Username   string
	Status     string
	Discovered int
	Total      int
	GenDisc    int
	AuthType   string
	GroupOwner bool
}

type Parameters struct {
	Groupid string
}

type Response struct {
	Success bool
	Group   *Group
	Users   []*User
}

type User struct {
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
