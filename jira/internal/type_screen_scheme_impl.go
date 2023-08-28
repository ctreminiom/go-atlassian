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
	"strings"
)

func NewTypeScreenSchemeService(client service.Connector, version string) (*TypeScreenSchemeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &TypeScreenSchemeService{
		internalClient: &internalTypeScreenSchemeImpl{c: client, version: version},
	}, nil
}

type TypeScreenSchemeService struct {
	internalClient jira.TypeScreenSchemeConnector
}

// Gets returns a paginated list of issue type screen schemes.
//
// Only issue type screen schemes used in classic projects are returned.
//
// GET /rest/api/{2-3}/issuetypescreenscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#get-issue-type-screen-schemes
func (t *TypeScreenSchemeService) Gets(ctx context.Context, options *model.ScreenSchemeParamsScheme, startAt, maxResults int) (*model.IssueTypeScreenSchemePageScheme, *model.ResponseScheme, error) {
	return t.internalClient.Gets(ctx, options, startAt, maxResults)
}

// Create creates an issue type screen scheme.
//
// POST /rest/api/{2-3}/issuetypescreenscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#create-issue-type-screen-scheme
func (t *TypeScreenSchemeService) Create(ctx context.Context, payload *model.IssueTypeScreenSchemePayloadScheme) (*model.IssueTypeScreenScreenCreatedScheme, *model.ResponseScheme, error) {
	return t.internalClient.Create(ctx, payload)
}

// Assign assigns an issue type screen scheme to a project.
//
// Issue type screen schemes can only be assigned to classic projects.
//
// PUT /rest/api/{2-3}/issuetypescreenscheme/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#assign-issue-type-screen-scheme-to-project
func (t *TypeScreenSchemeService) Assign(ctx context.Context, issueTypeScreenSchemeId, projectId string) (*model.ResponseScheme, error) {
	return t.internalClient.Assign(ctx, issueTypeScreenSchemeId, projectId)
}

// Projects returns a paginated list of issue type screen schemes and,
// for each issue type screen scheme, a list of the projects that use it.
//
// GET /rest/api/{2-3}/issuetypescreenscheme/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#assign-issue-type-screen-scheme-to-project
func (t *TypeScreenSchemeService) Projects(ctx context.Context, projectIds []int, startAt, maxResults int) (*model.IssueTypeProjectScreenSchemePageScheme, *model.ResponseScheme, error) {
	return t.internalClient.Projects(ctx, projectIds, startAt, maxResults)
}

// Mapping returns a paginated list of issue type screen scheme items.
//
// Only issue type screen schemes used in classic projects are returned.
//
// GET /rest/api/{2-3}/issuetypescreenscheme/mapping
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#get-issue-type-screen-scheme-items
func (t *TypeScreenSchemeService) Mapping(ctx context.Context, issueTypeScreenSchemeIds []int, startAt, maxResults int) (*model.IssueTypeScreenSchemeMappingScheme, *model.ResponseScheme, error) {
	return t.internalClient.Mapping(ctx, issueTypeScreenSchemeIds, startAt, maxResults)
}

// Update updates an issue type screen scheme.
//
// PUT /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#update-issue-type-screen-scheme
func (t *TypeScreenSchemeService) Update(ctx context.Context, issueTypeScreenSchemeId, name, description string) (*model.ResponseScheme, error) {
	return t.internalClient.Update(ctx, issueTypeScreenSchemeId, name, description)
}

// Delete deletes an issue type screen scheme.
//
// DELETE /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#delete-issue-type-screen-scheme
func (t *TypeScreenSchemeService) Delete(ctx context.Context, issueTypeScreenSchemeId string) (*model.ResponseScheme, error) {
	return t.internalClient.Delete(ctx, issueTypeScreenSchemeId)
}

// Append appends issue type to screen scheme mappings to an issue type screen scheme.
//
// PUT /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#append-mappings-to-issue-type-screen-scheme
func (t *TypeScreenSchemeService) Append(ctx context.Context, issueTypeScreenSchemeId string, payload *model.IssueTypeScreenSchemePayloadScheme) (*model.ResponseScheme, error) {
	return t.internalClient.Append(ctx, issueTypeScreenSchemeId, payload)
}

// UpdateDefault updates the default screen scheme of an issue type screen scheme.
// The default screen scheme is used for all unmapped issue types.
//
// PUT /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping/default
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#update-issue-type-screen-scheme-default-screen-scheme
func (t *TypeScreenSchemeService) UpdateDefault(ctx context.Context, issueTypeScreenSchemeId, screenSchemeId string) (*model.ResponseScheme, error) {
	return t.internalClient.UpdateDefault(ctx, issueTypeScreenSchemeId, screenSchemeId)
}

// Remove removes issue type to screen scheme mappings from an issue type screen scheme.
//
// POST /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping/remove
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#remove-mappings-from-issue-type-screen-scheme
func (t *TypeScreenSchemeService) Remove(ctx context.Context, issueTypeScreenSchemeId string, issueTypeIds []string) (*model.ResponseScheme, error) {
	return t.internalClient.Remove(ctx, issueTypeScreenSchemeId, issueTypeIds)
}

