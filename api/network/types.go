package network

const (
	None    Encryption = "None"    // Defined by WiGLE.
	WEP     Encryption = "WEP"     // Defined by WiGLE.
	WPA     Encryption = "WPA"     // Defined by WiGLE.
	WPA2    Encryption = "WPA2"    // Defined by WiGLE.
	WPA3    Encryption = "WPA3"    // Defined by WiGLE.
	Unknown Encryption = "Unknown" // Defined by WiGLE.
	CDMA    Network    = "CDMA"    // Defined by WiGLE.
	GSM     Network    = "GSM"     // Defined by WiGLE.
	LTE     Network    = "LTE"     // Defined by WiGLE.
	WIFI    Network    = "WIFI"    // Defined by WiGLE.
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
