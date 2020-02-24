package apiToken

const (
	ANDROID         TokenType = "ANDROID"
	API             TokenType = "API"
	COMMAPI         TokenType = "COMMAPI"
	COOKIE          TokenType = "COOKIE"
	STATUS_ACTIVE   Status    = "STATUS_ACTIVE"
	STATUS_DISABLED Status    = "STATUS_DISABLED"
)

type Parameters struct {
	Type *TokenType
}

type AuthTokenResponse struct {
	Success bool
	Result  []*AuthToken
}

type AuthToken struct {
	AuthName string
	Token    *TokenType
	Status   string
	Type     string
	PersonId int
}

type Status string

type TokenType string
