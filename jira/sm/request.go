package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestService struct {
	client      *Client
	Type        *RequestTypeService
	Approval    *RequestApprovalService
	Attachment  *RequestAttachmentService
	Comment     *RequestCommentService
	Participant *RequestParticipantService
	SLA         *RequestSLAService
}

type RequestGetOptionsScheme struct {
	SearchTerm        string
	RequestOwnerships []string
	RequestStatus     string
	ApprovalStatus    string
	OrganizationId    int
	ServiceDeskID     int
	RequestTypeID     int
	Expand            []string
}

// This method returns all customer requests for the user executing the query.
// The returned customer requests are ordered chronologically by the latest activity on each request. For example, the latest status transition or comment.
func (r *RequestService) Gets(ctx context.Context, opts *RequestGetOptionsScheme, start, limit int) (result *CustomerRequestsScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if opts != nil {

		if len(opts.SearchTerm) != 0 {
			params.Add("searchTerm", opts.SearchTerm)
		}

		for _, requestOwner := range opts.RequestOwnerships {
			params.Add("requestOwnership", requestOwner)
		}

		if len(opts.RequestStatus) != 0 {
			params.Add("requestStatus", opts.RequestStatus)
		}

		if len(opts.ApprovalStatus) != 0 {
			params.Add("approvalStatus", opts.ApprovalStatus)
		}

		if opts.OrganizationId != 0 {
			params.Add("organizationId", strconv.Itoa(opts.OrganizationId))
		}

		if opts.ServiceDeskID != 0 {
			params.Add("serviceDeskId", strconv.Itoa(opts.ServiceDeskID))
		}

		if opts.RequestTypeID != 0 {
			params.Add("requestTypeId", strconv.Itoa(opts.RequestTypeID))
		}

		var expand string
		for index, value := range opts.Expand {

			if index == 0 {
				expand = value
				continue
			}

			expand += "," + value
		}

		if len(expand) != 0 {
			params.Add("expand", expand)
		}

	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request?%v", params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomerRequestsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// This method returns a customer request.
func (r *RequestService) Get(ctx context.Context, issueKeyOrID string, expands []string) (result *CustomerRequestScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	params := url.Values{}

	if len(expands) != 0 {

		var expand string
		for index, value := range expands {

			if index == 0 {
				expand = value
				continue
			}

			expand += "," + value
		}

		if len(expand) != 0 {
			params.Add("expand", expand)
		}
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v?%v", issueKeyOrID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v", issueKeyOrID)
	}

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomerRequestScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (r *RequestService) Subscribe(ctx context.Context, issueKeyOrID string) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/notification", issueKeyOrID)

	request, err := r.client.newRequest(ctx, http.MethodPut, endpoint, nil)
	if err != nil {
		return
	}

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	return
}

func (r *RequestService) Unsubscribe(ctx context.Context, issueKeyOrID string) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/notification", issueKeyOrID)

	request, err := r.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	return
}

type CustomerRequestsScheme struct {
	Size       int                      `json:"size"`
	Start      int                      `json:"start"`
	Limit      int                      `json:"limit"`
	IsLastPage bool                     `json:"isLastPage"`
	Values     []*CustomerRequestScheme `json:"values"`
	Expands    []string                 `json:"_expands"`
	Links      struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}

type CustomerRequestScheme struct {
	IssueID       string `json:"issueId"`
	IssueKey      string `json:"issueKey"`
	RequestTypeID string `json:"requestTypeId"`
	RequestType   struct {
		ID            string   `json:"id"`
		Name          string   `json:"name"`
		Description   string   `json:"description"`
		HelpText      string   `json:"helpText"`
		IssueTypeID   string   `json:"issueTypeId"`
		ServiceDeskID string   `json:"serviceDeskId"`
		GroupIds      []string `json:"groupIds"`
		Icon          struct {
		} `json:"icon"`
		Fields struct {
		} `json:"fields"`
		Expands []string `json:"_expands"`
		Links   struct {
		} `json:"_links"`
	} `json:"requestType"`
	ServiceDeskID string `json:"serviceDeskId"`
	ServiceDesk   struct {
		ID          string `json:"id"`
		ProjectID   string `json:"projectId"`
		ProjectName string `json:"projectName"`
		ProjectKey  string `json:"projectKey"`
		Links       struct {
		} `json:"_links"`
	} `json:"serviceDesk"`
	CreatedDate struct {
		Iso8601     string `json:"iso8601"`
		Jira        string `json:"jira"`
		Friendly    string `json:"friendly"`
		EpochMillis int    `json:"epochMillis"`
	} `json:"createdDate"`
	Reporter struct {
		AccountID    string `json:"accountId"`
		Name         string `json:"name"`
		Key          string `json:"key"`
		EmailAddress string `json:"emailAddress"`
		DisplayName  string `json:"displayName"`
		Active       bool   `json:"active"`
		TimeZone     string `json:"timeZone"`
		Links        struct {
		} `json:"_links"`
	} `json:"reporter"`
	RequestFieldValues []struct {
		FieldID       string `json:"fieldId"`
		Label         string `json:"label"`
		RenderedValue struct {
		} `json:"renderedValue"`
	} `json:"requestFieldValues"`
	CurrentStatus struct {
		Status         string `json:"status"`
		StatusCategory string `json:"statusCategory"`
		StatusDate     struct {
		} `json:"statusDate"`
	} `json:"currentStatus"`
	Status struct {
		Size       int  `json:"size"`
		Start      int  `json:"start"`
		Limit      int  `json:"limit"`
		IsLastPage bool `json:"isLastPage"`
		Values     []struct {
		} `json:"values"`
		Expands []string `json:"_expands"`
		Links   struct {
		} `json:"_links"`
	} `json:"status"`
	Participants struct {
		Size       int  `json:"size"`
		Start      int  `json:"start"`
		Limit      int  `json:"limit"`
		IsLastPage bool `json:"isLastPage"`
		Values     []struct {
		} `json:"values"`
		Expands []string `json:"_expands"`
		Links   struct {
		} `json:"_links"`
	} `json:"participants"`
	SLA struct {
		Size       int  `json:"size"`
		Start      int  `json:"start"`
		Limit      int  `json:"limit"`
		IsLastPage bool `json:"isLastPage"`
		Values     []struct {
		} `json:"values"`
		Expands []string `json:"_expands"`
		Links   struct {
		} `json:"_links"`
	} `json:"sla"`
	Attachments struct {
		Size       int  `json:"size"`
		Start      int  `json:"start"`
		Limit      int  `json:"limit"`
		IsLastPage bool `json:"isLastPage"`
		Values     []struct {
		} `json:"values"`
		Expands []string `json:"_expands"`
		Links   struct {
		} `json:"_links"`
	} `json:"attachments"`
	Comments struct {
		Size       int  `json:"size"`
		Start      int  `json:"start"`
		Limit      int  `json:"limit"`
		IsLastPage bool `json:"isLastPage"`
		Values     []struct {
		} `json:"values"`
		Expands []string `json:"_expands"`
		Links   struct {
		} `json:"_links"`
	} `json:"comments"`
	Actions struct {
		AddAttachment struct {
		} `json:"addAttachment"`
		AddComment struct {
		} `json:"addComment"`
		AddParticipant struct {
		} `json:"addParticipant"`
		RemoveParticipant struct {
		} `json:"removeParticipant"`
	} `json:"actions"`
	Expands []string `json:"_expands"`
	Links   struct {
		Self     string `json:"self"`
		JiraRest string `json:"jiraRest"`
		Web      string `json:"web"`
		Agent    string `json:"agent"`
	} `json:"_links"`
}
