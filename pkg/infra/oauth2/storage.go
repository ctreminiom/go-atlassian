package oauth2

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// TokenStore provides an interface for persistent storage of OAuth2 tokens.
// Implementations can use any storage backend (Redis, database, file system, etc.)
//
// Error Handling:
// - SetRefreshToken errors are critical and will fail the token refresh operation
// - SetToken errors are logged but ignored to prioritize availability over consistency
// - GetToken/GetRefreshToken errors should return nil token and let the system refresh
type TokenStore interface {
	// GetToken retrieves the current OAuth2 token from storage
	GetToken(ctx context.Context) (*common.OAuth2Token, error)
	
	// SetToken stores the OAuth2 token in the storage backend.
	// Implementations should be fast as this is called frequently.
	// Errors from this method are ignored by the library.
	SetToken(ctx context.Context, token *common.OAuth2Token) error
	
	// GetRefreshToken retrieves only the refresh token from storage
	GetRefreshToken(ctx context.Context) (string, error)
	
	// SetRefreshToken stores only the refresh token in the storage backend.
	// This is critical - errors will cause the token refresh to fail.
	// Implementations MUST reliably store refresh tokens or return an error.
	SetRefreshToken(ctx context.Context, refreshToken string) error
}

// TokenCallback is called when tokens are refreshed successfully.
// This allows applications to react to token changes, log events, update caches, etc.
type TokenCallback interface {
	// OnTokenRefreshed is called after a token has been successfully refreshed.
	// oldToken may be nil if no previous token was available.
	// The implementation should be non-blocking and handle errors gracefully.
	OnTokenRefreshed(ctx context.Context, oldToken, newToken *common.OAuth2Token) error
}

// CompositeTokenCallback allows multiple callbacks to be executed
type CompositeTokenCallback struct {
	callbacks []TokenCallback
}

// NewCompositeTokenCallback creates a new composite callback
func NewCompositeTokenCallback(callbacks ...TokenCallback) *CompositeTokenCallback {
	return &CompositeTokenCallback{callbacks: callbacks}
}

// OnTokenRefreshed calls all registered callbacks
func (c *CompositeTokenCallback) OnTokenRefreshed(ctx context.Context, oldToken, newToken *common.OAuth2Token) error {
	for _, callback := range c.callbacks {
		if err := callback.OnTokenRefreshed(ctx, oldToken, newToken); err != nil {
			// Log error but continue with other callbacks
			// In production, you might want to use a logger here
			continue
		}
	}
	return nil
}