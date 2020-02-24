package countries

type Country struct {
	Country string
	Count   int
}

type Parameters struct {
}

type Response struct {
	Success   bool
	Countries []*Country
}
