package models

// BitbucketGroupScheme represents a Bitbucket group.
type BitbucketGroupScheme struct {
	Links     *BitbucketGroupLinksScheme `json:"links"`
	Owner     *BitbucketAccountScheme    `json:"owner"`
	Workspace *BitbucketWorkspaceScheme  `json:"workspace"`
	Name      string                     `json:"name"`
	Slug      string                     `json:"slug"`
	FullSlug  string                     `json:"full_slug"`
}

// BitbucketGroupLinksScheme represents a collection of links related to a Bitbucket group.
type BitbucketGroupLinksScheme struct {
	Self *BitbucketLinkScheme `json:"self,omitempty"`
	HTML *BitbucketLinkScheme `json:"html,omitempty"`
}
