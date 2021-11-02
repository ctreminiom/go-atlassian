package v3

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
)

type IssueLinkService struct {
	client *Client
	Type   *IssueLinkTypeService
}

type LinkPayloadScheme struct {
	Comment      *CommentPayloadScheme `json:"comment,omitempty"`
	InwardIssue  *LinkedIssueScheme    `json:"inwardIssue,omitempty"`
	OutwardIssue *LinkedIssueScheme    `json:"outwardIssue,omitempty"`
	Type         *LinkTypeScheme       `json:"type,omitempty"`
}

type CommentPayloadScheme struct {
	Visibility *CommentVisibilityScheme `json:"visibility,omitempty"`
	Body       *CommentNodeScheme       `json:"body,omitempty"`
}

type CommentVisibilityScheme struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type LinkedIssueScheme struct {
	ID     string                    `json:"id,omitempty"`
	Key    string                    `json:"key,omitempty"`
	Self   string                    `json:"self,omitempty"`
	Fields *models.IssueFieldsScheme `json:"fields,omitempty"`
}

type LinkTypeScheme struct {
	Self    string `json:"self,omitempty"`
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Inward  string `json:"inward,omitempty"`
	Outward string `json:"outward,omitempty"`
}

// Create creates a link between two issues. Use this operation to indicate a relationship between two issues
// and optionally add a comment to the from (outward) issue.
// To use this resource the site must have Issue Linking enabled.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link#create-issue-link
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-post
func (i *IssueLinkService) Create(ctx context.Context, payload *LinkPayloadScheme) (response *ResponseScheme, err error) {

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

type IssueLinkScheme struct {
	ID           string             `json:"id,omitempty"`
	Type         *LinkTypeScheme    `json:"type,omitempty"`
	InwardIssue  *LinkedIssueScheme `json:"inwardIssue,omitempty"`
	OutwardIssue *LinkedIssueScheme `json:"outwardIssue,omitempty"`
}

// Get returns an issue link.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-link
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-linkid-get
func (i *IssueLinkService) Get(ctx context.Context, linkID string) (result *IssueLinkScheme,
	response *ResponseScheme, err error) {

	if len(linkID) == 0 {
		return nil, nil, notLinkIDError
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

type IssueLinkPageScheme struct {
	Expand string `json:"expand,omitempty"`
	ID     string `json:"id,omitempty"`
	Self   string `json:"self,omitempty"`
	Key    string `json:"key,omitempty"`
	Fields struct {
		IssueLinks []*IssueLinkScheme `json:"issuelinks,omitempty"`
	} `json:"fields,omitempty"`
}

// Gets get the issue links ID's associated with a Jira Issue
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-links
func (i *IssueLinkService) Gets(ctx context.Context, issueKeyOrID string) (result *IssueLinkPageScheme,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueKeyOrIDError
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
		return nil, notLinkIDError
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

var (
	notLinkIDError = fmt.Errorf("error!, please provide a valid linkID value")
)
