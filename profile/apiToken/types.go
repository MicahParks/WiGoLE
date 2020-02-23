package apiToken

const (
	ANDROID TokenType = "ANDROID"
	API     TokenType = "API"
	COMMAPI TokenType = "COMMAPI"
)

type Parameters struct {
	Type TokenType
}

type Response struct {
	Success bool
	Result  []*Result
}

type Result struct {
	AuthName string
	Token    string
	Status   string
	Type     string
	PersonId int
}

type TokenType string
