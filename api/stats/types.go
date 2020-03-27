package stats

// UserStandings is used to deserialize information returned from the WiGLE API. Used for inheritance.
type UserStandings struct {
	Rank                     int
	MonthRank                int
	UserName                 string
	DiscoveredWiFiGPS        int
	DiscoveredWiFiGPSPercent float64
	DiscoveredWiFi           int
	DiscoveredCellGPS        int
	DiscoveredCell           int
	DiscoveredBtGPS          int
	DiscoveredBt             int
	EventMonthCount          int
	EventPrevMonthCount      int
	PrevRank                 int
	PrevMonthRank            int
	TotalWiFiLocations       int
	First                    string
	Last                     string
	Self                     bool
}

// Region is used to deserialize information returned from the WiGLE API. Used for inheritance.
type Region struct {
	Region string
	Count  int
}

// SsidStatistic is used to deserialize information returned from the WiGLE API. Used for inheritance.
type SsidStatistic struct {
	Name  string
	value int
}
