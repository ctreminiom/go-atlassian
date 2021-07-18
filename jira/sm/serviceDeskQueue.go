package sm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ServiceDeskQueueService struct{ client *Client }

// Gets returns the queues in a service desk
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-queues
func (s *ServiceDeskQueueService) Gets(ctx context.Context, serviceDeskID int, includeCount bool, start, limit int) (
	result *ServiceDeskQueuePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if includeCount {
		params.Add("includeCount", "true")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue?%v", serviceDeskID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns a specific queues in a service desk.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-queue
func (s *ServiceDeskQueueService) Get(ctx context.Context, serviceDeskID, queueID int, includeCount bool) (
	result *ServiceDeskQueueScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	if includeCount {
		params.Add("includeCount", "true")
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue/%v", serviceDeskID, queueID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Issues returns the customer requests in a queue
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-issues-in-queue
func (s *ServiceDeskQueueService) Issues(ctx context.Context, serviceDeskID, queueID, start, limit int) (
	result *ServiceDeskIssueQueueScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue/%v/issue?%v", serviceDeskID, queueID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

type ServiceDeskQueuePageScheme struct {
	Size       int                             `json:"size,omitempty"`
	Start      int                             `json:"start,omitempty"`
	Limit      int                             `json:"limit,omitempty"`
	IsLastPage bool                            `json:"isLastPage,omitempty"`
	Values     []*ServiceDeskQueueScheme       `json:"values,omitempty"`
	Expands    []string                        `json:"_expands,omitempty"`
	Links      *ServiceDeskQueuePageLinkScheme `json:"_links,omitempty"`
}

type ServiceDeskQueuePageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type ServiceDeskQueueScheme struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Jql        string   `json:"jql,omitempty"`
	Fields     []string `json:"fields,omitempty"`
	IssueCount int      `json:"issueCount,omitempty"`
	Links      struct {
		Self string `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type ServiceDeskIssueQueueScheme struct {
	Size       int  `json:"size"`
	Start      int  `json:"start"`
	Limit      int  `json:"limit"`
	IsLastPage bool `json:"isLastPage"`
	Values     []struct {
		Expand      string `json:"expand"`
		ID          string `json:"id"`
		Self        string `json:"self"`
		Key         string `json:"key"`
		Transitions []struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			HasScreen     bool   `json:"hasScreen"`
			IsGlobal      bool   `json:"isGlobal"`
			IsInitial     bool   `json:"isInitial"`
			IsAvailable   bool   `json:"isAvailable"`
			IsConditional bool   `json:"isConditional"`
			Expand        string `json:"expand"`
			Looped        bool   `json:"looped"`
		} `json:"transitions"`

		Changelog struct {
			StartAt    int `json:"startAt"`
			MaxResults int `json:"maxResults"`
			Total      int `json:"total"`
			Histories  []struct {
			} `json:"histories"`
		} `json:"changelog"`

		FieldsToInclude struct {
			Included         []string `json:"included"`
			ActuallyIncluded []string `json:"actuallyIncluded"`
			Excluded         []string `json:"excluded"`
		} `json:"fieldsToInclude"`
		Fields struct {
		} `json:"fields"`
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
