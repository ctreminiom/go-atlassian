package internal

import (
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/common"
)

// NewAuthenticationService returns a new instance of the AuthenticationService.
func NewAuthenticationService(client service.Connector) common.Authentication {
	return &AuthenticationService{c: client}
}

// AuthenticationService handles communication with the authentication related.
type AuthenticationService struct {
	c service.Connector

	basicAuthProvided bool
	mail, token       string

	userAgentProvided bool
	agent             string
}

// SetBearerToken sets the token to be used in the Authorization header.
func (a *AuthenticationService) SetBearerToken(token string) {
	a.token = token
}

// GetBearerToken returns the token used in the Authorization header.
func (a *AuthenticationService) GetBearerToken() string {
	return a.token
}

// SetExperimentalFlag sets the experimental flag to be used in the Authorization header.
func (a *AuthenticationService) SetExperimentalFlag() {}

// HasSetExperimentalFlag returns if the experimental flag was set.
func (a *AuthenticationService) HasSetExperimentalFlag() bool {
	return false
}

// SetBasicAuth sets the mail and token to be used in the Authorization header.
func (a *AuthenticationService) SetBasicAuth(mail, token string) {
	a.mail = mail
	a.token = token

	a.basicAuthProvided = true
}

// GetBasicAuth returns the mail and token used in the Authorization header.
func (a *AuthenticationService) GetBasicAuth() (string, string) {
	return a.mail, a.token
}

// HasBasicAuth returns if the mail and token were set.
func (a *AuthenticationService) HasBasicAuth() bool {
	return a.basicAuthProvided
}

// SetUserAgent sets the agent to be used in the User-Agent header.
func (a *AuthenticationService) SetUserAgent(agent string) {
	a.agent = agent
	a.userAgentProvided = true
}

// GetUserAgent returns the agent used in the User-Agent header.
func (a *AuthenticationService) GetUserAgent() string {
	return a.agent
}

// HasUserAgent returns if the agent was set.
func (a *AuthenticationService) HasUserAgent() bool {
	return a.userAgentProvided
}
