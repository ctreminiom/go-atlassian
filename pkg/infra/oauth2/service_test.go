package oauth2

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/ctreminiom/go-atlassian/v2/service/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHTTPClient is a mock implementation of common.HTTPClient
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	if resp, ok := args.Get(0).(*http.Response); ok {
		return resp, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestNewOAuth2Service(t *testing.T) {
	tests := []struct {
		name       string
		httpClient common.HTTPClient
		config     *common.OAuth2Config
		wantErr    bool
	}{
		{
			name:       "valid configuration",
			httpClient: &MockHTTPClient{},
			config: &common.OAuth2Config{
				ClientID:     "test-client-id",
				ClientSecret: "test-client-secret",
				RedirectURI:  "https://example.com/callback",
			},
			wantErr: false,
		},
		{
			name:       "nil http client uses default",
			httpClient: nil,
			config: &common.OAuth2Config{
				ClientID:     "test-client-id",
				ClientSecret: "test-client-secret",
				RedirectURI:  "https://example.com/callback",
			},
			wantErr: false,
		},
		{
			name:       "nil config",
			httpClient: &MockHTTPClient{},
			config:     nil,
			wantErr:    true,
		},
		{
			name:       "missing client ID",
			httpClient: &MockHTTPClient{},
			config: &common.OAuth2Config{
				ClientID:     "",
				ClientSecret: "test-client-secret",
				RedirectURI:  "https://example.com/callback",
			},
			wantErr: true,
		},
		{
			name:       "missing client secret",
			httpClient: &MockHTTPClient{},
			config: &common.OAuth2Config{
				ClientID:     "test-client-id",
				ClientSecret: "",
				RedirectURI:  "https://example.com/callback",
			},
			wantErr: true,
		},
		{
			name:       "missing redirect URI",
			httpClient: &MockHTTPClient{},
			config: &common.OAuth2Config{
				ClientID:     "test-client-id",
				ClientSecret: "test-client-secret",
				RedirectURI:  "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, err := NewOAuth2Service(tt.httpClient, tt.config)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, service)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, service)
			}
		})
	}
}

func TestService_GetAuthorizationURL(t *testing.T) {
	service := &Service{
		config: &common.OAuth2Config{
			ClientID:    "test-client-id",
			RedirectURI: "https://example.com/callback",
		},
	}

	tests := []struct {
		name    string
		scopes  []string
		state   string
		wantErr bool
		verify  func(t *testing.T, u *url.URL)
	}{
		{
			name:    "basic authorization URL",
			scopes:  []string{"read:jira-work", "write:jira-work"},
			state:   "test-state",
			wantErr: false,
			verify: func(t *testing.T, u *url.URL) {
				assert.Equal(t, "auth.atlassian.com", u.Host)
				assert.Equal(t, "/authorize", u.Path)
				
				q := u.Query()
				assert.Equal(t, "api.atlassian.com", q.Get("audience"))
				assert.Equal(t, "test-client-id", q.Get("client_id"))
				assert.Equal(t, "read:jira-work write:jira-work offline_access", q.Get("scope"))
				assert.Equal(t, "https://example.com/callback", q.Get("redirect_uri"))
				assert.Equal(t, "test-state", q.Get("state"))
				assert.Equal(t, "code", q.Get("response_type"))
				assert.Equal(t, "consent", q.Get("prompt"))
			},
		},
		{
			name:    "empty scopes still includes offline_access",
			scopes:  []string{},
			state:   "test-state",
			wantErr: false,
			verify: func(t *testing.T, u *url.URL) {
				q := u.Query()
				assert.Equal(t, "offline_access", q.Get("scope"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := service.GetAuthorizationURL(tt.scopes, tt.state)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, u)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, u)
				if tt.verify != nil {
					tt.verify(t, u)
				}
			}
		})
	}
}

