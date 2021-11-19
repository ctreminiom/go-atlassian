package v2

import (
	"context"
	"fmt"
	models2 "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type ProjectTypeService struct{ client *Client }

// Gets returns all project types, whether the instance has a valid license for each type.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-all-project-types
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-types/#api-rest-api-2-project-type-get
func (p *ProjectTypeService) Gets(ctx context.Context) (result []*models2.ProjectTypeScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/2/project/type"

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

// Licensed returns all project types with a valid license.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-licensed-project-types
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-types/#api-rest-api-2-project-type-accessible-get
func (p *ProjectTypeService) Licensed(ctx context.Context) (result []*models2.ProjectTypeScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/2/project/type/accessible"

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

// Get returns a project type.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-project-type-by-key
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-types/#api-rest-api-2-project-type-projecttypekey-get
func (p *ProjectTypeService) Get(ctx context.Context, projectTypeKey string) (result *models2.ProjectTypeScheme,
	response *ResponseScheme, err error) {

	if len(projectTypeKey) == 0 {
		return nil, nil, models2.ErrProjectTypeKeyError
	}

	var endpoint = fmt.Sprintf("rest/api/2/project/type/%v", projectTypeKey)

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

// Accessible returns a project type if it is accessible to the user.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-accessible-project-type-by-key
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-types/#api-rest-api-2-project-type-projecttypekey-accessible-get
func (p *ProjectTypeService) Accessible(ctx context.Context, projectTypeKey string) (result *models2.ProjectTypeScheme,
	response *ResponseScheme, err error) {

	if len(projectTypeKey) == 0 {
		return nil, nil, models2.ErrProjectTypeKeyError
	}

	var endpoint = fmt.Sprintf("rest/api/2/project/type/%v/accessible", projectTypeKey)

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
