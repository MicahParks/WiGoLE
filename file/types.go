package file

import (
	"time"
)

type TransLog struct {
	Transid           string
	Username          string
	FirstTime         time.Time
	Lastupdt          time.Time
	FileName          string
	FileSize          int
	FileLines         int
	Status            string
	DiscoveredGps     int
	Discovered        int
	Total             int
	TotalGps          int
	TotalLocations    int
	PercentDone       float64
	TimeParsing       int
	GenDiscovered     int
	GenDiscoveredGps  int
	GenTotal          int
	GenTotalGps       int
	GenTotalLocations int
	BtDiscovered      int
	BtDiscoveredGps   int
	BtTotal           int
	BtTotalGps        int
	BtTotalLocations  int
	Wait              int
}
