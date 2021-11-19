package v3

import (
	"context"
	"fmt"
	models2 "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type IssueLinkService struct {
	client *Client
	Type   *IssueLinkTypeService
}

// Create creates a link between two issues. Use this operation to indicate a relationship between two issues
// and optionally add a comment to the from (outward) issue.
// To use this resource the site must have Issue Linking enabled.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link#create-issue-link
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-post
func (i *IssueLinkService) Create(ctx context.Context, payload *models2.LinkPayloadScheme) (response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/issueLink"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Get returns an issue link.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-link
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-linkid-get
func (i *IssueLinkService) Get(ctx context.Context, linkID string) (result *models2.IssueLinkScheme,
	response *ResponseScheme, err error) {

	if len(linkID) == 0 {
		return nil, nil, models2.ErrNoTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issueLink/%v", linkID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Gets get the issue links ID's associated with a Jira Issue
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-links
func (i *IssueLinkService) Gets(ctx context.Context, issueKeyOrID string) (result *models2.IssueLinkPageScheme,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models2.ErrNoIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v?fields=issuelinks", issueKeyOrID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes an issue link.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link#delete-issue-link
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-linkid-delete
func (i *IssueLinkService) Delete(ctx context.Context, linkID string) (response *ResponseScheme, err error) {

	if len(linkID) == 0 {
		return nil, models2.ErrNoTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issueLink/%v", linkID)

	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
