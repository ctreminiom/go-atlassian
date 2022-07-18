package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type SharedIssue interface {

	// Delete deletes an issue.
	// An issue cannot be deleted if it has one or more subtasks.
	// To delete an issue with subtasks, set deleteSubtasks.
	// This causes the issue's subtasks to be deleted with the issue.
	// DELETE /rest/api/{2-3}/issue/{issueIdOrKey}
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#delete-issue
	Delete(ctx context.Context, issueKeyOrId string, deleteSubTasks bool) (*model.ResponseScheme, error)

	// Assign assigns an issue to a user.
	// Use this operation when the calling user does not have the Edit Issues permission but has the
	// Assign issue permission for the project that the issue is in.
	// If accountId is set to:
	//  1. "-1", the issue is assigned to the default assignee for the project.
	//  2. null, the issue is set to unassigned.
	// PUT /rest/api/{2-3}/issue/{issueIdOrKey}/assignee
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#assign-issue
	Assign(ctx context.Context, issueKeyOrId, accountId string) (*model.ResponseScheme, error)

	// Notify creates an email notification for an issue and adds it to the mail queue.
	// POST /rest/api/{2-3}/issue/{issueIdOrKey}/notify
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#send-notification-for-issue
	Notify(ctx context.Context, issueKeyOrId string, options *model.IssueNotifyOptionsScheme) (*model.ResponseScheme, error)

	// Transitions returns either all transitions or a transition that can be performed by the user on an issue, based on the issue's status.
	// Note, if a request is made for a transition that does not exist or cannot be performed on the issue,
	// given its status, the response will return any empty transitions list.
	// GET /rest/api/{2-3}/issue/{issueIdOrKey}/transitions
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#get-transitions
	Transitions(ctx context.Context, issueKeyOrId string) (*model.IssueTransitionsScheme, *model.ResponseScheme, error)
	// TODO The Transitions methods requires more parameters such as expand, transitionId, and more
	// The parameters are documented on this [page](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issues/#api-rest-api-3-issue-issueidorkey-transitions-get)
}

type RichTextIssue interface {
	SharedIssue

	// Create creates an issue or, where the option to create subtasks is enabled in Jira, a subtask.
	// POST /rest/api/2/issue
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#create-issue
	Create(ctx context.Context, payload *model.IssueSchemeV2, customFields *model.CustomFields) (*model.IssueResponseScheme, *model.ResponseScheme, error)

	// Creates issues and, where the option to create subtasks is enabled in Jira, subtasks.
	// Creates upto 50 issues and, where the option to create subtasks is enabled in Jira, subtasks.
	// Transitions may be applied, to move the issues or subtasks to a workflow step other than the default start step, and issue properties set.
	// POST /rest/api/2/issue/bulk
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#bulk-create-issue
	Creates(ctx context.Context, payload []*model.IssueBulkSchemeV2) (*model.IssueBulkResponseScheme, *model.ResponseScheme, error)

	// Get returns the details for an issue.
	// The issue is identified by its ID or key, however, if the identifier doesn't match an issue, a case-insensitive search
	// and check for moved issues is performed. If a matching issue is found its details are returned, a 302 or other redirect is not returned.
	// The issue key returned to the response is the key of the issue found.
	// GET /rest/api/2/issue/{issueIdOrKey}
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#get-issue
	Get(ctx context.Context, issueKeyOrId string, fields, expand []string) (*model.IssueSchemeV2, *model.ResponseScheme, error)

	// Update edits an issue.
	// Edits an issue. A transition may be applied and issue properties updated as part of the edit.
	// The edits to the issue's fields are defined using update and fields
	// PUT /rest/api/2/issue/{issueIdOrKey}
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#edit-issue
	Update(ctx context.Context, issueKeyOrId string, notify bool, payload *model.IssueSchemeV2, customFields *model.CustomFields,
		operations *model.UpdateOperations) (*model.ResponseScheme, error)

	// Move performs an issue transition and, if the transition has a screen, updates the fields from the transition screen.
	// POST /rest/api/2/issue/{issueIdOrKey}/transitions
	// sortByCategory To update the fields on the transition screen, specify the fields in the fields or update parameters in the request body. Get details about the fields using Get transitions with the transitions.fields expand.
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#transition-issue
	Move(ctx context.Context, issueKeyOrId, transitionId string, options *model.IssueMoveOptionsV2) (*model.ResponseScheme, error)
}

type AdfIssue interface {
	SharedIssue

	// Create creates an issue or, where the option to create subtasks is enabled in Jira, a subtask.
	// POST /rest/api/3/issue
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#create-issue
	Create(ctx context.Context, payload *model.IssueScheme, customFields *model.CustomFields) (*model.IssueResponseScheme, *model.ResponseScheme, error)

	// Creates issues and, where the option to create subtasks is enabled in Jira, subtasks.
	// Creates upto 50 issues and, where the option to create subtasks is enabled in Jira, subtasks.
	// Transitions may be applied, to move the issues or subtasks to a workflow step other than the default start step, and issue properties set.
	// POST /rest/api/3/issue/bulk
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#bulk-create-issue
	Creates(ctx context.Context, payload []*model.IssueBulkSchemeV3) (*model.IssueBulkResponseScheme, *model.ResponseScheme, error)

	// Get returns the details for an issue.
	// The issue is identified by its ID or key, however, if the identifier doesn't match an issue, a case-insensitive search
	// and check for moved issues is performed. If a matching issue is found its details are returned, a 302 or other redirect is not returned.
	// The issue key returned to the response is the key of the issue found.
	// GET /rest/api/3/issue/{issueIdOrKey}
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#get-issue
	Get(ctx context.Context, issueKeyOrId string, fields, expand []string) (*model.IssueScheme, *model.ResponseScheme, error)

	// Update edits an issue.
	// Edits an issue. A transition may be applied and issue properties updated as part of the edit.
	// The edits to the issue's fields are defined using update and fields
	// PUT /rest/api/2/issue/{issueIdOrKey}
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#edit-issue
	Update(ctx context.Context, issueKeyOrId string, notify bool, payload *model.IssueScheme, customFields *model.CustomFields,
		operations *model.UpdateOperations) (*model.ResponseScheme, error)

	// Move performs an issue transition and, if the transition has a screen, updates the fields from the transition screen.
	// POST /rest/api/2/issue/{issueIdOrKey}/transitions
	// sortByCategory To update the fields on the transition screen, specify the fields in the fields or update parameters in the request body. Get details about the fields using Get transitions with the transitions.fields expand.
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#transition-issue
	Move(ctx context.Context, issueKeyOrId, transitionId string, options *model.IssueMoveOptionsV3) (*model.ResponseScheme, error)
}
