package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewApplicationRoleService(client service.Client, version string) (jira.AppRoleConnector, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ApplicationRoleService{client, version}, nil
}

type ApplicationRoleService struct {
	c       service.Client
	version string
}

func (a ApplicationRoleService) Gets(ctx context.Context) ([]*model.ApplicationRoleScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/applicationrole", a.version)

	request, err := a.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var roles []*model.ApplicationRoleScheme
	response, err := a.c.Call(request, roles)
	if err != nil {
		return nil, nil, err
	}

	return roles, response, nil
}

func (a ApplicationRoleService) Get(ctx context.Context, key string) (*model.ApplicationRoleScheme, *model.ResponseScheme, error) {

	if key == "" {
		return nil, nil, model.ErrNoApplicationRoleError
	}

	endpoint := fmt.Sprintf("rest/api/%v/applicationrole/%v", a.version, key)

	request, err := a.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	role := new(model.ApplicationRoleScheme)
	response, err := a.c.Call(request, role)
	if err != nil {
		return nil, nil, err
	}

	return role, response, nil
}
