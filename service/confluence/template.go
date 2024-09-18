package confluence

import (
	"context"

	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// TemplateConnector provides methods to interact with template operations in Confluence.
type TemplateConnector interface {

	// Update updates a template.
	//
	// PUT /wiki/rest/api/template
	//
	// https://developer.atlassian.com/cloud/confluence/rest/v1/api-group-template/#api-wiki-rest-api-template-put
	Update(ctx context.Context, payload *models.UpdateTemplateScheme) (*models.ContentTemplateScheme, *models.ResponseScheme, error)

	// Create creates a new template.
	//
	// POST /wiki/rest/api/template
	//
	// https://developer.atlassian.com/cloud/confluence/rest/v1/api-group-template/#api-wiki-rest-api-template-post
	Create(ctx context.Context, payload *models.CreateTemplateScheme) (*models.ContentTemplateScheme, *models.ResponseScheme, error)

	// Get content template by ID.
	//
	// GET /wiki/rest/api/template/{id}
	//
	// https://developer.atlassian.com/cloud/confluence/rest/v1/api-group-template/#api-wiki-rest-api-template-contenttemplateid-get
	Get(ctx context.Context, templateID string) (*models.ContentTemplateScheme, *models.ResponseScheme, error)
}
