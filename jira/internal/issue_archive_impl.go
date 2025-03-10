// Package internal jira/internal/issue_archive_impl.go
package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
	"net/http"
	"path"
)

// NewIssueArchivalService creates a new instance of IssueArchivalService.
// It initializes an internal client with the given service connector and API version,
// which will be used to perform issue archival operations.
//
// Parameters:
//   - client: The service connector used to communicate with the underlying API.
//   - version: The API version to be used by the archival service.
//
// Returns:
//   - A pointer to an IssueArchivalService configured with the provided client and version.
//
// Example usage:
//
//	client := myConnectorInstance // your implementation of service.Connector
//	version := "v3"
//	archiveService := NewIssueArchivalService(client, version)
func NewIssueArchivalService(client service.Connector, version string) *IssueArchivalService {
	return &IssueArchivalService{
		internalClient: &internalIssueArchivalImpl{c: client, version: version},
	}
}

// IssueArchivalService provides methods to manage issue archival operations, including preserving, restoring, and exporting archived issues.
type IssueArchivalService struct {
	internalClient jira.ArchiveService
}

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
func (i *IssueArchivalService) Preserve(ctx context.Context, issueIDsOrKeys []string) (*model.IssueArchivalSyncResponseScheme, *model.ResponseScheme, error) {
	return i.internalClient.Preserve(ctx, issueIDsOrKeys)
}

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
func (i *IssueArchivalService) PreserveByJQL(ctx context.Context, jql string) (string, *model.ResponseScheme, error) {
	return i.internalClient.PreserveByJQL(ctx, jql)
}

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
func (i *IssueArchivalService) Restore(ctx context.Context, issueIDsOrKeys []string) (*model.IssueArchivalSyncResponseScheme, *model.ResponseScheme, error) {
	return i.internalClient.Restore(ctx, issueIDsOrKeys)
}

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
func (i *IssueArchivalService) Export(ctx context.Context, payload *model.IssueArchivalExportPayloadScheme) (*model.IssueArchiveExportResultScheme, *model.ResponseScheme, error) {
	return i.internalClient.Export(ctx, payload)
}

type internalIssueArchivalImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueArchivalImpl) Preserve(ctx context.Context, issueIDsOrKeys []string) (result *model.IssueArchivalSyncResponseScheme, response *model.ResponseScheme, err error) {

	if len(issueIDsOrKeys) == 0 {
		return nil, nil, model.ErrNoIssuesSlice
	}

	payload := make(map[string]interface{})
	payload["issueIdsOrKeys"] = issueIDsOrKeys

	endpoint := fmt.Sprintf("rest/api/%s/issue/archive", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	report := new(model.IssueArchivalSyncResponseScheme)
	response, err = i.c.Call(request, report)
	if err != nil {
		return nil, response, err
	}

	return report, response, nil
}

func (i *internalIssueArchivalImpl) PreserveByJQL(ctx context.Context, jql string) (taskID string, response *model.ResponseScheme, err error) {

	if jql == "" {
		return "", nil, model.ErrNoJQL
	}

	payload := make(map[string]interface{})
	payload["jql"] = jql

	endpoint := fmt.Sprintf("rest/api/%s/issue/archive", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return "", nil, err
	}

	response, err = i.c.Call(request, nil)
	if err != nil {
		return "", response, err
	}

	return path.Base(response.Bytes.String()), response, nil
}

func (i *internalIssueArchivalImpl) Restore(ctx context.Context, issueIDsOrKeys []string) (result *model.IssueArchivalSyncResponseScheme, response *model.ResponseScheme, err error) {

	if len(issueIDsOrKeys) == 0 {
		return nil, nil, model.ErrNoIssuesSlice
	}

	payload := make(map[string]interface{})
	payload["issueIdsOrKeys"] = issueIDsOrKeys

	endpoint := fmt.Sprintf("rest/api/%s/issue/unarchive", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	report := new(model.IssueArchivalSyncResponseScheme)
	response, err = i.c.Call(request, report)
	if err != nil {
		return nil, response, err
	}

	return report, response, nil
}

func (i *internalIssueArchivalImpl) Export(ctx context.Context, payload *model.IssueArchivalExportPayloadScheme) (task *model.IssueArchiveExportResultScheme, response *model.ResponseScheme, err error) {

	endpoint := fmt.Sprintf("rest/api/%s/issues/archive/export", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	result := new(model.IssueArchiveExportResultScheme)
	response, err = i.c.Call(request, result)
	if err != nil {
		return nil, response, err
	}

	return result, response, nil
}
