package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
	"net/http"
	"net/url"
	"strings"
)

func NewBulkOperationsService(client service.Connector, version string) *BulkOperationsService {

	return &BulkOperationsService{
		internalClient: &internalBulkOperationsImpl{
			c:       client,
			version: version,
		},
	}
}

type BulkOperationsService struct {
	internalClient jira.BulkOperationsConnector
}

// Delete executes a bulk delete operation for the specified issues in Jira.
//
// This method allows deleting up to 1000 issues in a single bulk operation. The deletion process is asynchronous,
// and a task is returned to track the operation's progress. Optionally, notifications can be sent to affected users.
//
// POST /rest/api/{2-3}/bulk/issues/delete
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellations.
//   - issueKeysOrIDs: A slice of issue keys or IDs to be deleted in bulk.
//   - sendNotification: A boolean flag indicating whether notifications should be sent.
//
// Returns:
//   - task: A SubmittedBulkOperationScheme containing details about the bulk delete operation.
//   - response: The API response metadata.
//   - err: An error if the request fails or the API response is invalid.
//
// Example Usage:
//
//	task, resp, err := connector.Delete(ctx, []string{"ISSUE-123", "ISSUE-456"}, true)
//	if err != nil {
//	    log.Fatalf("Error deleting issues: %v", err)
//	}
//	fmt.Println("Bulk delete task ID:", task.ID)
//
// Notes:
//   - The maximum number of issues allowed per request is 1000.
//   - Deletion is permanent and cannot be undone.
//   - The returned task ID can be used to track the operation's progress.
//
// Method Documentation:
// https://docs.go-atlassian.io/jira-software-cloud/issues/bulk#bulk-delete-issues
func (b *BulkOperationsService) Delete(ctx context.Context, issueKeysOrIDs []string, sendNotification bool) (*model.SubmittedBulkOperationScheme, *model.ResponseScheme, error) {
	return b.internalClient.Delete(ctx, issueKeysOrIDs, sendNotification)
}

// Edit performs a bulk edit operation on the specified issues in Jira.
//
// This method allows updating multiple issues in a single request by modifying their fields
// and applying specific actions. The edit process is asynchronous, and a task is returned
// to track the progress of the operation.
//
// POST /rest/api/{2-3}/bulk/issues/fields
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellations.
//   - payload: An IssueBulkEditPayloadScheme struct containing issue fields, selected actions, issue IDs, and notification settings.
//
// Returns:
//   - task: A SubmittedBulkOperationScheme containing details about the bulk edit operation.
//   - response: The API response metadata.
//   - err: An error if the request fails or the API response is invalid.
//
// Jira API Documentation:
// https://docs.go-atlassian.io/jira-software-cloud/issues/bulk#bulk-edit-issues
func (b *BulkOperationsService) Edit(ctx context.Context, payload *model.IssueBulkEditPayloadScheme) (*model.SubmittedBulkOperationScheme, *model.ResponseScheme, error) {
	return b.internalClient.Edit(ctx, payload)
}

// GetFields retrieves a list of editable fields available for bulk edit operations in Jira.
//
// This method allows querying one or multiple issues to determine which fields the user can modify in bulk edits.
// The API supports filtering fields based on a search text and provides cursor-based pagination for navigating results.
//
// GET /rest/api/{2-3}/bulk/issues/fields
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellations.
//   - issueKeyOrIDs: A slice of issue keys or IDs to check for editable fields.
//   - searchText: An optional filter string to narrow down field results.
//   - endingBefore: An optional cursor to fetch results before the specified field ID.
//   - startingAfter: An optional cursor to fetch results after the specified field ID.
//
// Returns:
//   - fields: A BulkEditGetFieldsScheme containing a list of editable fields.
//   - response: The API response metadata.
//   - err: An error if the request fails or the API response is invalid.
//
// Example Usage:
//
//	fields, resp, err := connector.GetFields(ctx, []string{"ISSUE-123", "ISSUE-456"}, "", "", "")
//	if err != nil {
//	    log.Fatalf("Error retrieving editable fields: %v", err)
//	}
//	fmt.Println("Editable Fields:", fields)
//
// Notes:
//   - The method ensures user permissions before returning editable fields.
//   - Cursor-based pagination allows efficient traversal of large field lists.
//   - Use the `searchText` parameter to filter fields dynamically.
//
// Method Documentation:
// https://docs.go-atlassian.io/jira-software-cloud/issues/bulk#bulk-edit-get-fields
func (b *BulkOperationsService) GetFields(ctx context.Context, issueKeyOrIDs []string, searchText, endingBefore, startingAfter string) (*model.BulkEditGetFieldsScheme, *model.ResponseScheme, error) {
	return b.internalClient.GetFields(ctx, issueKeyOrIDs, searchText, endingBefore, startingAfter)
}

// GetStatus retrieves the status of a bulk operation task in Jira.
//
// This method allows checking the progress of a previously submitted bulk operation.
//
// GET /rest/api/{2-3}/bulk/queue/{taskId}
//
// Method Documentation:
// https://docs.go-atlassian.io/jira-software-cloud/issues/bulk#get-bulk-operation-status
func (b *BulkOperationsService) GetStatus(ctx context.Context, taskID string) (*model.BulkOperationProgressScheme, *model.ResponseScheme, error) {
	return b.internalClient.GetStatus(ctx, taskID)
}