func TestService_ExchangeAuthorizationCode(t *testing.T) {
	tests := []struct {
		name           string
		code           string
		mockResponse   *http.Response
		mockError      error
		expectedToken  *common.OAuth2Token
		expectedError  bool
	}{
		{
			name: "successful token exchange",
			code: "test-auth-code",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(strings.NewReader(`{
					"access_token": "test-access-token",
					"token_type": "Bearer",
					"expires_in": 3600,
					"refresh_token": "test-refresh-token",
					"scope": "read:jira-work write:jira-work offline_access"
				}`)),
			},
			expectedToken: &common.OAuth2Token{
				AccessToken:  "test-access-token",
				TokenType:    "Bearer",
				ExpiresIn:    3600,
				RefreshToken: "test-refresh-token",
				Scope:        "read:jira-work write:jira-work offline_access",
			},
			expectedError: false,
		},
		{
			name: "token exchange error response",
			code: "invalid-code",
			mockResponse: &http.Response{
				StatusCode: http.StatusBadRequest,
				Body: io.NopCloser(strings.NewReader(`{
					"error": "invalid_grant",
					"error_description": "The provided authorization code is invalid"
				}`)),
			},
			expectedToken: nil,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(MockHTTPClient)
			service := &Service{
				httpClient: mockClient,
				config: &common.OAuth2Config{
					ClientID:     "test-client-id",
					ClientSecret: "test-client-secret",
					RedirectURI:  "https://example.com/callback",
				},
			}

			mockClient.On("Do", mock.MatchedBy(func(req *http.Request) bool {
				// Verify basic request properties
				return req.Method == http.MethodPost &&
					req.URL.String() == TokenURL &&
					req.Header.Get("Content-Type") == "application/x-www-form-urlencoded" &&
					req.Header.Get("Accept") == "application/json"
			})).Return(tt.mockResponse, tt.mockError)

			token, err := service.ExchangeAuthorizationCode(context.Background(), tt.code)
			
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, token)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedToken, token)
			}

			mockClient.AssertExpectations(t)
		})
	}
}

func TestService_GetAccessibleResources(t *testing.T) {
	tests := []struct {
		name              string
		accessToken       string
		mockResponse      *http.Response
		mockError         error
		expectedResources []*common.AccessibleResource
		expectedError     bool
	}{
		{
			name:        "successful get accessible resources",
			accessToken: "test-access-token",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(strings.NewReader(`[
					{
						"id": "resource-1",
						"url": "https://site1.atlassian.net",
						"name": "Site 1",
						"scopes": ["read:jira-work", "write:jira-work"],
						"avatarUrl": "https://site1.atlassian.net/avatar.png"
					},
					{
						"id": "resource-2",
						"url": "https://site2.atlassian.net",
						"name": "Site 2",
						"scopes": ["read:jira-work"],
						"avatarUrl": "https://site2.atlassian.net/avatar.png"
					}
				]`)),
			},
			expectedResources: []*common.AccessibleResource{
				{
					ID:        "resource-1",
					URL:       "https://site1.atlassian.net",
					Name:      "Site 1",
					Scopes:    []string{"read:jira-work", "write:jira-work"},
					AvatarURL: "https://site1.atlassian.net/avatar.png",
				},
				{
					ID:        "resource-2",
					URL:       "https://site2.atlassian.net",
					Name:      "Site 2",
					Scopes:    []string{"read:jira-work"},
					AvatarURL: "https://site2.atlassian.net/avatar.png",
				},
			},
			expectedError: false,
		},
		{
			name:        "unauthorized access",
			accessToken: "invalid-token",
			mockResponse: &http.Response{
				StatusCode: http.StatusUnauthorized,
				Body:       io.NopCloser(strings.NewReader("Unauthorized")),
			},
			expectedResources: nil,
			expectedError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(MockHTTPClient)
			service := &Service{
				httpClient: mockClient,
			}

			mockClient.On("Do", mock.MatchedBy(func(req *http.Request) bool {
				// Verify the request
				return req.Method == http.MethodGet &&
					req.URL.String() == ResourcesURL &&
					req.Header.Get("Authorization") == "Bearer "+tt.accessToken &&
					req.Header.Get("Accept") == "application/json"
			})).Return(tt.mockResponse, tt.mockError)

			resources, err := service.GetAccessibleResources(context.Background(), tt.accessToken)
			
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resources)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResources, resources)
			}

			mockClient.AssertExpectations(t)
		})
	}
}