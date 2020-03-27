package apitoken

const (
	ANDROID         TokenType = "ANDROID"         // Defined by WiGLE.
	API             TokenType = "API"             // Defined by WiGLE.
	COMMAPI         TokenType = "COMMAPI"         // Defined by WiGLE.
	COOKIE          TokenType = "COOKIE"          // Defined by WiGLE.
	STATUS_ACTIVE   Status    = "STATUS_ACTIVE"   // Defined by WiGLE.
	STATUS_DISABLED Status    = "STATUS_DISABLED" // Defined by WiGLE.
)

type Parameters struct {
	Type TokenType
}

type AuthTokenResponse struct {
	Success bool
	Result  []*AuthToken
}

type AuthToken struct {
	AuthName string
	Token    TokenType
	Status   string
	Type     string
	PersonId int
}

type Status string

type TokenType string