// GetTransitions retrieves available transitions for multiple issues in Jira.
//
// This method returns a list of possible transitions that can be applied to the provided issues.
// The API supports cursor-based pagination for navigating results.
//
// GET /rest/api/{2-3}/bulk/issues/transition
//
// Method Documentation:
// https://docs.go-atlassian.io/jira-software-cloud/issues/bulk#get-bulk-available-transitions
func (b *BulkOperationsService) GetTransitions(ctx context.Context, issueKeysOrIDs []string, startAt, cursor string) (*model.BulkTransitionGetAvailableTransitionsScheme, *model.ResponseScheme, error) {
	return b.internalClient.GetTransitions(ctx, issueKeysOrIDs, startAt, cursor)
}

// Transition applies transitions to multiple issues in Jira.
//
// This method allows performing bulk transitions on a list of issues by submitting the required
// transition inputs. The transition process is asynchronous, and a task is returned to track the operation.
//
// POST /rest/api/{2-3}/bulk/issues/transition
//
// Method Documentation:
// https://docs.go-atlassian.io/jira-software-cloud/issues/bulk#bulk-transition-issues
func (b *BulkOperationsService) Transition(ctx context.Context, inputs []*model.BulkTransitionSubmitInputScheme, sendNotification bool) (*model.SubmittedBulkOperationScheme, *model.ResponseScheme, error) {
	return b.internalClient.Transition(ctx, inputs, sendNotification)
}

type internalBulkOperationsImpl struct {
	c       service.Connector
	version string
}

func (i *internalBulkOperationsImpl) Edit(ctx context.Context, payload *model.IssueBulkEditPayloadScheme) (task *model.SubmittedBulkOperationScheme, response *model.ResponseScheme, err error) {

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%s/bulk/issues/fields", i.version))

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	taskResult := new(model.SubmittedBulkOperationScheme)
	response, err = i.c.Call(request, taskResult)
	if err != nil {
		return nil, response, err
	}

	return taskResult, response, nil
}

func (i *internalBulkOperationsImpl) Delete(ctx context.Context, issueKeysOrIDs []string, sendNotification bool) (task *model.SubmittedBulkOperationScheme, response *model.ResponseScheme, err error) {

	if len(issueKeysOrIDs) == 0 {
		return nil, nil, model.ErrNoIssuesSlice
	}

	payload := map[string]interface{}{
		"selectedIssueIdsOrKeys": issueKeysOrIDs,
		"sendBulkNotification":   sendNotification,
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%s/bulk/issues/delete", i.version))

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	taskResult := new(model.SubmittedBulkOperationScheme)
	response, err = i.c.Call(request, taskResult)
	if err != nil {
		return nil, response, err
	}

	return taskResult, response, nil
}

func (i *internalBulkOperationsImpl) GetFields(ctx context.Context, issueKeyOrIDs []string, searchText, endingBefore, startingAfter string) (fields *model.BulkEditGetFieldsScheme, response *model.ResponseScheme, err error) {

	if len(issueKeyOrIDs) == 0 {
		return nil, nil, model.ErrNoIssuesSlice
	}

	params := url.Values{}
	if searchText != "" {
		params.Add("searchText", searchText)
	}
	if endingBefore != "" {
		params.Add("endingBefore", endingBefore)
	}
	if startingAfter != "" {
		params.Add("startingAfter", startingAfter)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%s/bulk/edit/fields", i.version))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	fieldsResult := new(model.BulkEditGetFieldsScheme)
	response, err = i.c.Call(request, fieldsResult)
	if err != nil {
		return nil, response, err
	}

	return fieldsResult, response, nil
}

func (i *internalBulkOperationsImpl) GetTransitions(ctx context.Context, issueKeysOrIDs []string, startAt, cursor string) (transitions *model.BulkTransitionGetAvailableTransitionsScheme, response *model.ResponseScheme, err error) {

	if len(issueKeysOrIDs) == 0 {
		return nil, nil, model.ErrNoIssuesSlice
	}

	params := url.Values{}
	if startAt != "" {
		params.Add("startAt", startAt)
	}

	if cursor != "" {
		params.Add("cursor", cursor)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%s/bulk/issues/transition", i.version))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	transitionsResult := new(model.BulkTransitionGetAvailableTransitionsScheme)
	response, err = i.c.Call(request, transitionsResult)
	if err != nil {
		return nil, response, err
	}

	return transitionsResult, response, nil
}

func (i *internalBulkOperationsImpl) Transition(ctx context.Context, inputs []*model.BulkTransitionSubmitInputScheme, sendNotification bool) (task *model.SubmittedBulkOperationScheme, response *model.ResponseScheme, err error) {

	payload := map[string]interface{}{
		"transitions":      inputs,
		"sendNotification": sendNotification,
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%s/bulk/issues/transition", i.version))

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	taskResult := new(model.SubmittedBulkOperationScheme)
	response, err = i.c.Call(request, taskResult)
	if err != nil {
		return nil, response, err
	}

	return taskResult, response, nil
}

func (i *internalBulkOperationsImpl) GetStatus(ctx context.Context, taskID string) (status *model.BulkOperationProgressScheme, response *model.ResponseScheme, err error) {

	if taskID == "" {
		return nil, nil, model.ErrNoTaskID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%s/bulk/queue/%v", i.version, taskID))

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	statusResult := new(model.BulkOperationProgressScheme)
	response, err = i.c.Call(request, statusResult)
	if err != nil {
		return nil, response, err
	}

	return statusResult, response, nil

}
