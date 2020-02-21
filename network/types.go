package network

import (
	"time"
)

type Encryption string
type Network string

const (
	None    Encryption = "None"
	WEP     Encryption = "WEP"
	WPA     Encryption = "WPA"
	WPA2    Encryption = "WPA2"
	WPA3    Encryption = "WPA3"
	Unknown Encryption = "Unknown"
	CDMA    Network    = "CDMA"
	GSM     Network    = "GSM"
	LTE     Network    = "LTE"
	WIFI    Network    = "WIFI"
)

type LocationData struct {
	Alt             float64
	Accuracy        float64
	Lastupdt        time.Time
	Latitude        float64
	Longitude       float64
	Month           string
	Ssid            string
	Time            time.Time
	Signal          float64
	Name            string
	NetId           string
	Noise           float64
	snr             float64
	Wep             string
	Channel         int
	EncryptionValue string
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
	LocationData []LocationData
	Encryption   string
	Country      string
	Region       string
	City         string
	Housenumber  string
	Road         string
	Postalcode   string
}
