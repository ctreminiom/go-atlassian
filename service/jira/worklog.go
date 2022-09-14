package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type WorklogSharedConnector interface {

	// Gets returns worklog details for a list of worklog IDs.
	//
	// The returned list of worklogs is limited to 1000 items.
	//
	// POST /rest/api/{2-3}/worklog/list
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-worklogs
	Gets(ctx context.Context, worklogIds []int, expand []string) ([]*model.IssueWorklogScheme, *model.ResponseScheme, error)

	// Get returns a worklog.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// GET /rest/api/{2-3}/issue/{issueIdOrKey}/worklog/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-worklog
	Get(ctx context.Context, issueKeyOrId, worklogId string, expand []string) (*model.IssueWorklogScheme, *model.ResponseScheme, error)

	// Issue returns worklogs for an issue, starting from the oldest worklog or from the worklog started on or after a date and time.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// GET /rest/api/{2-3}/issue/{issueIdOrKey}/worklog
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-issue-worklogs
	Issue(ctx context.Context, issueKeyOrId string, startAt, maxResults, after int, expand []string) (*model.IssueWorklogPageScheme, *model.ResponseScheme, error)

	// Delete deletes a worklog from an issue.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// DELETE /rest/api/{2-3}/issue/{issueIdOrKey}/worklog/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#delete-worklog
	Delete(ctx context.Context, issueKeyOrId, worklogId string, options *model.WorklogOptionsScheme) (*model.ResponseScheme, error)

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

	// Add adds a worklog to an issue.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// POST /rest/api/2/issue/{issueIdOrKey}/worklog
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#add-worklog
	Add(ctx context.Context, issueKeyOrID string, payload *model.WorklogPayloadSchemeV2, options *model.WorklogOptionsScheme) (*model.IssueWorklogScheme,
		*model.ResponseScheme, error)

	// Update updates a worklog.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// PUT /rest/api/2/issue/{issueIdOrKey}/worklog/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#update-worklog
	Update(ctx context.Context, issueKeyOrId, worklogId string, payload *model.WorklogPayloadSchemeV2, options *model.WorklogOptionsScheme) (
		*model.IssueWorklogScheme, *model.ResponseScheme, error)
}

type WorklogADFConnector interface {
	WorklogSharedConnector

	// Add adds a worklog to an issue.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// POST /rest/api/3/issue/{issueIdOrKey}/worklog
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#add-worklog
	Add(ctx context.Context, issueKeyOrID string, payload *model.WorklogPayloadSchemeV3, options *model.WorklogOptionsScheme) (*model.IssueWorklogScheme,
		*model.ResponseScheme, error)

	// Update updates a worklog.
	//
	// Time tracking must be enabled in Jira, otherwise this operation returns an error.
	//
	// PUT /rest/api/3/issue/{issueIdOrKey}/worklog/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#update-worklog
	Update(ctx context.Context, issueKeyOrId, worklogId string, payload *model.WorklogPayloadSchemeV3, options *model.WorklogOptionsScheme) (
		*model.IssueWorklogScheme, *model.ResponseScheme, error)
}
