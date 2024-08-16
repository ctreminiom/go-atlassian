package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/tidwall/gjson"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
)

// NewMetadataService creates a new instance of MetadataService.
func NewMetadataService(client service.Connector, version string) (*MetadataService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &MetadataService{
		internalClient: &internalMetadataImpl{c: client, version: version},
	}, nil
}

// MetadataService provides methods to manage metadata in Jira Service Management.
type MetadataService struct {
	// internalClient is the connector interface for metadata operations.
	internalClient jira.MetadataConnector
}

// Get edit issue metadata returns the edit screen fields for an issue that are visible to and editable by the user.
//
// Use the information to populate the requests in Edit issue.
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}/editmeta
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/metadata#get-edit-issue-metadata
func (m *MetadataService) Get(ctx context.Context, issueKeyOrID string, overrideScreenSecurity, overrideEditableFlag bool) (gjson.Result, *model.ResponseScheme, error) {
	return m.internalClient.Get(ctx, issueKeyOrID, overrideScreenSecurity, overrideEditableFlag)
}

// Create returns details of projects, issue types within projects, and, when requested,
//
// the create screen fields for each issue type for the user.
//
// GET /rest/api/{2-3}/issue/createmeta
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/metadata#get-create-issue-metadata
func (m *MetadataService) Create(ctx context.Context, opts *model.IssueMetadataCreateOptions) (gjson.Result, *model.ResponseScheme, error) {
	return m.internalClient.Create(ctx, opts)
}

type internalMetadataImpl struct {
	c       service.Connector
	version string
}

func (i *internalMetadataImpl) Get(ctx context.Context, issueKeyOrID string, overrideScreenSecurity, overrideEditableFlag bool) (gjson.Result, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return gjson.Result{}, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("overrideEditableFlag", fmt.Sprintf("%v", overrideEditableFlag))
	params.Add("overrideScreenSecurity", fmt.Sprintf("%v", overrideScreenSecurity))

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/editmeta?%v", i.version, issueKeyOrID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return gjson.Result{}, nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		return gjson.Result{}, response, err
	}

	return gjson.ParseBytes(response.Bytes.Bytes()), response, nil
}

func (i *internalMetadataImpl) Create(ctx context.Context, opts *model.IssueMetadataCreateOptions) (gjson.Result, *model.ResponseScheme, error) {

	params := url.Values{}

	for _, id := range opts.IssueTypeIDs {
		params.Add("issuetypeIds", id)
	}

	for _, name := range opts.IssueTypeNames {
		params.Add("issuetypeNames", name)
	}

	for _, id := range opts.ProjectIDs {
		params.Add("projectIds", id)
	}

	for _, key := range opts.ProjectKeys {
		params.Add("projectKeys", key)
	}

	if opts.Expand != "" {
		params.Add("expand", opts.Expand)
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/createmeta?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return gjson.Result{}, nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		return gjson.Result{}, response, err
	}

	return gjson.ParseBytes(response.Bytes.Bytes()), response, nil
}
