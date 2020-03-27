package network

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

// Encryption includes all the WiGLE defined encryption constants for making API calls to network/*.
type Encryption string

// GeocodingResponse is used to deserialize information returned from the WiGLE API. Used for inheritance.
type GeocodingResponse struct {
	Address      map[string]string // TODO Is this the same for all of them?
	Lat          float64
	Lon          float64
	Importance   float64
	Place_id     int
	Licence      string
	Osm_type     string
	Display_name string
	Boundingbox  []float64
}

// Network includes all the WiGLE defined network constants for making API calls to network/*.
type Network string
