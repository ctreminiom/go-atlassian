package models

// CustomerApprovalPageScheme represents a page of customer approvals in a system.
type CustomerApprovalPageScheme struct {
	Size       int                             `json:"size,omitempty"`       // The number of customer approvals on the page.
	Start      int                             `json:"start,omitempty"`      // The index of the first customer approval on the page.
	Limit      int                             `json:"limit,omitempty"`      // The maximum number of customer approvals that can be on the page.
	IsLastPage bool                            `json:"isLastPage,omitempty"` // Indicates if this is the last page of customer approvals.
	Values     []*CustomerApprovalScheme       `json:"values,omitempty"`     // The customer approvals on the page.
	Expands    []string                        `json:"_expands,omitempty"`   // Additional data related to the customer approvals.
	Links      *CustomerApprovalPageLinkScheme `json:"_links,omitempty"`     // Links related to the page of customer approvals.
}

// CustomerApprovalPageLinkScheme represents links related to a page of customer approvals.
type CustomerApprovalPageLinkScheme struct {
	Self    string `json:"self,omitempty"`    // The URL of the page itself.
	Base    string `json:"base,omitempty"`    // The base URL for the links.
	Context string `json:"context,omitempty"` // The context for the links.
	Next    string `json:"next,omitempty"`    // The URL for the next page of customer approvals.
	Prev    string `json:"prev,omitempty"`    // The URL for the previous page of customer approvals.
}

// CustomerApprovalScheme represents a customer approval in a system.
type CustomerApprovalScheme struct {
	ID                string                      `json:"id,omitempty"`                // The ID of the customer approval.
	Name              string                      `json:"name,omitempty"`              // The name of the customer approval.
	FinalDecision     string                      `json:"finalDecision,omitempty"`     // The final decision of the customer approval.
	CanAnswerApproval bool                        `json:"canAnswerApproval,omitempty"` // Indicates if the customer approval can be answered.
	Approvers         []*CustomerApproveScheme    `json:"approvers,omitempty"`         // The approvers of the customer approval.
	CreatedDate       *CustomerRequestDateScheme  `json:"createdDate,omitempty"`       // The created date of the customer approval.
	CompletedDate     *CustomerRequestDateScheme  `json:"completedDate,omitempty"`     // The completed date of the customer approval.
	Links             *CustomerApprovalLinkScheme `json:"_links,omitempty"`            // Links related to the customer approval.
}

// CustomerApprovalLinkScheme represents links related to a customer approval.
type CustomerApprovalLinkScheme struct {
	Self string `json:"self"` // The URL of the customer approval itself.
}

// CustomerApproveScheme represents an approver of a customer approval.
type CustomerApproveScheme struct {
	Approver         *ApproverScheme `json:"approver,omitempty"`         // The approver of the customer approval.
	ApproverDecision string          `json:"approverDecision,omitempty"` // The decision of the approver.
}

// ApproverScheme represents an approver in a system.
type ApproverScheme struct {
	AccountID    string              `json:"accountId,omitempty"`    // The account ID of the approver.
	Name         string              `json:"name,omitempty"`         // The name of the approver.
	Key          string              `json:"key,omitempty"`          // The key of the approver.
	EmailAddress string              `json:"emailAddress,omitempty"` // The email address of the approver.
	DisplayName  string              `json:"displayName,omitempty"`  // The display name of the approver.
	Active       bool                `json:"active,omitempty"`       // Indicates if the approver is active.
	TimeZone     string              `json:"timeZone,omitempty"`     // The time zone of the approver.
	Links        *ApproverLinkScheme `json:"_links,omitempty"`       // Links related to the approver.
}

// ApproverLinkScheme represents links related to an approver.
type ApproverLinkScheme struct {
	JiraREST   string           `json:"jiraRest,omitempty"`   // The Jira REST API link for the approver.
	AvatarURLs *AvatarURLScheme `json:"avatarUrls,omitempty"` // The avatar URLs of the approver.
	Self       string           `json:"self,omitempty"`       // The URL of the approver itself.
}
