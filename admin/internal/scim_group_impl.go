package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/admin"
	"net/http"
	"net/url"
	"strconv"
)

// NewSCIMGroupService creates a new instance of SCIMGroupService.
// It takes a service.Connector as input and returns a pointer to SCIMGroupService.
func NewSCIMGroupService(client service.Connector) *SCIMGroupService {
	return &SCIMGroupService{internalClient: &internalSCIMGroupImpl{c: client}}
}

// SCIMGroupService provides methods to interact with SCIM groups in Atlassian Administration.
type SCIMGroupService struct {
	// internalClient is the connector interface for SCIM group operations.
	internalClient admin.SCIMGroupConnector
}

// Gets get groups from a directory.
//
// Filtering is supported with a single exact match (eq) against the displayName attribute.
//
// Pagination is supported.
//
// Sorting is not supported.
//
// GET /scim/directory/{directoryId}/Groups
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#get-groups
func (s *SCIMGroupService) Gets(ctx context.Context, directoryID, filter string, startAt, maxResults int) (*model.ScimGroupPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMGroupService).Gets")
	defer span.End()

	return s.internalClient.Gets(ctx, directoryID, filter, startAt, maxResults)
}

// Get returns a group from a directory by group ID.
//
// GET /scim/directory/{directoryId}/Groups/{id}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#get-a-group-by-id
func (s *SCIMGroupService) Get(ctx context.Context, directoryID, groupID string) (*model.ScimGroupScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMGroupService).Get")
	defer span.End()

	return s.internalClient.Get(ctx, directoryID, groupID)
}

// Update updates a group in a directory by group ID.
//
// PUT /scim/directory/{directoryId}/Groups/{id}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#update-a-group-by-id
func (s *SCIMGroupService) Update(ctx context.Context, directoryID, groupID string, newGroupName string) (*model.ScimGroupScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMGroupService).Update")
	defer span.End()

	return s.internalClient.Update(ctx, directoryID, groupID, newGroupName)
}

// Delete deletes a group from a directory.
//
// An attempt to delete a non-existent group fails with a 404 (Resource Not found) error.
//
// DELETE /scim/directory/{directoryId}/Groups/{id}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#delete-a-group-by-id
func (s *SCIMGroupService) Delete(ctx context.Context, directoryID, groupID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMGroupService).Delete")
	defer span.End()

	return s.internalClient.Delete(ctx, directoryID, groupID)
}

// Create creates a group in a directory.
//
// An attempt to create a group with an existing name fails with a 409 (Conflict) error.
//
// POST /scim/directory/{directoryId}/Groups
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#create-a-group
func (s *SCIMGroupService) Create(ctx context.Context, directoryID, groupName string) (*model.ScimGroupScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMGroupService).Create")
	defer span.End()

	return s.internalClient.Create(ctx, directoryID, groupName)
}

// Path update a group's information in a directory by groupId via PATCH.
//
// You can use this API to manage group membership.
//
// PATCH /scim/directory/{directoryId}/Groups/{id}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#update-a-group-by-id-patch
func (s *SCIMGroupService) Path(ctx context.Context, directoryID, groupID string, payload *model.SCIMGroupPathScheme) (*model.ScimGroupScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMGroupService).Path")
	defer span.End()

	return s.internalClient.Path(ctx, directoryID, groupID, payload)
}

type internalSCIMGroupImpl struct {
	c service.Connector
}

func (i *internalSCIMGroupImpl) Gets(ctx context.Context, directoryID, filter string, startAt, maxResults int) (*model.ScimGroupPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMGroupImpl).Gets")
	defer span.End()

	if directoryID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	params := url.Values{}
	params.Add("startIndex", strconv.Itoa(startAt))
	params.Add("count", strconv.Itoa(maxResults))

	if filter != "" {
		params.Add("filter", filter)
	}

	endpoint := fmt.Sprintf("scim/directory/%v/Groups?%v", directoryID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	groups := new(model.ScimGroupPageScheme)
	response, err := i.c.Call(request, groups)
	if err != nil {
		return nil, response, err
	}

	return groups, response, nil
}

func (i *internalSCIMGroupImpl) Get(ctx context.Context, directoryID, groupID string) (*model.ScimGroupScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMGroupImpl).Get")
	defer span.End()

	if directoryID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	if groupID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminGroupID)
	}

	endpoint := fmt.Sprintf("scim/directory/%v/Groups/%v", directoryID, groupID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(model.ScimGroupScheme)
	response, err := i.c.Call(request, group)
	if err != nil {
		return nil, response, err
	}

	return group, response, nil
}

func (i *internalSCIMGroupImpl) Update(ctx context.Context, directoryID, groupID string, newGroupName string) (*model.ScimGroupScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMGroupImpl).Update")
	defer span.End()

	if directoryID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	if groupID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminGroupID)
	}

	if newGroupName == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminGroupName)
	}

	endpoint := fmt.Sprintf("scim/directory/%v/Groups/%v", directoryID, groupID)

	payload := map[string]interface{}{"displayName": newGroupName}

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	group := new(model.ScimGroupScheme)
	response, err := i.c.Call(request, group)
	if err != nil {
		return nil, response, err
	}

	return group, response, nil
}

func (i *internalSCIMGroupImpl) Delete(ctx context.Context, directoryID, groupID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMGroupImpl).Delete")
	defer span.End()

	if directoryID == "" {
		return nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	if groupID == "" {
		return nil, fmt.Errorf("admin: %w", model.ErrNoAdminGroupID)
	}

	endpoint := fmt.Sprintf("scim/directory/%v/Groups/%v", directoryID, groupID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalSCIMGroupImpl) Create(ctx context.Context, directoryID, groupName string) (*model.ScimGroupScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMGroupImpl).Create")
	defer span.End()

	if directoryID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	if groupName == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminGroupName)
	}

	payload := map[string]interface{}{"displayName": groupName}

	endpoint := fmt.Sprintf("scim/directory/%v/Groups", directoryID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	group := new(model.ScimGroupScheme)
	response, err := i.c.Call(request, group)
	if err != nil {
		return nil, response, err
	}

	return group, response, nil
}

func (i *internalSCIMGroupImpl) Path(ctx context.Context, directoryID, groupID string, payload *model.SCIMGroupPathScheme) (*model.ScimGroupScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMGroupImpl).Path")
	defer span.End()

	if directoryID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	if groupID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminGroupID)
	}

	endpoint := fmt.Sprintf("scim/directory/%v/Groups/%v", directoryID, groupID)

	request, err := i.c.NewRequest(ctx, http.MethodPatch, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	group := new(model.ScimGroupScheme)
	response, err := i.c.Call(request, group)
	if err != nil {
		return nil, response, err
	}

	return group, response, nil
}
