package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestSLAService struct{ client *Client }

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/sla#get-sla-information
func (r *RequestSLAService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (result *RequestSLAPageScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/sla?%v", issueKeyOrID, params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestSLAPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/sla#get-sla-information-by-id
func (r *RequestSLAService) Get(ctx context.Context, issueKeyOrID string, slaMetricID int) (result *RequestSLAScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/sla/%v", issueKeyOrID, slaMetricID)

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestSLAScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type RequestSLAPageScheme struct {
	Size       int                 `json:"size"`
	Start      int                 `json:"start"`
	Limit      int                 `json:"limit"`
	IsLastPage bool                `json:"isLastPage"`
	Values     []*RequestSLAScheme `json:"values"`
	Expands    []string            `json:"_expands"`
	Links      struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}

type RequestSLAScheme struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	CompletedCycles []struct {
		StartTime struct {
		} `json:"startTime"`
		StopTime struct {
		} `json:"stopTime"`
		Breached     bool `json:"breached"`
		GoalDuration struct {
		} `json:"goalDuration"`
		ElapsedTime struct {
		} `json:"elapsedTime"`
		RemainingTime struct {
		} `json:"remainingTime"`
	} `json:"completedCycles"`
	OngoingCycle struct {
		StartTime struct {
		} `json:"startTime"`
		BreachTime struct {
		} `json:"breachTime"`
		Breached            bool `json:"breached"`
		Paused              bool `json:"paused"`
		WithinCalendarHours bool `json:"withinCalendarHours"`
		GoalDuration        struct {
		} `json:"goalDuration"`
		ElapsedTime struct {
		} `json:"elapsedTime"`
		RemainingTime struct {
		} `json:"remainingTime"`
	} `json:"ongoingCycle"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}
