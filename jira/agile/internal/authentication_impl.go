package internal

import (
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// NewAuthenticationService creates a new instance of AuthenticationService.
// It takes a service.Connector as input and returns a common.Authentication interface.
func NewAuthenticationService(client service.Connector) common.Authentication {
	return &AuthenticationService{c: client}
}

// AuthenticationService provides methods to handle authentication operations.
type AuthenticationService struct {
	// c is the connector interface for authentication operations.
	c service.Connector

	// basicAuthProvided indicates if basic authentication credentials have been provided.
	basicAuthProvided bool
	// mail is the email used for basic authentication.
	// token is the token used for basic authentication.
	mail, token string

	// userAgentProvided indicates if a user agent has been provided.
	userAgentProvided bool
	// agent is the user agent string.
	agent string
}

// SetBearerToken sets the bearer token for authentication.
func (a *AuthenticationService) SetBearerToken(token string) {
	a.token = token
}

// GetBearerToken returns the bearer token used for authentication.
func (a *AuthenticationService) GetBearerToken() string {
	return a.token
}

// SetExperimentalFlag is a placeholder method for setting an experimental flag.
func (a *AuthenticationService) SetExperimentalFlag() {}

// HasSetExperimentalFlag returns false indicating the experimental flag is not set.
func (a *AuthenticationService) HasSetExperimentalFlag() bool {
	return false
}

// SetBasicAuth sets the basic authentication credentials.
func (a *AuthenticationService) SetBasicAuth(mail, token string) {
	a.mail = mail
	a.token = token

	a.basicAuthProvided = true
}

// GetBasicAuth returns the email and token used for basic authentication.
func (a *AuthenticationService) GetBasicAuth() (string, string) {
	return a.mail, a.token
}

// HasBasicAuth returns true if basic authentication credentials have been provided.
func (a *AuthenticationService) HasBasicAuth() bool {
	return a.basicAuthProvided
}

// SetUserAgent sets the user agent string.
func (a *AuthenticationService) SetUserAgent(agent string) {
	a.agent = agent
	a.userAgentProvided = true
}

// GetUserAgent returns the user agent string.
func (a *AuthenticationService) GetUserAgent() string {
	return a.agent
}

// HasUserAgent returns true if a user agent string has been provided.
func (a *AuthenticationService) HasUserAgent() bool {
	return a.userAgentProvided
}
