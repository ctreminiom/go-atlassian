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

// SpaceV2Service provides methods to interact with space operations in Confluence V2.
type SpaceV2Service struct {
	// internalClient is the connector interface for space operations.
	internalClient confluence.SpaceV2Connector
}

// Bulk returns all spaces.
//
// The results will be sorted by id ascending.
//
// The number of results is limited by the limit parameter and additional results (if available)
//
// will be available through the next URL present in the Link response header.
//
// GET /wiki/api/v2/spaces
//
// https://docs.go-atlassian.io/confluence-cloud/v2/space#get-spaces
func (s *SpaceV2Service) Bulk(ctx context.Context, options *model.GetSpacesOptionSchemeV2, cursor string, limit int) (*model.SpaceChunkV2Scheme, *model.ResponseScheme, error) {
	return s.internalClient.Bulk(ctx, options, cursor, limit)
}

// Get returns a specific space.
//
// GET /wiki/api/v2/spaces/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/space#get-space-by-id
func (s *SpaceV2Service) Get(ctx context.Context, spaceID int, descriptionFormat string) (*model.SpaceSchemeV2, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, spaceID, descriptionFormat)
}

// Permissions returns space permissions for a specific space.
//
// GET /wiki/api/v2/spaces/{id}/permissions
func (s *SpaceV2Service) Permissions(ctx context.Context, spaceID int, cursor string, limit int) (*model.SpacePermissionPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Permissions(ctx, spaceID, cursor, limit)
}

func NewSpaceV2Service(client service.Connector) *SpaceV2Service {

	return &SpaceV2Service{
		internalClient: &internalSpaceV2Impl{c: client},
	}
}

type internalSpaceV2Impl struct {
	c service.Connector
}

func (i *internalSpaceV2Impl) Permissions(ctx context.Context, spaceID int, cursor string, limit int) (*model.SpacePermissionPageScheme, *model.ResponseScheme, error) {

	if spaceID == 0 {
		return nil, nil, model.ErrNoSpaceID
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/spaces/%v/permissions?%v", spaceID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	chunk := new(model.SpacePermissionPageScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	return chunk, response, nil
}

func (i *internalSpaceV2Impl) Bulk(ctx context.Context, options *model.GetSpacesOptionSchemeV2, cursor string, limit int) (*model.SpaceChunkV2Scheme, *model.ResponseScheme, error) {

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	if options != nil {

		if len(options.IDs) != 0 {
			query.Add("ids", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(options.IDs)), ","), "[]"))
		}

		if len(options.Keys) != 0 {
			query.Add("keys", strings.Join(options.Keys, ","))
		}

		if options.Type != "" {
			query.Add("type", options.Type)
		}

		if options.Status != "" {
			query.Add("status", options.Status)
		}

		if len(options.Labels) != 0 {
			query.Add("labels", strings.Join(options.Labels, ","))
		}

		if options.Sort != "" {
			query.Add("sort", options.Sort)
		}

		if options.DescriptionFormat != "" {
			query.Add("description-format", options.DescriptionFormat)
		}

		if options.SerializeIDs {
			query.Add("serialize-ids-as-strings", "true")
		}

	}

	endpoint := fmt.Sprintf("wiki/api/v2/spaces?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	chunk := new(model.SpaceChunkV2Scheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	return chunk, response, nil
}

func (i *internalSpaceV2Impl) Get(ctx context.Context, spaceID int, descriptionFormat string) (*model.SpaceSchemeV2, *model.ResponseScheme, error) {

	if spaceID == 0 {
		return nil, nil, model.ErrNoSpaceID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/api/v2/spaces/%v", spaceID))

	if descriptionFormat != "" {
		query := url.Values{}
		query.Add("description-format", descriptionFormat)

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	space := new(model.SpaceSchemeV2)
	response, err := i.c.Call(request, space)
	if err != nil {
		return nil, response, err
	}

	return space, response, nil
}
