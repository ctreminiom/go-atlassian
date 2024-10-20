package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type VersionConnector interface {

	// Gets returns the versions for a piece of content in descending order.
	//
	// GET /wiki/rest/api/content/{id}/version
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/versions#get-content-versions
	Gets(ctx context.Context, contentID string, expand []string, start, limit int) (*model.ContentVersionPageScheme, *model.ResponseScheme, error)

	// Get returns a version for a piece of content.
	//
	// GET /wiki/rest/api/content/{id}/version/{versionNumber}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/versions#get-content-version
	Get(ctx context.Context, contentID string, versionNumber int, expand []string) (*model.ContentVersionScheme, *model.ResponseScheme, error)

	// Restore restores a historical version to be the latest version.
	//
	// That is, a new version is created with the content of the historical version.
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/versions#restore-content-version
	Restore(ctx context.Context, contentID string, payload *model.ContentRestorePayloadScheme, expand []string) (*model.ContentVersionScheme, *model.ResponseScheme, error)

	// Delete deletes a historical version.
	//
	// This does not delete the changes made to the content in that version, rather the changes for the deleted version
	//
	// are rolled up into the next version. Note, you cannot delete the current version.
	//
	// DELETE /wiki/rest/api/content/{id}/version/{versionNumber}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/versions#delete-content-version
	Delete(ctx context.Context, contentID string, versionNumber int) (*model.ResponseScheme, error)
}
