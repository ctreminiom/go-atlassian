package common

type Authentication interface {
	SetBasicAuth(mail, token string)
	GetBasicAuth() (string, string)
	HasBasicAuth() bool

	SetUserAgent(agent string)
	GetUserAgent() string
	HasUserAgent() bool

	SetExperimentalFlag()
	HasSetExperimentalFlag() bool

	SetBearerToken(token string)
	GetBearerToken() string
	
	// OAuth 2.0 (3LO) methods
	SetOAuth2Config(clientID, clientSecret, redirectURI string)
	GetOAuth2Config() (clientID, clientSecret, redirectURI string)
	HasOAuth2Config() bool
	
	SetOAuth2AccessToken(token string)
	GetOAuth2AccessToken() string
	HasOAuth2AccessToken() bool
	
	SetOAuth2RefreshToken(token string)
	GetOAuth2RefreshToken() string
	HasOAuth2RefreshToken() bool
}
