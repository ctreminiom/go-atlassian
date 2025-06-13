package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
)

// NewCustomContentService creates a new instance of CustomContentService.
// It takes a service.Connector as input and returns a pointer to CustomContentService.
func NewCustomContentService(client service.Connector) *CustomContentService {
	return &CustomContentService{
		internalClient: &internalCustomContentServiceImpl{c: client},
	}
}

// CustomContentService provides methods to interact with custom content operations in Confluence.
type CustomContentService struct {
	// internalClient is the connector interface for custom content operations.
	internalClient confluence.CustomContentConnector
}

// Gets returns all custom content for a given type.
//
// GET /wiki/api/v2/custom-content
//
// https://docs.go-atlassian.io/confluence-cloud/v2/custom-content#get-custom-content-by-type
func (c *CustomContentService) Gets(ctx context.Context, typ string, options *model.CustomContentOptionsScheme, cursor string, limit int) (*model.CustomContentPageScheme, *model.ResponseScheme, error) {
	return c.internalClient.Gets(ctx, typ, options, cursor, limit)
}

// Create creates a new custom content in the given space, page, blogpost or other custom content.
//
// POST /wiki/api/v2/custom-content
//
// https://docs.go-atlassian.io/confluence-cloud/v2/custom-content#create-custom-content
func (c *CustomContentService) Create(ctx context.Context, payload *model.CustomContentPayloadScheme) (*model.CustomContentScheme, *model.ResponseScheme, error) {
	return c.internalClient.Create(ctx, payload)
}

// Get returns a specific piece of custom content.
//
// GET /wiki/api/v2/custom-content/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/custom-content#get-custom-content-by-id
func (c *CustomContentService) Get(ctx context.Context, customContentID int, format string, versionID int) (*model.CustomContentScheme, *model.ResponseScheme, error) {
	return c.internalClient.Get(ctx, customContentID, format, versionID)
}

// Update updates a custom content by id.
//
// The spaceId is always required and maximum one of pageId, blogPostId,
//
// or customContentId is allowed in the request body
//
// PUT /wiki/api/v2/custom-content/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/custom-content#update-custom-content
func (c *CustomContentService) Update(ctx context.Context, customContentID int, payload *model.CustomContentPayloadScheme) (*model.CustomContentScheme, *model.ResponseScheme, error) {
	return c.internalClient.Update(ctx, customContentID, payload)
}

// Delete deletes a custom content by id.
//
// DELETE /wiki/api/v2/custom-content/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/custom-content#delete-custom-content
func (c *CustomContentService) Delete(ctx context.Context, customContentID int) (*model.ResponseScheme, error) {
	return c.internalClient.Delete(ctx, customContentID)
}

type internalCustomContentServiceImpl struct {
	c service.Connector
}

func (i *internalCustomContentServiceImpl) Gets(ctx context.Context, typ string, options *model.CustomContentOptionsScheme, cursor string, limit int) (*model.CustomContentPageScheme, *model.ResponseScheme, error) {

	if typ == "" {
		return nil, nil, model.ErrNoCustomContentType
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	if options != nil {

		if len(options.IDs) != 0 {

			var ids []string
			for _, id := range options.IDs {
				ids = append(ids, strconv.Itoa(id))
			}

			query.Add("id", strings.Join(ids, ","))
		}

		if len(options.SpaceIDs) != 0 {

			var ids []string
			for _, id := range options.IDs {
				ids = append(ids, strconv.Itoa(id))
			}

			query.Add("space-id", strings.Join(ids, ","))
		}

		if options.Sort != "" {
			query.Add("sort", options.Sort)
		}

		if options.BodyFormat != "" {
			query.Add("body-format", options.BodyFormat)
		}
	}

	endpoint := fmt.Sprintf("wiki/api/v2/custom-content?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.CustomContentPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalCustomContentServiceImpl) Create(ctx context.Context, payload *model.CustomContentPayloadScheme) (*model.CustomContentScheme, *model.ResponseScheme, error) {

	endpoint := "wiki/api/v2/custom-content"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	customContent := new(model.CustomContentScheme)
	response, err := i.c.Call(request, customContent)
	if err != nil {
		return nil, response, err
	}

	return customContent, response, nil
}

func (i *internalCustomContentServiceImpl) Get(ctx context.Context, customContentID int, format string, versionID int) (*model.CustomContentScheme, *model.ResponseScheme, error) {

	if customContentID == 0 {
		return nil, nil, model.ErrNoCustomContentID
	}

	query := url.Values{}

	if format != "" {
		query.Add("body-format", format)
	}

	if versionID != 0 {
		query.Add("version", strconv.Itoa(versionID))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/api/v2/custom-content/%v", customContentID))

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	customContent := new(model.CustomContentScheme)
	response, err := i.c.Call(request, customContent)
	if err != nil {
		return nil, response, err
	}

	return customContent, response, nil
}

func (i *internalCustomContentServiceImpl) Update(ctx context.Context, customContentID int, payload *model.CustomContentPayloadScheme) (*model.CustomContentScheme, *model.ResponseScheme, error) {

	if customContentID == 0 {
		return nil, nil, model.ErrNoCustomContentID
	}

	endpoint := fmt.Sprintf("wiki/api/v2/custom-content/%v", customContentID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	customContent := new(model.CustomContentScheme)
	response, err := i.c.Call(request, customContent)
	if err != nil {
		return nil, response, err
	}

	return customContent, response, nil

}

func (i *internalCustomContentServiceImpl) Delete(ctx context.Context, customContentID int) (*model.ResponseScheme, error) {

	if customContentID == 0 {
		return nil, model.ErrNoCustomContentID
	}

	endpoint := fmt.Sprintf("wiki/api/v2/custom-content/%v", customContentID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
