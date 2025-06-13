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

// NewSpaceService creates a new instance of SpaceService.
// It takes a service.Connector and a pointer to SpacePermissionService as input and returns a pointer to SpaceService.
func NewSpaceService(client service.Connector, permission *SpacePermissionService) *SpaceService {
	return &SpaceService{
		internalClient: &internalSpaceImpl{c: client},
		Permission:     permission,
	}
}

// SpaceService provides methods to interact with space operations in Confluence.
type SpaceService struct {
	// internalClient is the connector interface for space operations.
	internalClient confluence.SpaceConnector
	// Permission is a pointer to SpacePermissionService for additional permission operations.
	Permission *SpacePermissionService
}

// Gets returns all spaces.
//
// The returned spaces are ordered alphabetically in ascending order by space key.
//
// GET /wiki/rest/api/space
//
// https://docs.go-atlassian.io/confluence-cloud/space#get-spaces
func (s *SpaceService) Gets(ctx context.Context, options *model.GetSpacesOptionScheme, startAt, maxResults int) (result *model.SpacePageScheme, response *model.ResponseScheme, err error) {
	return s.internalClient.Gets(ctx, options, startAt, maxResults)
}

// Create creates a new space.
//
// Note, currently you cannot set space labels when creating a space.
//
// POST /wiki/rest/api/space
//
// https://docs.go-atlassian.io/confluence-cloud/space#create-space
func (s *SpaceService) Create(ctx context.Context, payload *model.CreateSpaceScheme, private bool) (*model.SpaceScheme, *model.ResponseScheme, error) {
	return s.internalClient.Create(ctx, payload, private)
}

// Get returns a space.
//
// This includes information like the name, description, and permissions, but not the content in the space.
//
// GET /wiki/rest/api/space/{spaceKey}
//
// https://docs.go-atlassian.io/confluence-cloud/space#get-space
func (s *SpaceService) Get(ctx context.Context, spaceKey string, expand []string) (*model.SpaceScheme, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, spaceKey, expand)
}

// Update updates the name, description, or homepage of a space.
//
// PUT /wiki/rest/api/space/{spaceKey}
//
// https://docs.go-atlassian.io/confluence-cloud/space#update-space
func (s *SpaceService) Update(ctx context.Context, spaceKey string, payload *model.UpdateSpaceScheme) (*model.SpaceScheme, *model.ResponseScheme, error) {
	return s.internalClient.Update(ctx, spaceKey, payload)
}

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
func (s *SpaceService) Delete(ctx context.Context, spaceKey string) (*model.ContentTaskScheme, *model.ResponseScheme, error) {
	return s.internalClient.Delete(ctx, spaceKey)
}

// Content returns all content in a space.
//
// The returned content is grouped by type (pages then blogposts), then ordered by content ID in ascending order.
//
// GET /wiki/rest/api/space/{spaceKey}/content
//
// https://docs.go-atlassian.io/confluence-cloud/space#get-content-for-space
func (s *SpaceService) Content(ctx context.Context, spaceKey, depth string, expand []string, startAt, maxResults int) (*model.ContentChildrenScheme, *model.ResponseScheme, error) {
	return s.internalClient.Content(ctx, spaceKey, depth, expand, startAt, maxResults)
}

// ContentByType returns all content of a given type, in a space.
//
// The returned content is ordered by content ID in ascending order.
//
// GET /wiki/rest/api/space/{spaceKey}/content/{type}
//
// https://docs.go-atlassian.io/confluence-cloud/space#get-content-by-type-for-space
func (s *SpaceService) ContentByType(ctx context.Context, spaceKey, contentType, depth string, expand []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.ContentByType(ctx, spaceKey, contentType, depth, expand, startAt, maxResults)
}

type internalSpaceImpl struct {
	c service.Connector
}

