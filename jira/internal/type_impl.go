package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewTypeService(client service.Connector, version string, scheme *TypeSchemeService, screenScheme *TypeScreenSchemeService) (
	*TypeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &TypeService{
		internalClient: &internalTypeImpl{c: client, version: version},
		Scheme:         scheme,
		ScreenScheme:   screenScheme,
	}, nil
}

type TypeService struct {
	internalClient jira.TypeConnector
	Scheme         *TypeSchemeService
	ScreenScheme   *TypeScreenSchemeService
}

// Gets returns all issue types.
//
// GET /rest/api/{2-3}/issuetype
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-all-issue-types-for-user
func (t *TypeService) Gets(ctx context.Context) ([]*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Gets(ctx)
}

// Create creates an issue type and adds it to the default issue type scheme.
//
// POST /rest/api/{2-3}/issuetype
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/type#create-issue-type
func (t *TypeService) Create(ctx context.Context, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Create(ctx, payload)
}

// Get returns an issue type.
//
// GET /rest/api/{2-3}/issuetype/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-issue-type
func (t *TypeService) Get(ctx context.Context, issueTypeId string) (*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Get(ctx, issueTypeId)
}

// Update updates the issue type.
//
// PUT /rest/api/{2-3}/issuetype/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/type#update-issue-type
func (t *TypeService) Update(ctx context.Context, issueTypeId string, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Update(ctx, issueTypeId, payload)
}

// Delete deletes the issue type.
//
// If the issue type is in use, all uses are updated with the alternative issue type (alternativeIssueTypeId).
// A list of alternative issue types are obtained from the Get alternative issue types resource.
//
// DELETE /rest/api/{2-3}/issuetype/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/type#delete-issue-type
func (t *TypeService) Delete(ctx context.Context, issueTypeId string) (*model.ResponseScheme, error) {
	return t.internalClient.Delete(ctx, issueTypeId)
}

// Alternatives returns a list of issue types that can be used to replace the issue type.
//
// The alternative issue types are those assigned to the same workflow scheme, field configuration scheme, and screen scheme.
//
// GET /rest/api/{2-3}/issuetype/{id}/alternatives
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-alternative-issue-types
func (t *TypeService) Alternatives(ctx context.Context, issueTypeId string) ([]*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Alternatives(ctx, issueTypeId)
}

type internalTypeImpl struct {
	c       service.Connector
	version string
}

func (i *internalTypeImpl) Gets(ctx context.Context) ([]*model.IssueTypeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/issuetype", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
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

	endpoint := fmt.Sprintf("rest/api/%v/issuetype", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
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

	endpoint := fmt.Sprintf("rest/api/%v/issuetype/%v", i.version, issueTypeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
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

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeImpl) Alternatives(ctx context.Context, issueTypeId string) ([]*model.IssueTypeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/issuetype/%v/alternatives", i.version, issueTypeId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
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
