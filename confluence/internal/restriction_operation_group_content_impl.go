package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/confluence"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func NewRestrictionOperationGroupService(client service.Connector) *RestrictionOperationGroupService {

	return &RestrictionOperationGroupService{
		internalClient: &internalRestrictionOperationGroupImpl{c: client},
	}
}

type RestrictionOperationGroupService struct {
	internalClient confluence.RestrictionGroupOperationConnector
}

// Get returns whether the specified content restriction applies to a group.
//
// Note that a response of true does not guarantee that the group can view the page,
//
// as it does not account for account-inherited restrictions, space permissions, or even product access.
//
// GET /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/byGroupId/{groupId}
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/group#get-content-restriction-status-for-group
func (r *RestrictionOperationGroupService) Get(ctx context.Context, contentID, operationKey, groupNameOrID string) (*model.ResponseScheme, error) {
	return r.internalClient.Get(ctx, contentID, operationKey, groupNameOrID)
}

// Add adds a group to a content restriction.
//
// That is, grant read or update permission to the group for a piece of content.
//
// PUT /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/byGroupId/{groupId}
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/group#add-group-to-content-restriction
func (r *RestrictionOperationGroupService) Add(ctx context.Context, contentID, operationKey, groupNameOrID string) (*model.ResponseScheme, error) {
	return r.internalClient.Add(ctx, contentID, operationKey, groupNameOrID)
}

// Remove removes a group from a content restriction.
//
// That is, remove read or update permission for the group for a piece of content.
//
// DELETE /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/byGroupId/{groupId}
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/group#remove-group-from-content-restriction
func (r *RestrictionOperationGroupService) Remove(ctx context.Context, contentID, operationKey, groupNameOrID string) (*model.ResponseScheme, error) {
	return r.internalClient.Remove(ctx, contentID, operationKey, groupNameOrID)
}

type internalRestrictionOperationGroupImpl struct {
	c service.Connector
}

func (i *internalRestrictionOperationGroupImpl) Get(ctx context.Context, contentID, operationKey, groupNameOrID string) (*model.ResponseScheme, error) {

	if contentID == "" {
		return nil, model.ErrNoContentIDError
	}

	if operationKey == "" {
		return nil, model.ErrNoContentRestrictionKeyError
	}

	if groupNameOrID == "" {
		return nil, model.ErrNoConfluenceGroupError
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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalRestrictionOperationGroupImpl) Add(ctx context.Context, contentID, operationKey, groupNameOrID string) (*model.ResponseScheme, error) {

	if contentID == "" {
		return nil, model.ErrNoContentIDError
	}

	if operationKey == "" {
		return nil, model.ErrNoContentRestrictionKeyError
	}

	if groupNameOrID == "" {
		return nil, model.ErrNoConfluenceGroupError
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

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint.String(), "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalRestrictionOperationGroupImpl) Remove(ctx context.Context, contentID, operationKey, groupNameOrID string) (*model.ResponseScheme, error) {

	if contentID == "" {
		return nil, model.ErrNoContentIDError
	}

	if operationKey == "" {
		return nil, model.ErrNoContentRestrictionKeyError
	}

	if groupNameOrID == "" {
		return nil, model.ErrNoConfluenceGroupError
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

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint.String(), "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
