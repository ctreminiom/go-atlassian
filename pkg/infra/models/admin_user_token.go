// Package models provides the data structures used in the user token management.
package models

import "time"

// UserTokensScheme represents a user token.
// ID is the unique identifier of the token.
// Label is the label assigned to the token.
// CreatedAt is the time when the token was created.
// LastAccess is the last time the token was accessed.
type UserTokensScheme struct {
	ID         string    `json:"id,omitempty"`         // The ID of the user token.
	Label      string    `json:"label,omitempty"`      // The label of the user token.
	CreatedAt  time.Time `json:"createdAt,omitempty"`  // The creation time of the user token.
	LastAccess time.Time `json:"lastAccess,omitempty"` // The last access time of the user token.
}
