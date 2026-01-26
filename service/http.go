package service

import (
	"context"
	"net/http"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type Connector interface {
	// NewRequest creates a new HTTP request with the given parameters.
	NewRequest(ctx context.Context, method, urlStr, contentType string, body interface{}) (*http.Request, error)

	// Call executes the request and unmarshals the response into the provided structure.
	// The response body is fully read into memory, making this unsuitable for large responses.
	Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error)

	// Do executes the request and returns the raw HTTP response.
	// Use this for streaming large responses (e.g., file downloads) where buffering
	// the entire response in memory is not desirable.
	// The caller is responsible for closing the response body.
	Do(request *http.Request) (*http.Response, error)
}
