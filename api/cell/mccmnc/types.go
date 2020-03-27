package mccmnc

type Parameters struct {
	Mcc int
	Mnc int
}

type Mm struct {
	Type        string
	CountryName string
	CountryCode string
	Brand       string
	Operator    string
	Status      string
	Bands       string
	Notes       string
}

type AllStrings struct { // Let's convert those strings to ints.
	Mcc string
	Mnc string
	Mm
}

type MccMnc struct { // The API should have made the response contain a list, not a nested dict.
	Mcc int
	Mnc int
	Mm
}
