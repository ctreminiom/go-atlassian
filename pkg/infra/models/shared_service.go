package models

import (
	"bytes"
	"net/http"
)

// ResponseScheme represents the response from an HTTP request.
type ResponseScheme struct {
	*http.Response // Embedding the http.Response struct from the net/http package.

	Code     int          // The HTTP status code of the response.
	Endpoint string       // The endpoint that the request was made to.
	Method   string       // The HTTP method used for the request.
	Bytes    bytes.Buffer // The response body.
}
