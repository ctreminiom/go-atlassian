package admin

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SCIMUserService struct{ client *Client }

// Create a user in a directory.
// An attempt to create an existing user fails with a 409 (Conflict) error.
// A user account can only be created if it has an email address on a verified domain.
// If a managed Atlassian account already exists on the Atlassian platform for the specified email address,
// the user in your identity provider is linked to the user in your Atlassian organization.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#create-a-user
func (s *SCIMUserService) Create(ctx context.Context, directoryID string, payload *model.SCIMUserScheme, attributes,
	excludedAttributes []string) (result *model.SCIMUserScheme, response *ResponseScheme, err error) {

	if len(directoryID) == 0 {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	params := url.Values{}

	if len(attributes) != 0 {
		params.Add("attributes", strings.Join(attributes, ","))
	}

	if len(excludedAttributes) != 0 {
		params.Add("excludedAttributes", strings.Join(excludedAttributes, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/scim/directory/%v/Users", directoryID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint.String(), payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Gets get users from the specified directory
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#get-users
func (s *SCIMUserService) Gets(ctx context.Context, directoryID string, opts *model.SCIMUserGetsOptionsScheme, startIndex,
	count int) (result *model.SCIMUserPageScheme, response *ResponseScheme, err error) {

	if len(directoryID) == 0 {
		return nil, nil, model.ErrNoAdminDirectoryIDError
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

	var endpoint = fmt.Sprintf("/scim/directory/%v/Users?%v", directoryID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get a user from a directory by userId.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#get-a-user-by-id
func (s *SCIMUserService) Get(ctx context.Context, directoryID, userID string, attributes, excludedAttributes []string) (
	result *model.SCIMUserScheme, response *ResponseScheme, err error) {

	if len(directoryID) == 0 {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	if len(userID) == 0 {
		return nil, nil, model.ErrNoAdminUserIDError
	}

	params := url.Values{}
	if len(attributes) != 0 {
		params.Add("attributes", strings.Join(attributes, ","))
	}

	if len(excludedAttributes) != 0 {
		params.Add("excludedAttributes", strings.Join(excludedAttributes, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/scim/directory/%v/Users/%v", directoryID, userID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Deactivate a user by userId.
// The user is not available for future requests until activated again.
// Any future operation for the deactivated user returns the 404 (resource not found) error.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#deactivate-a-user
func (s *SCIMUserService) Deactivate(ctx context.Context, directoryID, userID string) (response *ResponseScheme, err error) {

	if len(directoryID) == 0 {
		return nil, model.ErrNoAdminDirectoryIDError
	}

	if len(userID) == 0 {
		return nil, model.ErrNoAdminUserIDError
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Users/%v", directoryID, userID)

	request, err := s.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = s.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Path updates a user's information in a directory by userId via PATCH.
// Refer to GET /ServiceProviderConfig for details on the supported operations.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#update-user-by-id-patch
func (s *SCIMUserService) Path(ctx context.Context, directoryID, userID string, payload *model.SCIMUserToPathScheme, attributes,
	excludedAttributes []string) (result *model.SCIMUserScheme, response *ResponseScheme, err error) {

	if len(directoryID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	if len(userID) == 0 {
		return nil, nil, model.ErrNoAdminUserIDError
	}

	params := url.Values{}

	if len(attributes) != 0 {
		params.Add("attributes", strings.Join(attributes, ","))
	}

	if len(excludedAttributes) != 0 {
		params.Add("excludedAttributes", strings.Join(excludedAttributes, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/scim/directory/%v/Users/%v", directoryID, userID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := s.client.newRequest(ctx, http.MethodPatch, endpoint.String(), payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates a user's information in a directory by userId via user attributes.
// User information is replaced attribute-by-attribute, with the exception of immutable and read-only attributes.
// Existing values of unspecified attributes are cleaned.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#update-user-via-user-attributes
func (s *SCIMUserService) Update(ctx context.Context, directoryID, userID string, payload *model.SCIMUserScheme, attributes,
	excludedAttributes []string) (result *model.SCIMUserScheme, response *ResponseScheme, err error) {

	if len(directoryID) == 0 {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	if len(userID) == 0 {
		return nil, nil, model.ErrNoAdminUserIDError
	}

	params := url.Values{}
	if len(attributes) != 0 {
		params.Add("", strings.Join(attributes, ","))
	}

	if len(excludedAttributes) != 0 {
		params.Add("excludedAttributes", strings.Join(excludedAttributes, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/scim/directory/%v/Users/%v", directoryID, userID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := s.client.newRequest(ctx, http.MethodPut, endpoint.String(), payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
