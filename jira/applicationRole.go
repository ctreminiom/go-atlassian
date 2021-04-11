package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// This service represents the Jira Cloud application roles
// Use it to get details of an application role or all application roles.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/application-roles
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

// Returns all application roles, this func needs the following parameters:
// 1. ctx = it's the context.context value
// Docs: https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-all-application-roles
func (a *ApplicationRoleService) Gets(ctx context.Context) (result *[]ApplicationRoleScheme, response *Response, err error) {

	var endpoint = "rest/api/3/applicationrole"
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = a.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ApplicationRoleScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns an application role, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. key = The key of the application role, use Gets() method to get the key for each application role.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-application-role
func (a *ApplicationRoleService) Get(ctx context.Context, key string) (result *ApplicationRoleScheme, response *Response, err error) {

	if len(key) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid key application role value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/applicationrole/%v", key)
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = a.client.Do(request)
	if err != nil {
		return
	}

	result = new(ApplicationRoleScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