func (i *internalSpaceImpl) Gets(ctx context.Context, options *model.GetSpacesOptionScheme, startAt, maxResults int) (*model.SpacePageScheme, *model.ResponseScheme, error) {

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.SpaceKeys) != 0 {

			for _, key := range options.SpaceKeys {
				query.Add("spaceKey", key)
			}
		}

		if len(options.SpaceIDs) != 0 {

			for _, id := range options.SpaceIDs {
				query.Add("spaceID", strconv.Itoa(id))
			}
		}

		if len(options.SpaceType) != 0 {
			query.Add("type", options.SpaceType)
		}

		if len(options.Status) != 0 {
			query.Add("status", options.Status)
		}

		if len(options.Labels) != 0 {
			query.Add("label", strings.Join(options.Labels, ","))
		}

		if options.Favorite {
			query.Add("favorite", "true")
		}

		if len(options.FavoriteUserKey) != 0 {
			query.Add("favouriteUserKey", options.FavoriteUserKey)
		}

		if len(options.Expand) != 0 {
			query.Add("expand", strings.Join(options.Expand, ","))
		}
	}

	var endpoint = fmt.Sprintf("wiki/rest/api/space?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.SpacePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalSpaceImpl) Create(ctx context.Context, payload *model.CreateSpaceScheme, private bool) (*model.SpaceScheme, *model.ResponseScheme, error) {

	if payload != nil {

		if payload.Name == "" {
			return nil, nil, model.ErrNoSpaceName
		}

		if payload.Key == "" {
			return nil, nil, model.ErrNoSpaceKey
		}

	}

	var endpoint strings.Builder
	endpoint.WriteString("wiki/rest/api/space")

	if private {
		endpoint.WriteString("/_private")
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	space := new(model.SpaceScheme)
	response, err := i.c.Call(request, space)
	if err != nil {
		return nil, response, err
	}

	return space, response, nil
}

func (i *internalSpaceImpl) Get(ctx context.Context, spaceKey string, expand []string) (*model.SpaceScheme, *model.ResponseScheme, error) {

	if spaceKey == "" {
		return nil, nil, model.ErrNoSpaceKey
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/space/%v", spaceKey))

	if expand != nil {
		query := url.Values{}
		query.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	space := new(model.SpaceScheme)
	response, err := i.c.Call(request, space)
	if err != nil {
		return nil, response, err
	}

	return space, response, nil
}

func (i *internalSpaceImpl) Update(ctx context.Context, spaceKey string, payload *model.UpdateSpaceScheme) (*model.SpaceScheme, *model.ResponseScheme, error) {

	if spaceKey == "" {
		return nil, nil, model.ErrNoSpaceKey
	}

	endpoint := fmt.Sprintf("wiki/rest/api/space/%v", spaceKey)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	space := new(model.SpaceScheme)
	response, err := i.c.Call(request, space)
	if err != nil {
		return nil, response, err
	}

	return space, response, nil
}

func (i *internalSpaceImpl) Delete(ctx context.Context, spaceKey string) (*model.ContentTaskScheme, *model.ResponseScheme, error) {

	if spaceKey == "" {
		return nil, nil, model.ErrNoSpaceKey
	}

	endpoint := fmt.Sprintf("wiki/rest/api/space/%v", spaceKey)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	task := new(model.ContentTaskScheme)
	response, err := i.c.Call(request, task)
	if err != nil {
		return nil, response, err
	}

	return task, response, nil
}

func (i *internalSpaceImpl) Content(ctx context.Context, spaceKey, depth string, expand []string, startAt, maxResults int) (*model.ContentChildrenScheme, *model.ResponseScheme, error) {

	if spaceKey == "" {
		return nil, nil, model.ErrNoSpaceKey
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if len(depth) != 0 {
		query.Add("depth", depth)
	}

	endpoint := fmt.Sprintf("wiki/rest/api/space/%v/content?%v", spaceKey, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	children := new(model.ContentChildrenScheme)
	response, err := i.c.Call(request, children)
	if err != nil {
		return nil, response, err
	}

	return children, response, nil
}

func (i *internalSpaceImpl) ContentByType(ctx context.Context, spaceKey, contentType, depth string, expand []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {

	if spaceKey == "" {
		return nil, nil, model.ErrNoSpaceKey
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if len(depth) != 0 {
		query.Add("depth", depth)
	}

	endpoint := fmt.Sprintf("wiki/rest/api/space/%v/content/%v?%v", spaceKey, contentType, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
