package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewRestrictionOperationService creates a new instance of RestrictionOperationService.
// It takes a service.Connector, a pointer to RestrictionOperationGroupService, and a pointer to RestrictionOperationUserService as input,
// and returns a pointer to RestrictionOperationService.
func NewRestrictionOperationService(client service.Connector, group *RestrictionOperationGroupService, user *RestrictionOperationUserService) *RestrictionOperationService {
	return &RestrictionOperationService{
		internalClient: &internalRestrictionOperationImpl{c: client},
		Group:          group,
		User:           user,
	}
}

// RestrictionOperationService provides methods to interact with content restriction operations in Confluence.
type RestrictionOperationService struct {
	// internalClient is the connector interface for content restriction operations.
	internalClient confluence.RestrictionOperationConnector
	// Group is a pointer to RestrictionOperationGroupService for group-related restriction operations.
	Group *RestrictionOperationGroupService
	// User is a pointer to RestrictionOperationUserService for user-related restriction operations.
	User *RestrictionOperationUserService
}

// Gets returns restrictions on a piece of content by operation.
//
// # This method is similar to Get restrictions except that the operations are properties
//
// of the return object, rather than items in a results array.
//
// GET /wiki/rest/api/content/{id}/restriction
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations#get-restrictions-by-operation
func (r *RestrictionOperationService) Gets(ctx context.Context, contentID string, expand []string) (*model.ContentRestrictionByOperationScheme, *model.ResponseScheme, error) {
	return r.internalClient.Gets(ctx, contentID, expand)
}

// Get returns the restrictions on a piece of content for a given operation (read or update).
//
// GET /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations#get-restrictions-for-operation
func (r *RestrictionOperationService) Get(ctx context.Context, contentID, operationKey string, expand []string, startAt, maxResults int) (*model.ContentRestrictionScheme, *model.ResponseScheme, error) {
	return r.internalClient.Get(ctx, contentID, operationKey, expand, startAt, maxResults)
}

type internalRestrictionOperationImpl struct {
	c service.Connector
}

func (i *internalRestrictionOperationImpl) Gets(ctx context.Context, contentID string, expand []string) (*model.ContentRestrictionByOperationScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation", contentID))

	if len(expand) != 0 {
		query := url.Values{}
		query.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	operation := new(model.ContentRestrictionByOperationScheme)
	response, err := i.c.Call(request, operation)
	if err != nil {
		return nil, response, err
	}

	return operation, response, nil
}

func (i *internalRestrictionOperationImpl) Get(ctx context.Context, contentID, operationKey string, expand []string, startAt, maxResults int) (*model.ContentRestrictionScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	if operationKey == "" {
		return nil, nil, model.ErrNoContentRestrictionKey
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation/%v?%v", contentID, operationKey, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	restriction := new(model.ContentRestrictionScheme)
	response, err := i.c.Call(request, restriction)
	if err != nil {
		return nil, response, err
	}

	return restriction, response, nil
}
