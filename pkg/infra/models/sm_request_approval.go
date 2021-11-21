package models

type CustomerApprovalPageScheme struct {
	Size       int                             `json:"size,omitempty"`
	Start      int                             `json:"start,omitempty"`
	Limit      int                             `json:"limit,omitempty"`
	IsLastPage bool                            `json:"isLastPage,omitempty"`
	Values     []*CustomerApprovalScheme       `json:"values,omitempty"`
	Expands    []string                        `json:"_expands,omitempty"`
	Links      *CustomerApprovalPageLinkScheme `json:"_links,omitempty"`
}

type CustomerApprovalPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type CustomerApprovalScheme struct {
	ID                string                      `json:"id,omitempty"`
	Name              string                      `json:"name,omitempty"`
	FinalDecision     string                      `json:"finalDecision,omitempty"`
	CanAnswerApproval bool                        `json:"canAnswerApproval,omitempty"`
	Approvers         []*CustomerApproveScheme    `json:"approvers,omitempty"`
	CreatedDate       *CustomerRequestDateScheme  `json:"createdDate,omitempty"`
	CompletedDate     *CustomerRequestDateScheme  `json:"completedDate,omitempty"`
	Links             *CustomerApprovalLinkScheme `json:"_links,omitempty"`
}

type CustomerApprovalLinkScheme struct {
	Self string `json:"self"`
}

type CustomerApproveScheme struct {
	Approver         *ApproverScheme `json:"approver,omitempty"`
	ApproverDecision string          `json:"approverDecision,omitempty"`
}

type ApproverScheme struct {
	AccountID    string              `json:"accountId,omitempty"`
	Name         string              `json:"name,omitempty"`
	Key          string              `json:"key,omitempty"`
	EmailAddress string              `json:"emailAddress,omitempty"`
	DisplayName  string              `json:"displayName,omitempty"`
	Active       bool                `json:"active,omitempty"`
	TimeZone     string              `json:"timeZone,omitempty"`
	Links        *ApproverLinkScheme `json:"_links,omitempty"`
}

type ApproverLinkScheme struct {
	JiraRest   string           `json:"jiraRest,omitempty"`
	AvatarUrls *AvatarURLScheme `json:"avatarUrls,omitempty"`
	Self       string           `json:"self,omitempty"`
}
