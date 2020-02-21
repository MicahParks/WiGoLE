package detail

import (
	"gitlab.com/MicahParks/wigole/network"
)

type Parameters struct {
	NetId       string
	Operator    uint64
	Lac         uint64
	Cid         uint64
	Type        network.Network
	System      uint64
	Network     uint64
	Basestation uint64
}

type Address struct {
	Address      map[string]string
	Lat          float64
	Lon          float64
	Importance   int
	Place_id     int
	Licence      string
	Osm_type     string
	Display_name string
	Boundingbox  []float64
}

type Response struct {
	Success   bool
	Cdma      bool
	Gsm       bool
	Lte       bool
	Wcdma     bool
	Wifi      bool
	Addresses []Address
	Results   []network.Result
}
