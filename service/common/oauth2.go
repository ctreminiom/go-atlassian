package common

import (
	"context"
	"net/url"
)

// OAuth2Config holds OAuth 2.0 configuration
type OAuth2Config struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scopes       []string
}

// OAuth2Service handles OAuth 2.0 authentication flow
type OAuth2Service interface {
	// GetAuthorizationURL generates the authorization URL for the OAuth 2.0 flow
	GetAuthorizationURL(scopes []string, state string) (*url.URL, error)
	
	// ExchangeAuthorizationCode exchanges the authorization code for access and refresh tokens
	ExchangeAuthorizationCode(ctx context.Context, code string) (*OAuth2Token, error)
	
	// RefreshAccessToken uses the refresh token to get a new access token
	RefreshAccessToken(ctx context.Context, refreshToken string) (*OAuth2Token, error)
	
	// GetAccessibleResources returns the list of Atlassian sites accessible with the current token
	GetAccessibleResources(ctx context.Context, accessToken string) ([]*AccessibleResource, error)
}

// OAuth2Token represents OAuth 2.0 token response
type OAuth2Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope"`
}

// AccessibleResource represents an Atlassian site accessible with OAuth token
type AccessibleResource struct {
	ID       string   `json:"id"`
	URL      string   `json:"url"`
	Name     string   `json:"name"`
	Scopes   []string `json:"scopes"`
	AvatarURL string  `json:"avatarUrl"`
}