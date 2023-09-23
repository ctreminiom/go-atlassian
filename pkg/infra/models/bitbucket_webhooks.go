package models

type WebhookSubscriptionPayloadScheme struct {
	Description string   `json:"description,omitempty"`
	Url         string   `json:"url,omitempty"`
	Active      bool     `json:"active,omitempty"`
	Events      []string `json:"events,omitempty"`
}

type WebhookSubscriptionPageScheme struct {
	Size     int                          `json:"size,omitempty"`
	Page     int                          `json:"page,omitempty"`
	Pagelen  int                          `json:"pagelen,omitempty"`
	Next     string                       `json:"next,omitempty"`
	Previous string                       `json:"previous,omitempty"`
	Values   []*WebhookSubscriptionScheme `json:"values,omitempty"`
}

type WebhookSubscriptionScheme struct {
	UUID        string                            `json:"uuid,omitempty"`
	URL         string                            `json:"url,omitempty"`
	Description string                            `json:"description,omitempty"`
	SubjectType string                            `json:"subject_type,omitempty"`
	Subject     *WebhookSubscriptionSubjectScheme `json:"subject,omitempty"`
	Active      bool                              `json:"active,omitempty"`
	CreatedAt   string                            `json:"created_at,omitempty"`
	Events      []string                          `json:"events,omitempty"`
}

type WebhookSubscriptionSubjectScheme struct {
	Type string `json:"type,omitempty"`
}
