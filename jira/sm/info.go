package sm

import (
	"context"
	"net/http"
)

type InfoService struct{ client *Client }

// Get retrieves information about the Jira Service Management instance such as software version,
// builds, and related links.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/info#get-info
func (i *InfoService) Get(ctx context.Context) (result *InfoScheme, response *ResponseScheme, err error) {

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

type InfoScheme struct {
	Version          string               `json:"version"`
	PlatformVersion  string               `json:"platformVersion"`
	BuildDate        *InfoBuildDataScheme `json:"buildDate"`
	BuildChangeSet   string               `json:"buildChangeSet"`
	IsLicensedForUse bool                 `json:"isLicensedForUse"`
	Links            *InfoLinkScheme      `json:"_links"`
}

type InfoBuildDataScheme struct {
	Iso8601     string `json:"iso8601"`
	Jira        string `json:"jira"`
	Friendly    string `json:"friendly"`
	EpochMillis int64  `json:"epochMillis"`
}

type InfoLinkScheme struct {
	Self string `json:"self"`
}
