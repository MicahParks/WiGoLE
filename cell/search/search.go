package search

func New() *Parameters {
	params := Parameters{}
	params.ShowGsm = true
	params.ShowCdma = true
	params.ShowLte = true
	params.ShowWcdma = true
	return &params
}
