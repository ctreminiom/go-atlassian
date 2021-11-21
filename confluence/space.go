package confluence

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SpaceService struct {
	client *Client
}

// Gets returns all spaces. The returned spaces are ordered alphabetically in ascending order by space key.
func (s *SpaceService) Gets(ctx context.Context, options *model.GetSpacesOptionScheme, startAt, maxResults int) (
	result *model.SpacePageScheme, response *ResponseScheme, err error) {

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.SpaceKeys) != 0 {
			query.Add("spaceKey", strings.Join(options.SpaceKeys, ","))
		}

		if len(options.SpaceIDs) != 0 {

			for _, id := range options.SpaceIDs {
				query.Add("spaceId", strconv.Itoa(id))
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

	var endpoint = fmt.Sprintf("/wiki/rest/api/space?%v", query.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Create creates a new space.
// Note, currently you cannot set space labels when creating a space.
func (s *SpaceService) Create(ctx context.Context, payload *model.CreateSpaceScheme, private bool) (
	result *model.SpaceScheme, response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	if len(payload.Name) == 0 {
		return nil, nil, model.ErrNoSpaceNameError
	}

	if len(payload.Key) == 0 {
		return nil, nil, model.ErrNoSpaceKeyError
	}

	var endpoint strings.Builder
	endpoint.WriteString("/wiki/rest/api/space")

	if private {
		endpoint.WriteString("/_private")
	}

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint.String(), payloadAsReader)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Get returns a space.
// This includes information like the name, description, and permissions,
// but not the content in the space.
func (s *SpaceService) Get(ctx context.Context, spaceKey string, expand []string) (result *model.SpaceScheme,
	response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, nil, model.ErrNoSpaceKeyError
	}

	query := url.Values{}
	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/wiki/rest/api/space/%v", spaceKey))

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Update updates the name, description, or homepage of a space.
func (s *SpaceService) Update(ctx context.Context, spaceKey string, payload *model.UpdateSpaceScheme) (result *model.SpaceScheme,
	response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, nil, model.ErrNoSpaceKeyError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/space/%v", spaceKey)

	request, err := s.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Delete deletes a space.
// Note, the space will be deleted in a long running task.
// Therefore, the space may not be deleted yet when this method has returned.
// Clients should poll the status link that is returned in the response until the task completes.
func (s *SpaceService) Delete(ctx context.Context, spaceKey string) (result *model.ContentTaskScheme, response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, nil, model.ErrNoSpaceKeyError
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/space/%v", spaceKey)

	request, err := s.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err = s.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Content returns all content in a space.
// The returned content is grouped by type (pages then blogposts),
// then ordered by content ID in ascending order.
func (s *SpaceService) Content(ctx context.Context, spaceKey, depth string, expand []string, startAt, maxResults int) (
	result *model.ContentChildrenScheme, response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, nil, model.ErrNoSpaceKeyError
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

	var endpoint = fmt.Sprintf("/wiki/rest/api/space/%v/content?%v", spaceKey, query.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err = s.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// ContentByType returns all content of a given type, in a space.
// The returned content is ordered by content ID in ascending order.
func (s *SpaceService) ContentByType(ctx context.Context, spaceKey, contentType, depth string, expand []string, startAt,
	maxResults int) (result *model.ContentPageScheme, response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, nil, model.ErrNoSpaceKeyError
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

	var endpoint = fmt.Sprintf("/wiki/rest/api/space/%v/content/%v?%v", spaceKey, contentType, query.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err = s.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}
