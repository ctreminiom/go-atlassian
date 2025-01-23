package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewRemoteLinkService creates a new instance of RemoteLinkService.
func NewRemoteLinkService(client service.Connector, version string) (*RemoteLinkService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &RemoteLinkService{
		internalClient: &internalRemoteLinkImpl{c: client, version: version},
	}, nil
}

// RemoteLinkService provides methods to manage remote issue links in Jira Service Management.
type RemoteLinkService struct {
	// internalClient is the connector interface for remote link operations.
	internalClient jira.RemoteLinkConnector
}

// Gets returns the remote issue links for an issue.
//
// When a remote issue link global ID is provided the record with that global ID is returned,
//
// otherwise all remote issue links are returned.
//
// # Where a global ID includes reserved URL characters these must be escaped in the request
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}/remotelink
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#get-remote-issue-links
func (r *RemoteLinkService) Gets(ctx context.Context, issueKeyOrID, globalID string) ([]*model.RemoteLinkScheme, *model.ResponseScheme, error) {
	return r.internalClient.Gets(ctx, issueKeyOrID, globalID)
}

// Get returns a remote issue link for an issue.
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}/remotelink/{linkID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#get-remote-issue-link
func (r *RemoteLinkService) Get(ctx context.Context, issueKeyOrID, linkID string) (*model.RemoteLinkScheme, *model.ResponseScheme, error) {
	return r.internalClient.Get(ctx, issueKeyOrID, linkID)
}

// Create creates or updates a remote issue link for an issue.
//
// If a globalID is provided and a remote issue link with that global ID is found it is updated.
//
// Any fields without values in the request are set to null. Otherwise, the remote issue link is created.
//
// POST /rest/api/{2-3}/issue/{issueKeyOrID}/remotelink
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#create-remote-issue-link
func (r *RemoteLinkService) Create(ctx context.Context, issueKeyOrID string, payload *model.RemoteLinkScheme) (*model.RemoteLinkIdentify, *model.ResponseScheme, error) {
	return r.internalClient.Create(ctx, issueKeyOrID, payload)
}

// Update updates a remote issue link for an issue.
//
// Note: Fields without values in the request are set to null.
//
// PUT /rest/api/{2-3}/issue/{issueKeyOrID}/remotelink/{linkID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#update-remote-issue-link
func (r *RemoteLinkService) Update(ctx context.Context, issueKeyOrID, linkID string, payload *model.RemoteLinkScheme) (*model.ResponseScheme, error) {
	return r.internalClient.Update(ctx, issueKeyOrID, linkID, payload)
}

// DeleteByID deletes a remote issue link from an issue.
//
// DELETE /rest/api/{2-3}/issue/{issueKeyOrID}/remotelink/{linkID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#delete-remote-issue-link-by-id
func (r *RemoteLinkService) DeleteByID(ctx context.Context, issueKeyOrID, linkID string) (*model.ResponseScheme, error) {
	return r.internalClient.DeleteByID(ctx, issueKeyOrID, linkID)
}

// DeleteByGlobalID deletes the remote issue link from the issue using the link's global ID.
//
// Where the global ID includes reserved URL characters these must be escaped in the request.
//
// For example, pass system=http://www.mycompany.com/support&id=1 as system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1.
//
// DELETE /rest/api/{2-3}/issue/{issueKeyOrID}/remotelink
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#delete-remote-issue-link-by-global-id
func (r *RemoteLinkService) DeleteByGlobalID(ctx context.Context, issueKeyOrID, globalID string) (*model.ResponseScheme, error) {
	return r.internalClient.DeleteByGlobalID(ctx, issueKeyOrID, globalID)
}

type internalRemoteLinkImpl struct {
	c       service.Connector
	version string
}

func (i *internalRemoteLinkImpl) Gets(ctx context.Context, issueKeyOrID, globalID string) ([]*model.RemoteLinkScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/issue/%v/remotelink", i.version, issueKeyOrID))

	if globalID != "" {

		params := url.Values{}
		params.Add("globalId", globalID)
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	var remoteLinks []*model.RemoteLinkScheme
	response, err := i.c.Call(request, &remoteLinks)
	if err != nil {
		return nil, response, err
	}

	return remoteLinks, response, nil
}

func (i *internalRemoteLinkImpl) Get(ctx context.Context, issueKeyOrID, linkID string) (*model.RemoteLinkScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	if linkID == "" {
		return nil, nil, model.ErrNoRemoteLinkID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/remotelink/%v", i.version, issueKeyOrID, linkID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	remoteLink := new(model.RemoteLinkScheme)
	response, err := i.c.Call(request, remoteLink)
	if err != nil {
		return nil, response, err
	}

	return remoteLink, response, nil
}

func (i *internalRemoteLinkImpl) Create(ctx context.Context, issueKeyOrID string, payload *model.RemoteLinkScheme) (*model.RemoteLinkIdentify, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/remotelink", i.version, issueKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	identify := new(model.RemoteLinkIdentify)
	response, err := i.c.Call(request, identify)
	if err != nil {
		return nil, response, err
	}

	return identify, response, nil
}

func (i *internalRemoteLinkImpl) Update(ctx context.Context, issueKeyOrID, linkID string, payload *model.RemoteLinkScheme) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrID
	}

	if linkID == "" {
		return nil, model.ErrNoRemoteLinkID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/remotelink/%v", i.version, issueKeyOrID, linkID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalRemoteLinkImpl) DeleteByID(ctx context.Context, issueKeyOrID, linkID string) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrID
	}

	if linkID == "" {
		return nil, model.ErrNoRemoteLinkID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/remotelink/%v", i.version, issueKeyOrID, linkID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalRemoteLinkImpl) DeleteByGlobalID(ctx context.Context, issueKeyOrID, globalID string) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrID
	}

	if globalID == "" {
		return nil, model.ErrNoRemoteLinkGlobalID
	}

	params := url.Values{}
	params.Add("globalId", globalID)

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/remotelink?%v", i.version, issueKeyOrID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
