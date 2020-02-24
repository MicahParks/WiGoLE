package transactions

import (
	"gitlab.com/MicahParks/wigole/file"
)

type TransLogResponse struct {
	success              bool
	results              *file.TransLog
	processingQueueDepth int
	geoQueueDepth        int
}
