package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewTypeService creates a new instance of TypeService.
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

// TypeService provides methods to manage issue types in Jira Service Management.
type TypeService struct {
	// internalClient is the connector interface for issue type operations.
	internalClient jira.TypeConnector
	// Scheme is the service for managing type schemes.
	Scheme *TypeSchemeService
	// ScreenScheme is the service for managing type screen schemes.
	ScreenScheme *TypeScreenSchemeService
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
func (t *TypeService) Get(ctx context.Context, issueTypeID string) (*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Get(ctx, issueTypeID)
}

// Update updates the issue type.
//
// PUT /rest/api/{2-3}/issuetype/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/type#update-issue-type
func (t *TypeService) Update(ctx context.Context, issueTypeID string, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Update(ctx, issueTypeID, payload)
}

// Delete deletes the issue type.
//
// If the issue type is in use, all uses are updated with the alternative issue type (alternativeIssueTypeID).
// A list of alternative issue types are obtained from the Get alternative issue types resource.
//
// DELETE /rest/api/{2-3}/issuetype/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/type#delete-issue-type
func (t *TypeService) Delete(ctx context.Context, issueTypeID string) (*model.ResponseScheme, error) {
	return t.internalClient.Delete(ctx, issueTypeID)
}

// Alternatives returns a list of issue types that can be used to replace the issue type.
//
// The alternative issue types are those assigned to the same workflow scheme, field configuration scheme, and screen scheme.
//
// GET /rest/api/{2-3}/issuetype/{id}/alternatives
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-alternative-issue-types
func (t *TypeService) Alternatives(ctx context.Context, issueTypeID string) ([]*model.IssueTypeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Alternatives(ctx, issueTypeID)
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

func (i *internalTypeImpl) Get(ctx context.Context, issueTypeID string) (*model.IssueTypeScheme, *model.ResponseScheme, error) {

	if issueTypeID == "" {
		return nil, nil, model.ErrNoIssueTypeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetype/%v", i.version, issueTypeID)

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

func (i *internalTypeImpl) Update(ctx context.Context, issueTypeID string, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error) {

	if issueTypeID == "" {
		return nil, nil, model.ErrNoIssueTypeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetype/%v", i.version, issueTypeID)

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

func (i *internalTypeImpl) Delete(ctx context.Context, issueTypeID string) (*model.ResponseScheme, error) {

	if issueTypeID == "" {
		return nil, model.ErrNoIssueTypeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetype/%v", i.version, issueTypeID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeImpl) Alternatives(ctx context.Context, issueTypeID string) ([]*model.IssueTypeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/issuetype/%v/alternatives", i.version, issueTypeID)

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
