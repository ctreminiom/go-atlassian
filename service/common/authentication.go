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
}
