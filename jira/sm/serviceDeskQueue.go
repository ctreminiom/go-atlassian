package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type ServiceDeskQueueService struct{ client *Client }

func (s *ServiceDeskQueueService) Gets(ctx context.Context, serviceDeskID int, includeCount bool, start, limit int) (result *ServiceDeskQueuePageScheme, response *Response, err error) {

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

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ServiceDeskQueuePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (s *ServiceDeskQueueService) Get(ctx context.Context, serviceDeskID, queueID int, includeCount bool) (result *ServiceDeskQueueScheme, response *Response, err error) {

	params := url.Values{}
	if includeCount {
		params.Add("includeCount", "true")
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue/%v?%v", serviceDeskID, queueID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue/%v", serviceDeskID, queueID)
	}

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ServiceDeskQueueScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (s *ServiceDeskQueueService) Issues(ctx context.Context, serviceDeskID, queueID, start, limit int) (result *ServiceDeskIssueQueueScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue/%v/issue?%v", serviceDeskID, queueID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ServiceDeskIssueQueueScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ServiceDeskQueuePageScheme struct {
	Size       int                       `json:"size"`
	Start      int                       `json:"start"`
	Limit      int                       `json:"limit"`
	IsLastPage bool                      `json:"isLastPage"`
	Values     []*ServiceDeskQueueScheme `json:"values"`
	Expands    []string                  `json:"_expands"`
	Links      struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}

type ServiceDeskQueueScheme struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Jql        string   `json:"jql"`
	Fields     []string `json:"fields"`
	IssueCount int      `json:"issueCount"`
	Links      struct {
		Self string `json:"self"`
	} `json:"_links"`
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
