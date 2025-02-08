package jira

import (
	"context"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// BulkOperationsConnector defines the interface for performing bulk operations in Jira.
type BulkOperationsConnector interface {

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
	Delete(ctx context.Context, issueKeysOrIDs []string, sendNotification bool) (task *models.SubmittedBulkOperationScheme, response *models.ResponseScheme, err error)

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
	GetFields(ctx context.Context, issueKeyOrIDs []string, searchText, endingBefore, startingAfter string) (fields *models.BulkEditGetFieldsScheme, response *models.ResponseScheme, err error)

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
	Edit(ctx context.Context, payload *models.IssueBulkEditPayloadScheme) (task *models.SubmittedBulkOperationScheme, response *models.ResponseScheme, err error)

	// GetTransitions retrieves available transitions for multiple issues in Jira.
	//
	// This method returns a list of possible transitions that can be applied to the provided issues.
	// The API supports cursor-based pagination for navigating results.
	//
	// GET /rest/api/{2-3}/bulk/issues/transition
	//
	// Method Documentation:
	// https://docs.go-atlassian.io/jira-software-cloud/issues/bulk#get-bulk-available-transitions
	GetTransitions(ctx context.Context, issueKeysOrIDs []string, startAt, cursor string) (transitions *models.BulkTransitionGetAvailableTransitionsScheme, response *models.ResponseScheme, err error)

	// Transition applies transitions to multiple issues in Jira.
	//
	// This method allows performing bulk transitions on a list of issues by submitting the required
	// transition inputs. The transition process is asynchronous, and a task is returned to track the operation.
	//
	// POST /rest/api/{2-3}/bulk/issues/transition
	//
	// Method Documentation:
	// https://docs.go-atlassian.io/jira-software-cloud/issues/bulk#bulk-transition-issues
	Transition(ctx context.Context, inputs []*models.BulkTransitionSubmitInputScheme, sendNotification bool) (task *models.SubmittedBulkOperationScheme, response *models.ResponseScheme, err error)

	// GetStatus retrieves the status of a bulk operation task in Jira.
	//
	// This method allows checking the progress of a previously submitted bulk operation.
	//
	// GET /rest/api/{2-3}/bulk/queue/{taskId}
	//
	// Method Documentation:
	// https://docs.go-atlassian.io/jira-software-cloud/issues/bulk#get-bulk-operation-status
	GetStatus(ctx context.Context, taskID string) (status *models.BulkOperationProgressScheme, response *models.ResponseScheme, err error)
}
