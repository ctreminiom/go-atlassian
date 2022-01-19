package confluence

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ContentRestrictionOperationService struct {
	client *Client
	Group  *ContentRestrictionOperationGroupService
}

// Gets returns restrictions on a piece of content by operation.
// This method is similar to Get restrictions except that the operations are properties
// of the return object, rather than items in a results array.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations#get-restrictions-by-operation
func (c *ContentRestrictionOperationService) Gets(ctx context.Context, contentID string, expand []string) (
	result *models.ContentRestrictionByOperationScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, models.ErrNoContentIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation", contentID))

	query := url.Values{}
	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Get returns the restrictions on a piece of content for a given operation (read or update).
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations#get-restrictions-for-operation
func (c *ContentRestrictionOperationService) Get(ctx context.Context, contentID, operationKey string, expand []string,
	startAt, maxResults int) (result *models.ContentRestrictionScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, models.ErrNoContentIDError
	}

	if len(operationKey) == 0 {
		return nil, nil, models.ErrNoContentRestrictionKeyError
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation/%v?%v", contentID, operationKey, query.Encode())

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}
