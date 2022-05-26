package models

import (
	"bytes"
	"net/http"
)

type ResponseScheme struct {
	*http.Response

	Code     int
	Endpoint string
	Method   string
	Bytes    bytes.Buffer
}
