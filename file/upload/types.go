package upload

type TransidResponse struct {
	File    string
	Size    int
	TransId string
}

type UploadResponse struct {
	Success  bool
	Warning  string
	Results  *UploadResultsResponse
	Observer string
}

type UploadResultsResponse struct {
	TimeTaken string
	Filesize  int
	Filename  string
	Transids  []*TransidResponse
}
