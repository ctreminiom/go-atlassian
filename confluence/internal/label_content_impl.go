package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewContentLabelService creates a new instance of ContentLabelService.
// It takes a service.Connector as input and returns a pointer to ContentLabelService.
func NewContentLabelService(client service.Connector) *ContentLabelService {
	return &ContentLabelService{
		internalClient: &internalContentLabelImpl{c: client},
	}
}

// ContentLabelService provides methods to interact with content label operations in Confluence.
type ContentLabelService struct {
	// internalClient is the connector interface for label operations.
	internalClient confluence.LabelsConnector
}

// Gets returns the labels on a piece of content.
//
// GET /wiki/rest/api/content/{id}/label
//
// https://docs.go-atlassian.io/confluence-cloud/content/labels#get-labels-for-content
func (c *ContentLabelService) Gets(ctx context.Context, contentID, prefix string, startAt, maxResults int) (*model.ContentLabelPageScheme, *model.ResponseScheme, error) {
	return c.internalClient.Gets(ctx, contentID, prefix, startAt, maxResults)
}

// Add adds labels to a piece of content. Does not modify the existing labels.
//
// POST /wiki/rest/api/content/{id}/label
//
// https://docs.go-atlassian.io/confluence-cloud/content/labels#add-labels-to-content
func (c *ContentLabelService) Add(ctx context.Context, contentID string, payload []*model.ContentLabelPayloadScheme, want400Response bool) (*model.ContentLabelPageScheme, *model.ResponseScheme, error) {
	return c.internalClient.Add(ctx, contentID, payload, want400Response)
}

// Remove removes a label from a piece of content
//
// DELETE /wiki/rest/api/content/{id}/label/{label}
//
// https://docs.go-atlassian.io/confluence-cloud/content/labels#remove-label-from-content
func (c *ContentLabelService) Remove(ctx context.Context, contentID, labelName string) (*model.ResponseScheme, error) {
	return c.internalClient.Remove(ctx, contentID, labelName)
}

type internalContentLabelImpl struct {
	c service.Connector
}

func (i *internalContentLabelImpl) Gets(ctx context.Context, contentID, prefix string, startAt, maxResults int) (*model.ContentLabelPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(prefix) != 0 {
		query.Add("prefix", prefix)
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/label?%v", contentID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentLabelPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalContentLabelImpl) Add(ctx context.Context, contentID string, payload []*model.ContentLabelPayloadScheme, want400Response bool) (*model.ContentLabelPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/label", contentID))

	if want400Response {
		query := url.Values{}
		query.Add("use-400-error-response", "true")

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentLabelPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalContentLabelImpl) Remove(ctx context.Context, contentID, labelName string) (*model.ResponseScheme, error) {

	if contentID == "" {
		return nil, model.ErrNoContentID
	}

	if labelName == "" {
		return nil, model.ErrNoContentLabel
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/label/%v", contentID, labelName)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
