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

// ProjectPayloadScheme represents the payload for a project in Jira.
type ProjectPayloadScheme struct {
	NotificationScheme       int    `json:"notificationScheme,omitempty"`       // The ID of the notification scheme for the project.
	FieldConfigurationScheme int    `json:"fieldConfigurationScheme,omitempty"` // The ID of the field configuration scheme for the project.
	IssueSecurityScheme      int    `json:"issueSecurityScheme,omitempty"`      // The ID of the issue security scheme for the project.
	PermissionScheme         int    `json:"permissionScheme,omitempty"`         // The ID of the permission scheme for the project.
	IssueTypeScheme          int    `json:"issueTypeScheme,omitempty"`          // The ID of the issue type scheme for the project.
	IssueTypeScreenScheme    int    `json:"issueTypeScreenScheme,omitempty"`    // The ID of the issue type screen scheme for the project.
	WorkflowScheme           int    `json:"workflowScheme,omitempty"`           // The ID of the workflow scheme for the project.
	Description              string `json:"description,omitempty"`              // The description of the project.
	LeadAccountID            string `json:"leadAccountId,omitempty"`            // The account ID of the lead for the project.
	URL                      string `json:"url,omitempty"`                      // The URL of the project.
	ProjectTemplateKey       string `json:"projectTemplateKey,omitempty"`       // The key of the project template for the project.
	AvatarID                 int    `json:"avatarId,omitempty"`                 // The ID of the avatar for the project.
	Name                     string `json:"name,omitempty"`                     // The name of the project.
	AssigneeType             string `json:"assigneeType,omitempty"`             // The type of assignee for the project.
	ProjectTypeKey           string `json:"projectTypeKey,omitempty"`           // The key of the project type for the project.
	Key                      string `json:"key,omitempty"`                      // The key of the project.
	CategoryID               int    `json:"categoryId,omitempty"`               // The ID of the category for the project.
}

// NewProjectCreatedScheme represents a newly created project in Jira.
type NewProjectCreatedScheme struct {
	Self string `json:"self"` // The URL of the newly created project.
	ID   int    `json:"id"`   // The ID of the newly created project.
	Key  string `json:"key"`  // The key of the newly created project.
}

