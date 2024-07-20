package models

// RequestTypePageScheme represents a page of request types.
type RequestTypePageScheme struct {
	Size       int                        `json:"size,omitempty"`       // The size of the page.
	Start      int                        `json:"start,omitempty"`      // The start index of the page.
	Limit      int                        `json:"limit,omitempty"`      // The limit of the page.
	IsLastPage bool                       `json:"isLastPage,omitempty"` // Indicates if this is the last page.
	Values     []*RequestTypeScheme       `json:"values,omitempty"`     // The request types in the page.
	Expands    []string                   `json:"_expands,omitempty"`   // The fields to expand in the page.
	Links      *RequestTypePageLinkScheme `json:"_links,omitempty"`     // The links related to the page.
}

// RequestTypePageLinkScheme represents the links related to a page of request types.
type RequestTypePageLinkScheme struct {
	Self    string `json:"self,omitempty"`    // The self link of the page.
	Base    string `json:"base,omitempty"`    // The base link of the page.
	Context string `json:"context,omitempty"` // The context link of the page.
	Next    string `json:"next,omitempty"`    // The next link of the page.
	Prev    string `json:"prev,omitempty"`    // The previous link of the page.
}

// ProjectRequestTypePageScheme represents a page of project request types.
type ProjectRequestTypePageScheme struct {
	Expands    []string                          `json:"_expands,omitempty"`   // The fields to expand in the page.
	Size       int                               `json:"size,omitempty"`       // The size of the page.
	Start      int                               `json:"start,omitempty"`      // The start index of the page.
	Limit      int                               `json:"limit,omitempty"`      // The limit of the page.
	IsLastPage bool                              `json:"isLastPage,omitempty"` // Indicates if this is the last page.
	Values     []*RequestTypeScheme              `json:"values,omitempty"`     // The project request types in the page.
	Links      *ProjectRequestTypePageLinkScheme `json:"_links,omitempty"`     // The links related to the page.
}

// ProjectRequestTypePageLinkScheme represents the links related to a page of project request types.
type ProjectRequestTypePageLinkScheme struct {
	Base    string `json:"base,omitempty"`    // The base link of the page.
	Context string `json:"context,omitempty"` // The context link of the page.
	Next    string `json:"next,omitempty"`    // The next link of the page.
	Prev    string `json:"prev,omitempty"`    // The previous link of the page.
}

// RequestTypeScheme represents a request type.
type RequestTypeScheme struct {
	ID            string                  `json:"id,omitempty"`            // The ID of the request type.
	Name          string                  `json:"name,omitempty"`          // The name of the request type.
	Description   string                  `json:"description,omitempty"`   // The description of the request type.
	HelpText      string                  `json:"helpText,omitempty"`      // The help text of the request type.
	Practice      string                  `json:"practice,omitempty"`      // The practice of the request type.
	IssueTypeID   string                  `json:"issueTypeId,omitempty"`   // The issue type ID of the request type.
	ServiceDeskID string                  `json:"serviceDeskId,omitempty"` // The service desk ID of the request type.
	PortalID      string                  `json:"portalId,omitempty"`      // The portal ID of the request type.
	GroupIDs      []string                `json:"groupIds,omitempty"`      // The group IDs of the request type.
	Expands       []string                `json:"_expands,omitempty"`      // The fields to expand in the request type.
	Links         *RequestTypeLinksScheme `json:"_links,omitempty"`        // The links related to the request type.
}

// RequestTypeLinksScheme represents the links related to a request type.
type RequestTypeLinksScheme struct {
	Self string `json:"self,omitempty"` // The self link the request type.
}

// RequestTypePayloadScheme represents a request type payload.
type RequestTypePayloadScheme struct {
	Description string `json:"description,omitempty"` // The description of the request type payload.
	HelpText    string `json:"helpText,omitempty"`    // The help text of the request type payload.
	IssueTypeID string `json:"issueTypeId,omitempty"` // The issue type ID of the request type payload.
	Name        string `json:"name,omitempty"`        // The name of the request type payload.
}

