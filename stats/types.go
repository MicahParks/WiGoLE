package stats

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

type Region struct {
	Region string
	Count  int
}

type SsidStatistic struct {
	Name  string
	value int
}
