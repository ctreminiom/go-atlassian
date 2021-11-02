package jira

type ProjectScheme struct {
	Expand            string                 `json:"expand,omitempty"`
	Self              string                 `json:"self,omitempty"`
	ID                string                 `json:"id,omitempty"`
	Key               string                 `json:"key,omitempty"`
	Description       string                 `json:"description,omitempty"`
	URL               string                 `json:"url,omitempty"`
	Email             string                 `json:"email,omitempty"`
	AssigneeType      string                 `json:"assigneeType,omitempty"`
	Name              string                 `json:"name,omitempty"`
	ProjectTypeKey    string                 `json:"projectTypeKey,omitempty"`
	Simplified        bool                   `json:"simplified,omitempty"`
	Style             string                 `json:"style,omitempty"`
	Favourite         bool                   `json:"favourite,omitempty"`
	IsPrivate         bool                   `json:"isPrivate,omitempty"`
	UUID              string                 `json:"uuid,omitempty"`
	Lead              *UserScheme            `json:"lead,omitempty"`
	Components        []*ComponentScheme     `json:"components,omitempty"`
	IssueTypes        []*IssueTypeScheme     `json:"issueTypes,omitempty"`
	Versions          []*VersionScheme       `json:"versions,omitempty"`
	Roles             *ProjectRolesScheme    `json:"roles,omitempty"`
	AvatarUrls        *AvatarURLScheme       `json:"avatarUrls,omitempty"`
	ProjectKeys       []string               `json:"projectKeys,omitempty"`
	Insight           *ProjectInsightScheme  `json:"insight,omitempty"`
	Category          *ProjectCategoryScheme `json:"projectCategory,omitempty"`
	Deleted           bool                   `json:"deleted,omitempty"`
	RetentionTillDate string                 `json:"retentionTillDate,omitempty"`
	DeletedDate       string                 `json:"deletedDate,omitempty"`
	DeletedBy         *UserScheme            `json:"deletedBy,omitempty"`
	Archived          bool                   `json:"archived,omitempty"`
	ArchivedDate      string                 `json:"archivedDate,omitempty"`
	ArchivedBy        *UserScheme            `json:"archivedBy,omitempty"`
}

type ProjectInsightScheme struct {
	TotalIssueCount     int    `json:"totalIssueCount,omitempty"`
	LastIssueUpdateTime string `json:"lastIssueUpdateTime,omitempty"`
}

type ProjectCategoryScheme struct {
	Self        string `json:"self,omitempty"`
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type TeamManagedProjectScopeScheme struct {
	Type    string         `json:"type,omitempty"`
	Project *ProjectScheme `json:"project,omitempty"`
}
