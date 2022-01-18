package confluence

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type ContentRestrictionOperationGroupService struct{ client *Client }

// Get returns whether the specified content restriction applies to a group
// Note that a response of true does not guarantee that the group can view the page,
// as it does not account for account-inherited restrictions, space permissions, or even product access.
func (c *ContentRestrictionOperationGroupService) Get(ctx context.Context, contentID, operationKey, groupNameOrID string) (response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, models.ErrNoContentIDError
	}

	if len(operationKey) == 0 {
		return nil, models.ErrNoContentRestrictionKeyError
	}

	if len(groupNameOrID) == 0 {
		return nil, models.ErrNoConfluenceGroupError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation/%v/", contentID, operationKey))

	// check if the group id is an uuid type
	// if so, it's the group id
	groupID, err := uuid.Parse(groupNameOrID)

	if err == nil {
		endpoint.WriteString(fmt.Sprintf("byGroupId/%v", groupID.String()))
	} else {
		endpoint.WriteString(fmt.Sprintf("group/%v", groupNameOrID))
	}

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err = c.client.Call(request, nil)
	if err != nil {
		return response, err
	}

	return
}

// Add adds a group to a content restriction. That is, grant read or update permission to the group for a piece of content.
func (c *ContentRestrictionOperationGroupService) Add(ctx context.Context, contentID, operationKey, groupNameOrID string) (response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, models.ErrNoContentIDError
	}

	if len(operationKey) == 0 {
		return nil, models.ErrNoContentRestrictionKeyError
	}

	if len(groupNameOrID) == 0 {
		return nil, models.ErrNoConfluenceGroupError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation/%v/", contentID, operationKey))

	// check if the group id is an uuid type
	// if so, it's the group id
	groupID, err := uuid.Parse(groupNameOrID)

	if err == nil {
		endpoint.WriteString(fmt.Sprintf("byGroupId/%v", groupID.String()))
	} else {
		endpoint.WriteString(fmt.Sprintf("group/%v", groupNameOrID))
	}

	request, err := c.client.newRequest(ctx, http.MethodPut, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err = c.client.Call(request, nil)
	if err != nil {
		return response, err
	}

	return
}

// Remove removes a group from a content restriction. That is, remove read or update permission for the group for a piece of content.
func (c *ContentRestrictionOperationGroupService) Remove(ctx context.Context, contentID, operationKey, groupNameOrID string) (response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, models.ErrNoContentIDError
	}

	if len(operationKey) == 0 {
		return nil, models.ErrNoContentRestrictionKeyError
	}

	if len(groupNameOrID) == 0 {
		return nil, models.ErrNoConfluenceGroupError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation/%v/", contentID, operationKey))

	// check if the group id is an uuid type
	// if so, it's the group id
	groupID, err := uuid.Parse(groupNameOrID)

	if err == nil {
		endpoint.WriteString(fmt.Sprintf("byGroupId/%v", groupID.String()))
	} else {
		endpoint.WriteString(fmt.Sprintf("group/%v", groupNameOrID))
	}

	request, err := c.client.newRequest(ctx, http.MethodDelete, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err = c.client.Call(request, nil)
	if err != nil {
		return response, err
	}

	return
}
