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
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-all-request-types
func (r *RequestTypeService) Search(ctx context.Context, query string, start, limit int) (result *RequestTypePageScheme, response *Response, err error) {

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

	/*
		This API is experimental.
		Experimental APIs are not guaranteed to be stable within the preview period.
	*/
	request.Header.Set("X-ExperimentalApi", "opt-in")

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

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-request-types
func (r *RequestTypeService) Gets(ctx context.Context, serviceDeskID, groupID, start, limit int) (result *ProjectRequestTypePageScheme, response *Response, err error) {

	if serviceDeskID == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid serviceDeskID value")
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if groupID != 0 {
		params.Add("groupId", strconv.Itoa(groupID))
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/requesttype?%v", serviceDeskID, params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectRequestTypePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ProjectRequestTypePageScheme struct {
	Expands    []string `json:"_expands"`
	Size       int      `json:"size"`
	Start      int      `json:"start"`
	Limit      int      `json:"limit"`
	IsLastPage bool     `json:"isLastPage"`
	Links      struct {
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
	Values []*RequestTypeScheme `json:"values"`
}

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#create-request-type
func (r *RequestTypeService) Create(ctx context.Context, serviceDeskID int, issueTypeID, name, description, helpText string) (result *RequestTypeScheme, response *Response, err error) {

	if serviceDeskID == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid serviceDeskID value")
	}

	payload := struct {
		IssueTypeID string `json:"issueTypeId,omitempty"`
		HelpText    string `json:"helpText,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		IssueTypeID: issueTypeID,
		HelpText:    helpText,
		Name:        name,
		Description: description,
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/requesttype", serviceDeskID)

	request, err := r.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	/*
		This API is experimental.
		Experimental APIs are not guaranteed to be stable within the preview period.
	*/
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-request-type-by-id
func (r *RequestTypeService) Get(ctx context.Context, serviceDeskID, requestTypeID int) (result *RequestTypeScheme, response *Response, err error) {

	if serviceDeskID == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid serviceDeskID value")
	}

	if requestTypeID == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid requestTypeID value")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/requesttype/%v", serviceDeskID, requestTypeID)

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#delete-request-type
func (r *RequestTypeService) Delete(ctx context.Context, serviceDeskID, requestTypeID int) (response *Response, err error) {

	if serviceDeskID == 0 {
		return nil, fmt.Errorf("error, please provide a valid serviceDeskID value")
	}

	if requestTypeID == 0 {
		return nil, fmt.Errorf("error, please provide a valid requestTypeID value")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/requesttype/%v", serviceDeskID, requestTypeID)

	request, err := r.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	/*
		This API is experimental.
		Experimental APIs are not guaranteed to be stable within the preview period.
	*/
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-request-type-fields
func (r *RequestTypeService) Fields(ctx context.Context, serviceDeskID, requestTypeID int) (result *RequestTypeFieldsScheme, response *Response, err error) {

	if serviceDeskID == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid serviceDeskID value")
	}

	if requestTypeID == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid requestTypeID value")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/requesttype/%v/field", serviceDeskID, requestTypeID)

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestTypeFieldsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type RequestTypePageScheme struct {
	Size       int                  `json:"size"`
	Start      int                  `json:"start"`
	Limit      int                  `json:"limit"`
	IsLastPage bool                 `json:"isLastPage"`
	Values     []*RequestTypeScheme `json:"values"`
	Expands    []string             `json:"_expands"`
	Links      struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}

type RequestTypeScheme struct {
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
}

type RequestTypeFieldsScheme struct {
	RequestTypeFields []struct {
		FieldID       string `json:"fieldId"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		Required      bool   `json:"required"`
		DefaultValues []struct {
			Value    string        `json:"value"`
			Label    string        `json:"label"`
			Children []interface{} `json:"children"`
		} `json:"defaultValues"`
		ValidValues []struct {
			Value    string        `json:"value"`
			Label    string        `json:"label"`
			Children []interface{} `json:"children"`
		} `json:"validValues"`
		JiraSchema struct {
			Type          string `json:"type"`
			Items         string `json:"items"`
			System        string `json:"system"`
			Custom        string `json:"custom"`
			CustomID      int    `json:"customId"`
			Configuration struct {
			} `json:"configuration"`
		} `json:"jiraSchema"`
		Visible bool `json:"visible"`
	} `json:"requestTypeFields"`
	CanRaiseOnBehalfOf        bool `json:"canRaiseOnBehalfOf"`
	CanAddRequestParticipants bool `json:"canAddRequestParticipants"`
}
