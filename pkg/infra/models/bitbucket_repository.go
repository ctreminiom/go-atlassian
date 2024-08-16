package models

// RepositoryPermissionPageScheme represents a paginated list of repository permissions.
type RepositoryPermissionPageScheme struct {
	Size     int                           `json:"size,omitempty"`     // The number of permissions in the current page.
	Page     int                           `json:"page,omitempty"`     // The current page number.
	Pagelen  int                           `json:"pagelen,omitempty"`  // The total number of pages.
	Next     string                        `json:"next,omitempty"`     // The URL to the next page.
	Previous string                        `json:"previous,omitempty"` // The URL to the previous page.
	Values   []*RepositoryPermissionScheme `json:"values,omitempty"`   // The repository permissions in the current page.
}

// RepositoryPermissionScheme represents a repository permission.
type RepositoryPermissionScheme struct {
	Type       string                  `json:"type,omitempty"`       // The type of the permission.
	Permission string                  `json:"permission,omitempty"` // The level of the permission.
	User       *BitbucketAccountScheme `json:"user,omitempty"`       // The user who has the permission.
	Repository *RepositoryScheme       `json:"repository,omitempty"` // The repository to which the permission applies.
}

// RepositoryScheme represents a repository.
type RepositoryScheme struct {
	Type        string                  `json:"type,omitempty"`        // The type of the repository.
	UUID        string                  `json:"uuid,omitempty"`        // The unique identifier of the repository.
	FullName    string                  `json:"full_name,omitempty"`   // The full name of the repository.
	IsPrivate   bool                    `json:"is_private,omitempty"`  // Indicates if the repository is private.
	SCM         string                  `json:"scm,omitempty"`         // The source control management system used by the repository.
	Name        string                  `json:"name,omitempty"`        // The name of the repository.
	Description string                  `json:"description,omitempty"` // The description of the repository.
	CreatedOn   string                  `json:"created_on,omitempty"`  // The creation time of the repository.
	UpdatedOn   string                  `json:"updated_on,omitempty"`  // The update time of the repository.
	Size        int                     `json:"size,omitempty"`        // The size of the repository.
	Language    string                  `json:"language,omitempty"`    // The programming language used in the repository.
	HasIssues   bool                    `json:"has_issues,omitempty"`  // Indicates if the repository has issues enabled.
	HasWiki     bool                    `json:"has_wiki,omitempty"`    // Indicates if the repository has a wiki enabled.
	ForkPolicy  string                  `json:"fork_policy,omitempty"` // The fork policy of the repository.
	Owner       *BitbucketAccountScheme `json:"owner,omitempty"`       // The owner of the repository.
	Parent      *RepositoryScheme       `json:"parent,omitempty"`      // The parent repository, if the repository is a fork.
	Project     BitbucketProjectScheme  `json:"project,omitempty"`     // The project to which the repository belongs.
	Links       *RepositoryLinksScheme  `json:"links,omitempty"`       // A collection of links related to the repository.
}

// RepositoryLinksScheme represents a collection of links related to a repository.
type RepositoryLinksScheme struct {
	Self         *BitbucketLinkScheme   `json:"self,omitempty"`         // The link to the repository itself.
	HTML         *BitbucketLinkScheme   `json:"html,omitempty"`         // The link to the repository's HTML page.
	Avatar       *BitbucketLinkScheme   `json:"avatar,omitempty"`       // The link to the repository's avatar.
	PullRequests *BitbucketLinkScheme   `json:"pullrequests,omitempty"` // The link to the repository's pull requests.
	Commits      *BitbucketLinkScheme   `json:"commits,omitempty"`      // The link to the repository's commits.
	Forks        *BitbucketLinkScheme   `json:"forks,omitempty"`        // The link to the repository's forks.
	Watchers     *BitbucketLinkScheme   `json:"watchers,omitempty"`     // The link to the repository's watchers.
	Downloads    *BitbucketLinkScheme   `json:"downloads,omitempty"`    // The link to the repository's downloads.
	Clone        []*BitbucketLinkScheme `json:"clone,omitempty"`        // The links to clone the repository.
	Hooks        *BitbucketLinkScheme   `json:"hooks,omitempty"`        // The link to the repository's hooks.
}
