package models

import "time"

type UserTokensScheme struct {
	ID         string    `json:"id,omitempty"`
	Label      string    `json:"label,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	LastAccess time.Time `json:"lastAccess,omitempty"`
}
