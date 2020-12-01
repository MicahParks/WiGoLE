package search

import (
	"github.com/MicahParks/wigole"
	"github.com/MicahParks/wigole/api/bluetooth"
)

// Parameters holds all information that can be used for an API call to bluetooth/search.
type Parameters struct {
	Netid    string
	Name     string
	Namelike string
	ShowBt   bool // Default to true.
	ShowBle  bool // Default to true.
	wigole.SearchParameters
}

// Response is the response from an API call for bluetooth/search.
type Response struct {
	Success               bool
	TotalResults          int
	First                 int
	Last                  int
	ResultCount           int
	Results               []*bluetooth.Bluetooth
	SearchAfter           string
	SearchAfterDeprecated int `json:"search_after"`
}
