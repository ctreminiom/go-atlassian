package jira

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ApplicationRoleService struct{ client *Client }

type ApplicationRoleScheme struct {
	Key                  string   `json:"key,omitempty"`
	Groups               []string `json:"groups,omitempty"`
	Name                 string   `json:"name,omitempty"`
	DefaultGroups        []string `json:"defaultGroups,omitempty"`
	SelectedByDefault    bool     `json:"selectedByDefault,omitempty"`
	Defined              bool     `json:"defined,omitempty"`
	NumberOfSeats        int      `json:"numberOfSeats,omitempty"`
	RemainingSeats       int      `json:"remainingSeats,omitempty"`
	UserCount            int      `json:"userCount,omitempty"`
	UserCountDescription string   `json:"userCountDescription,omitempty"`
	HasUnlimitedSeats    bool     `json:"hasUnlimitedSeats,omitempty"`
	Platform             bool     `json:"platform,omitempty"`
}

// Returns all application roles.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-application-roles/#api-rest-api-3-applicationrole-get
func (a *ApplicationRoleService) Gets(ctx context.Context) (result *[]ApplicationRoleScheme, response *Response, err error) {

	if ctx == nil {
		return nil, nil, errors.New("the context param is nil, please provide a valid one")
	}

	var endpoint = "rest/api/3/applicationrole"
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = a.client.Do(request)
	if err != nil {
		return
	}

	if len(response.BodyAsBytes) == 0 {
		return nil, nil, errors.New("unable to marshall the response body, the HTTP callback did not return any bytes")
	}

	result = new([]ApplicationRoleScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns an application role.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-application-roles/#api-rest-api-3-applicationrole-key-get
func (a *ApplicationRoleService) Get(ctx context.Context, key string) (result *ApplicationRoleScheme, response *Response, err error) {

	if ctx == nil {
		return nil, nil, errors.New("the context param is nil, please provide a valid one")
	}

	var endpoint = fmt.Sprintf("rest/api/3/applicationrole/%v", key)
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = a.client.Do(request)
	if err != nil {
		return
	}

	if len(response.BodyAsBytes) == 0 {
		return nil, nil, errors.New("unable to marshall the response body, the HTTP callback did not return any bytes")
	}

	result = new(ApplicationRoleScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
