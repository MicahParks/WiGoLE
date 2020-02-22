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

type Network string
