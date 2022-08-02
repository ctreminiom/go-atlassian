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

func NewTypeSchemeService(client service.Client, version string) (*TypeSchemeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &TypeSchemeService{
		internalClient: &internalTypeSchemeImpl{c: client, version: version},
	}, nil
}

type TypeSchemeService struct {
	internalClient jira.TypeSchemeConnector
}

func (t *TypeSchemeService) Gets(ctx context.Context, issueTypeSchemeIds []int, startAt, maxResults int) (*model.IssueTypeSchemePageScheme, *model.ResponseScheme, error) {
	return t.internalClient.Gets(ctx, issueTypeSchemeIds, startAt, maxResults)
}

func (t *TypeSchemeService) Create(ctx context.Context, payload *model.IssueTypeSchemePayloadScheme) (*model.NewIssueTypeSchemeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Create(ctx, payload)
}

func (t *TypeSchemeService) Items(ctx context.Context, issueTypeSchemeIds []int, startAt, maxResults int) (*model.IssueTypeSchemeItemPageScheme, *model.ResponseScheme, error) {
	return t.internalClient.Items(ctx, issueTypeSchemeIds, startAt, maxResults)
}

func (t *TypeSchemeService) Projects(ctx context.Context, projectIds []int, startAt, maxResults int) (*model.ProjectIssueTypeSchemePageScheme, *model.ResponseScheme, error) {
	return t.internalClient.Projects(ctx, projectIds, startAt, maxResults)
}

func (t *TypeSchemeService) Assign(ctx context.Context, issueTypeSchemeId, projectId string) (*model.ResponseScheme, error) {
	return t.internalClient.Assign(ctx, issueTypeSchemeId, projectId)
}

func (t *TypeSchemeService) Update(ctx context.Context, issueTypeSchemeId int, payload *model.IssueTypeSchemePayloadScheme) (*model.ResponseScheme, error) {
	return t.internalClient.Update(ctx, issueTypeSchemeId, payload)
}

func (t *TypeSchemeService) Delete(ctx context.Context, issueTypeSchemeId int) (*model.ResponseScheme, error) {
	return t.internalClient.Delete(ctx, issueTypeSchemeId)
}

func (t *TypeSchemeService) Append(ctx context.Context, issueTypeSchemeId int, issueTypeIds []int) (*model.ResponseScheme, error) {
	return t.internalClient.Append(ctx, issueTypeSchemeId, issueTypeIds)
}

func (t *TypeSchemeService) Remove(ctx context.Context, issueTypeSchemeId, issueTypeId int) (*model.ResponseScheme, error) {
	return t.internalClient.Remove(ctx, issueTypeSchemeId, issueTypeId)
}

type internalTypeSchemeImpl struct {
	c       service.Client
	version string
}

func (i *internalTypeSchemeImpl) Gets(ctx context.Context, issueTypeSchemeIds []int, startAt, maxResults int) (*model.IssueTypeSchemePageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeSchemeIds {
		params.Add("id", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.IssueTypeSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalTypeSchemeImpl) Create(ctx context.Context, payload *model.IssueTypeSchemePayloadScheme) (*model.NewIssueTypeSchemeScheme, *model.ResponseScheme, error) {

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	issueType := new(model.NewIssueTypeSchemeScheme)
	response, err := i.c.Call(request, issueType)
	if err != nil {
		return nil, response, err
	}

	return issueType, response, nil
}

func (i *internalTypeSchemeImpl) Items(ctx context.Context, issueTypeSchemeIds []int, startAt, maxResults int) (*model.IssueTypeSchemeItemPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeSchemeIds {
		params.Add("issueTypeSchemeId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/mapping?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.IssueTypeSchemeItemPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalTypeSchemeImpl) Projects(ctx context.Context, projectIds []int, startAt, maxResults int) (*model.ProjectIssueTypeSchemePageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range projectIds {
		params.Add("projectId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/project?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ProjectIssueTypeSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalTypeSchemeImpl) Assign(ctx context.Context, issueTypeSchemeId, projectId string) (*model.ResponseScheme, error) {

	if issueTypeSchemeId == "" {
		return nil, model.ErrNoIssueTypeSchemeIDError
	}

	if projectId == "" {
		return nil, model.ErrNoProjectIDError
	}

	payload := struct {
		IssueTypeSchemeID string `json:"issueTypeSchemeId"`
		ProjectID         string `json:"projectId"`
	}{
		IssueTypeSchemeID: issueTypeSchemeId,
		ProjectID:         projectId,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeSchemeImpl) Update(ctx context.Context, issueTypeSchemeId int, payload *model.IssueTypeSchemePayloadScheme) (*model.ResponseScheme, error) {

	if issueTypeSchemeId == 0 {
		return nil, model.ErrNoIssueTypeSchemeIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v", i.version, issueTypeSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeSchemeImpl) Delete(ctx context.Context, issueTypeSchemeId int) (*model.ResponseScheme, error) {

	if issueTypeSchemeId == 0 {
		return nil, model.ErrNoIssueTypeSchemeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v", i.version, issueTypeSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeSchemeImpl) Append(ctx context.Context, issueTypeSchemeId int, issueTypeIds []int) (*model.ResponseScheme, error) {

	if len(issueTypeIds) == 0 {
		return nil, model.ErrNoIssueTypesError
	}

	var ids []string
	for _, issueTypeID := range issueTypeIds {
		ids = append(ids, strconv.Itoa(issueTypeID))
	}

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{
		IssueTypeIds: ids,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v/issuetype", i.version, issueTypeSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeSchemeImpl) Remove(ctx context.Context, issueTypeSchemeId, issueTypeId int) (*model.ResponseScheme, error) {

	if issueTypeSchemeId == 0 {
		return nil, model.ErrNoIssueTypeSchemeIDError
	}

	if issueTypeId == 0 {
		return nil, model.ErrNoIssueTypeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v/issuetype/%v", i.version, issueTypeSchemeId, issueTypeId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
