// Package jira ...
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
	// https://docs.go-atlassian.io/jira-software-cloud/issues/archiving#archive-issues-by-issue-id-key
	Preserve(ctx context.Context, issueIDsOrKeys []string) (result *models.IssueArchivalSyncResponseScheme, response *models.ResponseScheme, err error)

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
	// https://docs.go-atlassian.io/jira-software-cloud/issues/archiving#archive-issues-by-jql
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
	// https://docs.go-atlassian.io/jira-software-cloud/issues/archiving#restore-issues-by-issue-id-key
	Restore(ctx context.Context, issueIDsOrKeys []string) (result *models.IssueArchivalSyncResponseScheme, response *models.ResponseScheme, err error)

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
	// https://docs.go-atlassian.io/jira-software-cloud/issues/archiving#export-archived-issues
	Export(ctx context.Context, payload *models.IssueArchivalExportPayloadScheme) (task *models.IssueArchiveExportResultScheme, response *models.ResponseScheme, err error)
}
