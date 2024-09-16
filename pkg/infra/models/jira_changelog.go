package models

// IssueChangelogScheme represents the changelog of an issue in Jira.
type IssueChangelogScheme struct {
	StartAt    int                            `json:"startAt,omitempty"`    // The starting index of the changelog.
	MaxResults int                            `json:"maxResults,omitempty"` // The maximum number of results in the changelog.
	Total      int                            `json:"total,omitempty"`      // The total number of changes in the changelog.
	Histories  []*IssueChangelogHistoryScheme `json:"histories,omitempty"`  // The history of changes in the changelog.
}

// IssueChangelogHistoryScheme represents a history of changes in an issue's changelog in Jira.
type IssueChangelogHistoryScheme struct {
	ID      string                             `json:"id,omitempty"`      // The ID of the history.
	Author  *IssueChangelogAuthor              `json:"author,omitempty"`  // The author of the history.
	Created string                             `json:"created,omitempty"` // The creation time of the history.
	Items   []*IssueChangelogHistoryItemScheme `json:"items,omitempty"`   // The items in the history.
}

// IssueChangelogAuthor represents the author of a history in an issue's changelog in Jira.
type IssueChangelogAuthor struct {
	Self         string           `json:"self,omitempty"`         // The URL of the author's profile.
	AccountID    string           `json:"accountId,omitempty"`    // The account ID of the author.
	EmailAddress string           `json:"emailAddress,omitempty"` // The email address of the author.
	AvatarURLs   *AvatarURLScheme `json:"avatarUrls,omitempty"`   // The URLs for different sizes of the author's avatar.
	DisplayName  string           `json:"displayName,omitempty"`  // The display name of the author.
	Active       bool             `json:"active,omitempty"`       // Indicates if the author's account is active.
	TimeZone     string           `json:"timeZone,omitempty"`     // The time zone of the author.
	AccountType  string           `json:"accountType,omitempty"`  // The type of the author's account.
}

// IssueChangelogHistoryItemScheme represents an item in a history of an issue's changelog in Jira.
type IssueChangelogHistoryItemScheme struct {
	Field      string `json:"field,omitempty"`      // The field of the item.
	Fieldtype  string `json:"fieldtype,omitempty"`  // The type of the field.
	FieldID    string `json:"fieldId,omitempty"`    // The ID of the field.
	From       string `json:"from,omitempty"`       // The previous value of the field.
	FromString string `json:"fromString,omitempty"` // The previous value of the field as a string.
	To         string `json:"to,omitempty"`         // The new value of the field.
	ToString   string `json:"toString,omitempty"`   // The new value of the field as a string.
}
