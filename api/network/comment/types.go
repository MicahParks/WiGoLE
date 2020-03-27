package comment

// Parameters holds all information that can be used for an API call to network/comment.
type Parameters struct {
	Netid   string
	Comment string
}

// Response is the response from an API call for network/comment.
type Response struct {
	Success bool
	Comment string
	Netid   string
}
