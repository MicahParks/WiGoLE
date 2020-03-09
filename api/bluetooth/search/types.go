package search

import (
	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/bluetooth"
)

type Parameters struct {
	Netid    string
	Name     string
	Namelike string
	ShowBt   bool // Default to true.
	ShowBle  bool // Default to true.
	wigole.SearchParameters
}

type Response struct {
	Success      bool
	TotalResults int
	First        int
	Last         int
	ResultCount  int
	Results      []*bluetooth.Bluetooth
	SearchAfter  string
	Search_after int
}
