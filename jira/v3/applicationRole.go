package v3

import (
	"context"
	"fmt"
	"net/http"
)

// ApplicationRoleService represents the Jira Cloud application roles
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

// Gets returns all application roles
// Docs: https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-all-application-roles
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-application-roles/#api-rest-api-3-applicationrole-get
func (a *ApplicationRoleService) Gets(ctx context.Context) (result []*ApplicationRoleScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/applicationrole"

	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = a.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns an application role, this func needs the following parameters:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-application-role
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-application-roles/#api-rest-api-3-applicationrole-key-get
func (a *ApplicationRoleService) Get(ctx context.Context, key string) (result *ApplicationRoleScheme, response *ResponseScheme, err error) {

	if key == "" {
		return nil, nil, notApplicationRoleKeyError
	}

	var endpoint = fmt.Sprintf("rest/api/3/applicationrole/%v", key)
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = a.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

var (
	notApplicationRoleKeyError = fmt.Errorf("error, please provide a valid key application role value")
)
