package jira

import (
	"context"
	"net/http"
)

type ServerService struct{ client *Client }

type ServerInformationScheme struct {
	BaseURL        string `json:"baseUrl"`
	Version        string `json:"version"`
	VersionNumbers []int  `json:"versionNumbers"`
	DeploymentType string `json:"deploymentType"`
	BuildNumber    int    `json:"buildNumber"`
	BuildDate      string `json:"buildDate"`
	ServerTime     string `json:"serverTime"`
	ScmInfo        string `json:"scmInfo"`
	ServerTitle    string `json:"serverTitle"`
	HealthChecks   []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Passed      bool   `json:"passed"`
	} `json:"healthChecks"`
}

// Info returns information about the Jira instance.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/server#get-jira-instance-info
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-server-info/#api-rest-api-3-serverinfo-get
func (s *ServerService) Info(ctx context.Context) (result *ServerInformationScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/serverInfo"

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
