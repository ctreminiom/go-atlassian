package oauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

const (
	// AuthorizationURL is the OAuth 2.0 authorization endpoint
	AuthorizationURL = "https://auth.atlassian.com/authorize"
	
	// TokenURL is the OAuth 2.0 token endpoint
	TokenURL = "https://auth.atlassian.com/oauth/token"
	
	// ResourcesURL is the endpoint to get accessible resources
	ResourcesURL = "https://api.atlassian.com/oauth/token/accessible-resources"
	
	// Audience for Atlassian APIs
	Audience = "api.atlassian.com"
)

// Service implements OAuth 2.0 authentication for Atlassian
type Service struct {
	httpClient     common.HTTPClient
	clientID       string
	clientSecret   string
	redirectURI    string
}

// NewOAuth2Service creates a new OAuth 2.0 service
func NewOAuth2Service(httpClient common.HTTPClient, clientID, clientSecret, redirectURI string) (*Service, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	
	if clientID == "" || clientSecret == "" || redirectURI == "" {
		return nil, fmt.Errorf("oauth2: clientID, clientSecret and redirectURI are required")
	}
	
	return &Service{
		httpClient:   httpClient,
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
	}, nil
}

// GetAuthorizationURL generates the authorization URL for the OAuth 2.0 flow
func (s *Service) GetAuthorizationURL(scopes []string, state string) (*url.URL, error) {
	u, err := url.Parse(AuthorizationURL)
	if err != nil {
		return nil, fmt.Errorf("oauth2: failed to parse authorization URL: %w", err)
	}
	
	// Add offline_access scope to get refresh token
	scopesWithOffline := append(scopes, "offline_access")
	
	q := u.Query()
	q.Set("audience", Audience)
	q.Set("client_id", s.clientID)
	q.Set("scope", strings.Join(scopesWithOffline, " "))
	q.Set("redirect_uri", s.redirectURI)
	q.Set("state", state)
	q.Set("response_type", "code")
	q.Set("prompt", "consent")
	u.RawQuery = q.Encode()
	
	return u, nil
}

// ExchangeAuthorizationCode exchanges the authorization code for access and refresh tokens
func (s *Service) ExchangeAuthorizationCode(ctx context.Context, code string) (*common.OAuth2Token, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", s.clientID)
	data.Set("client_secret", s.clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", s.redirectURI)
	
	return s.requestToken(ctx, data)
}

// RefreshAccessToken uses the refresh token to get a new access token
func (s *Service) RefreshAccessToken(ctx context.Context, refreshToken string) (*common.OAuth2Token, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", s.clientID)
	data.Set("client_secret", s.clientSecret)
	data.Set("refresh_token", refreshToken)
	
	return s.requestToken(ctx, data)
}

// GetAccessibleResources returns the list of Atlassian sites accessible with the current token
func (s *Service) GetAccessibleResources(ctx context.Context, accessToken string) ([]*common.AccessibleResource, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ResourcesURL, nil)
	if err != nil {
		return nil, fmt.Errorf("oauth2: failed to create request: %w", err)
	}
	
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Accept", "application/json")
	
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("oauth2: failed to get accessible resources: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("oauth2: failed to get accessible resources, status: %d, body: %s", resp.StatusCode, string(body))
	}
	
	var resources []*common.AccessibleResource
	if err := json.NewDecoder(resp.Body).Decode(&resources); err != nil {
		return nil, fmt.Errorf("oauth2: failed to decode accessible resources: %w", err)
	}
	
	return resources, nil
}

// requestToken makes a token request to the OAuth 2.0 token endpoint
func (s *Service) requestToken(ctx context.Context, data url.Values) (*common.OAuth2Token, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, TokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("oauth2: failed to create token request: %w", err)
	}
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("oauth2: failed to request token: %w", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("oauth2: failed to read response body: %w", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		var errResp struct {
			Error            string `json:"error"`
			ErrorDescription string `json:"error_description"`
		}
		if err := json.Unmarshal(body, &errResp); err == nil && errResp.Error != "" {
			return nil, fmt.Errorf("oauth2: token request failed: %s - %s", errResp.Error, errResp.ErrorDescription)
		}
		return nil, fmt.Errorf("oauth2: token request failed with status %d: %s", resp.StatusCode, string(body))
	}
	
	var token common.OAuth2Token
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, fmt.Errorf("oauth2: failed to decode token response: %w", err)
	}
	
	return &token, nil
}