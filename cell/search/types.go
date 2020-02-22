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
	ShowGsm   bool
	ShowCdma  bool
	ShowLte   bool
	ShowWcdma bool
	wigole.Parameters
}
