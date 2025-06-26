package oauth2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// SetupTokenSourcesWithStorage creates token sources with optional storage support.
// It extracts storage configuration from the HTTP client if it's wrapped.
func SetupTokenSourcesWithStorage(
	ctx context.Context,
	token *common.OAuth2Token,
	oauthService common.OAuth2Service,
	httpClient common.HTTPClient,
) (*RefreshTokenSource, *ReuseTokenSource, error) {
	if token == nil {
		return nil, nil, fmt.Errorf("token cannot be nil")
	}
	if oauthService == nil {
		return nil, nil, fmt.Errorf("oauth service cannot be nil")
	}

	var refreshSource *RefreshTokenSource
	var reuseSource *ReuseTokenSource

	// Check if we have storage configuration
	if wrapper, ok := ExtractWrapper(httpClient); ok && (wrapper.Store != nil || wrapper.Callback != nil) {
		// Use storage-aware token sources
		refreshSource = NewRefreshTokenSourceWithStorage(
			ctx,
			token.RefreshToken,
			oauthService,
			wrapper.Store,
			wrapper.Callback,
		)
		reuseSource = NewReuseTokenSourceWithStore(token, refreshSource, wrapper.Store)
	} else {
		// Use regular token sources
		refreshSource = NewRefreshTokenSource(ctx, token.RefreshToken, oauthService)
		reuseSource = NewReuseTokenSource(token, refreshSource)
	}

	return refreshSource, reuseSource, nil
}

// ExtractBaseTransport extracts the base HTTP transport from an HTTP client
func ExtractBaseTransport(httpClient common.HTTPClient) http.RoundTripper {
	// If it's wrapped, get the original client
	if wrapper, ok := httpClient.(*HTTPWrapper); ok {
		httpClient = wrapper.OriginalClient
	}

	// Extract transport
	if transport, ok := httpClient.(*http.Client); ok && transport.Transport != nil {
		return transport.Transport
	} else if rt, ok := httpClient.(http.RoundTripper); ok {
		return rt
	}

	return nil
}

// CreateOAuthTransport creates an OAuth transport with the given token source
func CreateOAuthTransport(source TokenSource, base http.RoundTripper, auth common.Authentication) *Transport {
	return &Transport{
		Source: source,
		Base:   base,
		Auth:   auth,
	}
}