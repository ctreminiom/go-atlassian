package jira

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type RemoteLinkConnector interface {

	// Gets returns the remote issue links for an issue.
	//
	// When a remote issue link global ID is provided the record with that global ID is returned,
	//
	// otherwise all remote issue links are returned.
	//
	// Where a global ID includes reserved URL characters these must be escaped in the request
	//
	// GET /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#get-remote-issue-links
	Gets(ctx context.Context, issueKeyOrId, globalId string) ([]*models.RemoteLinkScheme, *models.ResponseScheme, error)

	// Get returns a remote issue link for an issue.
	//
	// GET /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink/{linkId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#get-remote-issue-link
	Get(ctx context.Context, issueKeyOrId, linkId string) (*models.RemoteLinkScheme, *models.ResponseScheme, error)

	// Create creates or updates a remote issue link for an issue.
	//
	// If a globalId is provided and a remote issue link with that global ID is found it is updated.
	//
	// Any fields without values in the request are set to null. Otherwise, the remote issue link is created.
	//
	// POST /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#create-remote-issue-link
	Create(ctx context.Context, issueKeyOrId string, payload *models.RemoteLinkScheme) (*models.RemoteLinkIdentify, *models.ResponseScheme, error)

	// Update updates a remote issue link for an issue.
	//
	// Note: Fields without values in the request are set to null.
	//
	// PUT /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink/{linkId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#update-remote-issue-link
	Update(ctx context.Context, issueKeyOrId, linkId string, payload *models.RemoteLinkScheme) (*models.ResponseScheme, error)

	// DeleteById deletes a remote issue link from an issue.
	//
	// DELETE /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink/{linkId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#delete-remote-issue-link-by-id
	DeleteById(ctx context.Context, issueKeyOrId, linkId string) (*models.ResponseScheme, error)

	// DeleteByGlobalId deletes the remote issue link from the issue using the link's global ID.
	//
	// Where the global ID includes reserved URL characters these must be escaped in the request.
	//
	// For example, pass system=http://www.mycompany.com/support&id=1 as system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1.
	//
	// DELETE /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#delete-remote-issue-link-by-global-id
	DeleteByGlobalId(ctx context.Context, issueKeyOrId, globalId string) (*models.ResponseScheme, error)
}
