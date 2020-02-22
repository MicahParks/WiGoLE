package wigole

import (
	"errors"
	"io"
	"time"

	"gitlab.com/MicahParks/wigole/network"
)

var ErrUpdate = errors.New("this wrapper needs to be updated")

type Builder interface {
	Body() (io.Reader, error)
	Url() (string, error)
}

type Response struct {
	Success      bool
	TotalResults int
	First        int
	Last         int
	ResultCount  int
	Results      []Result
	SearchAfter  string
	Search_after int
}

type Result struct {
	Trilat       float64
	Trilong      float64
	Ssid         string
	Qos          int
	Transid      string
	Firsttime    time.Time
	Lasttime     time.Time
	Lastupdt     time.Time
	Netid        string
	Name         string
	Type         string
	Comment      string
	Wep          string
	Bcninterval  int
	Freenet      string
	Dhcp         string
	Paynet       string
	Userfound    bool
	Channel      int
	LocationData []network.LocationData
	Encryption   string
	Country      string
	Region       string
	City         string
	Housenumber  string
	Road         string
	Postalcode   string
}
