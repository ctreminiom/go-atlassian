package confluence

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strings"
)

type ContentRestrictionOperationUserService struct{ client *Client }

// Get returns whether the specified content restriction applies to a user.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/user#get-content-restriction-status-for-user
func (c *ContentRestrictionOperationUserService) Get(ctx context.Context, contentID, operationKey, accountID string) (
	response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, models.ErrNoContentIDError
	}

	if len(operationKey) == 0 {
		return nil, models.ErrNoContentRestrictionKeyError
	}

	if len(accountID) == 0 {
		return nil, models.ErrNoAccountIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation/%v/user", contentID, operationKey))

	query := url.Values{}
	query.Add("accountId", accountID)

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
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

// Add adds a user to a content restriction. That is, grant read or update permission to the user for a piece of content.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/user#add-user-to-content-restriction
func (c *ContentRestrictionOperationUserService) Add(ctx context.Context, contentID, operationKey, accountID string) (
	response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, models.ErrNoContentIDError
	}

	if len(operationKey) == 0 {
		return nil, models.ErrNoContentRestrictionKeyError
	}

	if len(accountID) == 0 {
		return nil, models.ErrNoAccountIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation/%v/user", contentID, operationKey))

	query := url.Values{}
	query.Add("accountId", accountID)

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
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
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/user#remove-user-from-content-restriction
func (c *ContentRestrictionOperationUserService) Remove(ctx context.Context, contentID, operationKey, accountID string) (
	response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, models.ErrNoContentIDError
	}

	if len(operationKey) == 0 {
		return nil, models.ErrNoContentRestrictionKeyError
	}

	if len(accountID) == 0 {
		return nil, models.ErrNoAccountIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation/%v/user", contentID, operationKey))

	query := url.Values{}
	query.Add("accountId", accountID)

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
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
