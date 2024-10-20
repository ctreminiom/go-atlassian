package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type WorklogSharedConnector interface {

	// Delete deletes a worklog from an issue.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// DELETE /rest/api/{2-3}/issue/{issueKeyOrID}/worklog/{worklogID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#delete-worklog
	Delete(ctx context.Context, issueKeyOrID, worklogID string, options *model.WorklogOptionsScheme) (*model.ResponseScheme, error)

	// Deleted returns a list of IDs and delete timestamps for worklogs deleted after a date and time.
	//
	// This resource is paginated, with a limit of 1000 worklogs per page. Each page lists worklogs from oldest to youngest.
	// If the number of items in the date range exceeds 1000, until indicates the timestamp of the youngest item on the page.
	// Also, nextPage provides the URL for the next page of worklogs.
	// The lastPage parameter is set to true on the last page of worklogs.
	//
	// This resource does not return worklogs deleted during the minute preceding the request.
	//
	// GET /rest/api/{2-3}/worklog/deleted
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-ids-of-deleted-worklogs
	Deleted(ctx context.Context, since int) (result *model.ChangedWorklogPageScheme, response *model.ResponseScheme, err error)

	// Updated returns a list of IDs and update timestamps for worklogs updated after a date and time.
	//
	// This resource is paginated, with a limit of 1000 worklogs per page. Each page lists worklogs from oldest to youngest.
	// If the number of items in the date range exceeds 1000, until indicates the timestamp of the youngest item on the page.
	// Also, nextPage provides the URL for the next page of worklogs.
	// The lastPage parameter is set to true on the last page of worklogs.
	//
	// This resource does not return worklogs updated during the minute preceding the request.
	//
	// GET /rest/api/{2-3}/worklog/updated
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-ids-of-updated-worklogs
	Updated(ctx context.Context, since int, expand []string) (*model.ChangedWorklogPageScheme, *model.ResponseScheme, error)
}

type WorklogRichTextConnector interface {
	WorklogSharedConnector

	// Gets returns worklog details for a list of worklog IDs.
	//
	// The returned list of worklogs is limited to 1000 items.
	//
	// POST /rest/api/{2-3}/worklog/list
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-worklogs
	Gets(ctx context.Context, worklogIDs []int, expand []string) ([]*model.IssueWorklogRichTextScheme, *model.ResponseScheme, error)

	// Get returns a worklog.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// GET /rest/api/{2-3}/issue/{issueKeyOrID}/worklog/{worklogID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-worklog
	Get(ctx context.Context, issueKeyOrID, worklogID string, expand []string) (*model.IssueWorklogRichTextScheme, *model.ResponseScheme, error)

	// Issue returns worklogs for an issue, starting from the oldest worklog or from the worklog started on or after a date and time.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// GET /rest/api/{2-3}/issue/{issueKeyOrID}/worklog
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-issue-worklogs
	Issue(ctx context.Context, issueKeyOrID string, startAt, maxResults, after int, expand []string) (*model.IssueWorklogRichTextPageScheme, *model.ResponseScheme, error)

	// Add adds a worklog to an issue.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// POST /rest/api/2/issue/{issueKeyOrID}/worklog
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#add-worklog
	Add(ctx context.Context, issueKeyOrID string, payload *model.WorklogRichTextPayloadScheme, options *model.WorklogOptionsScheme) (*model.IssueWorklogRichTextScheme,
		*model.ResponseScheme, error)

	// Update updates a worklog.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// PUT /rest/api/2/issue/{issueKeyOrID}/worklog/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#update-worklog
	Update(ctx context.Context, issueKeyOrID, worklogID string, payload *model.WorklogRichTextPayloadScheme, options *model.WorklogOptionsScheme) (
		*model.IssueWorklogRichTextScheme, *model.ResponseScheme, error)
}

type WorklogADFConnector interface {
	WorklogSharedConnector

	// Gets returns worklog details for a list of worklog IDs.
	//
	// The returned list of worklogs is limited to 1000 items.
	//
	// POST /rest/api/{2-3}/worklog/list
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-worklogs
	Gets(ctx context.Context, worklogIDs []int, expand []string) ([]*model.IssueWorklogADFScheme, *model.ResponseScheme, error)

	// Get returns a worklog.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// GET /rest/api/{2-3}/issue/{issueKeyOrID}/worklog/{worklogID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-worklog
	Get(ctx context.Context, issueKeyOrID, worklogID string, expand []string) (*model.IssueWorklogADFScheme, *model.ResponseScheme, error)

	// Issue returns worklogs for an issue, starting from the oldest worklog or from the worklog started on or after a date and time.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// GET /rest/api/{2-3}/issue/{issueKeyOrID}/worklog
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-issue-worklogs
	Issue(ctx context.Context, issueKeyOrID string, startAt, maxResults, after int, expand []string) (*model.IssueWorklogADFPageScheme, *model.ResponseScheme, error)

	// Add adds a worklog to an issue.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// POST /rest/api/3/issue/{issueKeyOrID}/worklog
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#add-worklog
	Add(ctx context.Context, issueKeyOrID string, payload *model.WorklogADFPayloadScheme, options *model.WorklogOptionsScheme) (*model.IssueWorklogADFScheme,
		*model.ResponseScheme, error)

	// Update updates a worklog.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// PUT /rest/api/3/issue/{issueKeyOrID}/worklog/{worklogID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#update-worklog
	Update(ctx context.Context, issueKeyOrID, worklogID string, payload *model.WorklogADFPayloadScheme, options *model.WorklogOptionsScheme) (
		*model.IssueWorklogADFScheme, *model.ResponseScheme, error)
}
