package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProjectTypeService struct{ client *Client }

type ProjectTypeScheme struct {
	Key                string `json:"key"`
	FormattedKey       string `json:"formattedKey"`
	DescriptionI18NKey string `json:"descriptionI18nKey"`
	Icon               string `json:"icon"`
	Color              string `json:"color"`
}

// Returns all project types, whether or not the instance has a valid license for each type.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-types/#api-rest-api-3-project-type-get
func (p *ProjectTypeService) Gets(ctx context.Context) (result *[]ProjectTypeScheme, response *Response, err error) {

	var endpoint = "rest/api/3/project/type"

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ProjectTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns all project types with a valid license.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-types/#api-rest-api-3-project-type-accessible-get
func (p *ProjectTypeService) Licensed(ctx context.Context) (result *[]ProjectTypeScheme, response *Response, err error) {

	var endpoint = "rest/api/3/project/type/accessible"

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ProjectTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a project type.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-types/#api-rest-api-3-project-type-projecttypekey-get
func (p *ProjectTypeService) Get(ctx context.Context, projectTypeKey string) (result *ProjectTypeScheme, response *Response, err error) {

	if len(projectTypeKey) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectTypeKey value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/type/%v", projectTypeKey)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a project type if it is accessible to the user.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-types/#api-rest-api-3-project-type-projecttypekey-accessible-get
func (p *ProjectTypeService) Accessible(ctx context.Context, projectTypeKey string) (result *ProjectTypeScheme, response *Response, err error) {

	if len(projectTypeKey) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectTypeKey value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/type/%v/accessible", projectTypeKey)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
