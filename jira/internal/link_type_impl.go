package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewLinkTypeService creates a new instance of LinkTypeService.
func NewLinkTypeService(client service.Connector, version string) (*LinkTypeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &LinkTypeService{
		internalClient: &internalLinkTypeImpl{c: client, version: version},
	}, nil
}

// LinkTypeService provides methods to manage issue link types in Jira Service Management.
type LinkTypeService struct {
	// internalClient is the connector interface for issue link type operations.
	internalClient jira.LinkTypeConnector
}

// Gets returns a list of all issue link types.
//
// GET /rest/api/{2-3}/issueLinkType
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#get-issue-link-types
func (l *LinkTypeService) Gets(ctx context.Context) (*model.IssueLinkTypeSearchScheme, *model.ResponseScheme, error) {
	return l.internalClient.Gets(ctx)
}

// Get returns an issue link type.
//
// GET /rest/api/{2-3}/issueLinkType/{issueLinkTypeID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#get-issue-link-type
func (l *LinkTypeService) Get(ctx context.Context, issueLinkTypeID string) (*model.LinkTypeScheme, *model.ResponseScheme, error) {
	return l.internalClient.Get(ctx, issueLinkTypeID)
}

// Create creates an issue link type.
//
// Use this operation to create descriptions of the reasons why issues are linked.
//
// The issue link type consists of a name and descriptions for a link's inward and outward relationships.
//
// POST /rest/api/{2-3}/issueLinkType
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#create-issue-link-type
func (l *LinkTypeService) Create(ctx context.Context, payload *model.LinkTypeScheme) (*model.LinkTypeScheme, *model.ResponseScheme, error) {
	return l.internalClient.Create(ctx, payload)
}

// Update updates an issue link type.
//
// PUT /rest/api/{2-3}/issueLinkType/{issueLinkTypeID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#update-issue-link-type
func (l *LinkTypeService) Update(ctx context.Context, issueLinkTypeID string, payload *model.LinkTypeScheme) (*model.LinkTypeScheme, *model.ResponseScheme, error) {
	return l.internalClient.Update(ctx, issueLinkTypeID, payload)
}

// Delete deletes an issue link type.
//
// DELETE /rest/api/{2-3}/issueLinkType/{issueLinkTypeID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#delete-issue-link-type
func (l *LinkTypeService) Delete(ctx context.Context, issueLinkTypeID string) (*model.ResponseScheme, error) {
	return l.internalClient.Delete(ctx, issueLinkTypeID)
}

type internalLinkTypeImpl struct {
	c       service.Connector
	version string
}

func (i *internalLinkTypeImpl) Gets(ctx context.Context) (*model.IssueLinkTypeSearchScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/issueLinkType", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	types := new(model.IssueLinkTypeSearchScheme)
	response, err := i.c.Call(request, types)
	if err != nil {
		return nil, response, err
	}

	return types, response, nil
}

func (i *internalLinkTypeImpl) Get(ctx context.Context, issueLinkTypeID string) (*model.LinkTypeScheme, *model.ResponseScheme, error) {

	if issueLinkTypeID == "" {
		return nil, nil, model.ErrNoLinkTypeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issueLinkType/%v", i.version, issueLinkTypeID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	linkType := new(model.LinkTypeScheme)
	response, err := i.c.Call(request, linkType)
	if err != nil {
		return nil, response, err
	}

	return linkType, response, nil
}

func (i *internalLinkTypeImpl) Create(ctx context.Context, payload *model.LinkTypeScheme) (*model.LinkTypeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/issueLinkType", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	linkType := new(model.LinkTypeScheme)
	response, err := i.c.Call(request, linkType)
	if err != nil {
		return nil, response, err
	}

	return linkType, response, nil
}

func (i *internalLinkTypeImpl) Update(ctx context.Context, issueLinkTypeID string, payload *model.LinkTypeScheme) (*model.LinkTypeScheme, *model.ResponseScheme, error) {

	if issueLinkTypeID == "" {
		return nil, nil, model.ErrNoLinkTypeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issueLinkType/%v", i.version, issueLinkTypeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	linkType := new(model.LinkTypeScheme)
	response, err := i.c.Call(request, linkType)
	if err != nil {
		return nil, response, err
	}

	return linkType, response, nil
}

func (i *internalLinkTypeImpl) Delete(ctx context.Context, issueLinkTypeID string) (*model.ResponseScheme, error) {

	if issueLinkTypeID == "" {
		return nil, model.ErrNoLinkTypeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issueLinkType/%v", i.version, issueLinkTypeID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
