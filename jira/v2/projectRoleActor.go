package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strings"
)

type ProjectRoleActorService struct{ client *Client }

// Add adds actors to a project role for the project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles/actors#add-actors-to-project-role
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-role-actors/#api-rest-api-2-project-projectidorkey-role-id-post
func (p *ProjectRoleActorService) Add(ctx context.Context, projectKeyOrID string, projectRoleID int, accountIDs, groups []string) (
	result *models.ProjectRoleScheme, response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	payload := struct {
		Group []string `json:"group,omitempty"`
		Users []string `json:"user,omitempty"`
	}{
		Group: groups,
		Users: accountIDs,
	}

	var endpoint = fmt.Sprintf("rest/api/2/project/%v/role/%v", projectKeyOrID, projectRoleID)

	payloadAsReader, _ := transformStructToReader(&payload)

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

// Delete deletes actors from a project role for the project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles/actors#delete-actors-from-project-role
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-role-actors/#api-rest-api-2-project-projectidorkey-role-id-delete
func (p *ProjectRoleActorService) Delete(ctx context.Context, projectKeyOrID string, projectRoleID int, accountID, group string) (
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, models.ErrNoProjectIDError
	}

	params := url.Values{}

	if len(accountID) != 0 {
		params.Add("user", accountID)
	}
	if len(group) != 0 {
		params.Add("group", group)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/2/project/%v/role/%v", projectKeyOrID, projectRoleID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := p.client.newRequest(ctx, http.MethodDelete, endpoint.String(), nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
