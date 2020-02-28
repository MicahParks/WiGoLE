package search

import (
	"gitlab.com/MicahParks/wigole"
)

type Parameters struct {
	Netid    string
	Name     string
	Namelike string
	ShowBt   bool // Default to true.
	ShowBle  bool // Default to true.
	wigole.SearchParameters
}

type Bluetooth struct { // The description in the docs is wrong.
	Device       int
	Capabilities []string
	LocationData []*wigole.Location
	wigole.Network
}

type Response struct {
}
