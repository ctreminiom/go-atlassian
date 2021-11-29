package admin

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type SCIMGroupService struct{ client *Client }

// Gets gets groups from a directory.
// Filtering is supported with a single exact match (eq) against the displayName attribute.
// Pagination is supported. Sorting is not supported.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#get-groups
func (g *SCIMGroupService) Gets(ctx context.Context, directoryID, filter string, startAt, maxResults int) (
	result *model.ScimGroupPageScheme, response *ResponseScheme, err error) {

	if directoryID == "" {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	params := url.Values{}
	params.Add("startIndex", strconv.Itoa(startAt))
	params.Add("count", strconv.Itoa(maxResults))

	if filter != "" {
		params.Add("filter", filter)
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups?%v", directoryID, params.Encode())

	request, err := g.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = g.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get a group from a directory by group ID.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#get-a-group-by-id
func (g *SCIMGroupService) Get(ctx context.Context, directoryID, groupID string) (result *model.ScimGroupScheme,
	response *ResponseScheme, err error) {

	if directoryID == "" {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	if groupID == "" {
		return nil, nil, model.ErrNoAdminGroupIDError
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups/%v", directoryID, groupID)

	request, err := g.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = g.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update a group in a directory by group ID.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#update-a-group-by-id
func (g *SCIMGroupService) Update(ctx context.Context, directoryID, groupID string, newGroupName string) (result *model.ScimGroupScheme,
	response *ResponseScheme, err error) {

	if directoryID == "" {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	if groupID == "" {
		return nil, nil, model.ErrNoAdminGroupIDError
	}

	if newGroupName == "" {
		return nil, nil, model.ErrNoAdminGroupNameError
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups/%v", directoryID, groupID)

	payload := struct {
		DisplayName string `json:"displayName"`
	}{
		DisplayName: newGroupName,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	request, err := g.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = g.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete a group from a directory.
// An attempt to delete a non-existent group fails with a 404 (Resource Not found) error.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#delete-a-group-by-id
func (g *SCIMGroupService) Delete(ctx context.Context, directoryID, groupID string) (response *ResponseScheme, err error) {

	if directoryID == "" {
		return nil, model.ErrNoAdminDirectoryIDError
	}

	if groupID == "" {
		return nil, model.ErrNoAdminGroupIDError
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups/%v", directoryID, groupID)

	request, err := g.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = g.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Create a group in a directory. An attempt to create a group with an existing name fails with a 409 (Conflict) error.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#create-a-group
func (g *SCIMGroupService) Create(ctx context.Context, directoryID, groupName string) (result *model.ScimGroupScheme,
	response *ResponseScheme, err error) {

	if directoryID == "" {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	if groupName == "" {
		return nil, nil, model.ErrNoAdminGroupNameError
	}

	payload := struct {
		DisplayName string `json:"displayName"`
	}{
		DisplayName: groupName,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups", directoryID)

	request, err := g.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = g.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Path update a group's information in a directory by groupId via PATCH.
// You can use this API to manage group membership.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#update-a-group-by-id-patch
func (g *SCIMGroupService) Path(ctx context.Context, directoryID, groupID string, payload *model.SCIMGroupPathScheme) (
	result *model.ScimGroupScheme, response *ResponseScheme, err error) {

	if directoryID == "" {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	if groupID == "" {
		return nil, nil, model.ErrNoAdminGroupIDError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	if len(payload.Operations) == 0 {
		return nil, nil, fmt.Errorf("erro!, the SCIMGroupPathScheme value must contains operations")
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Groups/%v", directoryID, groupID)

	request, err := g.client.newRequest(ctx, http.MethodPatch, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = g.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
