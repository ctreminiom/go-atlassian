package admin

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// SCIMGroupConnector represents the cloud admin SCIM group actions.
// Use it to search, get, create, delete, and change groups.
type SCIMGroupConnector interface {

	// Gets get groups from a directory.
	//
	// Filtering is supported with a single exact match (eq) against the displayName attribute.
	//
	// Pagination is supported.
	//
	// Sorting is not supported.
	//
	// GET /scim/directory/{directoryID}/Groups
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#get-groups
	Gets(ctx context.Context, directoryID, filter string, startAt, maxResults int) (*model.ScimGroupPageScheme, *model.ResponseScheme, error)

	// Get returns a group from a directory by group ID.
	//
	// GET /scim/directory/{directoryID}/Groups/{groupID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#get-a-group-by-id
	Get(ctx context.Context, directoryID, groupID string) (*model.ScimGroupScheme, *model.ResponseScheme, error)

	// Update updates a group in a directory by group ID.
	//
	// PUT /scim/directory/{directoryID}/Groups/{groupID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#update-a-group-by-id
	Update(ctx context.Context, directoryID, groupID string, newGroupName string) (*model.ScimGroupScheme, *model.ResponseScheme, error)

	// Delete deletes a group from a directory.
	//
	// An attempt to delete a non-existent group fails with a 404 (Resource Not found) error.
	//
	// DELETE /scim/directory/{directoryID}/Groups/{groupID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#delete-a-group-by-id
	Delete(ctx context.Context, directoryID, groupID string) (*model.ResponseScheme, error)

	// Create creates a group in a directory.
	//
	// An attempt to create a group with an existing name fails with a 409 (Conflict) error.
	//
	// POST /scim/directory/{directoryID}/Groups
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#create-a-group
	Create(ctx context.Context, directoryID, groupName string) (*model.ScimGroupScheme, *model.ResponseScheme, error)

	// Path update a group's information in a directory by group id via PATCH.
	//
	// You can use this API to manage group membership.
	//
	// PATCH /scim/directory/{directoryID}/Groups/{groupID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/groups#update-a-group-by-id-patch
	Path(ctx context.Context, directoryID, groupID string, payload *model.SCIMGroupPathScheme) (*model.ScimGroupScheme, *model.ResponseScheme, error)
}
