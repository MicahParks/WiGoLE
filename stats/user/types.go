package user

import (
	"gitlab.com/MicahParks/wigole/stats"
)

type Parameters struct {
}

type UserStatsResponse struct {
	Success       bool
	ImageBadgeUrl string
	Statistics    *stats.UserStandings
	Rank          int
	MonthRank     int
	User          string
}
