package internal

import (
	"context"
	"encoding/json"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewPermissionService(client service.Client, version string) (*PermissionService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &PermissionService{
		internalClient: &internalPermissionImpl{c: client, version: version},
	}, nil
}

type PermissionService struct {
	internalClient jira.PermissionConnector
}

// Gets returns all permissions, including: global permissions, project permissions and global permissions added by plugins.
//
// GET /rest/api/{2-3}/permissions
//
// TODO: Add/Create documentation
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
// TODO: Add/Create documentation
func (p *PermissionService) Projects(ctx context.Context, permissions []string) (*model.PermittedProjectsScheme, *model.ResponseScheme, error) {
	return p.internalClient.Projects(ctx, permissions)
}

type internalPermissionImpl struct {
	c       service.Client
	version string
}

func (i *internalPermissionImpl) Gets(ctx context.Context) ([]*model.PermissionScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/permissions", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/permissions/check", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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

	payload := struct {
		Permissions []string `json:"permissions,omitempty"`
	}{
		Permissions: permissions,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/permissions/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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
