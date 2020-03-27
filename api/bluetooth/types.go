package bluetooth

import (
	"gitlab.com/MicahParks/wigole"
)

// Bluetooth is used to deserialize information returned from the WiGLE API. Used for inheritance.
type Bluetooth struct { // The description in the docs is wrong.
	Device       int
	Capabilities []string
	wigole.Network
}
