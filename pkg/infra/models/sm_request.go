package models

// ServiceRequestOptionScheme represents the options for a service request.
type ServiceRequestOptionScheme struct {
	ApprovalStatus, RequestStatus, SearchTerm string   // The approval status, request status, and search term for the service request.
	OrganizationID, ServiceDeskID             int      // The organization ID and service desk ID for the service request.
	RequestTypeID                             int      // The request type ID for the service request.
	Expand, RequestOwnerships                 []string // The fields to expand and the request ownerships for the service request.
}

// CustomerRequestTransitionPageScheme represents a page of customer request transitions.
type CustomerRequestTransitionPageScheme struct {
	Size       int                                      `json:"size,omitempty"`       // The number of customer request transitions on the page.
	Start      int                                      `json:"start,omitempty"`      // The index of the first customer request transition on the page.
	Limit      int                                      `json:"limit,omitempty"`      // The maximum number of customer request transitions that can be on the page.
	IsLastPage bool                                     `json:"isLastPage,omitempty"` // Indicates if this is the last page of customer request transitions.
	Values     []*CustomerRequestTransitionScheme       `json:"values,omitempty"`     // The customer request transitions on the page.
	Expands    []string                                 `json:"_expands,omitempty"`   // Additional data related to the customer request transitions.
	Links      *CustomerRequestTransitionPageLinkScheme `json:"_links,omitempty"`     // Links related to the page of customer request transitions.
}

// CustomerRequestTransitionScheme represents a customer request transition.
type CustomerRequestTransitionScheme struct {
	ID   string `json:"id,omitempty"`   // The ID of the customer request transition.
	Name string `json:"name,omitempty"` // The name of the customer request transition.
}

// CustomerRequestTransitionPageLinkScheme represents links related to a page of customer request transitions.
type CustomerRequestTransitionPageLinkScheme struct {
	Self    string `json:"self,omitempty"`    // The URL of the page itself.
	Base    string `json:"base,omitempty"`    // The base URL for the links.
	Context string `json:"context,omitempty"` // The context for the links.
	Next    string `json:"next,omitempty"`    // The URL for the next page of customer request transitions.
	Prev    string `json:"prev,omitempty"`    // The URL for the previous page of customer request transitions.
}

// CustomerRequestPageScheme represents a page of customer requests.
type CustomerRequestPageScheme struct {
	Size       int                          `json:"size,omitempty"`       // The number of customer requests on the page.
	Start      int                          `json:"start,omitempty"`      // The index of the first customer request on the page.
	Limit      int                          `json:"limit,omitempty"`      // The maximum number of customer requests that can be on the page.
	IsLastPage bool                         `json:"isLastPage,omitempty"` // Indicates if this is the last page of customer requests.
	Values     []*CustomerRequestScheme     `json:"values,omitempty"`     // The customer requests on the page.
	Expands    []string                     `json:"_expands,omitempty"`   // Additional data related to the customer requests.
	Links      *CustomerRequestsLinksScheme `json:"_links,omitempty"`     // Links related to the page of customer requests.
}

// CustomerRequestsLinksScheme represents links related to a page of customer requests.
type CustomerRequestsLinksScheme struct {
	Self    string `json:"self,omitempty"`    // The URL of the page itself.
	Base    string `json:"base,omitempty"`    // The base URL for the links.
	Context string `json:"context,omitempty"` // The context for the links.
	Next    string `json:"next,omitempty"`    // The URL for the next page of customer requests.
	Prev    string `json:"prev,omitempty"`    // The URL for the previous page of customer requests.
}

// CustomerRequestTypeScheme represents a type of customer request.
type CustomerRequestTypeScheme struct {
	ID            string   `json:"id,omitempty"`            // The ID of the customer request type.
	Name          string   `json:"name,omitempty"`          // The name of the customer request type.
	Description   string   `json:"description,omitempty"`   // The description of the customer request type.
	HelpText      string   `json:"helpText,omitempty"`      // The help text for the customer request type.
	IssueTypeID   string   `json:"issueTypeId,omitempty"`   // The issue type ID for the customer request type.
	ServiceDeskID string   `json:"serviceDeskId,omitempty"` // The service desk ID for the customer request type.
	GroupIDs      []string `json:"groupIds,omitempty"`      // The group IDs for the customer request type.
}

// CustomerRequestServiceDeskScheme represents a service desk for a customer request.
type CustomerRequestServiceDeskScheme struct {
	ID          string `json:"id,omitempty"`          // The ID of the service desk.
	ProjectID   string `json:"projectId,omitempty"`   // The project ID for the service desk.
	ProjectName string `json:"projectName,omitempty"` // The project name for the service desk.
	ProjectKey  string `json:"projectKey,omitempty"`  // The project key for the service desk.
}

