package search

import (
	"gitlab.com/MicahParks/wigole"
)

// Cell includes all the WiGLE defined Cell constants for making API calls to cell/search.
type Cell string

const (
	GSM   Cell = "GSM"
	LTE   Cell = "LTE"
	WCDMA Cell = "WCDMA"
	CDMA  Cell = "CDMA"
)

// Parameters holds all information that can be used for an API call to cell/search.
type Parameters struct {
	Cell_op   Cell
	Cell_net  Cell
	Cell_id   Cell
	ShowGsm   bool // Default to true.
	ShowCdma  bool // Default to true
	ShowLte   bool // Default to true.
	ShowWcdma bool // Default to true.
	wigole.SearchSsid
}

// Response is the response from an API call for cell/search.
type Response struct {
	Success      bool
	TotalResults int
	First        int
	Last         int
	ResultCount  int
	Results      []*wigole.Network
	SearchAfter  string
	Search_after int
}
