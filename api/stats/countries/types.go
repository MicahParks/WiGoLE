package countries

// country holds the country information for an API call to stats/countries.
type country struct {
	Country string
	Count   int
}

// Parameters holds all information that can be used for an API call to stats/countries.
type Parameters struct {
}

// Response is the response from an API call for stats/countries.
type Response struct {
	Success   bool
	Countries []*country
}
