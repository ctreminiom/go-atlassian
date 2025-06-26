package oauth2

import (
	"net/http"

	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// HTTPWrapper wraps an HTTP client with token storage configuration.
// This allows token storage and callbacks to be configured before the OAuth transport is created.
type HTTPWrapper struct {
	OriginalClient common.HTTPClient
	Store          TokenStore
	Callback       TokenCallback
}

// WrapHTTPClient creates a new HTTPWrapper or returns the existing one if already wrapped
func WrapHTTPClient(client common.HTTPClient) *HTTPWrapper {
	if wrapper, ok := client.(*HTTPWrapper); ok {
		return wrapper
	}
	return &HTTPWrapper{
		OriginalClient: client,
	}
}

// WithStore sets the token store on the wrapper (chainable)
func (w *HTTPWrapper) WithStore(store TokenStore) *HTTPWrapper {
	w.Store = store
	return w
}

// WithCallback sets the token callback on the wrapper (chainable)
func (w *HTTPWrapper) WithCallback(callback TokenCallback) *HTTPWrapper {
	w.Callback = callback
	return w
}

// Do implements the HTTPClient interface by delegating to the original client
func (w *HTTPWrapper) Do(req *http.Request) (*http.Response, error) {
	return w.OriginalClient.Do(req)
}

// ExtractWrapper extracts the HTTPWrapper from an HTTPClient if it's wrapped
func ExtractWrapper(client common.HTTPClient) (*HTTPWrapper, bool) {
	wrapper, ok := client.(*HTTPWrapper)
	return wrapper, ok
}