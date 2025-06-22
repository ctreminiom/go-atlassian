package oauth2

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// TokenSource provides OAuth2 tokens
type TokenSource interface {
	// Token returns a token or an error.
	// Token must be safe for concurrent use by multiple goroutines.
	Token() (*common.OAuth2Token, error)
}

// ReuseTokenSource returns a TokenSource that reuses the same token as long as it's valid.
type ReuseTokenSource struct {
	mu          sync.Mutex
	token       *common.OAuth2Token
	expiryTime  time.Time
	tokenSource TokenSource
}

// NewReuseTokenSource creates a new ReuseTokenSource
func NewReuseTokenSource(token *common.OAuth2Token, tokenSource TokenSource) *ReuseTokenSource {
	expiryTime := time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	return &ReuseTokenSource{
		token:       token,
		expiryTime:  expiryTime,
		tokenSource: tokenSource,
	}
}

// Token returns the current token if it's still valid, otherwise refreshes it
func (s *ReuseTokenSource) Token() (*common.OAuth2Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// If token is still valid (with 5 minute buffer), return it
	if s.token != nil && time.Now().Add(5*time.Minute).Before(s.expiryTime) {
		return s.token, nil
	}

	// Get new token from underlying source
	token, err := s.tokenSource.Token()
	if err != nil {
		return nil, err
	}

	s.token = token
	s.expiryTime = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	return token, nil
}

// RefreshTokenSource refreshes tokens using OAuth2Service
type RefreshTokenSource struct {
	ctx          context.Context
	refreshToken string
	oauth2       common.OAuth2Service
	mu           sync.Mutex
}

// NewRefreshTokenSource creates a new RefreshTokenSource
func NewRefreshTokenSource(ctx context.Context, refreshToken string, oauth2 common.OAuth2Service) *RefreshTokenSource {
	return &RefreshTokenSource{
		ctx:          ctx,
		refreshToken: refreshToken,
		oauth2:       oauth2,
	}
}

// Token refreshes and returns a new token
func (s *RefreshTokenSource) Token() (*common.OAuth2Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	token, err := s.oauth2.RefreshAccessToken(s.ctx, s.refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	// Update refresh token if a new one was provided
	if token.RefreshToken != "" {
		s.refreshToken = token.RefreshToken
	}

	return token, nil
}

// Transport is an http.RoundTripper that automatically adds OAuth2 tokens to requests
type Transport struct {
	Source TokenSource
	Base   http.RoundTripper
	Auth   common.Authentication
}

// RoundTrip implements http.RoundTripper
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.Source.Token()
	if err != nil {
		return nil, fmt.Errorf("oauth2: failed to get token: %w", err)
	}

	// Update the authentication with the new token
	if t.Auth != nil {
		t.Auth.SetBearerToken(token.AccessToken)
	}

	// Clone request to avoid modifying the original
	req2 := req.Clone(req.Context())
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	base := t.Base
	if base == nil {
		base = http.DefaultTransport
	}

	return base.RoundTrip(req2)
}

// Do implements the HTTPClient interface
func (t *Transport) Do(req *http.Request) (*http.Response, error) {
	return t.RoundTrip(req)
}