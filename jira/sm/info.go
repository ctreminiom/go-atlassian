package sm

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type InfoService struct{ client *Client }

// Get retrieves information about the Jira Service Management instance such as software version,
// builds, and related links.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/info#get-info
func (i *InfoService) Get(ctx context.Context) (result *model.InfoScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/servicedeskapi/info"

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}
