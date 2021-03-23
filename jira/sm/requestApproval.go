package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestApprovalService struct{ client *Client }

func (r *RequestApprovalService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (result *CustomerApprovalsScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/approval?%v", issueKeyOrID, params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomerApprovalsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (r *RequestApprovalService) Get(ctx context.Context, issueKeyOrID string, approvalID int) (result *CustomerApprovalScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/approval/%v", issueKeyOrID, approvalID)

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomerApprovalScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (r *RequestApprovalService) Answer(ctx context.Context, issueKeyOrID string, approvalID int, approve bool) (result *CustomerApprovalScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/approval/%v", issueKeyOrID, approvalID)

	//Transform the approve bool into the request payload
	var approveAsString string
	if approve {
		approveAsString = "approve"
	} else {
		approveAsString = "decline"
	}

	payload := struct {
		Decision string `json:"decision"`
	}{
		Decision: approveAsString,
	}

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

	result = new(CustomerApprovalScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type CustomerApprovalsScheme struct {
	Size       int                       `json:"size"`
	Start      int                       `json:"start"`
	Limit      int                       `json:"limit"`
	IsLastPage bool                      `json:"isLastPage"`
	Values     []*CustomerApprovalScheme `json:"values"`
	Expands    []string                  `json:"_expands"`
	Links      struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}

type CustomerApprovalScheme struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	FinalDecision     string `json:"finalDecision"`
	CanAnswerApproval bool   `json:"canAnswerApproval"`
	Approvers         []struct {
		Approver struct {
			AccountID    string `json:"accountId"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			DisplayName  string `json:"displayName"`
			Active       bool   `json:"active"`
			TimeZone     string `json:"timeZone"`
			Links        struct {
				JiraRest   string `json:"jiraRest"`
				AvatarUrls struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				Self string `json:"self"`
			} `json:"_links"`
		} `json:"approver"`
		ApproverDecision string `json:"approverDecision"`
	} `json:"approvers"`
	CreatedDate struct {
		Iso8601     string `json:"iso8601"`
		Jira        string `json:"jira"`
		Friendly    string `json:"friendly"`
		EpochMillis int64  `json:"epochMillis"`
	} `json:"createdDate"`
	CompletedDate struct {
		Iso8601     string `json:"iso8601"`
		Jira        string `json:"jira"`
		Friendly    string `json:"friendly"`
		EpochMillis int64  `json:"epochMillis"`
	} `json:"completedDate"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}
