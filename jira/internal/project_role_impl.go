package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewProjectRoleService creates a new instance of ProjectRoleService.
func NewProjectRoleService(client service.Connector, version string, actor *ProjectRoleActorService) (*ProjectRoleService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectRoleService{
		internalClient: &internalProjectRoleImpl{c: client, version: version},
		Actor:          actor,
	}, nil
}

// ProjectRoleService provides methods to manage project roles in Jira Service Management.
type ProjectRoleService struct {
	// internalClient is the connector interface for project role operations.
	internalClient jira.ProjectRoleConnector
	// Actor is the service for managing project role actors.
	Actor *ProjectRoleActorService
}

// Gets returns a list of project roles for the project returning the name and self URL for each role.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/role
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-roles-for-project
func (p *ProjectRoleService) Gets(ctx context.Context, projectKeyOrID string) (*map[string]int, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, projectKeyOrID)
}

// Get returns a project role's details and actors associated with the project.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/role/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-role-for-project
func (p *ProjectRoleService) Get(ctx context.Context, projectKeyOrID string, roleID int) (*model.ProjectRoleScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, projectKeyOrID, roleID)
}

// Details returns all project roles and the details for each role.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/roledetails
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-role-details
func (p *ProjectRoleService) Details(ctx context.Context, projectKeyOrID string) ([]*model.ProjectRoleDetailScheme, *model.ResponseScheme, error) {
	return p.internalClient.Details(ctx, projectKeyOrID)
}

// Global gets a list of all project roles, complete with project role details and default actors.
//
// GET /rest/api/{2-3}/role
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-all-project-roles
func (p *ProjectRoleService) Global(ctx context.Context) ([]*model.ProjectRoleScheme, *model.ResponseScheme, error) {
	return p.internalClient.Global(ctx)
}

// Create creates a new project role with no default actors.
//
// POST /rest/api/{2-3}/role
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#create-project-role
func (p *ProjectRoleService) Create(ctx context.Context, payload *model.ProjectRolePayloadScheme) (*model.ProjectRoleScheme, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, payload)
}

type internalProjectRoleImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectRoleImpl) Gets(ctx context.Context, projectKeyOrID string) (*map[string]int, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/role", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		return nil, response, err
	}

	roles, jsonArray := make(map[string]int), make(map[string]interface{})

	if err = json.Unmarshal(response.Bytes.Bytes(), &jsonArray); err != nil {
		return nil, response, err
	}

	for name, link := range jsonArray {

		uri, err := url.Parse(link.(string))
		if err != nil {
			return nil, response, err
		}

		uriAsSlice := strings.Split(uri.Path, "/") // "ctreminiom.atlassian.net,rest,api,3,project,10000,role,10002"
		uriRoleID := uriAsSlice[len(uriAsSlice)-1] // 10002

		roleID, err := strconv.Atoi(uriRoleID)
		if err != nil {
			return nil, response, err
		}

		roles[name] = roleID
	}

	return &roles, response, nil
}

func (i *internalProjectRoleImpl) Get(ctx context.Context, projectKeyOrID string, roleID int) (*model.ProjectRoleScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/role/%v", i.version, projectKeyOrID, roleID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
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

func (i *internalProjectRoleImpl) Details(ctx context.Context, projectKeyOrID string) ([]*model.ProjectRoleDetailScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/roledetails", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var details []*model.ProjectRoleDetailScheme
	response, err := i.c.Call(request, &details)
	if err != nil {
		return nil, response, err
	}

	return details, response, nil
}

func (i *internalProjectRoleImpl) Global(ctx context.Context) ([]*model.ProjectRoleScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/role", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var roles []*model.ProjectRoleScheme
	response, err := i.c.Call(request, &roles)
	if err != nil {
		return nil, response, err
	}

	return roles, response, nil
}

func (i *internalProjectRoleImpl) Create(ctx context.Context, payload *model.ProjectRolePayloadScheme) (*model.ProjectRoleScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/role", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
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