// RequestTypeFieldsScheme represents the fields of a request type.
type RequestTypeFieldsScheme struct {
	RequestTypeFields         []*RequestTypeFieldScheme `json:"requestTypeFields,omitempty"`         // The fields of the request type.
	CanRaiseOnBehalfOf        bool                      `json:"canRaiseOnBehalfOf,omitempty"`        // Indicates if the request type can be raised on behalf of.
	CanAddRequestParticipants bool                      `json:"canAddRequestParticipants,omitempty"` // Indicates if the request type can add request participants.
}

// RequestTypeFieldScheme represents a field of a request type.
type RequestTypeFieldScheme struct {
	FieldID       string                         `json:"fieldId,omitempty"`       // The field ID of the request type field.
	Name          string                         `json:"name,omitempty"`          // The name of the request type field.
	Description   string                         `json:"description,omitempty"`   // The description of the request type field.
	Required      bool                           `json:"required,omitempty"`      // Indicates if the request type field is required.
	DefaultValues []*RequestTypeFieldValueScheme `json:"defaultValues,omitempty"` // The default values of the request type field.
	ValidValues   []*RequestTypeFieldValueScheme `json:"validValues,omitempty"`   // The valid values of the request type field.
	PresetValues  []string                       `json:"presetValues,omitempty"`  // The preset values of the request type field.
	JiraSchema    *RequestTypeJiraSchema         `json:"jiraSchema,omitempty"`    // The Jira schema of the request type field.
	Visible       bool                           `json:"visible,omitempty"`       // Indicates if the request type field is visible.
}

// RequestTypeFieldValueScheme represents a field value of a request type field.
type RequestTypeFieldValueScheme struct {
	Value    string                         `json:"value,omitempty"`    // The value of the request type field value.
	Label    string                         `json:"label,omitempty"`    // The label of the request type field value.
	Children []*RequestTypeFieldValueScheme `json:"children,omitempty"` // The children of the request type field value.
}

// RequestTypeJiraSchema represents the Jira schema of a request type field.
type RequestTypeJiraSchema struct {
	Type          string            `json:"type,omitempty"`          // The type of the Jira schema.
	Items         string            `json:"items,omitempty"`         // The items of the Jira schema.
	System        string            `json:"system,omitempty"`        // The system of the Jira schema.
	Custom        string            `json:"custom,omitempty"`        // The custom of the Jira schema.
	CustomID      int               `json:"customId,omitempty"`      // The custom ID of the Jira schema.
	Configuration map[string]string `json:"configuration,omitempty"` // The configuration of the Jira schema.
}

// RequestTypeGroupPageScheme represents a page of project request type groups.
type RequestTypeGroupPageScheme struct {
	Expands    []string                        `json:"_expands,omitempty"`   // The fields to expand in the page.
	Size       int                             `json:"size,omitempty"`       // The size of the page.
	Start      int                             `json:"start,omitempty"`      // The start index of the page.
	Limit      int                             `json:"limit,omitempty"`      // The limit of the page.
	IsLastPage bool                            `json:"isLastPage,omitempty"` // Indicates if this is the last page.
	Values     []*RequestTypeGroupsScheme      `json:"values,omitempty"`     // The project request types in the page.
	Links      *RequestTypeGroupPageLinkScheme `json:"_links,omitempty"`     // The links related to the page.
}

// RequestTypeGroupPageLinkScheme represents the links related to a page of project request type groups.
type RequestTypeGroupPageLinkScheme struct {
	Base    string `json:"base,omitempty"`    // The base link of the page.
	Context string `json:"context,omitempty"` // The context link of the page.
	Next    string `json:"next,omitempty"`    // The next link of the page.
	Prev    string `json:"prev,omitempty"`    // The previous link of the page.
}

// RequestTypeGroupsScheme represents the groups for request types.
type RequestTypeGroupsScheme struct {
	ID   string `json:"id,omitempty"`   // The ID of the request type group.
	Name string `json:"name,omitempty"` // The name of the request type group.
}
