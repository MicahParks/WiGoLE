package mccmnc

// allStrings holds what from the response to an API call to cell/mmcmnc deserializes to, which is strings.
type allStrings struct { // Let's convert those strings to ints.
	Mcc string
	Mnc string
	Mm
}

// Parameters holds all information that can be used for an API call to cell/mmcmnc.
type Parameters struct {
	Mcc int
	Mnc int
}

// Mm holds the values that are meant to be strings for the response from an API call to cell/mmcmnc.
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

// MccMnc is part of the response from the API call for cell/mmcmnc.
type MccMnc struct { // The API should have made the response contain a list, not a nested dict.
	Mcc int
	Mnc int
	Mm
}
