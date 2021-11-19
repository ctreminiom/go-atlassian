package v2

import (
	"context"
	"fmt"
	models2 "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type IssueLinkTypeService struct{ client *Client }

// Gets returns a list of all issue link types.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#get-issue-link-types
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-link-types/#api-rest-api-2-issuelinktype-get
func (i *IssueLinkTypeService) Gets(ctx context.Context) (result *models2.IssueLinkTypeSearchScheme, response *ResponseScheme,
	err error) {

	var endpoint = "rest/api/2/issueLinkType"

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Get returns an issue link type.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#get-issue-link-type
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-link-types/#api-rest-api-2-issuelinktype-issuelinktypeid-get
func (i *IssueLinkTypeService) Get(ctx context.Context, issueLinkTypeID string) (result *models2.LinkTypeScheme,
	response *ResponseScheme, err error) {

	if len(issueLinkTypeID) == 0 {
		return nil, nil, models2.ErrNoLinkTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/issueLinkType/%v", issueLinkTypeID)

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

// Create creates an issue link type.
// Use this operation to create descriptions of the reasons why issues are linked.
// The issue link type consists of a name and descriptions for a link's inward and outward relationships.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#create-issue-link-type
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-link-types/#api-rest-api-2-issuelinktype-post
func (i *IssueLinkTypeService) Create(ctx context.Context, payload *models2.LinkTypeScheme) (result *models2.LinkTypeScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/2/issueLinkType"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates an issue link type.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#update-issue-link-type
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-link-types/#api-rest-api-2-issuelinktype-issuelinktypeid-put
func (i *IssueLinkTypeService) Update(ctx context.Context, issueLinkTypeID string, payload *models2.LinkTypeScheme) (
	result *models2.LinkTypeScheme, response *ResponseScheme, err error) {

	if len(issueLinkTypeID) == 0 {
		return nil, nil, models2.ErrNoLinkTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/issueLinkType/%v", issueLinkTypeID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes an issue link type.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#delete-issue-link-type
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-link-types/#api-rest-api-2-issuelinktype-issuelinktypeid-delete
func (i *IssueLinkTypeService) Delete(ctx context.Context, issueLinkTypeID string) (response *ResponseScheme, err error) {

	if len(issueLinkTypeID) == 0 {
		return nil, models2.ErrNoLinkTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/issueLinkType/%v", issueLinkTypeID)

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
