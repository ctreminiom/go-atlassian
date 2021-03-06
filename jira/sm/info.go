package sm

import (
	"context"
	"encoding/json"
	"net/http"
)

type InfoService struct{ client *Client }

// This method retrieves information about the Jira Service Management instance such as software version,
// builds, and related links.
func (i *InfoService) Get(ctx context.Context) (result *InfoScheme, response *Response, err error) {

	var endpoint = "rest/servicedeskapi/info"

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(InfoScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type InfoScheme struct {
	Version         string `json:"version"`
	PlatformVersion string `json:"platformVersion"`
	BuildDate       struct {
		Iso8601     string `json:"iso8601"`
		Jira        string `json:"jira"`
		Friendly    string `json:"friendly"`
		EpochMillis int64  `json:"epochMillis"`
	} `json:"buildDate"`
	BuildChangeSet   string `json:"buildChangeSet"`
	IsLicensedForUse bool   `json:"isLicensedForUse"`
	Links            struct {
		Self string `json:"self"`
	} `json:"_links"`
}
