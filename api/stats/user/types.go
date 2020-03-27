package user

import (
	"gitlab.com/MicahParks/wigole/api/stats"
)

// Parameters holds all information that can be used for an API call to stats/user.
type Parameters struct {
}

// Response is the response from an API call for stats/user.
type Response struct {
	Success       bool
	ImageBadgeUrl string
	Statistics    *stats.UserStandings
	Rank          int
	MonthRank     int
	User          string
}
