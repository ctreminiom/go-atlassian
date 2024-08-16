package models

// WebhookSubscriptionPayloadScheme represents the payload for a webhook subscription.
type WebhookSubscriptionPayloadScheme struct {
	Description string   `json:"description,omitempty"` // The description of the webhook subscription.
	URL         string   `json:"url,omitempty"`         // The URL of the webhook subscription.
	Active      bool     `json:"active,omitempty"`      // Indicates if the webhook subscription is active.
	Events      []string `json:"events,omitempty"`      // The events for the webhook subscription.
}

// WebhookSubscriptionPageScheme represents a paginated list of webhook subscriptions.
type WebhookSubscriptionPageScheme struct {
	Size     int                          `json:"size,omitempty"`     // The number of subscriptions in the current page.
	Page     int                          `json:"page,omitempty"`     // The current page number.
	Pagelen  int                          `json:"pagelen,omitempty"`  // The total number of pages.
	Next     string                       `json:"next,omitempty"`     // The URL to the next page.
	Previous string                       `json:"previous,omitempty"` // The URL to the previous page.
	Values   []*WebhookSubscriptionScheme `json:"values,omitempty"`   // The webhook subscriptions in the current page.
}

// WebhookSubscriptionScheme represents a webhook subscription.
type WebhookSubscriptionScheme struct {
	UUID        string                            `json:"uuid,omitempty"`         // The unique identifier of the webhook subscription.
	URL         string                            `json:"url,omitempty"`          // The URL of the webhook subscription.
	Description string                            `json:"description,omitempty"`  // The description of the webhook subscription.
	SubjectType string                            `json:"subject_type,omitempty"` // The type of the subject of the webhook subscription.
	Subject     *WebhookSubscriptionSubjectScheme `json:"subject,omitempty"`      // The subject of the webhook subscription.
	Active      bool                              `json:"active,omitempty"`       // Indicates if the webhook subscription is active.
	CreatedAt   string                            `json:"created_at,omitempty"`   // The creation time of the webhook subscription.
	Events      []string                          `json:"events,omitempty"`       // The events for the webhook subscription.
}

// WebhookSubscriptionSubjectScheme represents the subject of a webhook subscription.
type WebhookSubscriptionSubjectScheme struct {
	Type string `json:"type,omitempty"` // The type of the subject.
}
