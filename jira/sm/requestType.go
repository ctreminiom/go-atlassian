package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestTypeService struct{ client *Client }

// This method returns all customer request types used in the Jira Service Management instance,
// optionally filtered by a query string.
func (r *RequestTypeService) Gets(ctx context.Context, query string, start, limit int) (result *RequestTypePageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if len(query) != 0 {
		params.Add("searchQuery", query)
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/requesttype?%v", params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestTypePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type RequestTypePageScheme struct {
	Size       int  `json:"size"`
	Start      int  `json:"start"`
	Limit      int  `json:"limit"`
	IsLastPage bool `json:"isLastPage"`
	Values     []struct {
		ID            string   `json:"id"`
		Name          string   `json:"name"`
		Description   string   `json:"description"`
		HelpText      string   `json:"helpText"`
		IssueTypeID   string   `json:"issueTypeId"`
		ServiceDeskID string   `json:"serviceDeskId"`
		GroupIds      []string `json:"groupIds"`
		Icon          struct {
			ID    string `json:"id"`
			Links struct {
			} `json:"_links"`
		} `json:"icon"`
		Fields struct {
			RequestTypeFields []struct {
			} `json:"requestTypeFields"`
			CanRaiseOnBehalfOf        bool `json:"canRaiseOnBehalfOf"`
			CanAddRequestParticipants bool `json:"canAddRequestParticipants"`
		} `json:"fields"`
		Expands []string `json:"_expands"`
		Links   struct {
			Self string `json:"self"`
		} `json:"_links"`
	} `json:"values"`
	Expands []string `json:"_expands"`
	Links   struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}
