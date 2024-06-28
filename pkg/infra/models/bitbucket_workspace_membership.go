package models

import "time"

// WorkspaceMembershipPageScheme represents a paginated list of workspace memberships.
// Size is the number of memberships in the current page.
// Page is the current page number.
// Pagelen is the total number of pages.
// Next is the URL to the next page.
// Previous is the URL to the previous page.
// Values is a slice of the workspace memberships in the current page.
type WorkspaceMembershipPageScheme struct {
	Size     int                          `json:"size,omitempty"`     // The number of memberships in the current page.
	Page     int                          `json:"page,omitempty"`     // The current page number.
	Pagelen  int                          `json:"pagelen,omitempty"`  // The total number of pages.
	Next     string                       `json:"next,omitempty"`     // The URL to the next page.
	Previous string                       `json:"previous,omitempty"` // The URL to the previous page.
	Values   []*WorkspaceMembershipScheme `json:"values,omitempty"`   // The workspace memberships in the current page.
}

// WorkspaceMembershipScheme represents a workspace membership.
// Links is a collection of links related to the membership.
// User is the user who has the membership.
// Workspace is the workspace to which the membership applies.
// AddedOn is the time when the membership was added.
// Permission is the level of the membership.
// LastAccessed is the last time the membership was accessed.
type WorkspaceMembershipScheme struct {
	Links        *WorkspaceMembershipLinksScheme `json:"links,omitempty"`         // The links related to the membership.
	User         *BitbucketAccountScheme         `json:"user,omitempty"`          // The user who has the membership.
	Workspace    *WorkspaceScheme                `json:"workspace,omitempty"`     // The workspace to which the membership applies.
	AddedOn      time.Time                       `json:"added_on,omitempty"`      // The time when the membership was added.
	Permission   string                          `json:"permission,omitempty"`    // The level of the membership.
	LastAccessed time.Time                       `json:"last_accessed,omitempty"` // The last time the membership was accessed.
}

// WorkspaceMembershipLinksScheme represents a collection of links related to a workspace membership.
// Self is the link to the membership itself.
type WorkspaceMembershipLinksScheme struct {
	Self *BitbucketLinkScheme `json:"self,omitempty"` // The link to the membership itself.
}

// BitbucketAccountScheme represents a Bitbucket account.
// Links is a collection of links related to the account.
// CreatedOn is the creation time of the account.
// DisplayName is the display name of the account.
// Username is the username of the account.
// UUID is the unique identifier of the account.
// Type is the type of the account.
// AccountID is the account ID of the account.
// Nickname is the nickname of the account.
type BitbucketAccountScheme struct {
	Links       *BitbucketAccountLinksScheme `json:"links,omitempty"`        // The links related to the account.
	CreatedOn   string                       `json:"created_on,omitempty"`   // The creation time of the account.
	DisplayName string                       `json:"display_name,omitempty"` // The display name of the account.
	Username    string                       `json:"username,omitempty"`     // The username of the account.
	UUID        string                       `json:"uuid,omitempty"`         // The unique identifier of the account.
	Type        string                       `json:"type,omitempty"`         // The type of the account.
	AccountID   string                       `json:"account_id,omitempty"`   // The account ID of the account.
	Nickname    string                       `json:"nickname,omitempty"`     // The nickname of the account.
}

// BitbucketAccountLinksScheme represents a collection of links related to a Bitbucket account.
// Avatar is the link to the account's avatar.
// Self is the link to the account itself.
// HTML is the link to the account's HTML page.
type BitbucketAccountLinksScheme struct {
	Avatar *BitbucketLinkScheme `json:"avatar,omitempty"` // The link to the account's avatar.
	Self   *BitbucketLinkScheme `json:"self,omitempty"`   // The link to the account itself.
	HTML   *BitbucketLinkScheme `json:"html,omitempty"`   // The link to the account's HTML page.
}
