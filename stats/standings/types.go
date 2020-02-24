package group

import (
	"gitlab.com/MicahParks/wigole/stats"
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

type Parameters struct {
	Sort      Sort
	Pagestart int
	Pageend   int
}

type Response struct {
	Success    bool
	EventView  string
	MyUsername string
	PageStart  int
	PageEnd    int
	TotalUsers int
	Results    []*stats.UserStandings
	SortBy     string
}

type Sort string
