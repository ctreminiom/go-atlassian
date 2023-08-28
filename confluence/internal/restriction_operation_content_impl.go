package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/confluence"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewRestrictionOperationService(client service.Connector, group *RestrictionOperationGroupService, user *RestrictionOperationUserService) *RestrictionOperationService {

	return &RestrictionOperationService{
		internalClient: &internalRestrictionOperationImpl{c: client},
		Group:          group,
		User:           user,
	}
}

type RestrictionOperationService struct {
	internalClient confluence.RestrictionOperationConnector
	Group          *RestrictionOperationGroupService
	User           *RestrictionOperationUserService
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
		return nil, nil, model.ErrNoContentIDError
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
		return nil, nil, model.ErrNoContentIDError
	}

	if operationKey == "" {
		return nil, nil, model.ErrNoContentRestrictionKeyError
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
