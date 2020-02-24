package network

import (
	"time"
)

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

type Encryption string

type GeocodingResponse struct {
	Address      map[string]string // TODO Is this the same for all of them?`1
	Lat          float64
	Lon          float64
	Importance   float64
	Place_id     int
	Licence      string
	Osm_type     string
	Display_name string
	Boundingbox  []float64
}

type WiFiLocation struct {
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

type Network string
