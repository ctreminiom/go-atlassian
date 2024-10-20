package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/tidwall/gjson"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
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
// Deprecated. Please use Issue.Metadata.FetchIssueMappings() and Issue.Metadata.FetchFieldMappings() instead.
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
// Deprecated. Please use Issue.Metadata.FetchIssueMappings() and Issue.Metadata.FetchFieldMappings() instead.
//
// the create screen fields for each issue type for the user.
//
// GET /rest/api/{2-3}/issue/createmeta
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/metadata#get-create-issue-metadata
func (m *MetadataService) Create(ctx context.Context, opts *model.IssueMetadataCreateOptions) (gjson.Result, *model.ResponseScheme, error) {
	return m.internalClient.Create(ctx, opts)
}

// FetchIssueMappings returns a page of issue type metadata for a specified project.
//
// Use the information to populate the requests in Create issue and Create issues.
//
// This operation can be accessed anonymously.
//
// GET /rest/api/{2-3}/issue/createmeta/{projectIdOrKey}/issuetypes
//
// Parameters:
// - ctx: The context for the request.
// - projectKeyOrID: The key or ID of the project.
// - startAt: The starting index of the returned issues.
// - maxResults: The maximum number of results to return.
//
// Returns:
// - A gjson.Result containing the issue type metadata.
// - A pointer to the response scheme.
// - An error if the retrieval fails.
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/metadata#get-create-metadata-issue-types-for-a-project
func (m *MetadataService) FetchIssueMappings(ctx context.Context, projectKeyOrID string, startAt, maxResults int) (gjson.Result, *model.ResponseScheme, error) {
	return m.internalClient.FetchIssueMappings(ctx, projectKeyOrID, startAt, maxResults)
}

// FetchFieldMappings returns a page of field metadata for a specified project and issue type.
//
// Use the information to populate the requests in Create issue and Create issues.
//
// This operation can be accessed anonymously.
//
// GET /rest/api/{2-3}/issue/createmeta/{projectIdOrKey}/fields/{issueTypeId}
//
// Parameters:
// - ctx: The context for the request.
// - projectKeyOrID: The key or ID of the project.
// - issueTypeID: The ID of the issue type whose metadata is to be retrieved.
// - startAt: The starting index of the returned fields.
// - maxResults: The maximum number of results to return.
//
// Returns:
// - A gjson.Result containing the field metadata.
// - A pointer to the response scheme.
// - An error if the retrieval fails.
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/metadata#get-create-field-metadata-for-a-project-and-issue-type-id
func (m *MetadataService) FetchFieldMappings(ctx context.Context, projectKeyOrID, issueTypeID string, startAt, maxResults int) (gjson.Result, *model.ResponseScheme, error) {
	return m.internalClient.FetchFieldMappings(ctx, projectKeyOrID, issueTypeID, startAt, maxResults)
}

type internalMetadataImpl struct {
	c       service.Connector
	version string
}

func (i *internalMetadataImpl) FetchIssueMappings(ctx context.Context, projectKeyOrID string, startAt, maxResults int) (gjson.Result, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return gjson.Result{}, nil, model.ErrNoProjectIDOrKey
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("rest/api/%v/issue/createmeta/%v/issuetypes?%v", i.version, projectKeyOrID, params.Encode())

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

func (i *internalMetadataImpl) FetchFieldMappings(ctx context.Context, projectKeyOrID, issueTypeID string, startAt, maxResults int) (gjson.Result, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return gjson.Result{}, nil, model.ErrNoProjectIDOrKey
	}

	if issueTypeID == "" {
		return gjson.Result{}, nil, model.ErrNoIssueTypeID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	endpoint := fmt.Sprintf("rest/api/%v/issue/createmeta/%v/issuetypes/%v?%v", i.version, projectKeyOrID, issueTypeID, params.Encode())

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

func (i *internalMetadataImpl) Get(ctx context.Context, issueKeyOrID string, overrideScreenSecurity, overrideEditableFlag bool) (gjson.Result, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return gjson.Result{}, nil, model.ErrNoIssueKeyOrID
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
