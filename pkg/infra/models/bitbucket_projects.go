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

type BitbucketProjectPayloadScheme struct {
	Type                    string                              `json:"type,omitempty"`
	Links                   *BitbucketProjectLinksScheme        `json:"links,omitempty"`
	Uuid                    string                              `json:"uuid,omitempty"`
	Key                     string                              `json:"key,omitempty"`
	Owner                   *BitbucketProjectPayloadOwnerScheme `json:"owner,omitempty"`
	Name                    string                              `json:"name,omitempty"`
	Description             string                              `json:"description,omitempty"`
	IsPrivate               bool                                `json:"is_private,omitempty"`
	CreatedOn               string                              `json:"created_on,omitempty"`
	UpdatedOn               string                              `json:"updated_on,omitempty"`
	HasPubliclyVisibleRepos bool                                `json:"has_publicly_visible_repos,omitempty"`
}

type BitbucketProjectPayloadOwnerScheme struct {
	Type string `json:"type,omitempty"`
}

type ProjectReviewersPageScheme struct {
	Pagelen int                               `json:"pagelen,omitempty"`
	Values  []*ProjectReviewersUserPageScheme `json:"values,omitempty"`
	Page    int                               `json:"page,omitempty"`
	Size    int                               `json:"size,omitempty"`
}

type ProjectReviewersUserPageScheme struct {
	User         *ProjectReviewerScheme `json:"user,omitempty"`
	ReviewerType string                 `json:"reviewer_type,omitempty"`
	Type         string                 `json:"type,omitempty"`
}

type ProjectReviewerScheme struct {
	DisplayName string `json:"display_name,omitempty"`
	Type        string `json:"type,omitempty"`
	Uuid        string `json:"uuid,omitempty"`
	AccountId   string `json:"account_id,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
}
