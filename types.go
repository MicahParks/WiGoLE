package wigole

import (
	"io"
	"net/url"
	"time"
)

type Builder interface {
	Body() (io.Reader, error)
	Url() (values url.Values, err error)
}

type Location struct {
	Alt             int
	Accuracy        float64
	Lastupdt        time.Time
	Latitude        float64
	Longitude       float64
	Month           string
	Ssid            string
	Time            time.Time
	Signal          int
	Name            string
	NetId           string
	Noise           float64
	Snr             float64
	Wep             string
	Channel         int
	EncryptionValue string
}

type NetSearchResponse struct {
	Success      bool
	TotalResults int
	First        int
	Last         int
	ResultCount  int
	Results      []*WiFiNetworkWithLocation
	SearchAfter  string
	Search_after int
}

type Network struct {
	Trilat      float64
	Trilong     float64
	Ssid        string
	Qos         int
	Transid     string
	Firsttime   time.Time
	Lasttime    time.Time
	Lastupdt    time.Time
	Netid       string
	Type        string
	Name        string
	Userfound   bool
	Country     string
	Region      string
	City        string
	Housenumber string
	Road        string
	Postalcode  string
}

type SearchParameters struct {
	Onlymine       bool // Defaults to false.
	Notmine        bool
	Latrange1      float64 // Makes Latrange2 used, even if 0.
	Latrange2      float64 // Makes Latrange1 used, even if 0.
	Longrange1     float64 // Makes Longrange2 used, even if 0.
	Longrange2     float64 // Makes Longrange1 used, even if 0.
	Lastupdt       time.Time
	StartTransID   time.Time // Year level precision only.
	EndTransID     time.Time // Year level precision only.
	MinQoS         uint8     // Between 0-7.
	Variance       float64   // Between 0.001 and 0.2.
	HouseNumber    string
	Road           string
	City           string
	Region         string
	PostalCode     string
	Country        string
	ResultsPerPage uint16 // Defaults to 25 for COMMAPI, 100 for site. Bounded at 1000 for COMMAPI, 100 for site.
	SearchAfter    string // What is this?
}

type SearchSsid struct {
	Ssid     string // SSID exact match.
	Ssidlike string // SSID with '%' as a wildcard and '_' as any character.
	SearchParameters
}

type WiFiNetwork struct {
	Comment     string
	Wep         string
	Bcninterval int
	Freenet     string
	Dhcp        string
	Paynet      string
	Channel     int
	Encryption  string
	Network
}

type WiFiNetworkWithLocation struct {
	LocationData []*Location
	WiFiNetwork
}
