package jira

import (
	"context"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// ArchiveService provides methods to manage issue archival operations, including preserving, restoring, and exporting archived issues.
type ArchiveService interface {

	// Preserve archives the given issues based on their issue IDs or keys.
	//
	// Parameters:
	//   - ctx: The context for controlling request lifecycle and deadlines.
	//   - issueIdsOrKeys: A list of issue IDs or keys to be archived.
	//
	// Returns:
	//   - result: A structure containing details of the archival synchronization process.
	//   - response: The HTTP response scheme for the request.
	//   - err: An error if the operation fails.
	//
	// Example Usage:
	//   result, response, err := issue.archive.Preserve(ctx, []string{"ISSUE-123", "ISSUE-456"})
	//
	// https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-application-role
	Preserve(ctx context.Context, issueIdsOrKeys []string) (result *models.IssueArchivalSyncResponseScheme, response *models.ResponseScheme, err error)

	// PreserveByJQL archives issues that match the provided JQL query.
	//
	// Parameters:
	//   - ctx: The context for request lifecycle management.
	//   - jql: The JQL query to select issues for archival.
	//
	// Returns:
	//   - taskID: A unique identifier for the asynchronous archival task.
	//   - response: The HTTP response scheme for the request.
	//   - err: An error if the operation fails.
	//
	// Example Usage:
	//   taskID, response, err := issue.Archive.PreserveByJQL(ctx, "project = ABC AND status = 'Resolved'")
	//
	// https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-application-role
	PreserveByJQL(ctx context.Context, jql string) (taskID string, response *models.ResponseScheme, err error)

	// Restore brings back the given archived issues using their issue IDs or keys.
	//
	// Parameters:
	//   - ctx: The context for controlling request execution.
	//   - issueIdsOrKeys: A list of issue IDs or keys to be restored from the archive.
	//
	// Returns:
	//   - result: A structure containing details of the restoration process.
	//   - response: The HTTP response scheme for the request.
	//   - err: An error if the operation fails.
	//
	// Example Usage:
	//   result, response, err := issue.Archive.Restore(ctx, []string{"ISSUE-789"})
	//
	// https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-application-role
	Restore(ctx context.Context, issueIdsOrKeys []string) (result *models.IssueArchivalSyncResponseScheme, response *models.ResponseScheme, err error)

	// Export generates an export of archived issues based on the provided payload.
	//
	// Parameters:
	//   - ctx: The context for controlling request execution.
	//   - payload: The export configuration, including filters and format specifications.
	//
	// Returns:
	//   - taskID: A unique identifier for the asynchronous export task.
	//   - response: The HTTP response scheme for the request.
	//   - err: An error if the operation fails.
	//
	// Example Usage:
	//   exportPayload := &models.IssueArchivalExportPayloadScheme{Format: "CSV", Fields: []string{"summary", "status"}}
	//   taskID, response, err := issue.Archive.Export(ctx, exportPayload)
	//
	// https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-application-role
	Export(ctx context.Context, payload *models.IssueArchivalExportPayloadScheme) (taskID string, response *models.ResponseScheme, err error)
}
