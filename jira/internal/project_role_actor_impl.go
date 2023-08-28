package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strings"
)

func NewProjectRoleActorService(client service.Connector, version string) (*ProjectRoleActorService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectRoleActorService{
		internalClient: &internalProjectRoleActorImpl{c: client, version: version},
	}, nil
}

type ProjectRoleActorService struct {
	internalClient jira.ProjectRoleActorConnector
}

// Add adds actors to a project role for the project.
//
// POST /rest/api/{2-3}/project/{projectIdOrKey}/role/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/roles/actors#add-actors-to-project-role
func (p *ProjectRoleActorService) Add(ctx context.Context, projectKeyOrId string, roleId int, accountIds, groups []string) (*model.ProjectRoleScheme, *model.ResponseScheme, error) {
	return p.internalClient.Add(ctx, projectKeyOrId, roleId, accountIds, groups)
}

// Delete deletes actors from a project role for the project.
//
// DELETE /rest/api/{2-3}/project/{projectIdOrKey}/role/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/roles/actors#delete-actors-from-project-role
func (p *ProjectRoleActorService) Delete(ctx context.Context, projectKeyOrId string, roleId int, accountId, group string) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, projectKeyOrId, roleId, accountId, group)
}

type internalProjectRoleActorImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectRoleActorImpl) Add(ctx context.Context, projectKeyOrId string, roleId int, accountIds, groups []string) (*model.ProjectRoleScheme, *model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, nil, model.ErrNoProjectIDOrKeyError
	}

	if roleId == 0 {
		return nil, nil, model.ErrNoProjectRoleIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/role/%v", i.version, projectKeyOrId, roleId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"group": groups, "user": accountIds})
	if err != nil {
		return nil, nil, err
	}

	role := new(model.ProjectRoleScheme)
	response, err := i.c.Call(request, role)
	if err != nil {
		return nil, response, err
	}

	return role, response, nil
}

func (i *internalProjectRoleActorImpl) Delete(ctx context.Context, projectKeyOrId string, roleId int, accountId, group string) (*model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, model.ErrNoProjectIDOrKeyError
	}

	if roleId == 0 {
		return nil, model.ErrNoProjectRoleIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/project/%v/role/%v", i.version, projectKeyOrId, roleId))

	params := url.Values{}

	if len(accountId) != 0 {
		params.Add("user", accountId)
	}

	if len(group) != 0 {
		params.Add("group", group)
	}

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint.String(), "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
