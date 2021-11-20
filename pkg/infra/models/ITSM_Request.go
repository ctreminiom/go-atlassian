package models

type RequestGetOptionsScheme struct {
	SearchTerm        string
	RequestOwnerships []string
	RequestStatus     string
	ApprovalStatus    string
	OrganizationId    int
	ServiceDeskID     int
	RequestTypeID     int
	Expand            []string
}

type CustomerRequestTransitionPageScheme struct {
	Size       int                                      `json:"size,omitempty"`
	Start      int                                      `json:"start,omitempty"`
	Limit      int                                      `json:"limit,omitempty"`
	IsLastPage bool                                     `json:"isLastPage,omitempty"`
	Values     []*CustomerRequestTransitionScheme       `json:"values,omitempty"`
	Expands    []string                                 `json:"_expands,omitempty"`
	Links      *CustomerRequestTransitionPageLinkScheme `json:"_links,omitempty"`
}

type CustomerRequestTransitionScheme struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CustomerRequestTransitionPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type CustomerRequestsScheme struct {
	Size       int                          `json:"size,omitempty"`
	Start      int                          `json:"start,omitempty"`
	Limit      int                          `json:"limit,omitempty"`
	IsLastPage bool                         `json:"isLastPage,omitempty"`
	Values     []*CustomerRequestScheme     `json:"values,omitempty"`
	Expands    []string                     `json:"_expands,omitempty"`
	Links      *CustomerRequestsLinksScheme `json:"_links,omitempty"`
}

type CustomerRequestsLinksScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type CustomerRequestTypeScheme struct {
	ID            string   `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	HelpText      string   `json:"helpText,omitempty"`
	IssueTypeID   string   `json:"issueTypeId,omitempty"`
	ServiceDeskID string   `json:"serviceDeskId,omitempty"`
	GroupIds      []string `json:"groupIds,omitempty"`
}

type CustomerRequestServiceDeskScheme struct {
	ID          string `json:"id,omitempty"`
	ProjectID   string `json:"projectId,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	ProjectKey  string `json:"projectKey,omitempty"`
}

type CustomerRequestDateScheme struct {
	Iso8601     string `json:"iso8601,omitempty"`
	Jira        string `json:"jira,omitempty"`
	Friendly    string `json:"friendly,omitempty"`
	EpochMillis int    `json:"epochMillis,omitempty"`
}

type CustomerRequestReporterScheme struct {
	AccountID    string `json:"accountId,omitempty"`
	Name         string `json:"name,omitempty"`
	Key          string `json:"key,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Active       bool   `json:"active,omitempty"`
	TimeZone     string `json:"timeZone,omitempty"`
}

type CustomerRequestRequestFieldValueScheme struct {
	FieldID string `json:"fieldId,omitempty"`
	Label   string `json:"label,omitempty"`
}

type CustomerRequestCurrentStatusScheme struct {
	Status         string `json:"status,omitempty"`
	StatusCategory string `json:"statusCategory,omitempty"`
	StatusDate     struct {
	} `json:"statusDate,omitempty"`
}

type CustomerRequestLinksScheme struct {
	Self     string `json:"self,omitempty"`
	JiraRest string `json:"jiraRest,omitempty"`
	Web      string `json:"web,omitempty"`
	Agent    string `json:"agent,omitempty"`
}

type CustomerRequestScheme struct {
	IssueID            string                                    `json:"issueId,omitempty"`
	IssueKey           string                                    `json:"issueKey,omitempty"`
	RequestTypeID      string                                    `json:"requestTypeId,omitempty"`
	RequestType        *CustomerRequestTypeScheme                `json:"requestType,omitempty"`
	ServiceDeskID      string                                    `json:"serviceDeskId,omitempty"`
	ServiceDesk        *CustomerRequestServiceDeskScheme         `json:"serviceDesk,omitempty"`
	CreatedDate        *CustomerRequestDateScheme                `json:"createdDate,omitempty"`
	Reporter           *CustomerRequestReporterScheme            `json:"reporter,omitempty"`
	RequestFieldValues []*CustomerRequestRequestFieldValueScheme `json:"requestFieldValues,omitempty"`
	CurrentStatus      *CustomerRequestCurrentStatusScheme       `json:"currentStatus,omitempty"`
	Expands            []string                                  `json:"_expands,omitempty"`
	Links              *CustomerRequestLinksScheme               `json:"_links,omitempty"`
}
