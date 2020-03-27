package site

// Parameters holds all information that can be used for an API call to stats/site.
type Parameters struct {
}

// Response is the response from an API call for stats/site.
type Response struct {
	GeoQueue    int
	Netwpa2     int
	Netwpa3     int
	Gentotal    int
	Netnowep    int
	Dfltssid    int
	Dfltwpkn    int
	Trans2da    int
	Netwpa      int
	Trans1da    int
	Nettotal    int
	Bttotal     int
	Nettoday    int
	NetwepQuest int `json:"netwep?"`
	Loctotal    int
	Netwep      int
	Netwwwd3    int
	Transtot    int
	WaitQueue   int
	Size        int
	Btloc       int
	Success     bool
	Dfltnowp    int
	Genloc      int
	Netlocdy    int
	Netloc      int
	Transtdy    int
	Userstot    int
	Netlocd2    int
}
