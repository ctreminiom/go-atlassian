package sm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestAttachmentService struct{ client *Client }

// Gets returns all the attachments for a customer requests.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/attachment#get-attachments-for-request
func (r *RequestAttachmentService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (
	result *RequestAttachmentPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/attachment?%v", issueKeyOrID, params.Encode())

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

// Create adds one or more temporary files (attached to the request's service desk using
// servicedesk/{serviceDeskId}/attachTemporaryFile) as attachments to a customer request
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/attachment#create-attachment
func (r *RequestAttachmentService) Create(ctx context.Context, issueKeyOrID string, temporaryAttachmentIDs []string,
	public bool) (result *RequestAttachmentCreationScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
	}

	if len(temporaryAttachmentIDs) == 0 {
		return nil, nil, notAttachmentsError
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/attachment", issueKeyOrID)

	payload := struct {
		TemporaryAttachmentIds []string `json:"temporaryAttachmentIds,omitempty"`
		Public                 bool     `json:"public,omitempty"`
	}{
		TemporaryAttachmentIds: temporaryAttachmentIDs,
		Public:                 public,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	request, err := r.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

type RequestAttachmentPageScheme struct {
	Size       int                              `json:"size,omitempty"`
	Start      int                              `json:"start,omitempty"`
	Limit      int                              `json:"limit,omitempty"`
	IsLastPage bool                             `json:"isLastPage,omitempty"`
	Values     []*RequestAttachmentScheme       `json:"values,omitempty"`
	Expands    []string                         `json:"_expands,omitempty"`
	Links      *RequestAttachmentPageLinkScheme `json:"_links,omitempty"`
}

type RequestAttachmentPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type RequestAttachmentScheme struct {
	Filename string                       `json:"filename,omitempty"`
	Author   *RequestAuthorScheme         `json:"author,omitempty"`
	Created  *CustomerRequestDateScheme   `json:"created,omitempty"`
	Size     int                          `json:"size,omitempty"`
	MimeType string                       `json:"mimeType,omitempty"`
	Links    *RequestAttachmentLinkScheme `json:"_links,omitempty"`
}

type RequestAttachmentLinkScheme struct {
	Self      string `json:"self,omitempty"`
	JiraRest  string `json:"jiraRest,omitempty"`
	Content   string `json:"content,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type RequestAuthorScheme struct {
	AccountID    string `json:"accountId,omitempty"`
	Name         string `json:"name,omitempty"`
	Key          string `json:"key,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Active       bool   `json:"active,omitempty"`
	TimeZone     string `json:"timeZone,omitempty"`
}

type RequestAttachmentCreationCommentScheme struct {
	Expands []string                   `json:"_expands,omitempty"`
	ID      string                     `json:"id,omitempty"`
	Body    string                     `json:"body,omitempty"`
	Public  bool                       `json:"public,omitempty"`
	Author  RequestAuthorScheme        `json:"author,omitempty"`
	Created *CustomerRequestDateScheme `json:"created,omitempty"`
	Links   struct {
		Self string `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type RequestAttachmentCreationScheme struct {
	Comment     *RequestAttachmentCreationCommentScheme `json:"comment,omitempty"`
	Attachments *RequestAttachmentPageScheme            `json:"attachments,omitempty"`
}

var (
	notAttachmentsError = fmt.Errorf("error, please provide a valid temporaryAttachmentIDs slice value")
)
