package models

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
	NotificationScheme       int    `json:"notificationScheme,omitempty"`
	FieldConfigurationScheme int    `json:"fieldConfigurationScheme,omitempty"`
	IssueSecurityScheme      int    `json:"issueSecurityScheme,omitempty"`
	PermissionScheme         int    `json:"permissionScheme,omitempty"`
	IssueTypeScheme          int    `json:"issueTypeScheme,omitempty"`
	IssueTypeScreenScheme    int    `json:"issueTypeScreenScheme,omitempty"`
	WorkflowScheme           int    `json:"workflowScheme,omitempty"`
	Description              string `json:"description,omitempty"`
	LeadAccountID            string `json:"leadAccountId,omitempty"`
	URL                      string `json:"url,omitempty"`
	ProjectTemplateKey       string `json:"projectTemplateKey,omitempty"`
	AvatarID                 int    `json:"avatarId,omitempty"`
	Name                     string `json:"name,omitempty"`
	AssigneeType             string `json:"assigneeType,omitempty"`
	ProjectTypeKey           string `json:"projectTypeKey,omitempty"`
	Key                      string `json:"key,omitempty"`
	CategoryID               int    `json:"categoryId,omitempty"`
}

type NewProjectCreatedScheme struct {
	Self string `json:"self"`
	ID   int    `json:"id"`
	Key  string `json:"key"`
}

type ProjectSearchOptionsScheme struct {
	OrderBy string

	// The project IDs to filter the results by.
	// To include multiple IDs, provide an ampersand-separated list.
	IDs []int

	// The project keys to filter the results by.
	// To include multiple keys, provide an ampersand-separated list.
	Keys []string

	// Filter the results using a literal string.
	// Projects with a matching key or name are returned (case-insensitive).
	Query string

	// Orders results by the project type.
	// This parameter accepts a comma-separated list.
	// Valid values are business, service_desk, and software.
	TypeKeys []string

	// The ID of the project's category.
	// A complete list of category IDs is found using the Get all project categories operation.
	CategoryID int

	Action string

	// EXPERIMENTAL. Filter results by project status:
	// 1. live: Search live projects.
	// 2. archived: Search archived projects.
	// 3. deleted: Search deleted projects, those in the recycle bin.
	Status []string

	// Use expand to include additional information in the response.
	// This parameter accepts a comma-separated list.
	Expand []string

	// EXPERIMENTAL. A list of project properties to return for the project.
	// This parameter accepts a comma-separated list.
	Properties []string

	// EXPERIMENTAL. A query string used to search properties.
	// The query string cannot be specified using a JSON object.
	// For example, to search for the value of nested from {"something":{"nested":1,"other":2}}
	// use [thepropertykey].something.nested=1.
	// Note that thepropertykey is only returned when included in properties.
	PropertyQuery string
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

type ProjectIdentifierScheme struct {
	ID  int    `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
}
