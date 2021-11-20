package sm

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type RequestCommentService struct{ client *Client }

// Gets returns all comments on a customer request.
// No permissions error is provided if, for example, the user doesn't have access to the service desk or request,
// the method simply returns an empty response.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-request-comments
func (r *RequestCommentService) Gets(ctx context.Context, issueKeyOrID string, public bool, expand []string, start,
	limit int) (result *model.RequestCommentPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if !public {
		params.Add("public", "false")
	}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/comment?%v", issueKeyOrID, params.Encode())

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

// Get returns details of a customer request's comment.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-request-comment-by-id
func (r *RequestCommentService) Get(ctx context.Context, issueKeyOrID string, commentID int, expand []string) (
	result *model.RequestCommentScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/servicedeskapi/request/%v/comment/%v", issueKeyOrID, commentID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
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

// Create creates a public or private (internal) comment on a customer request, with the comment visibility set by public.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#create-request-comment
func (r *RequestCommentService) Create(ctx context.Context, issueKeyOrID, body string, public bool) (
	result *model.RequestCommentScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if len(body) == 0 {
		return nil, nil, model.ErrNoCommentBodyError
	}

	payload := struct {
		Public bool   `json:"public"`
		Body   string `json:"body"`
	}{
		Public: public,
		Body:   body,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/comment", issueKeyOrID)

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

// Attachments  returns the attachments referenced in a comment.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-comment-attachments
func (r *RequestCommentService) Attachments(ctx context.Context, issueKeyOrID string, commentID, start, limit int) (
	result *model.RequestAttachmentPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
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

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}
