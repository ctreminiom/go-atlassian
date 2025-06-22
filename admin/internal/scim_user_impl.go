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
	"strings"
)

// NewSCIMUserService creates a new instance of SCIMUserService.
// It takes a service.Connector as input and returns a pointer to SCIMUserService.
func NewSCIMUserService(client service.Connector) *SCIMUserService {
	return &SCIMUserService{internalClient: &internalSCIMUserImpl{c: client}}
}

// SCIMUserService provides methods to interact with SCIM users in Atlassian Administration.
type SCIMUserService struct {
	// internalClient is the connector interface for SCIM user operations.
	internalClient admin.SCIMUserConnector
}

// Create creates a user in a directory.
//
// An attempt to create an existing user fails with a 409 (Conflict) error.
//
// A user account can only be created if it has an email address on a verified domain.
//
// If a managed Atlassian account already exists on the Atlassian platform for the specified email address,
//
// the user in your identity provider is linked to the user in your Atlassian organization.
//
// POST /scim/directory/{directoryId}/Users
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#create-a-user
func (s *SCIMUserService) Create(ctx context.Context, directoryID string, payload *model.SCIMUserScheme, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMUserService).Create")
	defer span.End()

	return s.internalClient.Create(ctx, directoryID, payload, attributes, excludedAttributes)
}

// Gets get users from the specified directory.
//
// Filtering is supported with a single exact match (eq) against the userName and externalId attributes.
//
// Pagination is supported.
//
// Sorting is not supported.
//
// GET /scim/directory/{directoryId}/Users
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#get-users
func (s *SCIMUserService) Gets(ctx context.Context, directoryID string, opts *model.SCIMUserGetsOptionsScheme, startIndex, count int) (*model.SCIMUserPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMUserService).Gets")
	defer span.End()

	return s.internalClient.Gets(ctx, directoryID, opts, startIndex, count)
}

// Get gets a user from a directory by userId.
//
// GET /scim/directory/{directoryId}/Users/{userId}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#get-a-user-by-id
func (s *SCIMUserService) Get(ctx context.Context, directoryID, userID string, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMUserService).Get")
	defer span.End()

	return s.internalClient.Get(ctx, directoryID, userID, attributes, excludedAttributes)
}

// Deactivate deactivates a user by userId.
//
// The user is not available for future requests until activated again.
//
// Any future operation for the deactivated user returns the 404 (resource not found) error.
//
// DELETE /scim/directory/{directoryId}/Users/{userId}
func (s *SCIMUserService) Deactivate(ctx context.Context, directoryID, userID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMUserService).Deactivate")
	defer span.End()

	return s.internalClient.Deactivate(ctx, directoryID, userID)
}

// Path updates a user's information in a directory by userId via PATCH.
//
// Refer to GET /ServiceProviderConfig for details on the supported operations.
//
// PATCH /scim/directory/{directoryId}/Users/{userId}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#update-user-by-id-patch
func (s *SCIMUserService) Path(ctx context.Context, directoryID, userID string, payload *model.SCIMUserToPathScheme, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMUserService).Path")
	defer span.End()

	return s.internalClient.Path(ctx, directoryID, userID, payload, attributes, excludedAttributes)
}

// Update updates a user's information in a directory by userId via user attributes.
//
// User information is replaced attribute-by-attribute, except immutable and read-only attributes.
//
// Existing values of unspecified attributes are cleaned.
//
// PUT /scim/directory/{directoryId}/Users/{userId}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#update-user-via-user-attributes
func (s *SCIMUserService) Update(ctx context.Context, directoryID, userID string, payload *model.SCIMUserScheme, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*SCIMUserService).Update")
	defer span.End()

	return s.internalClient.Update(ctx, directoryID, userID, payload, attributes, excludedAttributes)
}

type internalSCIMUserImpl struct {
	c service.Connector
}

