package v3

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
	"net/url"
)

type ProjectValidationService struct{ client *Client }

// Validate validates a project key by confirming the key is a valid string and not in use.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/validation#validate-project-key
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-key-and-name-validation/#api-rest-api-3-projectvalidate-key-get
func (p *ProjectValidationService) Validate(ctx context.Context, projectKey string) (result *models.ProjectValidationMessageScheme,
	response *ResponseScheme, err error) {

	if len(projectKey) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	params := url.Values{}
	params.Add("key", projectKey)

	var endpoint = fmt.Sprintf("rest/api/3/projectvalidate/key?%v", params.Encode())

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

// Key validates a project key and, if the key is invalid or in use, generates a valid random string for the project key.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/validation#get-valid-project-key
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-key-and-name-validation/#api-rest-api-3-projectvalidate-validprojectkey-get
func (p *ProjectValidationService) Key(ctx context.Context, projectKey string) (randomKey string, response *ResponseScheme,
	err error) {

	if len(projectKey) == 0 {
		return "", nil, models.ErrNoProjectIDError
	}

	params := url.Values{}
	params.Add("key", projectKey)

	var endpoint = fmt.Sprintf("rest/api/3/projectvalidate/validProjectKey?%v", params.Encode())

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	return response.Bytes.String(), response, nil
}

// Name checks that a project name isn't in use. If the name isn't in use, the passed string is returned.
// If the name is in use, this operation attempts to generate a valid project name based on the one supplied,
// usually by adding a sequence number. If a valid project name cannot be generated, a 404 response is returned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/validation#get-valid-project-name
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-key-and-name-validation/#api-rest-api-3-projectvalidate-validprojectname-get
func (p *ProjectValidationService) Name(ctx context.Context, projectName string) (randomName string,
	response *ResponseScheme, err error) {

	if len(projectName) == 0 {
		return "", nil, models.ErrNoProjectNameError
	}

	params := url.Values{}
	params.Add("name", projectName)

	var endpoint = fmt.Sprintf("rest/api/3/projectvalidate/validProjectName?%v", params.Encode())

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	return response.Bytes.String(), response, nil
}
