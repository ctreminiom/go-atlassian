package models

import (
	"fmt"
	"strconv"
	"time"
)

// ConfluenceDateTimeScheme is a custom time type for Confluence v2 API timestamps.
//
// The Confluence v2 API returns timestamps as ISO 8601 strings
// (e.g. "2024-09-23T20:17:35.607Z"), but some endpoints may return
// Unix epoch milliseconds as numbers.
//
// This type handles both formats during JSON unmarshaling and marshals
// back to RFC 3339 format.
type ConfluenceDateTimeScheme time.Time

// MarshalJSON marshals the ConfluenceDateTimeScheme to JSON as an RFC 3339 string.
func (d *ConfluenceDateTimeScheme) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(*d).Format(time.RFC3339))), nil
}

// UnmarshalJSON unmarshals the ConfluenceDateTimeScheme from JSON.
//
// It accepts:
//   - RFC 3339 strings with optional fractional seconds (e.g. "2024-09-23T20:17:35.607Z")
//   - Unix epoch millisecond numbers (e.g. 1727122655607)
func (d *ConfluenceDateTimeScheme) UnmarshalJSON(data []byte) error {
	s := string(data)

	if s == "null" {
		return nil
	}

	// Quoted string: parse as RFC 3339 (with optional fractional seconds).
	if len(data) >= 2 && data[0] == '"' {
		raw := s[1 : len(s)-1]

		parsed, err := time.Parse(time.RFC3339Nano, raw)
		if err != nil {
			return fmt.Errorf("cannot parse %q as Confluence timestamp: %w", raw, err)
		}

		*d = ConfluenceDateTimeScheme(parsed)
		return nil
	}

	// Unquoted number: parse as Unix epoch milliseconds.
	ms, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("cannot parse %s as Confluence timestamp: %w", s, err)
	}

	*d = ConfluenceDateTimeScheme(time.UnixMilli(ms))
	return nil
}
