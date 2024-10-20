package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewIssueFieldConfigurationSchemeService creates a new instance of IssueFieldConfigSchemeService.
// It takes a service.Connector and a version string as input.
// Returns a pointer to IssueFieldConfigSchemeService and an error if the version is not provided.
func NewIssueFieldConfigurationSchemeService(client service.Connector, version string) (*IssueFieldConfigSchemeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssueFieldConfigSchemeService{
		internalClient: &internalIssueFieldConfigSchemeServiceImpl{c: client, version: version},
	}, nil
}

// IssueFieldConfigSchemeService provides methods to manage field configuration schemes in Jira Service Management.
type IssueFieldConfigSchemeService struct {
	// internalClient is the connector interface for field configuration scheme operations.
	internalClient jira.FieldConfigSchemeConnector
}

// Gets returns a paginated list of field configuration schemes.
//
// Only field configuration schemes used in classic projects are returned.
//
// GET /rest/api/{2-3}/fieldconfigurationscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#get-field-configuration-schemes
func (i *IssueFieldConfigSchemeService) Gets(ctx context.Context, ids []int, startAt, maxResults int) (*model.FieldConfigurationSchemePageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, ids, startAt, maxResults)
}

// Create creates a field configuration scheme.
//
// This operation can only create field configuration schemes used in company-managed (classic) projects.
//
// POST /rest/api/{2-3}/fieldconfigurationscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#create-field-configuration-scheme
func (i *IssueFieldConfigSchemeService) Create(ctx context.Context, name, description string) (*model.FieldConfigurationSchemeScheme, *model.ResponseScheme, error) {
	return i.internalClient.Create(ctx, name, description)
}

// Mapping returns a paginated list of field configuration issue type items.
//
// Only items used in classic projects are returned.
//
// GET /rest/api/{2-3}/fieldconfigurationscheme/mapping
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#get-field-configuration-scheme-mapping
func (i *IssueFieldConfigSchemeService) Mapping(ctx context.Context, fieldConfigIDs []int, startAt, maxResults int) (*model.FieldConfigurationIssueTypeItemPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Mapping(ctx, fieldConfigIDs, startAt, maxResults)
}

// Project returns a paginated list of field configuration schemes and, for each scheme, a list of the projects that use it.
//
// 1. The list is sorted by field configuration scheme ID. The first item contains the list of project IDs assigned to the default field configuration scheme.
//
// 2. Only field configuration schemes used in classic projects are returned.\
//
// GET /rest/api/{2-3}/fieldconfigurationscheme/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#get-field-configuration-schemes-by-project
func (i *IssueFieldConfigSchemeService) Project(ctx context.Context, projectIDs []int, startAt, maxResults int) (*model.FieldConfigurationSchemeProjectPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Project(ctx, projectIDs, startAt, maxResults)
}

// Assign assigns a field configuration scheme to a project. If the field configuration scheme ID is null,
//
// the operation assigns the default field configuration scheme.
//
// Field configuration schemes can only be assigned to classic projects.
//
// PUT /rest/api/{2-3}/fieldconfigurationscheme/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#assign-field-configuration-scheme
func (i *IssueFieldConfigSchemeService) Assign(ctx context.Context, payload *model.FieldConfigurationSchemeAssignPayload) (*model.ResponseScheme, error) {
	return i.internalClient.Assign(ctx, payload)
}

// Update updates a field configuration scheme.
//
// This operation can only update field configuration schemes used in company-managed (classic) projects.
//
// PUT /rest/api/{2-3}/fieldconfigurationscheme/{schemeID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#update-field-configuration-scheme
func (i *IssueFieldConfigSchemeService) Update(ctx context.Context, schemeID int, name, description string) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, schemeID, name, description)
}

// Delete deletes a field configuration scheme.
//
// This operation can only delete field configuration schemes used in company-managed (classic) projects.
//
// DELETE /rest/api/{2-3}/fieldconfigurationscheme/{schemeID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#delete-field-configuration-scheme
func (i *IssueFieldConfigSchemeService) Delete(ctx context.Context, schemeID int) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, schemeID)
}

