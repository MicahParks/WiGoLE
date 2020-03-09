package bluetooth

import (
	"gitlab.com/MicahParks/wigole"
)

type Bluetooth struct { // The description in the docs is wrong.
	Device       int
	Capabilities []string
	wigole.Network
}
