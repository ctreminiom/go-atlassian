package jira

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ProjectPropertyService struct{ client *Client }

type ProjectPropertiesScheme struct {
	Keys []struct {
		Self string `json:"self"`
		Key  string `json:"key"`
	} `json:"keys"`
}

// Returns all project property keys for the project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-properties/#api-rest-api-3-project-projectidorkey-properties-get
func (p *ProjectPropertyService) Keys(ctx context.Context, projectKeyOrID string) (result *ProjectPropertiesScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/properties", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectPropertiesScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ProjectPropertyScheme struct {
	Key   string `json:"key"`
	Value struct {
		SystemConversationID string `json:"system.conversation.id"`
		SystemSupportTime    string `json:"system.support.time"`
	} `json:"value"`
}

// Returns the value of a project property.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-properties/#api-rest-api-3-project-projectidorkey-properties-propertykey-get
func (p *ProjectPropertyService) Get(ctx context.Context, projectKeyOrID, propertyKey string) (result *ProjectPropertyScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/properties/%v", projectKeyOrID, propertyKey)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectPropertyScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Sets the value of the project property.
// You can use project properties to store custom data against the project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-properties/#api-rest-api-3-project-projectidorkey-properties-propertykey-put
func (p *ProjectPropertyService) Set(ctx context.Context, projectKeyOrID, propertyKey string, propertyValue interface{}) (response *Response, err error) {

	payload := make(map[string]interface{})

	switch propertyValue.(type) {

	case int:
		payload["number"] = propertyValue.(int)
	case string:
		payload["string"] = propertyValue.(string)
	default:
		err = errors.New("unable string and int values are permitted")
		return
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/properties/%v", projectKeyOrID, propertyKey)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, &payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Deletes the property from a project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-properties/#api-rest-api-3-project-projectidorkey-properties-propertykey-delete
func (p *ProjectPropertyService) Delete(ctx context.Context, projectKeyOrID, propertyKey string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/properties/%v", projectKeyOrID, propertyKey)

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
