package oauth2

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/ctreminiom/go-atlassian/v2/service/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTokenSource is a mock implementation of TokenSource
type MockTokenSource struct {
	mock.Mock
}

func (m *MockTokenSource) Token() (*common.OAuth2Token, error) {
	args := m.Called()
	if token, ok := args.Get(0).(*common.OAuth2Token); ok {
		return token, args.Error(1)
	}
	return nil, args.Error(1)
}

// MockOAuth2Service is a mock implementation of OAuth2Service
type MockOAuth2Service struct {
	mock.Mock
}

func (m *MockOAuth2Service) GetAuthorizationURL(scopes []string, state string) (*url.URL, error) {
	args := m.Called(scopes, state)
	if u, ok := args.Get(0).(*url.URL); ok {
		return u, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockOAuth2Service) ExchangeAuthorizationCode(ctx context.Context, code string) (*common.OAuth2Token, error) {
	args := m.Called(ctx, code)
	if token, ok := args.Get(0).(*common.OAuth2Token); ok {
		return token, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockOAuth2Service) RefreshAccessToken(ctx context.Context, refreshToken string) (*common.OAuth2Token, error) {
	args := m.Called(ctx, refreshToken)
	if token, ok := args.Get(0).(*common.OAuth2Token); ok {
		return token, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockOAuth2Service) GetAccessibleResources(ctx context.Context, accessToken string) ([]*common.AccessibleResource, error) {
	args := m.Called(ctx, accessToken)
	if resources, ok := args.Get(0).([]*common.AccessibleResource); ok {
		return resources, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestReuseTokenSource_Token(t *testing.T) {
	tests := []struct {
		name           string
		initialToken   *common.OAuth2Token
		mockSetup      func(*MockTokenSource)
		expectedCalls  int
		waitDuration   time.Duration
		wantErr        bool
	}{
		{
			name: "returns cached token when still valid",
			initialToken: &common.OAuth2Token{
				AccessToken:  "valid-token",
				TokenType:    "Bearer",
				ExpiresIn:    3600, // 1 hour
				RefreshToken: "refresh-token",
			},
			mockSetup: func(m *MockTokenSource) {
				// Should not be called since token is still valid
			},
			expectedCalls: 0,
			waitDuration: 0,
			wantErr:      false,
		},
		{
			name: "refreshes token when expired",
			initialToken: &common.OAuth2Token{
				AccessToken:  "expired-token",
				TokenType:    "Bearer",
				ExpiresIn:    1, // 1 second - will expire immediately
				RefreshToken: "refresh-token",
			},
			mockSetup: func(m *MockTokenSource) {
				newToken := &common.OAuth2Token{
					AccessToken:  "new-token",
					TokenType:    "Bearer",
					ExpiresIn:    3600,
					RefreshToken: "new-refresh-token",
				}
				m.On("Token").Return(newToken, nil).Once()
			},
			expectedCalls: 1,
			waitDuration: 2 * time.Second,
			wantErr:      false,
		},
		{
			name: "returns error when refresh fails",
			initialToken: &common.OAuth2Token{
				AccessToken:  "expired-token",
				TokenType:    "Bearer",
				ExpiresIn:    1, // 1 second
				RefreshToken: "refresh-token",
			},
			mockSetup: func(m *MockTokenSource) {
				m.On("Token").Return(nil, fmt.Errorf("refresh failed")).Once()
			},
			expectedCalls: 1,
			waitDuration: 2 * time.Second,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSource := new(MockTokenSource)
			if tt.mockSetup != nil {
				tt.mockSetup(mockSource)
			}

			source := NewReuseTokenSource(tt.initialToken, mockSource)

			// Wait if needed to let token expire
			if tt.waitDuration > 0 {
				time.Sleep(tt.waitDuration)
			}

			token, err := source.Token()

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, token)
			}

			mockSource.AssertNumberOfCalls(t, "Token", tt.expectedCalls)
		})
	}
}

func TestRefreshTokenSource_Token(t *testing.T) {
	tests := []struct {
		name          string
		refreshToken  string
		mockResponse  *common.OAuth2Token
		mockError     error
		expectedToken *common.OAuth2Token
		wantErr       bool
	}{
		{
			name:         "successful token refresh",
			refreshToken: "refresh-token",
			mockResponse: &common.OAuth2Token{
				AccessToken:  "new-access-token",
				TokenType:    "Bearer",
				ExpiresIn:    3600,
				RefreshToken: "new-refresh-token",
			},
			expectedToken: &common.OAuth2Token{
				AccessToken:  "new-access-token",
				TokenType:    "Bearer",
				ExpiresIn:    3600,
				RefreshToken: "new-refresh-token",
			},
			wantErr: false,
		},
		{
			name:         "refresh token error",
			refreshToken: "invalid-refresh-token",
			mockError:    fmt.Errorf("invalid refresh token"),
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOAuth := new(MockOAuth2Service)
			ctx := context.Background()

			mockOAuth.On("RefreshAccessToken", ctx, tt.refreshToken).Return(tt.mockResponse, tt.mockError)

			source := NewRefreshTokenSource(ctx, tt.refreshToken, mockOAuth)
			token, err := source.Token()

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, token)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedToken, token)
			}

			mockOAuth.AssertExpectations(t)
		})
	}
}

