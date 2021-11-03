package jira

const (
	BusinessContentManagementProjectTemplate    = "com.atlassian.jira-core-project-templates:jira-core-simplified-content-management"
	BusinessDocumentApprovalProjectTemplate     = "com.atlassian.jira-core-project-templates:jira-core-simplified-document-approval"
	BusinessLeadTrackingProjectTemplate         = "com.atlassian.jira-core-project-templates:jira-core-simplified-lead-tracking"
	BusinessProcessControlProjectTemplate       = "com.atlassian.jira-core-project-templates:jira-core-simplified-process-control"
	BusinessProcurementProjectTemplate          = "com.atlassian.jira-core-project-templates:jira-core-simplified-procurement"
	BusinessProjectManagementProjectTemplate    = "com.atlassian.jira-core-project-templates:jira-core-simplified-project-management"
	BusinessRecruitmentProjectTemplate          = "com.atlassian.jira-core-project-templates:jira-core-simplified-recruitment"
	BusinessTaskTrackingProjectTemplate         = "com.atlassian.jira-core-project-templates:jira-core-simplified-task-tracking"
	ITSMServiceDeskProjectTemplate              = "com.atlassian.servicedesk:simplified-it-service-desk"
	ITSMInternalServiceDeskProjectTemplate      = "com.atlassian.servicedesk:simplified-internal-service-desk"
	ITSMExternalServiceDeskProjectTemplate      = "com.atlassian.servicedesk:simplified-external-service-desk"
	SoftwareTeamManagedKanbanProjectTemplate    = "com.pyxis.greenhopper.jira:gh-simplified-agility-kanban"
	SoftwareTeamManagedScrumProjectTemplate     = "com.pyxis.greenhopper.jira:gh-simplified-agility-scrum"
	SoftwareCompanyManagedKanbanProjectTemplate = "com.pyxis.greenhopper.jira:gh-simplified-kanban-classic"
	SoftwareCompanyManagedScrumProjectTemplate  = "com.pyxis.greenhopper.jira:gh-simplified-scrum-classic"
)

type ProjectPayloadScheme struct {
	NotificationScheme  int    `json:"notificationScheme"`
	Description         string `json:"description"`
	LeadAccountID       string `json:"leadAccountId"`
	URL                 string `json:"url"`
	ProjectTemplateKey  string `json:"projectTemplateKey"`
	AvatarID            int    `json:"avatarId"`
	IssueSecurityScheme int    `json:"issueSecurityScheme"`
	Name                string `json:"name"`
	PermissionScheme    int    `json:"permissionScheme"`
	AssigneeType        string `json:"assigneeType"`
	ProjectTypeKey      string `json:"projectTypeKey"`
	Key                 string `json:"key"`
	CategoryID          int    `json:"categoryId"`
}

type NewProjectCreatedScheme struct {
	Self string `json:"self"`
	ID   int    `json:"id"`
	Key  string `json:"key"`
}

type ProjectSearchOptionsScheme struct {
	OrderBy        string
	Query          string
	Action         string
	ProjectKeyType string
	CategoryID     int
	Expand         []string
}

type ProjectSearchScheme struct {
	Self       string           `json:"self,omitempty"`
	NextPage   string           `json:"nextPage,omitempty"`
	MaxResults int              `json:"maxResults,omitempty"`
	StartAt    int              `json:"startAt,omitempty"`
	Total      int              `json:"total,omitempty"`
	IsLast     bool             `json:"isLast,omitempty"`
	Values     []*ProjectScheme `json:"values,omitempty"`
}

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

type ProjectUpdateScheme struct {
	NotificationScheme  int    `json:"notificationScheme,omitempty"`
	Description         string `json:"description,omitempty"`
	Lead                string `json:"lead,omitempty"`
	URL                 string `json:"url,omitempty"`
	ProjectTemplateKey  string `json:"projectTemplateKey,omitempty"`
	AvatarID            int    `json:"avatarId,omitempty"`
	IssueSecurityScheme int    `json:"issueSecurityScheme,omitempty"`
	Name                string `json:"name,omitempty"`
	PermissionScheme    int    `json:"permissionScheme,omitempty"`
	AssigneeType        string `json:"assigneeType,omitempty"`
	ProjectTypeKey      string `json:"projectTypeKey,omitempty"`
	Key                 string `json:"key,omitempty"`
	CategoryID          int    `json:"categoryId,omitempty"`
}

type ProjectStatusPageScheme struct {
	Self     string                        `json:"self,omitempty"`
	ID       string                        `json:"id,omitempty"`
	Name     string                        `json:"name,omitempty"`
	Subtask  bool                          `json:"subtask,omitempty"`
	Statuses []*ProjectStatusDetailsScheme `json:"statuses,omitempty"`
}

type ProjectStatusDetailsScheme struct {
	Self           string                `json:"self,omitempty"`
	Description    string                `json:"description,omitempty"`
	IconURL        string                `json:"iconUrl,omitempty"`
	Name           string                `json:"name,omitempty"`
	ID             string                `json:"id,omitempty"`
	StatusCategory *StatusCategoryScheme `json:"statusCategory,omitempty"`
}

type ProjectHierarchyScheme struct {
	EntityID   string                             `json:"entityId,omitempty"`
	Level      int                                `json:"level,omitempty"`
	Name       string                             `json:"name,omitempty"`
	IssueTypes []*ProjectHierarchyIssueTypeScheme `json:"issueTypes,omitempty"`
}

type ProjectHierarchyIssueTypeScheme struct {
	ID       int    `json:"id,omitempty"`
	EntityID string `json:"entityId,omitempty"`
	Name     string `json:"name,omitempty"`
	AvatarID int    `json:"avatarId,omitempty"`
}
