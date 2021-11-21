package models

type PriorityScheme struct {
	Self        string `json:"self,omitempty"`
	StatusColor string `json:"statusColor,omitempty"`
	Description string `json:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty"`
	ID          string `json:"id,omitempty"`
}
