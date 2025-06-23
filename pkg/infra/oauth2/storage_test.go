package oauth2

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ctreminiom/go-atlassian/v2/service/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTokenStore is a mock implementation of TokenStore
type MockTokenStore struct {
	mock.Mock
}

func (m *MockTokenStore) GetToken(ctx context.Context) (*common.OAuth2Token, error) {
	args := m.Called(ctx)
	if token, ok := args.Get(0).(*common.OAuth2Token); ok {
		return token, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTokenStore) SetToken(ctx context.Context, token *common.OAuth2Token) error {
	args := m.Called(ctx, token)
	return args.Error(0)
}

func (m *MockTokenStore) GetRefreshToken(ctx context.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *MockTokenStore) SetRefreshToken(ctx context.Context, refreshToken string) error {
	args := m.Called(ctx, refreshToken)
	return args.Error(0)
}

func TestRefreshTokenSource_StorageError(t *testing.T) {
	t.Run("refresh token storage error causes token refresh to fail", func(t *testing.T) {
		// Mock OAuth service
		mockOAuth := new(MockOAuth2Service)
		mockStore := new(MockTokenStore)
		
		// Setup successful token refresh from OAuth
		newToken := &common.OAuth2Token{
			AccessToken:  "new-access-token",
			RefreshToken: "new-refresh-token",
			ExpiresIn:    3600,
		}
		mockOAuth.On("RefreshAccessToken", mock.Anything, "old-refresh-token").Return(newToken, nil)
		
		// Setup storage - the constructor tries to load refresh token from storage
		mockStore.On("GetRefreshToken", mock.Anything).Return("", errors.New("not found"))
		
		// Setup storage error for refresh token
		mockStore.On("SetRefreshToken", mock.Anything, "new-refresh-token").Return(errors.New("storage failed"))
		
		// Create source with storage
		source := NewRefreshTokenSourceWithStorage(
			context.Background(),
			"old-refresh-token",
			mockOAuth,
			mockStore,
			nil,
		)
		
		// Try to refresh token
		token, err := source.Token()
		
		// Should fail due to storage error
		assert.Error(t, err)
		assert.Nil(t, token)
		assert.Contains(t, err.Error(), "failed to store refresh token")
		assert.Contains(t, err.Error(), "storage failed")
		
		// Verify OAuth was called
		mockOAuth.AssertCalled(t, "RefreshAccessToken", mock.Anything, "old-refresh-token")
		
		// Verify storage was attempted
		mockStore.AssertCalled(t, "SetRefreshToken", mock.Anything, "new-refresh-token")
		
		// Verify SetToken was NOT called (should fail before that)
		mockStore.AssertNotCalled(t, "SetToken", mock.Anything, mock.Anything)
	})
	
	t.Run("access token storage error is ignored", func(t *testing.T) {
		// Mock OAuth service
		mockOAuth := new(MockOAuth2Service)
		mockStore := new(MockTokenStore)
		
		// Setup successful token refresh from OAuth (no new refresh token)
		newToken := &common.OAuth2Token{
			AccessToken:  "new-access-token",
			RefreshToken: "", // No new refresh token
			ExpiresIn:    3600,
		}
		mockOAuth.On("RefreshAccessToken", mock.Anything, "old-refresh-token").Return(newToken, nil)
		
		// Setup storage - the constructor tries to load refresh token from storage
		mockStore.On("GetRefreshToken", mock.Anything).Return("", errors.New("not found"))
		
		// Setup storage error for access token (this should be ignored)
		mockStore.On("SetToken", mock.Anything, newToken).Return(errors.New("storage failed"))
		
		// Create source with storage
		source := NewRefreshTokenSourceWithStorage(
			context.Background(),
			"old-refresh-token",
			mockOAuth,
			mockStore,
			nil,
		)
		
		// Try to refresh token
		token, err := source.Token()
		
		// Should succeed despite storage error
		assert.NoError(t, err)
		assert.NotNil(t, token)
		assert.Equal(t, "new-access-token", token.AccessToken)
		
		// Wait a bit for async operation
		// In real code, this would be handled differently
		// but for testing we need to ensure the goroutine runs
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		<-ctx.Done()
		
		// Verify storage was attempted (eventually)
		mockStore.AssertCalled(t, "SetToken", mock.Anything, newToken)
	})
}