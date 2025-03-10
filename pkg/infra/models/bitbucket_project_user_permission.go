package models

type BitbucketProjectUserPermissionPageScheme struct {
	Size     int                                     `json:"size"`
	Page     int                                     `json:"page"`
	Pagelen  int                                     `json:"pagelen"`
	Next     string                                  `json:"next"`
	Previous string                                  `json:"previous"`
	Values   []*BitbucketProjectUserPermissionScheme `json:"values"`
}

type BitbucketProjectUserPermissionScheme struct {
	Type       string                                     `json:"type"`
	Links      *BitbucketProjectUserPermissionLinksScheme `json:"links"`
	Permission string                                     `json:"permission"`
	User       *BitbucketAccountScheme                    `json:"user"`
	Project    *BitbucketProjectScheme                    `json:"project"`
}

type BitbucketProjectUserPermissionLinksScheme struct {
	Self *BitbucketLinkScheme `json:"self,omitempty"`
}
