package transactions

import (
	"gitlab.com/MicahParks/wigole/file"
)

type TransLogResponse struct {
	Success              bool
	Results              *file.TransLog
	ProcessingQueueDepth int
	GeoQueueDepth        int
}
