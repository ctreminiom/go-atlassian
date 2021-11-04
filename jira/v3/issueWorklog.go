package v3

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type IssueWorklogService struct{ client *Client }

// Get returns a worklog.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-worklogs/#api-rest-api-3-issue-issueidorkey-worklog-id-get
func (w *IssueWorklogService) Get(ctx context.Context, issueKeyOrID, worklogID string, expand []string) (result *models.IssueWorklogScheme,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models.ErrNoIssueKeyOrIDError
	}

	if len(worklogID) == 0 {
		return nil, nil, models.ErrNoWorklogIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/issue/%v/worklog/%v", issueKeyOrID, worklogID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Issue returns worklogs for an issue, starting from the oldest worklog or from the worklog started
// on or after a date and time.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-worklogs/#api-rest-api-3-issue-issueidorkey-worklog-get
func (w *IssueWorklogService) Issue(ctx context.Context, issueKeyOrID string, startAt, maxResults, after int,
	expand []string) (result *models.IssueWorklogPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models.ErrNoIssueKeyOrIDError
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

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/worklog?%v", issueKeyOrID, params.Encode())

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Add adds a worklog to an issue.
// Docs: N/A
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-worklogs/#api-rest-api-3-issue-issueidorkey-worklog-post
func (w *IssueWorklogService) Add(ctx context.Context, issueKeyOrID string, options *models.WorklogOptionsScheme) (
	result *models.IssueWorklogScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	if !options.Notify {
		params.Add("notifyUsers", "false")
	}

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

	if options.OverrideEditableFlag {
		params.Add("overrideEditableFlag", "true")
	}

	payloadAsReader, err := transformStructToReader(options.Payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/issue/%v/worklog", issueKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := w.client.newRequest(ctx, http.MethodPost, endpoint.String(), payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = w.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates a worklog.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-worklogs/#api-rest-api-3-issue-issueidorkey-worklog-id-put
func (w *IssueWorklogService) Update(ctx context.Context, issueKeyOrID, worklogID string, options *models.WorklogOptionsScheme) (
	result *models.IssueWorklogScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models.ErrNoIssueKeyOrIDError
	}

	if len(worklogID) == 0 {
		return nil, nil, models.ErrNoWorklogIDError
	}

	params := url.Values{}

	if !options.Notify {
		params.Add("notifyUsers", "false")
	}

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

	if options.OverrideEditableFlag {
		params.Add("overrideEditableFlag", "true")
	}

	payloadAsReader, err := transformStructToReader(options.Payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/issue/%v/worklog/%v", issueKeyOrID, worklogID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := w.client.newRequest(ctx, http.MethodPut, endpoint.String(), payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = w.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes a worklog from an issue.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-worklogs/#api-rest-api-3-issue-issueidorkey-worklog-id-delete
func (w *IssueWorklogService) Delete(ctx context.Context, issueKeyOrID, worklogID string, options *models.WorklogOptionsScheme) (
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, models.ErrNoIssueTypeIDError
	}

	if len(worklogID) == 0 {
		return nil, models.ErrNoWorklogIDError
	}

	params := url.Values{}
	if options != nil {

		if !options.Notify {
			params.Add("notifyUsers", "false")
		}

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

		if options.OverrideEditableFlag {
			params.Add("overrideEditableFlag", "true")
		}
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/issue/%v/worklog/%v", issueKeyOrID, worklogID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := w.client.newRequest(ctx, http.MethodDelete, endpoint.String(), nil)
	if err != nil {
		return
	}

	response, err = w.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Deleted returns a list of IDs and delete timestamps for worklogs deleted after a date and time.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-worklogs/#api-rest-api-3-worklog-deleted-get
func (w *IssueWorklogService) Deleted(ctx context.Context, since int) (result *models.ChangedWorklogPageScheme, response *ResponseScheme,
	err error) {

	params := url.Values{}
	if since != 0 {
		params.Add("since", strconv.Itoa(since))
	}

	var endpoint strings.Builder
	endpoint.WriteString("rest/api/3/worklog/deleted")

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Gets returns worklog details for a list of worklog IDs.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-worklogs/#api-rest-api-3-worklog-list-post
func (w *IssueWorklogService) Gets(ctx context.Context, worklogIDs []int, expand []string) (result []*models.IssueWorklogScheme,
	response *ResponseScheme, err error) {

	if len(worklogIDs) == 0 {
		return nil, nil, models.ErrNpWorklogsError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString("rest/api/3/worklog/list")

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	payload := struct {
		Ids []int `json:"ids"`
	}{
		Ids: worklogIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	request, err := w.client.newRequest(ctx, http.MethodPost, endpoint.String(), payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = w.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Updated returns a list of IDs and update timestamps for worklogs updated after a date and time.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-worklogs/#api-rest-api-3-worklog-updated-get
func (w *IssueWorklogService) Updated(ctx context.Context, since int, expand []string) (result *models.ChangedWorklogPageScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	if since != 0 {
		params.Add("since", strconv.Itoa(since))
	}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString("rest/api/3/worklog/updated")

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
