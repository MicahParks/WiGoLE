package network

type Encryption string

const (
	None    Encryption = "None"
	WEP     Encryption = "WEP"
	WPA     Encryption = "WPA"
	WPA2    Encryption = "WPA2"
	WPA3    Encryption = "WPA3"
	Unknown Encryption = "Unknown"
)
