package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type ContentPropertyConnector interface {

	// Gets returns the properties for a piece of content.
	//
	// GET /wiki/rest/api/content/{id}/property
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/properties#get-content-properties
	Gets(ctx context.Context, contentID string, expand []string, startAt, maxResults int) (*model.ContentPropertyPageScheme, *model.ResponseScheme, error)

	// Create creates a property for an existing piece of content.
	//
	// POST /wiki/rest/api/content/{id}/property
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/properties#create-content-property
	Create(ctx context.Context, contentID string, payload *model.ContentPropertyPayloadScheme) (*model.ContentPropertyScheme, *model.ResponseScheme, error)

	// Get returns a content property for a piece of content.
	//
	// GET /wiki/rest/api/content/{id}/property/{key}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/properties#get-content-property
	Get(ctx context.Context, contentID, key string) (*model.ContentPropertyScheme, *model.ResponseScheme, error)

	// Delete deletes a content property.
	//
	// DELETE /wiki/rest/api/content/{id}/property/{key}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/properties#delete-content-property
	Delete(ctx context.Context, contentID, key string) (*model.ResponseScheme, error)
}
