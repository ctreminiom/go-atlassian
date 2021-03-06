package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ProjectRoleActorService struct{ client *Client }

// Adds actors to a project role for the project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles/actors#add-actors-to-project-role
func (p *ProjectRoleActorService) Add(ctx context.Context, projectKeyOrID string, projectRoleID int, accountIDs, groups []string) (result *ProjectRoleScheme, response *Response, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKeyOrID value")
	}

	payload := struct {
		Group []string `json:"group,omitempty"`
		Users []string `json:"user,omitempty"`
	}{
		Group: groups,
		Users: accountIDs,
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/role/%v", projectKeyOrID, projectRoleID)

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectRoleScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes actors from a project role for the project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles/actors#delete-actors-from-project-role
func (p *ProjectRoleActorService) Delete(ctx context.Context, projectKeyOrID string, projectRoleID int, accountID, group string) (response *Response, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid projectKeyOrID value")
	}

	params := url.Values{}

	if len(accountID) != 0 {
		params.Add("user", accountID)
	}

	if len(group) != 0 {
		params.Add("group", group)
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/role/%v?%v", projectKeyOrID, projectRoleID, params.Encode())

	request, err := p.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	return
}
