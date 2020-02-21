package search

import (
	"time"

	"gitlab.com/MicahParks/wigole/network"
)

type Parameters struct {
	Onlymine       bool // Defaults to false.
	Notmine        bool
	Latrange1      float64 // Makes Latrange2 used, even if 0.
	Latrange2      float64 // Makes Latrange1 used, even if 0.
	Longrange1     float64 // Makes Longrange2 used, even if 0.
	Longrange2     float64 // Makes Longrange1 used, even if 0.
	Lastupdt       time.Time
	StartTransID   time.Time          // Year level precision only.
	EndTransID     time.Time          // Year level precision only.
	Encryption     network.Encryption // "None", "WEP", "WPA", "WPA2", "WPA3", "Unknown". Case insensitive.
	Freenet        bool               // Default to false.
	Paynet         bool               // Default to false.
	Netid          string             // BSSID exact match. Need first three octets. Format "0A:2C:EF:3D:25:1B".
	Ssid           string             // SSID exact match.
	Ssidlike       string             // SSID with '%' as a wildcard and '_' as any character.
	MinQoS         uint8              // Between 0-7.
	Variance       float64            // Between 0.001 and 0.2.
	HouseNumber    string
	Road           string
	City           string
	Region         string
	PostalCode     string
	Country        string
	ResultsPerPage uint16 // Defaults to 25 for COMMAPI, 100 for site. Bounded at 1000 for COMMAPI, 100 for site.
	SearchAfter    string // What is this?
}

type Result struct {
	Trilat      float64
	Trilong     float64
	Ssid        string
	Qos         int
	Transid     string
	Firsttime   time.Time
	Lasttime    time.Time
	Lastupdt    time.Time
	Netid       string
	Name        string
	Type        string
	Comment     string
	Wep         string
	Bcninterval int
	Freenet     string
	Dhcp        string
	Paynet      string
	Userfound   bool
	Channel     int
	Encryption  string
	Country     string
	Region      string
	City        string
	Housenumber string
	Road        string
	Postalcode  string
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
