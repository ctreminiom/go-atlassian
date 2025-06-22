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

// AuthenticationService provides methods to manage authentication in Jira Service Management.
type AuthenticationService struct {
	// c is the connector interface for authentication operations.
	c service.Connector

	// basicAuthProvided indicates if basic authentication credentials have been provided.
	basicAuthProvided bool
	// mail is the email address used for basic authentication.
	// token is the token used for basic authentication.
	mail, token string

	// userAgentProvided indicates if a user agent has been provided.
	userAgentProvided bool
	// agent is the user agent string.
	agent string
	
	// OAuth 2.0 fields
	oauth2ConfigProvided bool
	clientID, clientSecret, redirectURI string
	
	oauth2AccessTokenProvided bool
	oauth2AccessToken string
	
	oauth2RefreshTokenProvided bool
	oauth2RefreshToken string
}

// SetBearerToken sets the bearer token for authentication.
func (a *AuthenticationService) SetBearerToken(token string) {
	a.token = token
}

// GetBearerToken returns the bearer token used for authentication.
func (a *AuthenticationService) GetBearerToken() string {
	return a.token
}

// SetExperimentalFlag is a placeholder for setting an experimental flag.
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

// HasUserAgent returns true if a user agent has been provided.
func (a *AuthenticationService) HasUserAgent() bool {
	return a.userAgentProvided
}

// SetOAuth2Config sets the OAuth 2.0 configuration.
func (a *AuthenticationService) SetOAuth2Config(clientID, clientSecret, redirectURI string) {
	a.clientID = clientID
	a.clientSecret = clientSecret
	a.redirectURI = redirectURI
	a.oauth2ConfigProvided = true
}

// GetOAuth2Config returns the OAuth 2.0 configuration.
func (a *AuthenticationService) GetOAuth2Config() (string, string, string) {
	return a.clientID, a.clientSecret, a.redirectURI
}

// HasOAuth2Config returns true if OAuth 2.0 configuration has been provided.
func (a *AuthenticationService) HasOAuth2Config() bool {
	return a.oauth2ConfigProvided
}

// SetOAuth2AccessToken sets the OAuth 2.0 access token.
func (a *AuthenticationService) SetOAuth2AccessToken(token string) {
	a.oauth2AccessToken = token
	a.oauth2AccessTokenProvided = true
}

// GetOAuth2AccessToken returns the OAuth 2.0 access token.
func (a *AuthenticationService) GetOAuth2AccessToken() string {
	return a.oauth2AccessToken
}

// HasOAuth2AccessToken returns true if an OAuth 2.0 access token has been provided.
func (a *AuthenticationService) HasOAuth2AccessToken() bool {
	return a.oauth2AccessTokenProvided
}

// SetOAuth2RefreshToken sets the OAuth 2.0 refresh token.
func (a *AuthenticationService) SetOAuth2RefreshToken(token string) {
	a.oauth2RefreshToken = token
	a.oauth2RefreshTokenProvided = true
}

// GetOAuth2RefreshToken returns the OAuth 2.0 refresh token.
func (a *AuthenticationService) GetOAuth2RefreshToken() string {
	return a.oauth2RefreshToken
}

// HasOAuth2RefreshToken returns true if an OAuth 2.0 refresh token has been provided.
func (a *AuthenticationService) HasOAuth2RefreshToken() bool {
	return a.oauth2RefreshTokenProvided
}
