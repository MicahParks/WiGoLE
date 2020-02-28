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

type Network string
