package models

type RequestTypePageScheme struct {
	Size       int                        `json:"size,omitempty"`
	Start      int                        `json:"start,omitempty"`
	Limit      int                        `json:"limit,omitempty"`
	IsLastPage bool                       `json:"isLastPage,omitempty"`
	Values     []*RequestTypeScheme       `json:"values,omitempty"`
	Expands    []string                   `json:"_expands,omitempty"`
	Links      *RequestTypePageLinkScheme `json:"_links,omitempty"`
}

type RequestTypePageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type ProjectRequestTypePageScheme struct {
	Expands    []string                          `json:"_expands,omitempty"`
	Size       int                               `json:"size,omitempty"`
	Start      int                               `json:"start,omitempty"`
	Limit      int                               `json:"limit,omitempty"`
	IsLastPage bool                              `json:"isLastPage,omitempty"`
	Values     []*RequestTypeScheme              `json:"values,omitempty"`
	Links      *ProjectRequestTypePageLinkScheme `json:"_links,omitempty"`
}

type ProjectRequestTypePageLinkScheme struct {
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type RequestTypeScheme struct {
	ID            string   `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	HelpText      string   `json:"helpText,omitempty"`
	IssueTypeID   string   `json:"issueTypeId,omitempty"`
	ServiceDeskID string   `json:"serviceDeskId,omitempty"`
	GroupIds      []string `json:"groupIds,omitempty"`
	Expands       []string `json:"_expands,omitempty"`
}

type RequestTypePayloadScheme struct {
	Description string `json:"description,omitempty"`
	HelpText    string `json:"helpText,omitempty"`
	IssueTypeId string `json:"issueTypeId,omitempty"`
	Name        string `json:"name,omitempty"`
}

type RequestTypeFieldsScheme struct {
	RequestTypeFields         []*RequestTypeFieldScheme `json:"requestTypeFields,omitempty"`
	CanRaiseOnBehalfOf        bool                      `json:"canRaiseOnBehalfOf,omitempty"`
	CanAddRequestParticipants bool                      `json:"canAddRequestParticipants,omitempty"`
}

type RequestTypeFieldScheme struct {
	FieldID       string                         `json:"fieldId,omitempty"`
	Name          string                         `json:"name,omitempty"`
	Description   string                         `json:"description,omitempty"`
	Required      bool                           `json:"required,omitempty"`
	DefaultValues []*RequestTypeFieldValueScheme `json:"defaultValues,omitempty"`
	ValidValues   []*RequestTypeFieldValueScheme `json:"validValues,omitempty"`
	JiraSchema    *RequestTypeJiraSchema         `json:"jiraSchema,omitempty"`
	Visible       bool                           `json:"visible,omitempty"`
}

type RequestTypeFieldValueScheme struct {
	Value    string        `json:"value,omitempty"`
	Label    string        `json:"label,omitempty"`
	Children []interface{} `json:"children,omitempty"`
}

type RequestTypeJiraSchema struct {
	Type     string `json:"type,omitempty"`
	Items    string `json:"items,omitempty"`
	System   string `json:"system,omitempty"`
	Custom   string `json:"custom,omitempty"`
	CustomID int    `json:"customId,omitempty"`
}
