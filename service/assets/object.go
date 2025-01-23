package assets

import (
	"context"

	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// ObjectConnector represents the Assets objects.
// Use it to search, get, create, delete, and change objects.
type ObjectConnector interface {

	// Get loads one object.
	//
	// GET /jsm/assets/workspace/{workspaceID}/v1/object/{objectID}
	//
	// https://docs.go-atlassian.io/jira-assets/object#get-object-by-id
	Get(ctx context.Context, workspaceID, objectID string) (*models.ObjectScheme, *models.ResponseScheme, error)

	// Update updates an existing object in Assets.
	//
	// PUT /jsm/assets/workspace/{workspaceID}/v1/object/{objectID}
	//
	// https://docs.go-atlassian.io/jira-assets/object#update-object-by-id
	Update(ctx context.Context, workspaceID, objectID string, payload *models.ObjectPayloadScheme) (*models.ObjectScheme, *models.ResponseScheme, error)

	// Delete deletes the referenced object
	//
	// DELETE /jsm/assets/workspace/{workspaceID}/v1/object/{objectID}
	//
	// https://docs.go-atlassian.io/jira-assets/object#delete-object-by-id
	Delete(ctx context.Context, workspaceID, objectID string) (*models.ResponseScheme, error)

	// Attributes list all attributes for the given object.
	//
	// GET /jsm/assets/workspace/{workspaceID}/v1/object/{objectID}/attributes
	//
	// https://docs.go-atlassian.io/jira-assets/object#get-object-attributes
	Attributes(ctx context.Context, workspaceID, objectID string) ([]*models.ObjectAttributeScheme, *models.ResponseScheme, error)

	// History retrieves the history entries for this object.
	//
	// GET /jsm/assets/workspace/{workspaceID}/v1/object/{objectID}/history
	//
	// https://docs.go-atlassian.io/jira-assets/object#get-object-changelogs
	History(ctx context.Context, workspaceID, objectID string, ascOrder bool) ([]*models.ObjectHistoryScheme, *models.ResponseScheme, error)

	// References finds all references for an object.
	//
	// GET /jsm/assets/workspace/{workspaceID}/v1/object/{objectID}/referenceinfo
	//
	// https://docs.go-atlassian.io/jira-assets/object#get-object-references
	References(ctx context.Context, workspaceID, objectID string) ([]*models.ObjectReferenceTypeInfoScheme, *models.ResponseScheme, error)

	// Create creates a new object in Assets.
	//
	// POST /jsm/assets/workspace/{workspaceID}/v1/object/create
	//
	// https://docs.go-atlassian.io/jira-assets/object#create-object
	Create(ctx context.Context, workspaceID string, payload *models.ObjectPayloadScheme) (*models.ObjectScheme, *models.ResponseScheme, error)

	// Relation returns the relation between Jira issues and Assets objects
	//
	// GET /jsm/assets/workspace/{workspaceID}/v1/objectconnectedtickets/{objectID}/tickets
	//
	// https://docs.go-atlassian.io/jira-assets/object#get-object-tickets
	Relation(ctx context.Context, workspaceID, objectID string) (*models.TicketPageScheme, *models.ResponseScheme, error)

	// Filter fetch Objects by AQL.
	//
	// POST /jsm/assets/workspace/{workspaceID}/v1/object/aql
	//
	// https://docs.go-atlassian.io/jira-assets/object#filter-objects
	Filter(ctx context.Context, workspaceID, aql string, attributes bool, startAt, maxResults int) (*models.ObjectListResultScheme, *models.ResponseScheme, error)

	// Search retrieve a list of objects based on an AQL.
	//
	// Note that the preferred endpoint is /aql
	//
	// POST /jsm/assets/workspace/{workspaceID}/v1/object/navlist/aql
	//
	// https://docs.go-atlassian.io/jira-assets/object#search-objects
	Search(ctx context.Context, workspaceID string, payload *models.ObjectSearchParamsScheme) (*models.ObjectListScheme, *models.ResponseScheme, error)
}
