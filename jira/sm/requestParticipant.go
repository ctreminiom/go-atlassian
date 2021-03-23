package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestParticipantService struct{ client *Client }

func (r *RequestParticipantService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (result *RequestParticipantPageScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/participant?%v", issueKeyOrID, params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestParticipantPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (r *RequestParticipantService) Add(ctx context.Context, issueKeyOrID string, accountIDs []string) (result *RequestParticipantPageScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	if len(accountIDs) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid accountIDs slice value")
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/participant", issueKeyOrID)

	request, err := r.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestParticipantPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (r *RequestParticipantService) Remove(ctx context.Context, issueKeyOrID string, accountIDs []string) (result *RequestParticipantPageScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	if len(accountIDs) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid accountIDs slice value")
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/participant", issueKeyOrID)

	request, err := r.client.newRequest(ctx, http.MethodDelete, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestParticipantPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type RequestParticipantPageScheme struct {
	Size       int                         `json:"size"`
	Start      int                         `json:"start"`
	Limit      int                         `json:"limit"`
	IsLastPage bool                        `json:"isLastPage"`
	Values     []*RequestParticipantScheme `json:"values"`
	Expands    []string                    `json:"_expands"`
	Links      struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}

type RequestParticipantScheme struct {
	AccountID    string `json:"accountId"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
	TimeZone     string `json:"timeZone"`
	Links        struct {
		Self       string `json:"self"`
		JiraRest   string `json:"jiraRest"`
		AvatarUrls struct {
		} `json:"avatarUrls"`
	} `json:"_links"`
}
