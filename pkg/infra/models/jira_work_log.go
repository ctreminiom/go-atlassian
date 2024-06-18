package models

// WorklogOptionsScheme represents the options for a worklog in Jira.
type WorklogOptionsScheme struct {
	Notify               bool     // Indicates if notifications should be sent for the worklog.
	AdjustEstimate       string   // The method for adjusting the estimate of the issue.
	NewEstimate          string   // The new estimate of the issue if the adjust estimate is "new".
	ReduceBy             string   // The amount to reduce the estimate by if the adjust estimate is "manual".
	OverrideEditableFlag bool     // Indicates if the editable flag of the worklog should be overridden.
	Expand               []string // The fields that should be expanded in the response.
}

// WorklogADFPayloadScheme represents the payload for a worklog with Atlassian Document Format (ADF) content in Jira.
type WorklogADFPayloadScheme struct {
	Comment          *CommentNodeScheme            `json:"comment,omitempty"`          // The comment for the worklog in ADF format.
	Visibility       *IssueWorklogVisibilityScheme `json:"visibility,omitempty"`       // The visibility of the worklog.
	Started          string                        `json:"started,omitempty"`          // The date and time when the work started.
	TimeSpent        string                        `json:"timeSpent,omitempty"`        // The time spent on the work.
	TimeSpentSeconds int                           `json:"timeSpentSeconds,omitempty"` // The time spent on the work in seconds.
}

// WorklogRichTextPayloadScheme represents the payload for a worklog with rich text content in Jira.
type WorklogRichTextPayloadScheme struct {
	Comment          *CommentPayloadSchemeV2       `json:"comment,omitempty"`          // The comment for the worklog in rich text format.
	Visibility       *IssueWorklogVisibilityScheme `json:"visibility,omitempty"`       // The visibility of the worklog.
	Started          string                        `json:"started,omitempty"`          // The date and time when the work started.
	TimeSpent        string                        `json:"timeSpent,omitempty"`        // The time spent on the work.
	TimeSpentSeconds int                           `json:"timeSpentSeconds,omitempty"` // The time spent on the work in seconds.
}

// ChangedWorklogPageScheme represents a page of changed worklogs in Jira.
type ChangedWorklogPageScheme struct {
	Since    int                     `json:"since,omitempty"`    // The timestamp of the start of the period.
	Until    int                     `json:"until,omitempty"`    // The timestamp of the end of the period.
	Self     string                  `json:"self,omitempty"`     // The URL of the page.
	NextPage string                  `json:"nextPage,omitempty"` // The URL of the next page.
	LastPage bool                    `json:"lastPage,omitempty"` // Indicates if this is the last page of results.
	Values   []*ChangedWorklogScheme `json:"values,omitempty"`   // The changed worklogs on the page.
}

// ChangedWorklogScheme represents a changed worklog in Jira.
type ChangedWorklogScheme struct {
	WorklogID   int                             `json:"worklogId,omitempty"`   // The ID of the worklog.
	UpdatedTime int                             `json:"updatedTime,omitempty"` // The timestamp when the worklog was updated.
	Properties  []*ChangedWorklogPropertyScheme `json:"properties,omitempty"`  // The properties of the worklog.
}

// ChangedWorklogPropertyScheme represents a property of a changed worklog in Jira.
type ChangedWorklogPropertyScheme struct {
	Key string `json:"key,omitempty"` // The key of the property.
}

// IssueWorklogRichTextPageScheme represents a page of worklogs with rich text content in Jira.
type IssueWorklogRichTextPageScheme struct {
	StartAt    int                           `json:"startAt,omitempty"`    // The index of the first result returned.
	MaxResults int                           `json:"maxResults,omitempty"` // The maximum number of results returned.
	Total      int                           `json:"total,omitempty"`      // The total number of results available.
	Worklogs   []*IssueWorklogRichTextScheme `json:"worklogs,omitempty"`   // The worklogs on the page.
}

// IssueWorklogADFPageScheme represents a page of worklogs with Atlassian Document Format (ADF) content in Jira.
type IssueWorklogADFPageScheme struct {
	StartAt    int                      `json:"startAt,omitempty"`    // The index of the first result returned.
	MaxResults int                      `json:"maxResults,omitempty"` // The maximum number of results returned.
	Total      int                      `json:"total,omitempty"`      // The total number of results available.
	Worklogs   []*IssueWorklogADFScheme `json:"worklogs,omitempty"`   // The worklogs on the page.
}

// IssueWorklogVisibilityScheme represents the visibility of a worklog in Jira.
type IssueWorklogVisibilityScheme struct {
	Type       string `json:"type,omitempty"`       // The type of the visibility.
	Value      string `json:"value,omitempty"`      // The value of the visibility.
	Identifier string `json:"identifier,omitempty"` // The identifier of the visibility.
}

// IssueWorklogRichTextScheme represents a worklog with rich text content in Jira.
type IssueWorklogRichTextScheme struct {
	Self             string                        `json:"self,omitempty"`             // The URL of the worklog.
	Author           *UserDetailScheme             `json:"author,omitempty"`           // The author of the worklog.
	UpdateAuthor     *UserDetailScheme             `json:"updateAuthor,omitempty"`     // The user who last updated the worklog.
	Comment          string                        `json:"comment,omitempty"`          // The comment of the worklog.
	Updated          string                        `json:"updated,omitempty"`          // The date and time when the worklog was last updated.
	Visibility       *IssueWorklogVisibilityScheme `json:"visibility,omitempty"`       // The visibility of the worklog.
	Started          string                        `json:"started,omitempty"`          // The date and time when the work started.
	TimeSpent        string                        `json:"timeSpent,omitempty"`        // The time spent on the work.
	TimeSpentSeconds int                           `json:"timeSpentSeconds,omitempty"` // The time spent on the work in seconds.
	ID               string                        `json:"id,omitempty"`               // The ID of the worklog.
	IssueID          string                        `json:"issueId,omitempty"`          // The ID of the issue the worklog is associated with.
}

// IssueWorklogADFScheme represents a worklog with Atlassian Document Format (ADF) content in Jira.
type IssueWorklogADFScheme struct {
	Self             string                        `json:"self,omitempty"`             // The URL of the worklog.
	Author           *UserDetailScheme             `json:"author,omitempty"`           // The author of the worklog.
	UpdateAuthor     *UserDetailScheme             `json:"updateAuthor,omitempty"`     // The user who last updated the worklog.
	Comment          *CommentNodeScheme            `json:"comment,omitempty"`          // The comment of the worklog in ADF format.
	Updated          string                        `json:"updated,omitempty"`          // The date and time when the worklog was last updated.
	Visibility       *IssueWorklogVisibilityScheme `json:"visibility,omitempty"`       // The visibility of the worklog.
	Started          string                        `json:"started,omitempty"`          // The date and time when the work started.
	TimeSpent        string                        `json:"timeSpent,omitempty"`        // The time spent on the work.
	TimeSpentSeconds int                           `json:"timeSpentSeconds,omitempty"` // The time spent on the work in seconds.
	ID               string                        `json:"id,omitempty"`               // The ID of the worklog.
	IssueID          string                        `json:"issueId,omitempty"`          // The ID of the issue the worklog is associated with.
}
