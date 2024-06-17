// Package models provides the data structures used in the agile management.
package models

// EpicScheme represents an agile epic.
type EpicScheme struct {
	ID      int              `json:"id,omitempty"`      // The ID of the epic.
	Key     string           `json:"key,omitempty"`     // The key of the epic.
	Self    string           `json:"self,omitempty"`    // The self URL of the epic.
	Name    string           `json:"name,omitempty"`    // The name of the epic.
	Summary string           `json:"summary,omitempty"` // The summary of the epic.
	Color   *EpicColorScheme `json:"color,omitempty"`   // The color scheme of the epic.
	Done    bool             `json:"done,omitempty"`    // The status of the epic.
}

// EpicColorScheme represents the color scheme of an epic.
// Key is the key of the color scheme.
type EpicColorScheme struct {
	Key string `json:"key,omitempty"` // The key of the color scheme.
}
