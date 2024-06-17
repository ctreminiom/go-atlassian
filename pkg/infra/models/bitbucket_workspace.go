package models

// WorkspaceScheme represents a workspace.
// Type is the type of the workspace.
// Links is a collection of links related to the workspace.
// Uuid is the unique identifier of the workspace.
// Name is the name of the workspace.
// Slug is the slug of the workspace.
// IsPrivate indicates if the workspace is private.
// CreatedOn is the creation time of the workspace.
// UpdatedOn is the update time of the workspace.
type WorkspaceScheme struct {
	Type      string                `json:"type,omitempty"`       // The type of the workspace.
	Links     *WorkspaceLinksScheme `json:"links,omitempty"`      // The links related to the workspace.
	Uuid      string                `json:"uuid,omitempty"`       // The unique identifier of the workspace.
	Name      string                `json:"name,omitempty"`       // The name of the workspace.
	Slug      string                `json:"slug,omitempty"`       // The slug of the workspace.
	IsPrivate bool                  `json:"is_private,omitempty"` // Indicates if the workspace is private.
	CreatedOn string                `json:"created_on,omitempty"` // The creation time of the workspace.
	UpdatedOn string                `json:"updated_on,omitempty"` // The update time of the workspace.
}

// WorkspaceLinksScheme represents a collection of links related to a workspace.
// Avatar is the link to the workspace's avatar.
// Html is the link to the workspace's HTML page.
// Members is the link to the workspace's members.
// Owners is the link to the workspace's owners.
// Projects is the link to the workspace's projects.
// Repositories is the link to the workspace's repositories.
// Snippets is the link to the workspace's snippets.
// Self is the link to the workspace itself.
type WorkspaceLinksScheme struct {
	Avatar       *BitbucketLinkScheme `json:"avatar,omitempty"`       // The link to the workspace's avatar.
	Html         *BitbucketLinkScheme `json:"html,omitempty"`         // The link to the workspace's HTML page.
	Members      *BitbucketLinkScheme `json:"members,omitempty"`      // The link to the workspace's members.
	Owners       *BitbucketLinkScheme `json:"owners,omitempty"`       // The link to the workspace's owners.
	Projects     *BitbucketLinkScheme `json:"projects,omitempty"`     // The link to the workspace's projects.
	Repositories *BitbucketLinkScheme `json:"repositories,omitempty"` // The link to the workspace's repositories.
	Snippets     *BitbucketLinkScheme `json:"snippets,omitempty"`     // The link to the workspace's snippets.
	Self         *BitbucketLinkScheme `json:"self,omitempty"`         // The link to the workspace itself.
}

// BitbucketLinkScheme represents a link in Bitbucket.
// Href is the URL of the link.
// Name is the name of the link.
type BitbucketLinkScheme struct {
	Href string `json:"href,omitempty"` // The URL of the link.
	Name string `json:"name,omitempty"` // The name of the link.
}