func TestTransport_RoundTrip(t *testing.T) {
	tests := []struct {
		name          string
		tokenResponse *common.OAuth2Token
		tokenError    error
		checkRequest  func(*testing.T, *http.Request)
		wantErr       bool
	}{
		{
			name: "adds bearer token to request",
			tokenResponse: &common.OAuth2Token{
				AccessToken: "test-token",
				TokenType:   "Bearer",
			},
			checkRequest: func(t *testing.T, r *http.Request) {
				assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
			},
			wantErr: false,
		},
		{
			name:       "returns error when token source fails",
			tokenError: fmt.Errorf("token source error"),
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSource := new(MockTokenSource)
			mockSource.On("Token").Return(tt.tokenResponse, tt.tokenError)

			// Create test server
			var capturedRequest *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r.Clone(context.Background())
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			transport := &Transport{
				Source: mockSource,
			}

			req, _ := http.NewRequest("GET", server.URL, nil)
			resp, err := transport.RoundTrip(req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				resp.Body.Close()

				if tt.checkRequest != nil && capturedRequest != nil {
					tt.checkRequest(t, capturedRequest)
				}
			}

			mockSource.AssertExpectations(t)
		})
	}
}

func TestRefreshTokenSource_UpdatesRefreshToken(t *testing.T) {
	mockOAuth := new(MockOAuth2Service)
	ctx := context.Background()
	
	// First refresh returns a new refresh token
	firstResponse := &common.OAuth2Token{
		AccessToken:  "first-access-token",
		RefreshToken: "new-refresh-token",
		ExpiresIn:    3600,
	}
	mockOAuth.On("RefreshAccessToken", ctx, "old-refresh-token").Return(firstResponse, nil).Once()
	
	// Second refresh should use the new refresh token
	secondResponse := &common.OAuth2Token{
		AccessToken:  "second-access-token",
		RefreshToken: "newer-refresh-token",
		ExpiresIn:    3600,
	}
	mockOAuth.On("RefreshAccessToken", ctx, "new-refresh-token").Return(secondResponse, nil).Once()
	
	source := NewRefreshTokenSource(ctx, "old-refresh-token", mockOAuth)
	
	// First token request
	token1, err := source.Token()
	assert.NoError(t, err)
	assert.Equal(t, "first-access-token", token1.AccessToken)
	
	// Second token request should use updated refresh token
	token2, err := source.Token()
	assert.NoError(t, err)
	assert.Equal(t, "second-access-token", token2.AccessToken)
	
	mockOAuth.AssertExpectations(t)
}