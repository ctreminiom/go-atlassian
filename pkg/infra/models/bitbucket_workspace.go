package models

type WorkspaceScheme struct {
	Type      string                `json:"type,omitempty"`
	Links     *WorkspaceLinksScheme `json:"links,omitempty"`
	Uuid      string                `json:"uuid,omitempty"`
	Name      string                `json:"name,omitempty"`
	Slug      string                `json:"slug,omitempty"`
	IsPrivate bool                  `json:"is_private,omitempty"`
	CreatedOn string                `json:"created_on,omitempty"`
	UpdatedOn string                `json:"updated_on,omitempty"`
}

type WorkspaceLinksScheme struct {
	Avatar       *BitbucketLinkScheme `json:"avatar,omitempty"`
	Html         *BitbucketLinkScheme `json:"html,omitempty"`
	Members      *BitbucketLinkScheme `json:"members,omitempty"`
	Owners       *BitbucketLinkScheme `json:"owners,omitempty"`
	Projects     *BitbucketLinkScheme `json:"projects,omitempty"`
	Repositories *BitbucketLinkScheme `json:"repositories,omitempty"`
	Snippets     *BitbucketLinkScheme `json:"snippets,omitempty"`
	Self         *BitbucketLinkScheme `json:"self,omitempty"`
}

type BitbucketLinkScheme struct {
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
}