// CustomerRequestDateScheme represents a date for a customer request.
type CustomerRequestDateScheme struct {
	ISO8601     DateTimeScheme `json:"iso8601,omitempty"`     // The ISO 8601 format of the date.
	Jira        string         `json:"jira,omitempty"`        // The Jira format of the date.
	Friendly    string         `json:"friendly,omitempty"`    // The friendly format of the date.
	EpochMillis int            `json:"epochMillis,omitempty"` // The epoch milliseconds of the date.
}

// CustomerRequestReporterScheme represents a reporter for a customer request.
type CustomerRequestReporterScheme struct {
	AccountID    string `json:"accountId,omitempty"`    // The account ID of the reporter.
	Name         string `json:"name,omitempty"`         // The name of the reporter.
	Key          string `json:"key,omitempty"`          // The key of the reporter.
	EmailAddress string `json:"emailAddress,omitempty"` // The email address of the reporter.
	DisplayName  string `json:"displayName,omitempty"`  // The display name of the reporter.
	Active       bool   `json:"active,omitempty"`       // Indicates if the reporter is active.
	TimeZone     string `json:"timeZone,omitempty"`     // The time zone of the reporter.
}

// CustomerRequestRequestFieldValueScheme represents the field value of a customer request.
type CustomerRequestRequestFieldValueScheme struct {
	FieldID string      `json:"fieldId,omitempty"` // The ID of the field.
	Label   string      `json:"label,omitempty"`   // The label of the field.
	Value   interface{} `json:"value,omitempty"`   // The value of the field.
}

// CustomerRequestCurrentStatusScheme represents the current status of a customer request.
type CustomerRequestCurrentStatusScheme struct {
	Status         string                                  `json:"status,omitempty"`         // The status of the customer request.
	StatusCategory string                                  `json:"statusCategory,omitempty"` // The category of the status.
	StatusDate     *CustomerRequestCurrentStatusDateScheme `json:"statusDate,omitempty"`     // The date of the status.
}

// CustomerRequestCurrentStatusDateScheme represents a date for a customer request current status.
type CustomerRequestCurrentStatusDateScheme struct {
	ISO8601     DateTimeScheme `json:"iso8601,omitempty"`     // The ISO 8601 format of the date.
	Jira        string         `json:"jira,omitempty"`        // The Jira format of the date.
	Friendly    string         `json:"friendly,omitempty"`    // The friendly format of the date.
	EpochMillis int            `json:"epochMillis,omitempty"` // The epoch milliseconds of the date.
}

// CustomerRequestLinksScheme represents the links related to a customer request.
type CustomerRequestLinksScheme struct {
	Self     string `json:"self,omitempty"`     // The URL of the customer request itself.
	JiraREST string `json:"jiraRest,omitempty"` // The Jira REST API link for the customer request.
	Web      string `json:"web,omitempty"`      // The web link for the customer request.
	Agent    string `json:"agent,omitempty"`    // The agent link for the customer request.
}

// CustomerRequestScheme represents a customer request.
type CustomerRequestScheme struct {
	IssueID            string                                    `json:"issueId,omitempty"`            // The issue ID of the customer request.
	IssueKey           string                                    `json:"issueKey,omitempty"`           // The issue key of the customer request.
	RequestTypeID      string                                    `json:"requestTypeId,omitempty"`      // The request type ID of the customer request.
	RequestType        *CustomerRequestTypeScheme                `json:"requestType,omitempty"`        // The request type of the customer request.
	ServiceDeskID      string                                    `json:"serviceDeskId,omitempty"`      // The service desk ID of the customer request.
	ServiceDesk        *CustomerRequestServiceDeskScheme         `json:"serviceDesk,omitempty"`        // The service desk of the customer request.
	CreatedDate        *CustomerRequestDateScheme                `json:"createdDate,omitempty"`        // The created date of the customer request.
	Reporter           *CustomerRequestReporterScheme            `json:"reporter,omitempty"`           // The reporter of the customer request.
	RequestFieldValues []*CustomerRequestRequestFieldValueScheme `json:"requestFieldValues,omitempty"` // The field values of the customer request.
	CurrentStatus      *CustomerRequestCurrentStatusScheme       `json:"currentStatus,omitempty"`      // The current status of the customer request.
	Expands            []string                                  `json:"_expands,omitempty"`           // The fields to expand in the customer request.
	Links              *CustomerRequestLinksScheme               `json:"_links,omitempty"`             // The links related to the customer request.
}
