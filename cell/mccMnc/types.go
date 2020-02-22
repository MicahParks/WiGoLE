package mccMnc

type Parameters struct {
	Mcc uint
	Mnc uint
}

type Response struct { // The API should have made the response contain a list, not a nested dict.
	Key         int
	Type        string
	CountryName string
	CountryCode string
	Mcc         int
	Mnc         int
	Brand       string
	Operator    string
	Status      string
	Bands       string
	Notes       string
}
