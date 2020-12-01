package standings

import (
	"github.com/MicahParks/wigole/api/stats"
)

const (
	// Discovered is a sorting constant defined by WiGLE.
	Discovered Sort = "discovered"
	// Total is a sorting constant defined by WiGLE.
	Total Sort = "total"
	// MonthCount is a sorting constant defined by WiGLE.
	MonthCount Sort = "monthcount"
	// PrevMonthCount is a sorting constant defined by WiGLE.
	PrevMonthCount Sort = "prevmonthcount"
	// GenDisc is a sorting constant defined by WiGLE.
	GenDisc Sort = "gendisc"
	// GenTotal is a sorting constant defined by WiGLE.
	GenTotal Sort = "gentotal"
	// FirstTransId is a sorting constant defined by WiGLE.
	FirstTransId Sort = "firsttransid"
	// LastTransId is a sorting constant defined by WiGLE.
	LastTransId Sort = "lasttransid"
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
