package general

import (
	"gitlab.com/MicahParks/wigole/stats"
)

type Parameters struct {
}

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
