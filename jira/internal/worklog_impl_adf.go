package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewWorklogADFService creates a new instance of WorklogADFService.
func NewWorklogADFService(client service.Connector, version string) (*WorklogADFService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &WorklogADFService{
		internalClient: &internalWorklogAdfImpl{c: client, version: version},
	}, nil
}

// WorklogADFService provides methods to manage worklogs in Jira Service Management.
type WorklogADFService struct {
	// internalClient is the connector interface for worklog operations.
	internalClient jira.WorklogADFConnector
}

// Gets returns worklog details for a list of worklog IDs.
//
// The returned list of worklogs is limited to 1000 items.
//
// POST /rest/api/{2-3}/worklog/list
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-worklogs
func (w *WorklogADFService) Gets(ctx context.Context, worklogIDs []int, expand []string) ([]*model.IssueWorklogADFScheme, *model.ResponseScheme, error) {
	return w.internalClient.Gets(ctx, worklogIDs, expand)
}

// Get returns a worklog.
//
// Time tracking must be enabled in Jira, otherwise this operation returns an error.
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}/worklog/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-worklog
func (w *WorklogADFService) Get(ctx context.Context, issueKeyOrID, worklogID string, expand []string) (*model.IssueWorklogADFScheme, *model.ResponseScheme, error) {
	return w.internalClient.Get(ctx, issueKeyOrID, worklogID, expand)
}

// Issue returns worklogs for an issue, starting from the oldest worklog or from the worklog started on or after a date and time.
//
// Time tracking must be enabled in Jira, otherwise this operation returns an error.
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}/worklog
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#get-issue-worklogs
func (w *WorklogADFService) Issue(ctx context.Context, issueKeyOrID string, startAt, maxResults, after int, expand []string) (*model.IssueWorklogADFPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Issue(ctx, issueKeyOrID, startAt, maxResults, after, expand)
}

// Delete deletes a worklog from an issue.
//
// Time tracking must be enabled in Jira, otherwise this operation returns an error.
//
// DELETE /rest/api/{2-3}/issue/{issueKeyOrID}/worklog/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#delete-worklog
func (w *WorklogADFService) Delete(ctx context.Context, issueKeyOrID, worklogID string, options *model.WorklogOptionsScheme) (*model.ResponseScheme, error) {
	return w.internalClient.Delete(ctx, issueKeyOrID, worklogID, options)
}

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
func (w *WorklogADFService) Deleted(ctx context.Context, since int) (result *model.ChangedWorklogPageScheme, response *model.ResponseScheme, err error) {
	return w.internalClient.Deleted(ctx, since)
}

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
func (w *WorklogADFService) Updated(ctx context.Context, since int, expand []string) (*model.ChangedWorklogPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Updated(ctx, since, expand)
}

// Add adds a worklog to an issue.
//
// Time tracking must be enabled in Jira, otherwise this operation returns an error.
//
// POST /rest/api/3/issue/{issueKeyOrID}/worklog
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#add-worklog
func (w *WorklogADFService) Add(ctx context.Context, issueKeyOrID string, payload *model.WorklogADFPayloadScheme, options *model.WorklogOptionsScheme) (*model.IssueWorklogADFScheme, *model.ResponseScheme, error) {
	return w.internalClient.Add(ctx, issueKeyOrID, payload, options)
}

// Update updates a worklog.
//
// Time tracking must be enabled in Jira, otherwise this operation returns an error.
//
// PUT /rest/api/3/issue/{issueKeyOrID}/worklog/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/worklogs#update-worklog
func (w *WorklogADFService) Update(ctx context.Context, issueKeyOrID, worklogID string, payload *model.WorklogADFPayloadScheme, options *model.WorklogOptionsScheme) (*model.IssueWorklogADFScheme, *model.ResponseScheme, error) {
	return w.internalClient.Update(ctx, issueKeyOrID, worklogID, payload, options)
}

type internalWorklogAdfImpl struct {
	c       service.Connector
	version string
}

func (i *internalWorklogAdfImpl) Gets(ctx context.Context, worklogIDs []int, expand []string) ([]*model.IssueWorklogADFScheme, *model.ResponseScheme, error) {

	if len(worklogIDs) == 0 {
		return nil, nil, model.ErrNpWorklogs
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/worklog/list", i.version))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", map[string]interface{}{"ids": worklogIDs})
	if err != nil {
		return nil, nil, err
	}

	var worklogs []*model.IssueWorklogADFScheme
	response, err := i.c.Call(request, &worklogs)
	if err != nil {
		return nil, response, err
	}

	return worklogs, response, nil
}

