package models

type IssueChangelogScheme struct {
	StartAt    int                            `json:"startAt,omitempty"`
	MaxResults int                            `json:"maxResults,omitempty"`
	Total      int                            `json:"total,omitempty"`
	Histories  []*IssueChangelogHistoryScheme `json:"histories,omitempty"`
}

type IssueChangelogHistoryScheme struct {
	ID      string                             `json:"id,omitempty"`
	Author  *IssueChangelogAuthor              `json:"author,omitempty"`
	Created string                             `json:"created,omitempty"`
	Items   []*IssueChangelogHistoryItemScheme `json:"items,omitempty"`
}

type IssueChangelogAuthor struct {
	Self         string           `json:"self,omitempty"`
	AccountID    string           `json:"accountId,omitempty"`
	EmailAddress string           `json:"emailAddress,omitempty"`
	AvatarUrls   *AvatarURLScheme `json:"avatarUrls,omitempty"`
	DisplayName  string           `json:"displayName,omitempty"`
	Active       bool             `json:"active,omitempty"`
	TimeZone     string           `json:"timeZone,omitempty"`
	AccountType  string           `json:"accountType,omitempty"`
}

type IssueChangelogHistoryItemScheme struct {
	Field      string `json:"field,omitempty"`
	Fieldtype  string `json:"fieldtype,omitempty"`
	FieldID    string `json:"fieldId,omitempty"`
	From       string `json:"from,omitempty"`
	FromString string `json:"fromString,omitempty"`
	To         string `json:"to,omitempty"`
	ToString   string `json:"toString,omitempty"`
}
