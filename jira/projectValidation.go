package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ProjectValidationService struct{ client *Client }

type ProjectValidationMessageScheme struct {
	ErrorMessages []string `json:"errorMessages"`
	Errors        struct {
		ProjectKey string `json:"projectKey"`
	} `json:"errors"`
}

// Validates a project key by confirming the key is a valid string and not in use.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/validation#validate-project-key
func (p *ProjectValidationService) Validate(ctx context.Context, projectKey string) (result *ProjectValidationMessageScheme, response *Response, err error) {

	if len(projectKey) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKey value")
	}

	params := url.Values{}
	params.Add("key", projectKey)

	var endpoint = fmt.Sprintf("rest/api/3/projectvalidate/key?%v", params.Encode())

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectValidationMessageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Validates a project key and, if the key is invalid or in use, generates a valid random string for the project key.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/validation#get-valid-project-key
func (p *ProjectValidationService) Key(ctx context.Context, projectKey string) (randomKey string, response *Response, err error) {

	if len(projectKey) == 0 {
		return "", nil, fmt.Errorf("error, please provide a valid projectKey value")
	}

	params := url.Values{}
	params.Add("key", projectKey)

	var endpoint = fmt.Sprintf("rest/api/3/projectvalidate/validProjectKey?%v", params.Encode())

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	randomKey = string(response.BodyAsBytes)

	return
}

// Checks that a project name isn't in use. If the name isn't in use, the passed string is returned.
// If the name is in use, this operation attempts to generate a valid project name based on the one supplied,
// usually by adding a sequence number. If a valid project name cannot be generated, a 404 response is returned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/validation#get-valid-project-name
func (p *ProjectValidationService) Name(ctx context.Context, projectName string) (randomName string, response *Response, err error) {

	if len(projectName) == 0 {
		return "", nil, fmt.Errorf("error, please provide a valid projectKey value")
	}

	params := url.Values{}
	params.Add("name", projectName)

	var endpoint = fmt.Sprintf("rest/api/3/projectvalidate/validProjectName?%v", params.Encode())

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	randomName = string(response.BodyAsBytes)

	return
}
