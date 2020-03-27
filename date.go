package wigole

import (
	"fmt"
	"time"
)

func TimeString(t time.Time) string {
	// yyyyMMdd[hhmm[ss]]
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
