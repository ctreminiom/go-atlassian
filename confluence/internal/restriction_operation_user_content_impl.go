package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"net/http"
	"net/url"
)

// NewRestrictionOperationUserService creates a new instance of RestrictionOperationUserService.
// It takes a service.Connector as input and returns a pointer to RestrictionOperationUserService.
func NewRestrictionOperationUserService(client service.Connector) *RestrictionOperationUserService {
	return &RestrictionOperationUserService{
		internalClient: &internalRestrictionOperationUserImpl{c: client},
	}
}

// RestrictionOperationUserService provides methods to interact with content restriction operations for users in Confluence.
type RestrictionOperationUserService struct {
	// internalClient is the connector interface for content restriction operations for users.
	internalClient confluence.RestrictionUserOperationConnector
}

// Get returns whether the specified content restriction applies to a user.
//
// GET /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/user
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/user#get-content-restriction-status-for-user
func (r *RestrictionOperationUserService) Get(ctx context.Context, contentID, operationKey, accountID string) (*model.ResponseScheme, error) {
	return r.internalClient.Get(ctx, contentID, operationKey, accountID)
}

// Add adds a user to a content restriction.
//
// That is, grant read or update permission to the user for a piece of content.
//
// PUT /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/user
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/user#add-user-to-content-restriction
func (r *RestrictionOperationUserService) Add(ctx context.Context, contentID, operationKey, accountID string) (*model.ResponseScheme, error) {
	return r.internalClient.Add(ctx, contentID, operationKey, accountID)
}

// Remove removes a group from a content restriction.
//
// That is, remove read or update permission for the group for a piece of content.
//
// DELETE /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/user
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/user#remove-user-from-content-restriction
func (r *RestrictionOperationUserService) Remove(ctx context.Context, contentID, operationKey, accountID string) (*model.ResponseScheme, error) {
	return r.internalClient.Remove(ctx, contentID, operationKey, accountID)
}

type internalRestrictionOperationUserImpl struct {
	c service.Connector
}

func (i *internalRestrictionOperationUserImpl) Get(ctx context.Context, contentID, operationKey, accountID string) (*model.ResponseScheme, error) {

	if contentID == "" {
		return nil, model.ErrNoContentID
	}

	if operationKey == "" {
		return nil, model.ErrNoContentRestrictionKey
	}

	if accountID == "" {
		return nil, model.ErrNoAccountID
	}

	query := url.Values{}
	query.Add("accountId", accountID)

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation/%v/user?%v", contentID, operationKey, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalRestrictionOperationUserImpl) Add(ctx context.Context, contentID, operationKey, accountID string) (*model.ResponseScheme, error) {

	if contentID == "" {
		return nil, model.ErrNoContentID
	}

	if operationKey == "" {
		return nil, model.ErrNoContentRestrictionKey
	}

	if accountID == "" {
		return nil, model.ErrNoAccountID
	}

	query := url.Values{}
	query.Add("accountId", accountID)

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation/%v/user?%v", contentID, operationKey, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalRestrictionOperationUserImpl) Remove(ctx context.Context, contentID, operationKey, accountID string) (*model.ResponseScheme, error) {

	if contentID == "" {
		return nil, model.ErrNoContentID
	}

	if operationKey == "" {
		return nil, model.ErrNoContentRestrictionKey
	}

	if accountID == "" {
		return nil, model.ErrNoAccountID
	}

	query := url.Values{}
	query.Add("accountId", accountID)

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/restriction/byOperation/%v/user?%v", contentID, operationKey, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
