package user

import (
	"time"
)

type Parameters struct {
}

type Response struct {
	Userid    string
	Email     string
	Donate    string
	Joindate  time.Time
	Lastlogin time.Time
	Session   string
	Success   string
}