// SchemesByProject returns a paginated list of projects associated with an issue type screen scheme.
//
// GET /rest/api/{2-3}/issuetypescreenscheme/{issueTypeScreenSchemeId}/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#get-issue-type-screen-scheme-projects
func (t *TypeScreenSchemeService) SchemesByProject(ctx context.Context, issueTypeScreenSchemeId, startAt, maxResults int) (*model.IssueTypeScreenSchemeByProjectPageScheme, *model.ResponseScheme, error) {
	return t.internalClient.SchemesByProject(ctx, issueTypeScreenSchemeId, startAt, maxResults)
}

type internalTypeScreenSchemeImpl struct {
	c       service.Connector
	version string
}

func (i *internalTypeScreenSchemeImpl) Gets(ctx context.Context, options *model.ScreenSchemeParamsScheme, startAt, maxResults int) (*model.IssueTypeScreenSchemePageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		for _, id := range options.IDs {
			params.Add("id", strconv.Itoa(id))
		}

		if options.QueryString != "" {
			params.Add("queryString", options.QueryString)
		}

		if options.OrderBy != "orderBy" {
			params.Add("", options.OrderBy)
		}

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescreenscheme?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.IssueTypeScreenSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalTypeScreenSchemeImpl) Create(ctx context.Context, payload *model.IssueTypeScreenSchemePayloadScheme) (*model.IssueTypeScreenScreenCreatedScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescreenscheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	scheme := new(model.IssueTypeScreenScreenCreatedScheme)
	response, err := i.c.Call(request, scheme)
	if err != nil {
		return nil, response, err
	}

	return scheme, response, nil
}

func (i *internalTypeScreenSchemeImpl) Assign(ctx context.Context, issueTypeScreenSchemeId, projectId string) (*model.ResponseScheme, error) {

	if issueTypeScreenSchemeId == "" {
		return nil, model.ErrNoIssueTypeScreenSchemeIDError
	}

	if projectId == "" {
		return nil, model.ErrNoProjectIDError
	}

	payload := map[string]interface{}{
		"issueTypeScreenSchemeId": issueTypeScreenSchemeId,
		"projectId":               projectId,
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescreenscheme/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeScreenSchemeImpl) Projects(ctx context.Context, projectIds []int, startAt, maxResults int) (*model.IssueTypeProjectScreenSchemePageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range projectIds {
		params.Add("projectId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescreenscheme/project?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.IssueTypeProjectScreenSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalTypeScreenSchemeImpl) Mapping(ctx context.Context, issueTypeScreenSchemeIds []int, startAt, maxResults int) (*model.IssueTypeScreenSchemeMappingScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeScreenSchemeIds {
		params.Add("issueTypeScreenSchemeId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescreenscheme/mapping?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	mapping := new(model.IssueTypeScreenSchemeMappingScheme)
	response, err := i.c.Call(request, mapping)
	if err != nil {
		return nil, response, err
	}

	return mapping, response, nil
}

func (i *internalTypeScreenSchemeImpl) Update(ctx context.Context, issueTypeScreenSchemeId, name, description string) (*model.ResponseScheme, error) {

	if issueTypeScreenSchemeId == "" {
		return nil, model.ErrNoIssueTypeScreenSchemeIDError
	}

	payload := map[string]interface{}{"name": name}

	if description != "" {
		payload["description"] = description
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescreenscheme/%v", i.version, issueTypeScreenSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeScreenSchemeImpl) Delete(ctx context.Context, issueTypeScreenSchemeId string) (*model.ResponseScheme, error) {

	if issueTypeScreenSchemeId == "" {
		return nil, model.ErrNoIssueTypeScreenSchemeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescreenscheme/%v", i.version, issueTypeScreenSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeScreenSchemeImpl) Append(ctx context.Context, issueTypeScreenSchemeId string, payload *model.IssueTypeScreenSchemePayloadScheme) (*model.ResponseScheme, error) {

	if issueTypeScreenSchemeId == "" {
		return nil, model.ErrNoIssueTypeScreenSchemeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescreenscheme/%v/mapping", i.version, issueTypeScreenSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeScreenSchemeImpl) UpdateDefault(ctx context.Context, issueTypeScreenSchemeId, screenSchemeId string) (*model.ResponseScheme, error) {

	if issueTypeScreenSchemeId == "" {
		return nil, model.ErrNoIssueTypeScreenSchemeIDError
	}

	if screenSchemeId == "" {
		return nil, model.ErrNoScreenSchemeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescreenscheme/%v/mapping/default", i.version, issueTypeScreenSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"screenSchemeId": screenSchemeId})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeScreenSchemeImpl) Remove(ctx context.Context, issueTypeScreenSchemeId string, issueTypeIds []string) (*model.ResponseScheme, error) {

	if issueTypeScreenSchemeId == "" {
		return nil, model.ErrNoIssueTypeScreenSchemeIDError
	}

	if len(issueTypeIds) == 0 {
		return nil, model.ErrNoIssueTypesError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescreenscheme/%v/mapping/remove", i.version, issueTypeScreenSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"issueTypeIds": issueTypeIds})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeScreenSchemeImpl) SchemesByProject(ctx context.Context, issueTypeScreenSchemeId int, startAt, maxResults int) (*model.IssueTypeScreenSchemeByProjectPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescreenscheme/%v/project?%v", i.version, issueTypeScreenSchemeId, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.IssueTypeScreenSchemeByProjectPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
