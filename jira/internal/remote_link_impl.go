package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strings"
)

func NewRemoteLinkService(client service.Client, version string) (*RemoteLinkService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &RemoteLinkService{
		internalClient: &internalRemoteLinkImpl{c: client, version: version},
	}, nil
}

type RemoteLinkService struct {
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
// GET /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#get-remote-issue-links
func (r *RemoteLinkService) Gets(ctx context.Context, issueKeyOrId, globalId string) ([]*model.RemoteLinkScheme, *model.ResponseScheme, error) {
	return r.internalClient.Gets(ctx, issueKeyOrId, globalId)
}

// Get returns a remote issue link for an issue.
//
// GET /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink/{linkId}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#get-remote-issue-link
func (r *RemoteLinkService) Get(ctx context.Context, issueKeyOrId, linkId string) (*model.RemoteLinkScheme, *model.ResponseScheme, error) {
	return r.internalClient.Get(ctx, issueKeyOrId, linkId)
}

// Create creates or updates a remote issue link for an issue.
//
// If a globalId is provided and a remote issue link with that global ID is found it is updated.
//
// Any fields without values in the request are set to null. Otherwise, the remote issue link is created.
//
// POST /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#create-remote-issue-link
func (r *RemoteLinkService) Create(ctx context.Context, issueKeyOrId string, payload *model.RemoteLinkScheme) (*model.RemoteLinkIdentify, *model.ResponseScheme, error) {
	return r.internalClient.Create(ctx, issueKeyOrId, payload)
}

// Update updates a remote issue link for an issue.
//
// Note: Fields without values in the request are set to null.
//
// PUT /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink/{linkId}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#update-remote-issue-link
func (r *RemoteLinkService) Update(ctx context.Context, issueKeyOrId, linkId string, payload *model.RemoteLinkScheme) (*model.ResponseScheme, error) {
	return r.internalClient.Update(ctx, issueKeyOrId, linkId, payload)
}

// DeleteById deletes a remote issue link from an issue.
//
// DELETE /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink/{linkId}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#delete-remote-issue-link-by-id
func (r *RemoteLinkService) DeleteById(ctx context.Context, issueKeyOrId, linkId string) (*model.ResponseScheme, error) {
	return r.internalClient.DeleteById(ctx, issueKeyOrId, linkId)
}

// DeleteByGlobalId deletes the remote issue link from the issue using the link's global ID.
//
// Where the global ID includes reserved URL characters these must be escaped in the request.
//
// For example, pass system=http://www.mycompany.com/support&id=1 as system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1.
//
// DELETE /rest/api/{2-3}/issue/{issueIdOrKey}/remotelink
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/remote#delete-remote-issue-link-by-global-id
func (r *RemoteLinkService) DeleteByGlobalId(ctx context.Context, issueKeyOrId, globalId string) (*model.ResponseScheme, error) {
	return r.internalClient.DeleteByGlobalId(ctx, issueKeyOrId, globalId)
}

type internalRemoteLinkImpl struct {
	c       service.Client
	version string
}

func (i *internalRemoteLinkImpl) Gets(ctx context.Context, issueKeyOrId, globalId string) ([]*model.RemoteLinkScheme, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/issue/%v/remotelink", i.version, issueKeyOrId))

	if globalId != "" {

		params := url.Values{}
		params.Add("globalId", globalId)
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
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

func (i *internalRemoteLinkImpl) Get(ctx context.Context, issueKeyOrId, linkId string) (*model.RemoteLinkScheme, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if linkId == "" {
		return nil, nil, model.ErrNoRemoteLinkIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/remotelink/%v", i.version, issueKeyOrId, linkId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	remoteLink := new(model.RemoteLinkScheme)
	response, err := i.c.Call(request, &remoteLink)
	if err != nil {
		return nil, response, err
	}

	return remoteLink, response, nil
}

func (i *internalRemoteLinkImpl) Create(ctx context.Context, issueKeyOrId string, payload *model.RemoteLinkScheme) (*model.RemoteLinkIdentify, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/remotelink", i.version, issueKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	identify := new(model.RemoteLinkIdentify)
	response, err := i.c.Call(request, &identify)
	if err != nil {
		return nil, response, err
	}

	return identify, response, nil
}

func (i *internalRemoteLinkImpl) Update(ctx context.Context, issueKeyOrId, linkId string, payload *model.RemoteLinkScheme) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if linkId == "" {
		return nil, model.ErrNoRemoteLinkIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/remotelink/%v", i.version, issueKeyOrId, linkId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalRemoteLinkImpl) DeleteById(ctx context.Context, issueKeyOrId, linkId string) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if linkId == "" {
		return nil, model.ErrNoRemoteLinkIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/remotelink/%v", i.version, issueKeyOrId, linkId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalRemoteLinkImpl) DeleteByGlobalId(ctx context.Context, issueKeyOrId, globalId string) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if globalId == "" {
		return nil, model.ErrNoRemoteLinkGlobalIDError
	}

	params := url.Values{}
	params.Add("globalId", globalId)

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/remotelink?%v", i.version, issueKeyOrId, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
