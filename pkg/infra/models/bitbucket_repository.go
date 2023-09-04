package models

type RepositoryPermissionPageScheme struct {
	Size     int                           `json:"size,omitempty"`
	Page     int                           `json:"page,omitempty"`
	Pagelen  int                           `json:"pagelen,omitempty"`
	Next     string                        `json:"next,omitempty"`
	Previous string                        `json:"previous,omitempty"`
	Values   []*RepositoryPermissionScheme `json:"values,omitempty"`
}

type RepositoryPermissionScheme struct {
	Type       string                  `json:"type,omitempty"`
	Permission string                  `json:"permission,omitempty"`
	User       *BitbucketAccountScheme `json:"user,omitempty"`
	Repository *RepositoryScheme       `json:"repository,omitempty"`
}

type RepositoryScheme struct {
	Type        string                  `json:"type,omitempty"`
	Uuid        string                  `json:"uuid,omitempty"`
	FullName    string                  `json:"full_name,omitempty"`
	IsPrivate   bool                    `json:"is_private,omitempty"`
	Scm         string                  `json:"scm,omitempty"`
	Name        string                  `json:"name,omitempty"`
	Description string                  `json:"description,omitempty"`
	CreatedOn   string                  `json:"created_on,omitempty"`
	UpdatedOn   string                  `json:"updated_on,omitempty"`
	Size        int                     `json:"size,omitempty"`
	Language    string                  `json:"language,omitempty"`
	HasIssues   bool                    `json:"has_issues,omitempty"`
	HasWiki     bool                    `json:"has_wiki,omitempty"`
	ForkPolicy  string                  `json:"fork_policy,omitempty"`
	Owner       *BitbucketAccountScheme `json:"owner,omitempty"`
	Parent      *RepositoryScheme       `json:"parent,omitempty"`
	Project     BitbucketProjectScheme  `json:"project,omitempty"`
	Links       *RepositoryLinksScheme  `json:"links,omitempty"`
}

type RepositoryLinksScheme struct {
	Self         *BitbucketLinkScheme   `json:"self,omitempty"`
	Html         *BitbucketLinkScheme   `json:"html,omitempty"`
	Avatar       *BitbucketLinkScheme   `json:"avatar,omitempty"`
	PullRequests *BitbucketLinkScheme   `json:"pullrequests,omitempty"`
	Commits      *BitbucketLinkScheme   `json:"commits,omitempty"`
	Forks        *BitbucketLinkScheme   `json:"forks,omitempty"`
	Watchers     *BitbucketLinkScheme   `json:"watchers,omitempty"`
	Downloads    *BitbucketLinkScheme   `json:"downloads,omitempty"`
	Clone        []*BitbucketLinkScheme `json:"clone,omitempty"`
	Hooks        *BitbucketLinkScheme   `json:"hooks,omitempty"`
}
