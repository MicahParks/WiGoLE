package search

import (
	"gitlab.com/MicahParks/wigole"
)

type Cell string

const (
	GSM   Cell = "GSM"
	LTE   Cell = "LTE"
	WCDMA Cell = "WCDMA"
	CDMA  Cell = "CDMA"
)

type Parameters struct {
	Cell_op   Cell
	Cell_net  Cell
	Cell_id   Cell
	ShowGsm   bool // Default to true.
	ShowCdma  bool // Default to true
	ShowLte   bool // Default to true.
	ShowWcdma bool // Default to true.
	wigole.Parameters
}
