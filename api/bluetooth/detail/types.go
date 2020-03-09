package detail

import (
	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/api/bluetooth"
)

type Parameters struct {
	Netid          string
	ReverseAddress string
}

type Response struct {
	wigole.Detail
	Results []*bluetooth.Bluetooth
}
