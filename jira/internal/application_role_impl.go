package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewApplicationRoleService(client service.Connector, version string) (*ApplicationRoleService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ApplicationRoleService{
		internalClient: &internalApplicationRoleImpl{c: client, version: version},
	}, nil
}

type ApplicationRoleService struct {
	internalClient jira.AppRoleConnector
}

// Gets returns all application roles.
//
// In Jira, application roles are managed using the Application access configuration page.
//
// GET /rest/api/{2-3}/applicationrole
//
// https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-all-application-roles
func (a *ApplicationRoleService) Gets(ctx context.Context) ([]*model.ApplicationRoleScheme, *model.ResponseScheme, error) {
	return a.internalClient.Gets(ctx)
}

// Get returns an application role.
//
// GET /rest/api/{2-3}/applicationrole/{key}
//
// https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-application-role
func (a *ApplicationRoleService) Get(ctx context.Context, key string) (*model.ApplicationRoleScheme, *model.ResponseScheme, error) {
	return a.internalClient.Get(ctx, key)
}

type internalApplicationRoleImpl struct {
	c       service.Connector
	version string
}

func (i *internalApplicationRoleImpl) Gets(ctx context.Context) ([]*model.ApplicationRoleScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/applicationrole", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var roles []*model.ApplicationRoleScheme
	response, err := i.c.Call(request, &roles)
	if err != nil {
		return nil, nil, err
	}

	return roles, response, nil
}

func (i *internalApplicationRoleImpl) Get(ctx context.Context, key string) (*model.ApplicationRoleScheme, *model.ResponseScheme, error) {

	if key == "" {
		return nil, nil, model.ErrNoApplicationRoleError
	}

	endpoint := fmt.Sprintf("rest/api/%v/applicationrole/%v", i.version, key)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	role := new(model.ApplicationRoleScheme)
	response, err := i.c.Call(request, role)
	if err != nil {
		return nil, nil, err
	}

	return role, response, nil
}
