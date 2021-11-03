package v2

import (
	"context"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
)

type ServerService struct{ client *Client }

// Info returns information about the Jira instance.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/server#get-jira-instance-info
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-server-info/#api-rest-api-2-serverinfo-get
func (s *ServerService) Info(ctx context.Context) (result *models.ServerInformationScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/2/serverInfo"

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
