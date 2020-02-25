package comment

type Parameters struct {
	Netid   string
	Comment string
}

type NetCommentResponse struct {
	Success bool
	Comment string
	Netid   string
}
