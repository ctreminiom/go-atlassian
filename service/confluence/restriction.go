package confluence

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type ContentRestrictionConnector interface {

	// Gets returns the restrictions on a piece of content.
	//
	// GET /wiki/rest/api/content/{id}/restriction
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions#get-restrictions
	Gets(ctx context.Context, contentID string, expand []string, startAt, maxResults int) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error)

	// Add adds restrictions to a piece of content. Note, this does not change any existing restrictions on the content.
	//
	// POST /wiki/rest/api/content/{id}/restriction
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions#add-restrictions
	Add(ctx context.Context, contentID string, payload *model.ContentRestrictionUpdatePayloadScheme, expand []string) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error)

	// Delete removes all restrictions (read and update) on a piece of content.
	//
	// DELETE /wiki/rest/api/content/{id}/restriction
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions#delete-restrictions
	Delete(ctx context.Context, contentID string, expand []string) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error)

	// Update updates restrictions for a piece of content. This removes the existing restrictions and replaces them with the restrictions in the request.
	//
	// PUT /wiki/rest/api/content/{id}/restriction
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions#update-restrictions
	Update(ctx context.Context, contentID string, payload *model.ContentRestrictionUpdatePayloadScheme, expand []string) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error)
}

type RestrictionOperationConnector interface {

	// Gets returns restrictions on a piece of content by operation.
	//
	// This method is similar to Get restrictions except that the operations are properties
	//
	// of the return object, rather than items in a results array.
	//
	// GET /wiki/rest/api/content/{id}/restriction
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations#get-restrictions-by-operation
	Gets(ctx context.Context, contentID string, expand []string) (*model.ContentRestrictionByOperationScheme, *model.ResponseScheme, error)

	// Get returns the restrictions on a piece of content for a given operation (read or update).
	//
	// GET /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations#get-restrictions-for-operation
	Get(ctx context.Context, contentID, operationKey string, expand []string, startAt, maxResults int) (*model.ContentRestrictionScheme, *model.ResponseScheme, error)
}

type RestrictionGroupOperationConnector interface {

	// Get returns whether the specified content restriction applies to a group.
	//
	// Note that a response of true does not guarantee that the group can view the page,
	//
	// as it does not account for account-inherited restrictions, space permissions, or even product access.
	//
	// GET /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/byGroupId/{groupNameOrID}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/group#get-content-restriction-status-for-group
	Get(ctx context.Context, contentID, operationKey, groupNameOrID string) (*model.ResponseScheme, error)

	// Add adds a group to a content restriction.
	//
	// That is, grant read or update permission to the group for a piece of content.
	//
	// PUT /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/byGroupId/{groupNameOrID}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/group#add-group-to-content-restriction
	Add(ctx context.Context, contentID, operationKey, groupNameOrID string) (*model.ResponseScheme, error)

	// Remove removes a group from a content restriction.
	//
	// That is, remove read or update permission for the group for a piece of content.
	//
	// DELETE /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/byGroupId/{groupNameOrID}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/group#remove-group-from-content-restriction
	Remove(ctx context.Context, contentID, operationKey, groupNameOrID string) (*model.ResponseScheme, error)
}

type RestrictionUserOperationConnector interface {

	// Get returns whether the specified content restriction applies to a user.
	//
	// GET /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/user
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/user#get-content-restriction-status-for-user
	Get(ctx context.Context, contentID, operationKey, accountID string) (*model.ResponseScheme, error)

	// Add adds a user to a content restriction.
	//
	// That is, grant read or update permission to the user for a piece of content.
	//
	// PUT /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/user
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/user#add-user-to-content-restriction
	Add(ctx context.Context, contentID, operationKey, accountID string) (*model.ResponseScheme, error)

	// Remove removes a group from a content restriction.
	//
	// That is, remove read or update permission for the group for a piece of content.
	//
	// DELETE /wiki/rest/api/content/{id}/restriction/byOperation/{operationKey}/user
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/restrictions/operations/user#remove-user-from-content-restriction
	Remove(ctx context.Context, contentID, operationKey, accountID string) (*model.ResponseScheme, error)
}
