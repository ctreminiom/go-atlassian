package models

type BitbucketProjectGroupPermissionPageScheme struct {
	Size     int                             `json:"size"`
	Page     int                             `json:"page"`
	Pagelen  int                             `json:"pagelen"`
	Next     string                          `json:"next"`
	Previous string                          `json:"previous"`
	Values   []*ProjectGroupPermissionScheme `json:"values"`
}

type ProjectGroupPermissionScheme struct {
	Type       string                             `json:"type"`
	Links      *ProjectGroupPermissionLinksScheme `json:"links"`
	Permission string                             `json:"permission"`
	Group      *BitbucketGroupScheme              `json:"group"`
	Project    *BitbucketProjectScheme            `json:"project"`
}

type ProjectGroupPermissionLinksScheme struct {
	Self *BitbucketLinkScheme `json:"self,omitempty"`
}
