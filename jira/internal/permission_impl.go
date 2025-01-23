package internal

import (
	"context"
	"encoding/json"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
	"net/http"
)

// NewPermissionService creates a new instance of PermissionService.
func NewPermissionService(client service.Connector, version string, scheme *PermissionSchemeService) (*PermissionService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &PermissionService{
		internalClient: &internalPermissionImpl{c: client, version: version},
		Scheme:         scheme,
	}, nil
}

// PermissionService provides methods to manage permissions in Jira Service Management.
type PermissionService struct {
	// internalClient is the connector interface for permission operations.
	internalClient jira.PermissionConnector
	// Scheme is the service for managing permission schemes.
	Scheme *PermissionSchemeService
}

// Gets returns all permissions, including: global permissions, project permissions and global permissions added by plugins.
//
// GET /rest/api/{2-3}/permissions
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions#get-my-permissions
func (p *PermissionService) Gets(ctx context.Context) ([]*model.PermissionScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx)
}

// Check search the permissions linked to an accountID, then check if the user permissions.
//
// POST /rest/api/{2-3}/permissions/check
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions#check-permissions
func (p *PermissionService) Check(ctx context.Context, payload *model.PermissionCheckPayload) (*model.PermissionGrantsScheme, *model.ResponseScheme, error) {
	return p.internalClient.Check(ctx, payload)
}

// Projects returns all the projects where the user is granted a list of project permissions.
//
// POST /rest/api/{2-3}/permissions/project
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions#get-permitted-projects
func (p *PermissionService) Projects(ctx context.Context, permissions []string) (*model.PermittedProjectsScheme, *model.ResponseScheme, error) {
	return p.internalClient.Projects(ctx, permissions)
}

type internalPermissionImpl struct {
	c       service.Connector
	version string
}

func (i *internalPermissionImpl) Gets(ctx context.Context) ([]*model.PermissionScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/permissions", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		return nil, response, err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(response.Bytes.Bytes(), &jsonMap)
	if err != nil {
		return nil, nil, err
	}

	var permissions []*model.PermissionScheme
	for key, value := range jsonMap["permissions"].(map[string]interface{}) {
		data, ok := value.(map[string]interface{})

		if ok {
			permissions = append(permissions, &model.PermissionScheme{
				Key:         key,
				Name:        data["name"].(string),
				Type:        data["type"].(string),
				Description: data["description"].(string),
			})
		}
	}

	return permissions, response, nil
}

func (i *internalPermissionImpl) Check(ctx context.Context, payload *model.PermissionCheckPayload) (*model.PermissionGrantsScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/permissions/check", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	permissions := new(model.PermissionGrantsScheme)
	response, err := i.c.Call(request, permissions)
	if err != nil {
		return nil, response, err
	}

	return permissions, response, nil
}

func (i *internalPermissionImpl) Projects(ctx context.Context, permissions []string) (*model.PermittedProjectsScheme, *model.ResponseScheme, error) {

	if len(permissions) == 0 {
		return nil, nil, model.ErrNoPermissionKeys
	}

	endpoint := fmt.Sprintf("rest/api/%v/permissions/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"permissions": permissions})
	if err != nil {
		return nil, nil, err
	}

	projects := new(model.PermittedProjectsScheme)
	response, err := i.c.Call(request, projects)
	if err != nil {
		return nil, response, err
	}

	return projects, response, nil
}
