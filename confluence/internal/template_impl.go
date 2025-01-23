package internal

import (
	"context"
	"net/http"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
)

// NewTemplateService creates a new instance of TemplateService.
// It takes a service.Connector as inputs and returns a pointer to TemplateService.
func NewTemplateService(client service.Connector) *TemplateService {
	return &TemplateService{
		internalClient: &internalTemplateImpl{c: client},
	}
}

// TemplateService provides methods to interact with template operations in Confluence.
type TemplateService struct {
	// internalClient is the connector interface for content operations.
	internalClient confluence.TemplateConnector
}

// Create creates a new template.
//
// POST /wiki/rest/api/template
//
// https://docs.go-atlassian.io/confluence-cloud/template#create-content-template
func (t *TemplateService) Create(ctx context.Context, payload *models.CreateTemplateScheme) (*models.ContentTemplateScheme, *models.ResponseScheme, error) {
	return t.internalClient.Create(ctx, payload)
}

// Update updates a template.
//
// PUT /wiki/rest/api/template
//
// https://docs.go-atlassian.io/confluence-cloud/template#update-content-template
func (t *TemplateService) Update(ctx context.Context, payload *models.UpdateTemplateScheme) (*models.ContentTemplateScheme, *models.ResponseScheme, error) {
	return t.internalClient.Update(ctx, payload)
}

// Get content template by ID.
//
// GET /wiki/rest/api/template/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/template#get-content-template
func (t *TemplateService) Get(ctx context.Context, templateID string) (*models.ContentTemplateScheme, *models.ResponseScheme, error) {
	return t.internalClient.Get(ctx, templateID)
}

// internalTemplateImpl is the internal implementation of TemplateService.
type internalTemplateImpl struct {
	c service.Connector
}

// Create implements TemplateService.Create.
func (i *internalTemplateImpl) Create(ctx context.Context, payload *models.CreateTemplateScheme) (*models.ContentTemplateScheme, *models.ResponseScheme, error) {
	endpoint := "/wiki/rest/api/template"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	result := new(models.ContentTemplateScheme)
	response, err := i.c.Call(request, result)
	if err != nil {
		return nil, response, err
	}

	return result, response, nil
}

// Update implements TemplateService.Update.
func (i *internalTemplateImpl) Update(ctx context.Context, payload *models.UpdateTemplateScheme) (*models.ContentTemplateScheme, *models.ResponseScheme, error) {
	endpoint := "/wiki/rest/api/template"

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	result := new(models.ContentTemplateScheme)
	response, err := i.c.Call(request, result)
	if err != nil {
		return nil, response, err
	}

	return result, response, nil
}

// Get implements TemplateService.Get.
func (i *internalTemplateImpl) Get(ctx context.Context, templateID string) (*models.ContentTemplateScheme, *models.ResponseScheme, error) {
	endpoint := "/wiki/rest/api/template/" + templateID

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(models.ContentTemplateScheme)
	response, err := i.c.Call(request, result)
	if err != nil {
		return nil, response, err
	}

	return result, response, nil
}
