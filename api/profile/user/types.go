package user

import (
	"time"
)

// Parameters holds all information that can be used for an API call to profile/user.
type Parameters struct {
}

// Person is the response from an API call for profile/user.
// Look at "Models" at the bottom of https://api.wigle.net/swagger
type Person struct {
	Userid    string
	Email     string
	Donate    string
	Joindate  time.Time
	Lastlogin time.Time
	Session   string
	Success   string
}
