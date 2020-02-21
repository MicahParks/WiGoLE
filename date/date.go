package date

import (
	"fmt"
	"time"
)

func String(t time.Time) string {
	// yyyyMMdd[hhmm[ss]]
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
