package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewProjectPermissionSchemeService creates a new instance of ProjectPermissionSchemeService.
func NewProjectPermissionSchemeService(client service.Connector, version string) (*ProjectPermissionSchemeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectPermissionSchemeService{
		internalClient: &internalProjectPermissionSchemeImpl{c: client, version: version},
	}, nil
}

// ProjectPermissionSchemeService provides methods to manage project permission schemes in Jira Service Management.
type ProjectPermissionSchemeService struct {
	// internalClient is the connector interface for project permission scheme operations.
	internalClient jira.ProjectPermissionSchemeConnector
}

// Get search the permission scheme associated with the project.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/permissionscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#get-assigned-permission-scheme
func (p *ProjectPermissionSchemeService) Get(ctx context.Context, projectKeyOrID string, expand []string) (*model.PermissionSchemeScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, projectKeyOrID, expand)
}

// Assign assigns a permission scheme with a project.
//
// See Managing project permissions for more information about permission schemes.
//
// PUT /rest/api/{2-3}/project/{projectKeyOrID}/permissionscheme
func (p *ProjectPermissionSchemeService) Assign(ctx context.Context, projectKeyOrID string, permissionSchemeID int) (*model.PermissionSchemeScheme, *model.ResponseScheme, error) {
	return p.internalClient.Assign(ctx, projectKeyOrID, permissionSchemeID)
}

// SecurityLevels returns all issue security levels for the project that the user has access to.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/securitylevel
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#get-project-issue-security-levels
func (p *ProjectPermissionSchemeService) SecurityLevels(ctx context.Context, projectKeyOrID string) (*model.IssueSecurityLevelsScheme, *model.ResponseScheme, error) {
	return p.internalClient.SecurityLevels(ctx, projectKeyOrID)
}

type internalProjectPermissionSchemeImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectPermissionSchemeImpl) Get(ctx context.Context, projectKeyOrID string, expand []string) (*model.PermissionSchemeScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/project/%v/permissionscheme", i.version, projectKeyOrID))

	if expand != nil {
		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	permissionScheme := new(model.PermissionSchemeScheme)
	response, err := i.c.Call(request, permissionScheme)
	if err != nil {
		return nil, response, err
	}

	return permissionScheme, response, nil
}

func (i *internalProjectPermissionSchemeImpl) Assign(ctx context.Context, projectKeyOrID string, permissionSchemeID int) (*model.PermissionSchemeScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/permissionscheme", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"id": permissionSchemeID})
	if err != nil {
		return nil, nil, err
	}

	permissionScheme := new(model.PermissionSchemeScheme)
	response, err := i.c.Call(request, permissionScheme)
	if err != nil {
		return nil, response, err
	}

	return permissionScheme, response, nil
}

func (i *internalProjectPermissionSchemeImpl) SecurityLevels(ctx context.Context, projectKeyOrID string) (*model.IssueSecurityLevelsScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/securitylevel", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	securityLevel := new(model.IssueSecurityLevelsScheme)
	response, err := i.c.Call(request, securityLevel)
	if err != nil {
		return nil, response, err
	}

	return securityLevel, response, nil
}
