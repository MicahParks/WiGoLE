package mccMnc

type Parameters struct {
	Mcc int
	Mnc int
}

type AllStrings struct { // Let's convert those strings to ints.
	Type        string
	CountryName string
	CountryCode string
	Mcc         string
	Mnc         string
	Brand       string
	Operator    string
	Status      string
	Bands       string
	Notes       string
}

type MccMnc struct { // The API should have made the response contain a list, not a nested dict.
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
