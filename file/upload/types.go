package upload

type TransidResponse struct {
	file    string
	size    int
	transId string
}

type UploadResponse struct {
	success  bool
	warning  string
	results  *UploadResultsResponse
	observer string
}

type UploadResultsResponse struct {
	timeTaken string
	filesize  int
	filename  string
	transids  *TransidResponse
}
