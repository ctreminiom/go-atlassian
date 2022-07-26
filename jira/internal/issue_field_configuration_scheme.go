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

func (i *IssueFieldConfigSchemeService) Gets(ctx context.Context, ids []int, startAt, maxResults int) (*model.FieldConfigurationSchemePageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, ids, startAt, maxResults)
}

func (i *IssueFieldConfigSchemeService) Create(ctx context.Context, name, description string) (*model.FieldConfigurationSchemeScheme, *model.ResponseScheme, error) {
	return i.internalClient.Create(ctx, name, description)
}

func (i *IssueFieldConfigSchemeService) Mapping(ctx context.Context, fieldConfigIds []int, startAt, maxResults int) (*model.FieldConfigurationIssueTypeItemPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Mapping(ctx, fieldConfigIds, startAt, maxResults)
}

func (i *IssueFieldConfigSchemeService) Project(ctx context.Context, projectIds []int, startAt, maxResults int) (*model.FieldConfigurationSchemeProjectPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Project(ctx, projectIds, startAt, maxResults)
}

func (i *IssueFieldConfigSchemeService) Assign(ctx context.Context, payload *model.FieldConfigurationSchemeAssignPayload) (*model.ResponseScheme, error) {
	return i.internalClient.Assign(ctx, payload)
}

func (i *IssueFieldConfigSchemeService) Update(ctx context.Context, schemeId int, name, description string) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, schemeId, name, description)
}

func (i *IssueFieldConfigSchemeService) Delete(ctx context.Context, schemeId int) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, schemeId)
}

func (i *IssueFieldConfigSchemeService) Link(ctx context.Context, schemeId int, payload *model.FieldConfigurationToIssueTypeMappingPayloadScheme) (*model.ResponseScheme, error) {
	return i.internalClient.Link(ctx, schemeId, payload)
}

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
