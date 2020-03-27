package network

const (
	// None is a encryption constant defined by WiGLE.
	None Encryption = "None"
	// Wep is a encryption constant defined by WiGLE.
	Wep Encryption = "WEP"
	// Wpa is a encryption constant defined by WiGLE.
	Wpa Encryption = "WPA"
	// Wpa2 is a encryption constant defined by WiGLE.
	Wpa2 Encryption = "WPA2"
	// Wpa3 is a encryption constant defined by WiGLE.
	Wpa3 Encryption = "WPA3"
	// Unknown is a encryption constant defined by WiGLE.
	Unknown Encryption = "Unknown"
	// Cdma is a network constant defined by WiGLE.
	Cdma Network = "CDMA"
	// Gsm is a network constant defined by WiGLE.
	Gsm Network = "GSM"
	// Lte is a network constant defined by WiGLE.
	Lte Network = "LTE"
	// Wifi is a network constant defined by WiGLE.
	Wifi Network = "WIFI"
)

// Encryption includes all the WiGLE defined encryption constants for making API calls to network/*.
type Encryption string

// GeocodingResponse is used to deserialize information returned from the WiGLE API. Used for inheritance.
type GeocodingResponse struct {
	Address     map[string]string // TODO Is this the same for all of them?
	Lat         float64
	Lon         float64
	Importance  float64
	PlaceId     int `json:"place_id"`
	Licence     string
	OsmType     string    `json:"osm_type"`
	DisplayName string    `json:"display_name"`
	BoundingBox []float64 `json:"boundingbox"`
}

// Network includes all the WiGLE defined network constants for making API calls to network/*.
type Network string
