package admin

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// SCIMUserConnector represents the cloud admin SCIM user actions.
// Use it to search, get, create, delete, and change user.
type SCIMUserConnector interface {

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
	Create(ctx context.Context, directoryID string, payload *model.SCIMUserScheme, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error)

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
	Gets(ctx context.Context, directoryID string, opts *model.SCIMUserGetsOptionsScheme, startIndex, count int) (*model.SCIMUserPageScheme, *model.ResponseScheme, error)

	// Get gets a user from a directory by userId.
	//
	// GET /scim/directory/{directoryId}/Users/{userId}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#get-a-user-by-id
	Get(ctx context.Context, directoryID, userID string, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error)

	// Deactivate deactivates a user by userId.
	//
	// The user is not available for future requests until activated again.
	//
	// Any future operation for the deactivated user returns the 404 (resource not found) error.
	//
	// DELETE /scim/directory/{directoryId}/Users/{userId}
	Deactivate(ctx context.Context, directoryID, userID string) (*model.ResponseScheme, error)

	// Path updates a user's information in a directory by userId via PATCH.
	//
	// Refer to GET /ServiceProviderConfig for details on the supported operations.
	//
	// PATCH /scim/directory/{directoryId}/Users/{userId}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#update-user-by-id-patch
	Path(ctx context.Context, directoryID, userID string, payload *model.SCIMUserToPathScheme, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error)

	// Update updates a user's information in a directory by userId via user attributes.
	//
	// User information is replaced attribute-by-attribute, except immutable and read-only attributes.
	//
	// Existing values of unspecified attributes are cleaned.
	//
	// PUT /scim/directory/{directoryId}/Users/{userId}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/users#update-user-via-user-attributes
	Update(ctx context.Context, directoryID, userID string, payload *model.SCIMUserScheme, attributes, excludedAttributes []string) (*model.SCIMUserScheme, *model.ResponseScheme, error)
}
