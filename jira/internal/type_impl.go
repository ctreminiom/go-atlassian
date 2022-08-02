package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewTypeService(client service.Client, version string, scheme *TypeSchemeService) (*TypeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &TypeService{
		internalClient: &internalTypeImpl{c: client, version: version},
		Scheme:         scheme,
	}, nil
}

type TypeService struct {
	internalClient jira.TypeConnector
	Scheme         *TypeSchemeService
}

func (t *TypeService) Gets(ctx context.Context) ([]*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Gets(ctx)
}

func (t *TypeService) Create(ctx context.Context, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Create(ctx, payload)
}

func (t *TypeService) Get(ctx context.Context, issueTypeId string) (*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Get(ctx, issueTypeId)
}

func (t *TypeService) Update(ctx context.Context, issueTypeId string, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Update(ctx, issueTypeId, payload)
}

func (t *TypeService) Delete(ctx context.Context, issueTypeId string) (*model.ResponseScheme, error) {
	return t.internalClient.Delete(ctx, issueTypeId)
}

func (t *TypeService) Alternatives(ctx context.Context, issueTypeId string) ([]*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Alternatives(ctx, issueTypeId)
}

type internalTypeImpl struct {
	c       service.Client
	version string
}

func (i *internalTypeImpl) Gets(ctx context.Context) ([]*model.IssueTypeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/issuetype", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var issueTypes []*model.IssueTypeScheme
	response, err := i.c.Call(request, &issueTypes)
	if err != nil {
		return nil, response, err
	}

	return issueTypes, response, nil
}

func (i *internalTypeImpl) Create(ctx context.Context, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error) {

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetype", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	issueType := new(model.IssueTypeScheme)
	response, err := i.c.Call(request, issueType)
	if err != nil {
		return nil, response, err
	}

	return issueType, response, nil
}

func (i *internalTypeImpl) Get(ctx context.Context, issueTypeId string) (*model.IssueTypeScheme, *model.ResponseScheme, error) {

	if issueTypeId == "" {
		return nil, nil, model.ErrNoIssueTypeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetype/%v", i.version, issueTypeId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	issueType := new(model.IssueTypeScheme)
	response, err := i.c.Call(request, issueType)
	if err != nil {
		return nil, response, err
	}

	return issueType, response, nil
}

func (i *internalTypeImpl) Update(ctx context.Context, issueTypeId string, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error) {

	if issueTypeId == "" {
		return nil, nil, model.ErrNoIssueTypeIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetype/%v", i.version, issueTypeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	issueType := new(model.IssueTypeScheme)
	response, err := i.c.Call(request, issueType)
	if err != nil {
		return nil, response, err
	}

	return issueType, response, nil
}

func (i *internalTypeImpl) Delete(ctx context.Context, issueTypeId string) (*model.ResponseScheme, error) {

	if issueTypeId == "" {
		return nil, model.ErrNoIssueTypeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetype/%v", i.version, issueTypeId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeImpl) Alternatives(ctx context.Context, issueTypeId string) ([]*model.IssueTypeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/issuetype/%v/alternatives", i.version, issueTypeId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var issueTypes []*model.IssueTypeScheme
	response, err := i.c.Call(request, &issueTypes)
	if err != nil {
		return nil, response, err
	}

	return issueTypes, response, nil
}
