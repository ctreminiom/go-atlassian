package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strconv"
)

func NewIssueFieldConfigurationSchemeService(client service.Client, version string) (*IssueFieldConfigSchemeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssueFieldConfigSchemeService{
		internalClient: &internalIssueFieldConfigSchemeServiceImpl{c: client, version: version},
	}, nil
}

type IssueFieldConfigSchemeService struct {
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
func (i *IssueFieldConfigSchemeService) Mapping(ctx context.Context, fieldConfigIds []int, startAt, maxResults int) (*model.FieldConfigurationIssueTypeItemPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Mapping(ctx, fieldConfigIds, startAt, maxResults)
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
func (i *IssueFieldConfigSchemeService) Project(ctx context.Context, projectIds []int, startAt, maxResults int) (*model.FieldConfigurationSchemeProjectPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Project(ctx, projectIds, startAt, maxResults)
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
// PUT /rest/api/{2-3}/fieldconfigurationscheme/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#update-field-configuration-scheme
func (i *IssueFieldConfigSchemeService) Update(ctx context.Context, schemeId int, name, description string) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, schemeId, name, description)
}

// Delete deletes a field configuration scheme.
//
// This operation can only delete field configuration schemes used in company-managed (classic) projects.
//
// DELETE /rest/api/{2-3}/fieldconfigurationscheme/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#delete-field-configuration-scheme
func (i *IssueFieldConfigSchemeService) Delete(ctx context.Context, schemeId int) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, schemeId)
}

// Link assigns issue types to field configurations on field configuration scheme.
//
// This operation can only modify field configuration schemes used in company-managed (classic) projects.
//
// PUT /rest/api/{2-3}/fieldconfigurationscheme/{id}/mapping
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#assign-issue-types-to-field-configuration
func (i *IssueFieldConfigSchemeService) Link(ctx context.Context, schemeId int, payload *model.FieldConfigurationToIssueTypeMappingPayloadScheme) (*model.ResponseScheme, error) {
	return i.internalClient.Link(ctx, schemeId, payload)
}

// Unlink removes issue types from the field configuration scheme.
//
// This operation can only modify field configuration schemes used in company-managed (classic) projects.
//
// POST /rest/api/{2-3}/fieldconfigurationscheme/{id}/mapping/delete
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#remove-issue-types-to-field-configuration
func (i *IssueFieldConfigSchemeService) Unlink(ctx context.Context, schemeId int, issueTypeIDs []string) (*model.ResponseScheme, error) {
	return i.internalClient.Unlink(ctx, schemeId, issueTypeIDs)
}

type internalIssueFieldConfigSchemeServiceImpl struct {
	c       service.Client
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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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
		return nil, nil, model.ErrNoFieldConfigurationSchemeNameError
	}

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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

func (i *internalIssueFieldConfigSchemeServiceImpl) Mapping(ctx context.Context, fieldConfigIds []int, startAt, maxResults int) (*model.FieldConfigurationIssueTypeItemPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range fieldConfigIds {
		params.Add("fieldConfigurationSchemeId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/mapping?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalIssueFieldConfigSchemeServiceImpl) Project(ctx context.Context, projectIds []int, startAt, maxResults int) (*model.FieldConfigurationSchemeProjectPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, projectID := range projectIds {
		params.Add("projectId", strconv.Itoa(projectID))
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/project?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Update(ctx context.Context, schemeId int, name, description string) (*model.ResponseScheme, error) {

	if schemeId == 0 {
		return nil, model.ErrNoFieldConfigurationSchemeIDError
	}

	if name == "" {
		return nil, model.ErrNoFieldConfigurationSchemeNameError
	}

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/%v", i.version, schemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Delete(ctx context.Context, schemeId int) (*model.ResponseScheme, error) {

	if schemeId == 0 {
		return nil, model.ErrNoFieldConfigurationSchemeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/%v", i.version, schemeId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Link(ctx context.Context, schemeId int, payload *model.FieldConfigurationToIssueTypeMappingPayloadScheme) (*model.ResponseScheme, error) {

	if schemeId == 0 {
		return nil, model.ErrNoFieldConfigurationSchemeIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/%v/mapping", i.version, schemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldConfigSchemeServiceImpl) Unlink(ctx context.Context, schemeId int, issueTypeIDs []string) (*model.ResponseScheme, error) {

	if schemeId == 0 {
		return nil, model.ErrNoFieldConfigurationSchemeIDError
	}

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{
		IssueTypeIds: issueTypeIDs,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfigurationscheme/%v/mapping/delete", i.version, schemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
