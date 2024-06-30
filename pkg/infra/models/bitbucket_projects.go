// Package models provides the data structures used in the Bitbucket integration.
package models

// BitbucketProjectPageScheme represents a page of Bitbucket projects.
type BitbucketProjectPageScheme struct {
	Size     int                       `json:"size,omitempty"`     // The size of the page.
	Page     int                       `json:"page,omitempty"`     // The current page number.
	Pagelen  int                       `json:"pagelen,omitempty"`  // The length of the page.
	Next     string                    `json:"next,omitempty"`     // The link to the next page.
	Previous string                    `json:"previous,omitempty"` // The link to the previous page.
	Values   []*BitbucketProjectScheme `json:"values,omitempty"`   // The projects on the page.
}

// BitbucketProjectScheme represents a Bitbucket project.
type BitbucketProjectScheme struct {
	Links                   *BitbucketProjectLinksScheme `json:"links,omitempty"`                      // The links related to the project.
	UUID                    string                       `json:"uuid,omitempty"`                       // The UUID of the project.
	Key                     string                       `json:"key,omitempty"`                        // The key of the project.
	Name                    string                       `json:"name,omitempty"`                       // The name of the project.
	Description             string                       `json:"description,omitempty"`                // The description of the project.
	IsPrivate               bool                         `json:"is_private,omitempty"`                 // Whether the project is private.
	CreatedOn               string                       `json:"created_on,omitempty"`                 // The creation time of the project.
	UpdatedOn               string                       `json:"updated_on,omitempty"`                 // The last update time of the project.
	HasPubliclyVisibleRepos bool                         `json:"has_publicly_visible_repos,omitempty"` // Whether the project has publicly visible repositories.
}

// BitbucketProjectLinksScheme represents the links related to a Bitbucket project.
type BitbucketProjectLinksScheme struct {
	HTML   *BitbucketLinkScheme `json:"html,omitempty"`   // The HTML link of the project.
	Avatar *BitbucketLinkScheme `json:"avatar,omitempty"` // The avatar link of the project.
}
