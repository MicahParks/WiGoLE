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

// Parameters holds all information that can be used for an API call to profile/apitoken.
type Parameters struct {
	Type TokenType
}

// Response is the response from an API call for profile/apitoken.
type Response struct {
	Success bool
	Result  []*AuthToken
}

// AuthToken holds the authentication token information for an API response from profile/apitoken.
// Look at "Models" at the bottom of https://api.wigle.net/swagger
type AuthToken struct {
	AuthName string
	Token    TokenType
	Status   string
	Type     string
	PersonId int
}

// Status includes all the WiGLE defined status constants for making API calls to profile/apitoken.
type Status string

// TokenType includes all the WiGLE defined token types for making API calls to profile/apitoken.
type TokenType string