func (i *internalSCIMUserImpl) Create(ctx context.Context, directoryID string, payload *model.SCIMUserScheme, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMUserImpl).Create")
	defer span.End()

	if directoryID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	params := url.Values{}

	if len(attributes) != 0 {
		params.Add("attributes", strings.Join(attributes, ","))
	}

	if len(excludedAttributes) != 0 {
		params.Add("excludedAttributes", strings.Join(excludedAttributes, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("scim/directory/%v/Users", directoryID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	user := new(model.SCIMUserScheme)
	response, err := i.c.Call(request, user)
	if err != nil {
		return nil, response, err
	}

	return user, response, nil
}

func (i *internalSCIMUserImpl) Gets(ctx context.Context, directoryID string, opts *model.SCIMUserGetsOptionsScheme, startIndex, count int) (*model.SCIMUserPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMUserImpl).Gets")
	defer span.End()

	if directoryID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	params := url.Values{}
	params.Add("startIndex", strconv.Itoa(startIndex))
	params.Add("count", strconv.Itoa(count))

	if opts != nil {

		if len(opts.Attributes) != 0 {
			params.Add("attributes", strings.Join(opts.Attributes, ","))
		}

		if len(opts.ExcludedAttributes) != 0 {
			params.Add("excludedAttributes", strings.Join(opts.ExcludedAttributes, ","))
		}

		if len(opts.Filter) != 0 {
			params.Add("filter", opts.Filter)
		}
	}

	endpoint := fmt.Sprintf("scim/directory/%v/Users?%v", directoryID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	users := new(model.SCIMUserPageScheme)
	response, err := i.c.Call(request, users)
	if err != nil {
		return nil, response, err
	}

	return users, response, nil
}

func (i *internalSCIMUserImpl) Get(ctx context.Context, directoryID, userID string, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMUserImpl).Get")
	defer span.End()

	if directoryID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	if userID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminUserID)
	}

	params := url.Values{}
	if len(attributes) != 0 {
		params.Add("attributes", strings.Join(attributes, ","))
	}

	if len(excludedAttributes) != 0 {
		params.Add("excludedAttributes", strings.Join(excludedAttributes, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("scim/directory/%v/Users/%v", directoryID, userID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(model.SCIMUserScheme)
	response, err := i.c.Call(request, user)
	if err != nil {
		return nil, response, err
	}

	return user, response, nil
}

func (i *internalSCIMUserImpl) Deactivate(ctx context.Context, directoryID, userID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMUserImpl).Deactivate")
	defer span.End()

	if directoryID == "" {
		return nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	if userID == "" {
		return nil, fmt.Errorf("admin: %w", model.ErrNoAdminUserID)
	}

	endpoint := fmt.Sprintf("scim/directory/%v/Users/%v", directoryID, userID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalSCIMUserImpl) Path(ctx context.Context, directoryID, userID string, payload *model.SCIMUserToPathScheme, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMUserImpl).Path")
	defer span.End()

	if directoryID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	if userID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminUserID)
	}

	params := url.Values{}

	if len(attributes) != 0 {
		params.Add("attributes", strings.Join(attributes, ","))
	}

	if len(excludedAttributes) != 0 {
		params.Add("excludedAttributes", strings.Join(excludedAttributes, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("scim/directory/%v/Users/%v", directoryID, userID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPatch, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	user := new(model.SCIMUserScheme)
	response, err := i.c.Call(request, user)
	if err != nil {
		return nil, response, err
	}

	return user, response, nil
}

func (i *internalSCIMUserImpl) Update(ctx context.Context, directoryID, userID string, payload *model.SCIMUserScheme, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalSCIMUserImpl).Update")
	defer span.End()

	if directoryID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminDirectoryID)
	}

	if userID == "" {
		return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminUserID)
	}

	params := url.Values{}
	if len(attributes) != 0 {
		params.Add("attributes", strings.Join(attributes, ","))
	}

	if len(excludedAttributes) != 0 {
		params.Add("excludedAttributes", strings.Join(excludedAttributes, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("scim/directory/%v/Users/%v", directoryID, userID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	user := new(model.SCIMUserScheme)
	response, err := i.c.Call(request, user)
	if err != nil {
		return nil, response, err
	}

	return user, response, nil
}
