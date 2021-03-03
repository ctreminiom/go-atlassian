package jira

import (
	"context"
	"encoding/json"
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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/application-roles#application-roles
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

// Returns an application role.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/application-roles#application-role
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
