package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestCommentService struct{ client *Client }

func (r *RequestCommentService) Gets(ctx context.Context, issueKeyOrID string, public bool, expands []string, start, limit int) (result *RequestCommentPageScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if !public {
		params.Add("public", "false")
	}

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

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/comment?%v", issueKeyOrID, params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestCommentPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (r *RequestCommentService) Get(ctx context.Context, issueKeyOrID string, commentID int, expands []string) (result *RequestCommentScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	params := url.Values{}
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

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/comment/%v?%v", issueKeyOrID, commentID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/comment/%v", issueKeyOrID, commentID)
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

	result = new(RequestCommentScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (r *RequestCommentService) Create(ctx context.Context, issueKeyOrID, body string, public bool) (result *RequestCommentScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	if len(body) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid body value")
	}

	payload := struct {
		Public bool   `json:"public"`
		Body   string `json:"body"`
	}{
		Public: public,
		Body:   body,
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/comment", issueKeyOrID)

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

	result = new(RequestCommentScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (r *RequestCommentService) Attachments(ctx context.Context, issueKeyOrID string, commentID, start, limit int) (result *RequestCommentAttachmentPageScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/comment/%v/attachment?%v", issueKeyOrID, commentID, params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(RequestCommentAttachmentPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type RequestCommentAttachmentPageScheme struct {
	Size       int  `json:"size"`
	Start      int  `json:"start"`
	Limit      int  `json:"limit"`
	IsLastPage bool `json:"isLastPage"`
	Values     []struct {
		Filename string `json:"filename"`
		Author   struct {
			AccountID    string `json:"accountId"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			DisplayName  string `json:"displayName"`
			Active       bool   `json:"active"`
			TimeZone     string `json:"timeZone"`
			Links        struct {
			} `json:"_links"`
		} `json:"author"`
		Created struct {
			Iso8601     string `json:"iso8601"`
			Jira        string `json:"jira"`
			Friendly    string `json:"friendly"`
			EpochMillis int    `json:"epochMillis"`
		} `json:"created"`
		Size     int    `json:"size"`
		MimeType string `json:"mimeType"`
		Links    struct {
			Self      string `json:"self"`
			JiraRest  string `json:"jiraRest"`
			Content   string `json:"content"`
			Thumbnail string `json:"thumbnail"`
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

type RequestCommentPageScheme struct {
	Size       int                     `json:"size"`
	Start      int                     `json:"start"`
	Limit      int                     `json:"limit"`
	IsLastPage bool                    `json:"isLastPage"`
	Values     []*RequestCommentScheme `json:"values"`
	Expands    []string                `json:"_expands"`
	Links      struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}

type RequestCommentScheme struct {
	ID           string `json:"id"`
	Body         string `json:"body"`
	RenderedBody struct {
		HTML string `json:"html"`
	} `json:"renderedBody"`
	Author struct {
		AccountID    string `json:"accountId"`
		Name         string `json:"name"`
		Key          string `json:"key"`
		EmailAddress string `json:"emailAddress"`
		DisplayName  string `json:"displayName"`
		Active       bool   `json:"active"`
		TimeZone     string `json:"timeZone"`
		Links        struct {
		} `json:"_links"`
	} `json:"author"`
	Created struct {
		Iso8601     string `json:"iso8601"`
		Jira        string `json:"jira"`
		Friendly    string `json:"friendly"`
		EpochMillis int    `json:"epochMillis"`
	} `json:"created"`
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
	Expands []string `json:"_expands"`
	Public  bool     `json:"public"`
	Links   struct {
		Self string `json:"self"`
	} `json:"_links"`
}