func (i *internalWorklogAdfImpl) Get(ctx context.Context, issueKeyOrID, worklogID string, expand []string) (*model.IssueWorklogADFScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	if worklogID == "" {
		return nil, nil, model.ErrNoWorklogID
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/issue/%v/worklog/%v", i.version, issueKeyOrID, worklogID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	worklog := new(model.IssueWorklogADFScheme)
	response, err := i.c.Call(request, worklog)
	if err != nil {
		return nil, response, err
	}

	return worklog, response, nil
}

func (i *internalWorklogAdfImpl) Issue(ctx context.Context, issueKeyOrID string, startAt, maxResults, after int, expand []string) (*model.IssueWorklogADFPageScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if after != 0 {
		params.Add("startedAfter", strconv.Itoa(after))
	}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/worklog?%v", i.version, issueKeyOrID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	worklogs := new(model.IssueWorklogADFPageScheme)
	response, err := i.c.Call(request, worklogs)
	if err != nil {
		return nil, response, err
	}

	return worklogs, response, nil
}

func (i *internalWorklogAdfImpl) Delete(ctx context.Context, issueKeyOrID, worklogID string, options *model.WorklogOptionsScheme) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrID
	}

	if worklogID == "" {
		return nil, model.ErrNoWorklogID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/issue/%v/worklog/%v", i.version, issueKeyOrID, worklogID))

	if options != nil {

		params := url.Values{}
		params.Add("notifyUsers", fmt.Sprintf("%v", options.Notify))
		params.Add("overrideEditableFlag", fmt.Sprintf("%v", options.OverrideEditableFlag))

		if options.AdjustEstimate != "" {
			params.Add("adjustEstimate", options.AdjustEstimate)
		}

		if options.NewEstimate != "" {
			params.Add("newEstimate", options.NewEstimate)
		}

		if options.ReduceBy != "" {
			params.Add("reduceBy", options.ReduceBy)
		}

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint.String(), "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalWorklogAdfImpl) Deleted(ctx context.Context, since int) (*model.ChangedWorklogPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	if since != 0 {
		params.Add("since", strconv.Itoa(since))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/worklog/deleted", i.version))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	worklogs := new(model.ChangedWorklogPageScheme)
	response, err := i.c.Call(request, worklogs)
	if err != nil {
		return nil, response, err
	}

	return worklogs, response, nil
}

func (i *internalWorklogAdfImpl) Updated(ctx context.Context, since int, expand []string) (*model.ChangedWorklogPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	if since != 0 {
		params.Add("since", strconv.Itoa(since))
	}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/worklog/updated", i.version))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	worklogs := new(model.ChangedWorklogPageScheme)
	response, err := i.c.Call(request, worklogs)
	if err != nil {
		return nil, response, err
	}

	return worklogs, response, nil
}

func (i *internalWorklogAdfImpl) Add(ctx context.Context, issueKeyOrID string, payload *model.WorklogADFPayloadScheme, options *model.WorklogOptionsScheme) (*model.IssueWorklogADFScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/issue/%v/worklog", i.version, issueKeyOrID))

	if options != nil {

		params := url.Values{}

		params.Add("notifyUsers", fmt.Sprintf("%v", options.Notify))
		params.Add("overrideEditableFlag", fmt.Sprintf("%v", options.OverrideEditableFlag))

		if options.AdjustEstimate != "" {
			params.Add("adjustEstimate", options.AdjustEstimate)
		}

		if options.NewEstimate != "" {
			params.Add("newEstimate", options.NewEstimate)
		}

		if options.ReduceBy != "" {
			params.Add("reduceBy", options.ReduceBy)
		}

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	worklog := new(model.IssueWorklogADFScheme)
	response, err := i.c.Call(request, worklog)
	if err != nil {
		return nil, response, err
	}

	return worklog, response, nil
}

func (i *internalWorklogAdfImpl) Update(ctx context.Context, issueKeyOrID, worklogID string, payload *model.WorklogADFPayloadScheme, options *model.WorklogOptionsScheme) (*model.IssueWorklogADFScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	if worklogID == "" {
		return nil, nil, model.ErrNoWorklogID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/issue/%v/worklog/%v", i.version, issueKeyOrID, worklogID))

	if options != nil {

		params := url.Values{}

		params.Add("notifyUsers", fmt.Sprintf("%v", options.Notify))
		params.Add("overrideEditableFlag", fmt.Sprintf("%v", options.OverrideEditableFlag))

		if options.AdjustEstimate != "" {
			params.Add("adjustEstimate", options.AdjustEstimate)
		}

		if options.NewEstimate != "" {
			params.Add("newEstimate", options.NewEstimate)
		}

		if options.ReduceBy != "" {
			params.Add("reduceBy", options.ReduceBy)
		}

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	worklog := new(model.IssueWorklogADFScheme)
	response, err := i.c.Call(request, worklog)
	if err != nil {
		return nil, response, err
	}

	return worklog, response, nil
}
