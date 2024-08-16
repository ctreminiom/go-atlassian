package models

// WorkspaceScheme represents a workspace.
type WorkspaceScheme struct {
	Type      string                `json:"type,omitempty"`       // The type of the workspace.
	Links     *WorkspaceLinksScheme `json:"links,omitempty"`      // The links related to the workspace.
	UUID      string                `json:"uuid,omitempty"`       // The unique identifier of the workspace.
	Name      string                `json:"name,omitempty"`       // The name of the workspace.
	Slug      string                `json:"slug,omitempty"`       // The slug of the workspace.
	IsPrivate bool                  `json:"is_private,omitempty"` // Indicates if the workspace is private.
	CreatedOn string                `json:"created_on,omitempty"` // The creation time of the workspace.
	UpdatedOn string                `json:"updated_on,omitempty"` // The update time of the workspace.
}

// WorkspaceLinksScheme represents a collection of links related to a workspace.
type WorkspaceLinksScheme struct {
	Avatar       *BitbucketLinkScheme `json:"avatar,omitempty"`       // The link to the workspace's avatar.
	HTML         *BitbucketLinkScheme `json:"html,omitempty"`         // The link to the workspace's HTML page.
	Members      *BitbucketLinkScheme `json:"members,omitempty"`      // The link to the workspace's members.
	Owners       *BitbucketLinkScheme `json:"owners,omitempty"`       // The link to the workspace's owners.
	Projects     *BitbucketLinkScheme `json:"projects,omitempty"`     // The link to the workspace's projects.
	Repositories *BitbucketLinkScheme `json:"repositories,omitempty"` // The link to the workspace's repositories.
	Snippets     *BitbucketLinkScheme `json:"snippets,omitempty"`     // The link to the workspace's snippets.
	Self         *BitbucketLinkScheme `json:"self,omitempty"`         // The link to the workspace itself.
}

// BitbucketLinkScheme represents a link in Bitbucket.
type BitbucketLinkScheme struct {
	Href string `json:"href,omitempty"` // The URL of the link.
	Name string `json:"name,omitempty"` // The name of the link.
}
