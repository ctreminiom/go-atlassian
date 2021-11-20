package sm

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type RequestAttachmentService struct{ client *Client }

// Gets returns all the attachments for a customer requests.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/attachment#get-attachments-for-request
func (r *RequestAttachmentService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (
	result *model.RequestAttachmentPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
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
	public bool) (result *model.RequestAttachmentCreationScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if len(temporaryAttachmentIDs) == 0 {
		return nil, nil, model.ErrNoAttachmentIDError
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
