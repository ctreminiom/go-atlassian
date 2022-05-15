package models

import "bytes"

type ResponseScheme struct {
	Code     int
	Endpoint string
	Method   string
	Bytes    bytes.Buffer
	Headers  map[string][]string
}
