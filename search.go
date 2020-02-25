package wigole

import (
	"errors"
	"net/url"
	"strconv"

	"gitlab.com/MicahParks/wigole/date"
)

var errVariance = errors.New("variance must be between 0.001 and 0.2")

func (p *SearchParameters) SearchUrl() (values url.Values, err error) {
	values = url.Values{}
	values.Set("onlymine", strconv.FormatBool(p.Onlymine))
	if p.Notmine {
		values.Set("notmine", strconv.FormatBool(p.Notmine))
	}
	if p.Latrange1 != 0 || p.Latrange2 != 0 {
		values.Set("latrange1", strconv.FormatFloat(p.Latrange1, 'f', -1, 64))
		values.Set("latrange2", strconv.FormatFloat(p.Latrange2, 'f', -1, 64))
	}
	if p.Longrange1 != 0 || p.Longrange2 != 0 {
		values.Set("longrange1", strconv.FormatFloat(p.Longrange1, 'f', -1, 64))
		values.Set("longrange2", strconv.FormatFloat(p.Longrange2, 'f', -1, 64))
	}
	if !p.Lastupdt.IsZero() {
		values.Set("lastupdt", date.String(p.Lastupdt))
	}
	if !p.StartTransID.IsZero() {
		values.Set("startTransID", strconv.Itoa(p.StartTransID.Year()))
	}
	if !p.EndTransID.IsZero() {
		values.Set("endTransID", strconv.Itoa(p.EndTransID.Year()))
	}
	if p.MinQoS != 8 {
		values.Set("minQoS", strconv.Itoa(int(p.MinQoS)))
	}
	if p.Variance < 0.001 && p.Variance > 0.2 {
		values.Set("variance", strconv.FormatFloat(p.Variance, 'f', -1, 64))
	} else if p.Variance != 0 {
		return url.Values{}, errVariance
	}
	if len(p.HouseNumber) != 0 {
		values.Set("houseNumber", p.HouseNumber)
	}
	if len(p.Road) != 0 {
		values.Set("road", p.Road)
	}
	if len(p.City) != 0 {
		values.Set("city", p.City)
	}
	if len(p.Region) != 0 {
		values.Set("region", p.Region)
	}
	if len(p.PostalCode) != 0 {
		values.Set("postalCode", p.PostalCode)
	}
	if len(p.Country) != 0 {
		values.Set("country", p.Country)
	}
	if p.ResultsPerPage != 0 {
		values.Set("resultsPerPage", strconv.Itoa(int(p.ResultsPerPage)))
	}
	if len(p.SearchAfter) != 0 {
		values.Set("searchAfter", p.SearchAfter)
	}
	return values, nil
}

func NewSearch() *SearchParameters {
	return &SearchParameters{
		MinQoS: 8,
	}
}

func NewSsid() *SearchSsid {
	s := SearchSsid{}
	s.SearchParameters = *NewSearch()
	return &s
}

func (p *SearchSsid) SsidUrl() (values url.Values, err error) {
	values, err = p.SearchParameters.SearchUrl()
	if err != nil {
		return nil, err
	}
	if len(p.Ssid) != 0 {
		values.Set("ssid", p.Ssid)
	}
	if len(p.Ssidlike) != 0 {
		values.Set("ssidlike", p.Ssidlike)
	}
	return values, err
}
