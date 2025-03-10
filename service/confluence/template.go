package confluence

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// TemplateConnector provides methods to interact with template operations in Confluence.
type TemplateConnector interface {
	// Create creates a new template.
	//
	// POST /wiki/rest/api/template
	//
	// https://docs.go-atlassian.io/confluence-cloud/template#create-content-template
	Create(ctx context.Context, payload *models.CreateTemplateScheme) (*models.ContentTemplateScheme, *models.ResponseScheme, error)

	// Update updates a template.
	//
	// PUT /wiki/rest/api/template
	//
	// https://docs.go-atlassian.io/confluence-cloud/template#update-content-template
	Update(ctx context.Context, payload *models.UpdateTemplateScheme) (*models.ContentTemplateScheme, *models.ResponseScheme, error)

	// Get content template by ID.
	//
	// GET /wiki/rest/api/template/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/template#get-content-template
	Get(ctx context.Context, templateID string) (*models.ContentTemplateScheme, *models.ResponseScheme, error)
}
