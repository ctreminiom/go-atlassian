package sm

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type RequestTypeService struct{ client *Client }

// Search returns all customer request types used in the Jira Service Management instance,
// optionally filtered by a query string.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-all-request-types
func (r *RequestTypeService) Search(ctx context.Context, query string, start, limit int) (result *model.RequestTypePageScheme,
	response *ResponseScheme, err error) {

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

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Gets returns all customer request types from a service desk.
// There are two parameters for filtering the returned list:
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-request-types
func (r *RequestTypeService) Gets(ctx context.Context, serviceDeskID, groupID, start, limit int) (
	result *model.ProjectRequestTypePageScheme, response *ResponseScheme, err error) {

	if serviceDeskID == 0 {
		return nil, nil, model.ErrNoServiceDeskIDError
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

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Create enables a customer request type to be added to a service desk based on an issue type.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#create-request-type
func (r *RequestTypeService) Create(ctx context.Context, serviceDeskID int, issueTypeID, name, description,
	helpText string) (result *model.RequestTypeScheme, response *ResponseScheme, err error) {

	if serviceDeskID == 0 {
		return nil, nil, model.ErrNoServiceDeskIDError
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

	payloadAsReader, _ := transformStructToReader(&payload)
	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/requesttype", serviceDeskID)

	request, err := r.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
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

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns a customer request type from a service desk.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-request-type-by-id
func (r *RequestTypeService) Get(ctx context.Context, serviceDeskID, requestTypeID int) (result *model.RequestTypeScheme,
	response *ResponseScheme, err error) {

	if serviceDeskID == 0 {
		return nil, nil, model.ErrNoServiceDeskIDError
	}

	if requestTypeID == 0 {
		return nil, nil, model.ErrNoRequestTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/requesttype/%v", serviceDeskID, requestTypeID)

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes a customer request type from a service desk, and removes it from all customer requests.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#delete-request-type
func (r *RequestTypeService) Delete(ctx context.Context, serviceDeskID, requestTypeID int) (response *ResponseScheme, err error) {

	if serviceDeskID == 0 {
		return nil, model.ErrNoServiceDeskIDError
	}

	if requestTypeID == 0 {
		return nil, model.ErrNoRequestTypeIDError
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

	response, err = r.client.Call(request, nil)
	if err != nil {
		return
	}

	return
}

// Fields returns the fields for a service desk's customer request type.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-request-type-fields
func (r *RequestTypeService) Fields(ctx context.Context, serviceDeskID, requestTypeID int) (
	result *model.RequestTypeFieldsScheme, response *ResponseScheme, err error) {

	if serviceDeskID == 0 {
		return nil, nil, model.ErrNoServiceDeskIDError
	}

	if requestTypeID == 0 {
		return nil, nil, model.ErrNoRequestTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/requesttype/%v/field", serviceDeskID, requestTypeID)

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}
