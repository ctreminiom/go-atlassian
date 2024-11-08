package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type SpaceV2Connector interface {

	// Bulk returns all spaces.
	//
	// The results will be sorted by id ascending.
	//
	// The number of results is limited by the limit parameter and additional results (if available)
	//
	// will be available through the next URL present in the Link response header.
	//
	// GET /wiki/api/v2/spaces
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/space#get-spaces
	Bulk(ctx context.Context, options *model.GetSpacesOptionSchemeV2, cursor string, limit int) (*model.SpaceChunkV2Scheme, *model.ResponseScheme, error)

	// Get returns a specific space.
	//
	// GET /wiki/api/v2/spaces/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/space#get-space-by-id
	Get(ctx context.Context, spaceID int, descriptionFormat string) (*model.SpaceSchemeV2, *model.ResponseScheme, error)

	// Permissions returns space permissions for a specific space.
	//
	// GET /wiki/api/v2/spaces/{id}/permissions
	Permissions(ctx context.Context, spaceID int, cursor string, limit int) (*model.SpacePermissionPageScheme, *model.ResponseScheme, error)
}

type SpaceConnector interface {

	// Gets returns all spaces.
	//
	// The returned spaces are ordered alphabetically in ascending order by space key.
	//
	// GET /wiki/rest/api/space
	//
	// https://docs.go-atlassian.io/confluence-cloud/space#get-spaces
	Gets(ctx context.Context, options *model.GetSpacesOptionScheme, startAt, maxResults int) (result *model.SpacePageScheme, response *model.ResponseScheme, err error)

	// Create creates a new space.
	//
	// Note, currently you cannot set space labels when creating a space.
	//
	// POST /wiki/rest/api/space
	//
	// https://docs.go-atlassian.io/confluence-cloud/space#create-space
	Create(ctx context.Context, payload *model.CreateSpaceScheme, private bool) (*model.SpaceScheme, *model.ResponseScheme, error)

	// Get returns a space.
	//
	// This includes information like the name, description, and permissions, but not the content in the space.
	//
	// GET /wiki/rest/api/space/{spaceKey}
	//
	// https://docs.go-atlassian.io/confluence-cloud/space#get-space
	Get(ctx context.Context, spaceKey string, expand []string) (*model.SpaceScheme, *model.ResponseScheme, error)

	// Update updates the name, description, or homepage of a space.
	//
	// PUT /wiki/rest/api/space/{spaceKey}
	//
	// https://docs.go-atlassian.io/confluence-cloud/space#update-space
	Update(ctx context.Context, spaceKey string, payload *model.UpdateSpaceScheme) (*model.SpaceScheme, *model.ResponseScheme, error)

	// Delete deletes a space.
	//
	// Note, the space will be deleted in a long-running task.
	//
	// Therefore, the space may not be deleted yet when this method has returned.
	//
	// Clients should poll the status link that is returned to the response until the task completes.
	//
	// DELETE /wiki/rest/api/space/{spaceKey}
	//
	// https://docs.go-atlassian.io/confluence-cloud/space#delete-space
	Delete(ctx context.Context, spaceKey string) (*model.ContentTaskScheme, *model.ResponseScheme, error)

	// Content returns all content in a space.
	//
	// The returned content is grouped by type (pages then blogposts), then ordered by content ID in ascending order.
	//
	// GET /wiki/rest/api/space/{spaceKey}/content
	//
	// https://docs.go-atlassian.io/confluence-cloud/space#get-content-for-space
	Content(ctx context.Context, spaceKey, depth string, expand []string, startAt, maxResults int) (*model.ContentChildrenScheme, *model.ResponseScheme, error)

	// ContentByType returns all content of a given type, in a space.
	//
	// The returned content is ordered by content ID in ascending order.
	//
	// GET /wiki/rest/api/space/{spaceKey}/content/{type}
	//
	// https://docs.go-atlassian.io/confluence-cloud/space#get-content-by-type-for-space
	ContentByType(ctx context.Context, spaceKey, contentType, depth string, expand []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error)
}
