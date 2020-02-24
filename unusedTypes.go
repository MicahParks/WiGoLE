package wigole

import (
	"time"
)

type NetCommentResponse struct {
	Success bool
	Comment string
	Netid   string
}

type Person struct {
	Userid    string
	Email     string
	Donate    string
	Joindate  time.Time
	Lastlogin time.Time
	Session   string
	Success   string
}
