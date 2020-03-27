package apitoken

const (
	// Android is a type of token defined by WiGLE.
	Android TokenType = "ANDROID"
	// Api is a type of token defined by WiGLE.
	Api TokenType = "API"
	// CommApi is a type of token defined by WiGLE.
	CommApi TokenType = "COMMAPI"
	// Cookie is a type of token defined by WiGLE.
	Cookie TokenType = "COOKIE"
	// StatusActive is a status constant defined by WiGLE.
	StatusActive Status = "STATUS_ACTIVE"
	// StatusDisabled is a status constant defined by WiGLE.
	StatusDisabled Status = "STATUS_DISABLED"
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
