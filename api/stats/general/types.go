package general

import (
	"gitlab.com/MicahParks/wigole/api/stats"
)

// Parameters holds all information that can be used for an API call to stats/general.
type Parameters struct {
}

// Response is the response from an API call for stats/general.
type Response struct {
	Octet          bool
	Netwpa2        int
	Android        bool
	Netwpa3        int
	Gentotal       int
	Manufacturer   bool
	Netnowep       int
	Dfltssid       int
	Dfltwpkn       int
	Trans2da       int
	Netwpa         int
	Trans1da       int
	Nettotal       int
	Bttotal        int
	Nettoday       int
	NetwepQuest    int `json:"netwep?"`
	Loctotal       int
	Netwep         int
	SsidStatistics []*stats.SsidStatistic
	btloc          int
	success        bool
	dfltnowp       int
	genloc         int
	netlocdy       int
	netloc         int
	transtdy       int
	userstot       int
	netlocd2       int
}