// ProjectSearchOptionsScheme represents the search options for projects in Jira.
type ProjectSearchOptionsScheme struct {
	OrderBy string // The order by field for the search.

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

	Action string // The action to perform on the search results.

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

// ProjectSearchScheme represents the search results for projects in Jira.
type ProjectSearchScheme struct {
	Self       string           `json:"self,omitempty"`       // The URL of the search results.
	NextPage   string           `json:"nextPage,omitempty"`   // The URL of the next page of search results.
	MaxResults int              `json:"maxResults,omitempty"` // The maximum number of results per page.
	StartAt    int              `json:"startAt,omitempty"`    // The starting index of the results.
	Total      int              `json:"total,omitempty"`      // The total number of results.
	IsLast     bool             `json:"isLast,omitempty"`     // Indicates if this is the last page of results.
	Values     []*ProjectScheme `json:"values,omitempty"`     // The projects in the search results.
}

// ProjectScheme represents a project in Jira.
type ProjectScheme struct {
	Expand            string                 `json:"expand,omitempty"`            // The fields to expand in the project.
	Self              string                 `json:"self,omitempty"`              // The URL of the project.
	ID                string                 `json:"id,omitempty"`                // The ID of the project.
	Key               string                 `json:"key,omitempty"`               // The key of the project.
	Description       string                 `json:"description,omitempty"`       // The description of the project.
	URL               string                 `json:"url,omitempty"`               // The URL of the project.
	Email             string                 `json:"email,omitempty"`             // The email of the project.
	AssigneeType      string                 `json:"assigneeType,omitempty"`      // The type of assignee for the project.
	Name              string                 `json:"name,omitempty"`              // The name of the project.
	ProjectTypeKey    string                 `json:"projectTypeKey,omitempty"`    // The key of the project type for the project.
	Simplified        bool                   `json:"simplified,omitempty"`        // Indicates if the project is simplified.
	Style             string                 `json:"style,omitempty"`             // The style of the project.
	Favourite         bool                   `json:"favourite,omitempty"`         // Indicates if the project is a favourite.
	IsPrivate         bool                   `json:"isPrivate,omitempty"`         // Indicates if the project is private.
	UUID              string                 `json:"uuid,omitempty"`              // The UUID of the project.
	Lead              *UserScheme            `json:"lead,omitempty"`              // The lead of the project.
	Components        []*ComponentScheme     `json:"components,omitempty"`        // The components of the project.
	IssueTypes        []*IssueTypeScheme     `json:"issueTypes,omitempty"`        // The issue types of the project.
	Versions          []*VersionScheme       `json:"versions,omitempty"`          // The versions of the project.
	Roles             *ProjectRolesScheme    `json:"roles,omitempty"`             // The roles of the project.
	AvatarURLs        *AvatarURLScheme       `json:"avatarUrls,omitempty"`        // The avatar URLs of the project.
	ProjectKeys       []string               `json:"projectKeys,omitempty"`       // The keys of the project.
	Insight           *ProjectInsightScheme  `json:"insight,omitempty"`           // The insight of the project.
	Category          *ProjectCategoryScheme `json:"projectCategory,omitempty"`   // The category of the project.
	Deleted           bool                   `json:"deleted,omitempty"`           // Indicates if the project is deleted.
	RetentionTillDate *DateTimeScheme        `json:"retentionTillDate,omitempty"` // The retention till date of the project.
	DeletedDate       *DateTimeScheme        `json:"deletedDate,omitempty"`       // The date the project was deleted.
	DeletedBy         *UserScheme            `json:"deletedBy,omitempty"`         // The user who deleted the project.
	Archived          bool                   `json:"archived,omitempty"`          // Indicates if the project is archived.
	ArchivedDate      *DateTimeScheme        `json:"archivedDate,omitempty"`      // The date the project was archived.
	ArchivedBy        *UserScheme            `json:"archivedBy,omitempty"`        // The user who archived the project.
}

// ProjectInsightScheme represents the insight of a project in Jira.
type ProjectInsightScheme struct {
	TotalIssueCount     int    `json:"totalIssueCount,omitempty"`     // The total number of issues in the project.
	LastIssueUpdateTime string `json:"lastIssueUpdateTime,omitempty"` // The last time an issue was updated in the project.
}

// ProjectCategoryScheme represents a category of a project in Jira.
type ProjectCategoryScheme struct {
	Self        string `json:"self,omitempty"`        // The URL of the category.
	ID          string `json:"id,omitempty"`          // The ID of the category.
	Name        string `json:"name,omitempty"`        // The name of the category.
	Description string `json:"description,omitempty"` // The description of the category.
}

// TeamManagedProjectScopeScheme represents the scope of a team-managed project in Jira.
type TeamManagedProjectScopeScheme struct {
	Type    string         `json:"type,omitempty"`    // The type of the scope.
	Project *ProjectScheme `json:"project,omitempty"` // The project in the scope.
}

// ProjectUpdateScheme represents the update scheme for a project in Jira.
type ProjectUpdateScheme struct {
	AssigneeType        string `json:"assigneeType,omitempty"`        // The type of assignee for the project.
	AvatarID            int    `json:"avatarId,omitempty"`            // The ID of the avatar for the project.
	CategoryID          int    `json:"categoryId,omitempty"`          // The ID of the category for the project.
	Description         string `json:"description,omitempty"`         // The description of the project.
	IssueSecurityScheme int    `json:"issueSecurityScheme,omitempty"` // The ID of the issue security scheme for the project.
	Key                 string `json:"key,omitempty"`                 // The key of the project.
	Lead                string `json:"lead,omitempty"`                // Deprecated, use LeadAccountID instead. The lead of the project.
	LeadAccountID       string `json:"leadAccountId,omitempty"`       // The account ID of the lead for the project.
	Name                string `json:"name,omitempty"`                // The name of the project.
	NotificationScheme  int    `json:"notificationScheme,omitempty"`  // The ID of the notification scheme for the project.
	PermissionScheme    int    `json:"permissionScheme,omitempty"`    // The ID of the permission scheme for the project.
	URL                 string `json:"url,omitempty"`                 // The URL of the project.
	ProjectTemplateKey  string `json:"projectTemplateKey,omitempty"`  // The key of the project template for the project.
	ProjectTypeKey      string `json:"projectTypeKey,omitempty"`      // The key of the project type for the project.
}

// ProjectStatusPageScheme represents the status page scheme for a project in Jira.
type ProjectStatusPageScheme struct {
	Self     string                        `json:"self,omitempty"`     // The URL of the status page.
	ID       string                        `json:"id,omitempty"`       // The ID of the status page.
	Name     string                        `json:"name,omitempty"`     // The name of the status page.
	Subtask  bool                          `json:"subtask,omitempty"`  // Indicates if the status page is a subtask.
	Statuses []*ProjectStatusDetailsScheme `json:"statuses,omitempty"` // The statuses in the status page.
}

// ProjectStatusDetailsScheme represents the status details scheme for a project in Jira.
type ProjectStatusDetailsScheme struct {
	Self           string                `json:"self,omitempty"`           // The URL of the status details.
	Description    string                `json:"description,omitempty"`    // The description of the status.
	IconURL        string                `json:"iconUrl,omitempty"`        // The URL of the icon for the status.
	Name           string                `json:"name,omitempty"`           // The name of the status.
	ID             string                `json:"id,omitempty"`             // The ID of the status.
	StatusCategory *StatusCategoryScheme `json:"statusCategory,omitempty"` // The status category of the status.
}

// ProjectHierarchyScheme represents the hierarchy scheme for a project in Jira.
type ProjectHierarchyScheme struct {
	EntityID   string                             `json:"entityId,omitempty"`   // The entity ID of the hierarchy.
	Level      int                                `json:"level,omitempty"`      // The level of the hierarchy.
	Name       string                             `json:"name,omitempty"`       // The name of the hierarchy.
	IssueTypes []*ProjectHierarchyIssueTypeScheme `json:"issueTypes,omitempty"` // The issue types in the hierarchy.
}

// ProjectHierarchyIssueTypeScheme represents the hierarchy issue type scheme for a project in Jira.
type ProjectHierarchyIssueTypeScheme struct {
	ID       int    `json:"id,omitempty"`       // The ID of the issue type.
	EntityID string `json:"entityId,omitempty"` // The entity ID of the issue type.
	Name     string `json:"name,omitempty"`     // The name of the issue type.
	AvatarID int    `json:"avatarId,omitempty"` // The ID of the avatar for the issue type.
}

// ProjectIdentifierScheme represents the identifier scheme for a project in Jira.
type ProjectIdentifierScheme struct {
	ID  int    `json:"id,omitempty"`  // The ID of the project.
	Key string `json:"key,omitempty"` // The key of the project.
}
