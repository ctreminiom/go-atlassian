package v3

import (
	"context"
	"encoding/json"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ProjectRoleService struct {
	client *Client
	Actor  *ProjectRoleActorService
}

// Gets returns a list of project roles for the project returning the name and self URL for each role.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-roles-for-project
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-roles/#api-rest-api-3-project-projectidorkey-role-get
func (p *ProjectRoleService) Gets(ctx context.Context, projectKeyOrID string) (result *map[string]int, response *ResponseScheme,
	err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/role", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	var (
		roles          = make(map[string]int)
		resultAsObject map[string]interface{}
	)

	if err = json.Unmarshal(response.Bytes.Bytes(), &resultAsObject); err != nil {
		return
	}

	for roleName, roleURL := range resultAsObject {

		urlParsed, err := url.Parse(roleURL.(string))
		if err != nil {
			return nil, response, err
		}

		var (
			urlPart            = strings.Split(urlParsed.Path, "/")
			urlPartLastElement = urlPart[len(urlPart)-1]
		)

		projectRoleIDAsInt, err := strconv.Atoi(urlPartLastElement)
		if err != nil {
			return nil, response, err
		}

		roles[roleName] = projectRoleIDAsInt
	}

	result = &roles

	return
}

type ProjectRoleScheme struct {
	Self             string                                `json:"self,omitempty"`
	Name             string                                `json:"name,omitempty"`
	ID               int                                   `json:"id,omitempty"`
	Description      string                                `json:"description,omitempty"`
	Actors           []*RoleActorScheme                    `json:"actors,omitempty"`
	Scope            *models.TeamManagedProjectScopeScheme `json:"scope,omitempty"`
	TranslatedName   string                                `json:"translatedName,omitempty"`
	CurrentUserRole  bool                                  `json:"currentUserRole,omitempty"`
	Admin            bool                                  `json:"admin,omitempty"`
	RoleConfigurable bool                                  `json:"roleConfigurable,omitempty"`
	Default          bool                                  `json:"default,omitempty"`
}

type RoleActorScheme struct {
	ID          int    `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	AvatarURL   string `json:"avatarUrl,omitempty"`
	ActorUser   struct {
		AccountID string `json:"accountId,omitempty"`
	} `json:"actorUser,omitempty"`
	ActorGroup *models.GroupScheme `json:"actorGroup,omitempty"`
}

// Get returns a project role's details and actors associated with the project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-role-for-project
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-roles/#api-rest-api-3-project-projectidorkey-role-id-get
func (p *ProjectRoleService) Get(ctx context.Context, projectKeyOrID string, roleID int) (result *ProjectRoleScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/role/%v", projectKeyOrID, roleID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type ProjectRoleDetailScheme struct {
	Self             string                                `json:"self,omitempty"`
	Name             string                                `json:"name,omitempty"`
	ID               int                                   `json:"id,omitempty"`
	Description      string                                `json:"description,omitempty"`
	Admin            bool                                  `json:"admin,omitempty"`
	Scope            *models.TeamManagedProjectScopeScheme `json:"scope,omitempty"`
	RoleConfigurable bool                                  `json:"roleConfigurable,omitempty"`
	TranslatedName   string                                `json:"translatedName,omitempty"`
	Default          bool                                  `json:"default,omitempty"`
}

// Details returns all project roles and the details for each role.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-role-details
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-roles/#api-rest-api-3-project-projectidorkey-roledetails-get
func (p *ProjectRoleService) Details(ctx context.Context, projectKeyOrID string) (result []*ProjectRoleDetailScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/roledetails", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Global gets a list of all project roles, complete with project role details and default actors.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-all-project-roles
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-roles/#api-rest-api-3-role-get
func (p *ProjectRoleService) Global(ctx context.Context) (result []*ProjectRoleScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/role"

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type ProjectRolePayloadScheme struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// Create creates a new project role with no default actors.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles#create-project-role
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-roles/#api-rest-api-3-role-post
func (p *ProjectRoleService) Create(ctx context.Context, payload *ProjectRolePayloadScheme) (result *ProjectRoleScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/role"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
