package models

type BitbucketProjectPageScheme struct {
	Size     int                       `json:"size,omitempty"`
	Page     int                       `json:"page,omitempty"`
	Pagelen  int                       `json:"pagelen,omitempty"`
	Next     string                    `json:"next,omitempty"`
	Previous string                    `json:"previous,omitempty"`
	Values   []*BitbucketProjectScheme `json:"values,omitempty"`
}

type BitbucketProjectScheme struct {
	Links                   *BitbucketProjectLinksScheme `json:"links,omitempty"`
	Uuid                    string                       `json:"uuid,omitempty"`
	Key                     string                       `json:"key,omitempty"`
	Name                    string                       `json:"name,omitempty"`
	Description             string                       `json:"description,omitempty"`
	IsPrivate               bool                         `json:"is_private,omitempty"`
	CreatedOn               string                       `json:"created_on,omitempty"`
	UpdatedOn               string                       `json:"updated_on,omitempty"`
	HasPubliclyVisibleRepos bool                         `json:"has_publicly_visible_repos,omitempty"`
}

type BitbucketProjectLinksScheme struct {
	Html   *BitbucketLinkScheme `json:"html,omitempty"`
	Avatar *BitbucketLinkScheme `json:"avatar,omitempty"`
}
