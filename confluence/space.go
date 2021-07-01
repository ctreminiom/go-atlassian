package confluence

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SpaceService struct {
	client *Client
}

type SpacePageScheme struct {
	Results []*SpaceScheme `json:"results,omitempty"`
	Start   int            `json:"start"`
	Limit   int            `json:"limit"`
	Size    int            `json:"size"`
	Links   struct {
		Base    string `json:"base"`
		Context string `json:"context"`
		Self    string `json:"self"`
	} `json:"_links"`
}

type GetSpacesOptionScheme struct {
	SpaceKeys       []string
	SpaceIDs        []int
	SpaceType       string
	Status          string
	Labels          []string
	Favorite        bool
	FavoriteUserKey string
	Expand          []string
}

// Gets returns all spaces. The returned spaces are ordered alphabetically in ascending order by space key.
func (s *SpaceService) Gets(ctx context.Context, options *GetSpacesOptionScheme, startAt, maxResults int) (
	result *SpacePageScheme, response *ResponseScheme, err error) {

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

type CreateSpaceScheme struct {
	Key              string                        `json:"key,omitempty"`
	Name             string                        `json:"name,omitempty"`
	Description      *CreateSpaceDescriptionScheme `json:"description,omitempty"`
	AnonymousAccess  bool                          `json:"anonymousAccess,omitempty"`
	UnlicensedAccess bool                          `json:"unlicensedAccess,omitempty"`
}

type CreateSpaceDescriptionScheme struct {
	Plain *CreateSpaceDescriptionPlainScheme `json:"plain"`
}

type CreateSpaceDescriptionPlainScheme struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

type SpacePermissionScheme struct {
	Subject   *SubjectPermissionScheme   `json:"subject,omitempty"`
	Operation *OperationPermissionScheme `json:"operation,omitempty"`
}

type OperationPermissionScheme struct {
	Operation  string `json:"operation,omitempty"`
	TargetType string `json:"targetType,omitempty"`
}

type SubjectPermissionScheme struct {
	User       *UserPermissionScheme  `json:"user,omitempty"`
	Group      *GroupPermissionScheme `json:"group,omitempty"`
	Expandable struct {
		User  string `json:"user,omitempty"`
		Group string `json:"group,omitempty"`
	} `json:"_expandable,omitempty"`
}

type UserPermissionScheme struct {
	Results []*UserScheme `json:"results,omitempty"`
	Size    int           `json:"size,omitempty"`
}

type GroupPermissionScheme struct {
	Results []*GroupScheme `json:"results,omitempty"`
	Size    int            `json:"size,omitempty"`
}

type GroupScheme struct {
	Type  string      `json:"type,omitempty"`
	Name  string      `json:"name,omitempty"`
	ID    string      `json:"id,omitempty"`
	Links *LinkScheme `json:"_links,omitempty"`
}

// Create creates a new space.
// Note, currently you cannot set space labels when creating a space.
func (s *SpaceService) Create(ctx context.Context, payload *CreateSpaceScheme, private bool) (
	result *SpaceScheme, response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	if len(payload.Name) == 0 {
		return nil, nil, notSpaceNameProvidedError
	}

	if len(payload.Key) == 0 {
		return nil, nil, notSpaceKeyProvidedError
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
func (s *SpaceService) Get(ctx context.Context, spaceKey string, expand []string) (result *SpaceScheme,
	response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, nil, notSpaceKeyProvidedError
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

type UpdateSpaceScheme struct {
	Name        string                        `json:"name,omitempty"`
	Description *CreateSpaceDescriptionScheme `json:"description,omitempty"`
	Homepage    *UpdateSpaceHomepageScheme    `json:"homepage,omitempty"`
}

type UpdateSpaceHomepageScheme struct {
	ID string `json:"id"`
}

// Update updates the name, description, or homepage of a space.
func (s *SpaceService) Update(ctx context.Context, spaceKey string, payload *UpdateSpaceScheme) (result *SpaceScheme,
	response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, nil, notSpaceKeyProvidedError
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
func (s *SpaceService) Delete(ctx context.Context, spaceKey string) (result *TaskScheme, response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, nil, notSpaceKeyProvidedError
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
	result *ContentChildrenScheme, response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, nil, notSpaceKeyProvidedError
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
	maxResults int) (result *ContentPageScheme, response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, nil, notSpaceKeyProvidedError
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

var (
	notSpaceNameProvidedError = fmt.Errorf("error, space name parameter is required, please provide a valid value")
	notSpaceKeyProvidedError  = fmt.Errorf("error, space key parameter is required, please provide a valid value")
)
