package standings

import (
	"gitlab.com/MicahParks/wigole/api/stats"
)

const (
	Discovered     Sort = "discovered"
	Total          Sort = "total"
	Monthcount     Sort = "monthcount"
	Prevmonthcount Sort = "prevmonthcount"
	Gendisc        Sort = "gendisc"
	Gentotal       Sort = "gentotal"
	Firsttransid   Sort = "firsttransid"
	Lasttransid    Sort = "lasttransid"
)

// Parameters holds all information that can be used for an API call to stats/standings.
type Parameters struct {
	Sort      Sort
	Pagestart int
	Pageend   int
}

// Response is the response from an API call for stats/standings.
type Response struct {
	Success    bool
	EventView  bool
	MyUsername string
	PageStart  int
	PageEnd    int
	TotalUsers int
	Results    []*stats.UserStandings
	SortBy     string
}

// Sort includes all the WiGLE defined sorting constants for making API calls to stats/standings.
type Sort string