// Link assigns issue types to field configurations on field configuration scheme.
//
// This operation can only modify field configuration schemes used in company-managed (classic) projects.
//
// PUT /rest/api/{2-3}/fieldconfigurationscheme/{schemeID}/mapping
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#assign-issue-types-to-field-configuration
func (i *IssueFieldConfigSchemeService) Link(ctx context.Context, schemeID int, payload *model.FieldConfigurationToIssueTypeMappingPayloadScheme) (*model.ResponseScheme, error) {
	return i.internalClient.Link(ctx, schemeID, payload)
}

// Unlink removes issue types from the field configuration scheme.
//
// This operation can only modify field configuration schemes used in company-managed (classic) projects.
//
// POST /rest/api/{2-3}/fieldconfigurationscheme/{schemeID}/mapping/delete
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#remove-issue-types-to-field-configuration
func (i *IssueFieldConfigSchemeService) Unlink(ctx context.Context, schemeID int, issueTypeIDs []string) (*model.ResponseScheme, error) {
	return i.internalClient.Unlink(ctx, schemeID, issueTypeIDs)
}

type internalIssueFieldConfigSchemeServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Gets(ctx context.Context, ids []int, startAt, maxResults int) (*model.FieldConfigurationSchemePageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range ids {
		params.Add("id", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	scheme := new(model.FieldConfigurationSchemePageScheme)
	response, err := i.c.Call(request, scheme)
	if err != nil {
		return nil, response, err
	}

	return scheme, response, nil
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Create(ctx context.Context, name, description string) (*model.FieldConfigurationSchemeScheme, *model.ResponseScheme, error) {

	if name == "" {
		return nil, nil, model.ErrNoFieldConfigurationSchemeName
	}

	payload := map[string]interface{}{"name": name}

	if description != "" {
		payload["description"] = description
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	scheme := new(model.FieldConfigurationSchemeScheme)
	response, err := i.c.Call(request, scheme)
	if err != nil {
		return nil, response, err
	}

	return scheme, response, nil
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Mapping(ctx context.Context, fieldConfigIDs []int, startAt, maxResults int) (*model.FieldConfigurationIssueTypeItemPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range fieldConfigIDs {
		params.Add("fieldConfigurationSchemeId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/mapping?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.FieldConfigurationIssueTypeItemPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Project(ctx context.Context, projectIDs []int, startAt, maxResults int) (*model.FieldConfigurationSchemeProjectPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, projectID := range projectIDs {
		params.Add("projectId", strconv.Itoa(projectID))
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/project?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.FieldConfigurationSchemeProjectPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Assign(ctx context.Context, payload *model.FieldConfigurationSchemeAssignPayload) (*model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Update(ctx context.Context, schemeID int, name, description string) (*model.ResponseScheme, error) {

	if schemeID == 0 {
		return nil, model.ErrNoFieldConfigurationSchemeID
	}

	if name == "" {
		return nil, model.ErrNoFieldConfigurationSchemeName
	}

	payload := map[string]interface{}{"name": name}

	if description != "" {
		payload["description"] = description
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/%v", i.version, schemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Delete(ctx context.Context, schemeID int) (*model.ResponseScheme, error) {

	if schemeID == 0 {
		return nil, model.ErrNoFieldConfigurationSchemeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/%v", i.version, schemeID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Link(ctx context.Context, schemeID int, payload *model.FieldConfigurationToIssueTypeMappingPayloadScheme) (*model.ResponseScheme, error) {

	if schemeID == 0 {
		return nil, model.ErrNoFieldConfigurationSchemeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/%v/mapping", i.version, schemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Unlink(ctx context.Context, schemeID int, issueTypeIDs []string) (*model.ResponseScheme, error) {

	if schemeID == 0 {
		return nil, model.ErrNoFieldConfigurationSchemeID
	}

	if len(issueTypeIDs) == 0 {
		return nil, model.ErrNoIssueTypes
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/%v/mapping/delete", i.version, schemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"issueTypeIds": issueTypeIDs})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
