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

// NewFolderService creates a new instance of FolderService.
// It takes a service.Connector as input and returns a pointer to FolderService.
func NewFolderService(client service.Connector) *FolderService {
	return &FolderService{internalClient: &internalFolderImpl{c: client}}
}

// FolderService provides methods to interact with folder operations in Confluence.
type FolderService struct {
	// internalClient is the connector interface for folder operations.
	internalClient confluence.FolderConnector
}

// Get returns a specific folder.
//
// GET /wiki/api/v2/folders/{id}
//
// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
func (f *FolderService) Get(ctx context.Context, folderID string) (*model.FolderScheme, *model.ResponseScheme, error) {
	return f.internalClient.Get(ctx, folderID)
}

// Gets returns all folders that fit the filtering criteria.
//
// GET /wiki/api/v2/folders
//
// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
func (f *FolderService) Gets(ctx context.Context, options *model.FolderOptionsScheme, cursor string, limit int) (*model.FolderChunkScheme, *model.ResponseScheme, error) {
	return f.internalClient.Gets(ctx, options, cursor, limit)
}

// GetsBySpace returns all folders in a space.
//
// The number of results is limited by the limit parameter and additional results (if available)
//
// will be available through the next cursor
//
// GET /wiki/api/v2/spaces/{id}/folders
//
// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
func (f *FolderService) GetsBySpace(ctx context.Context, spaceID int, cursor string, limit int) (*model.FolderChunkScheme, *model.ResponseScheme, error) {
	return f.internalClient.GetsBySpace(ctx, spaceID, cursor, limit)
}

// GetsByParent returns all child folders of a parent folder.
//
// The number of results is limited by the limit parameter and additional results (if available)
//
// will be available through the next cursor
//
// GET /wiki/api/v2/folders/{id}/children
//
// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
func (f *FolderService) GetsByParent(ctx context.Context, parentID string, cursor string, limit int) (*model.FolderChunkScheme, *model.ResponseScheme, error) {
	return f.internalClient.GetsByParent(ctx, parentID, cursor, limit)
}

// Create creates a folder in the space.
//
// POST /wiki/api/v2/folders
//
// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
func (f *FolderService) Create(ctx context.Context, payload *model.FolderCreatePayloadScheme) (*model.FolderScheme, *model.ResponseScheme, error) {
	return f.internalClient.Create(ctx, payload)
}

// Update updates a folder by id.
//
// PUT /wiki/api/v2/folders/{id}
//
// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
func (f *FolderService) Update(ctx context.Context, folderID string, payload *model.FolderUpdatePayloadScheme) (*model.FolderScheme, *model.ResponseScheme, error) {
	return f.internalClient.Update(ctx, folderID, payload)
}

// Delete deletes a folder by id.
//
// DELETE /wiki/api/v2/folders/{id}
//
// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
func (f *FolderService) Delete(ctx context.Context, folderID string) (*model.ResponseScheme, error) {
	return f.internalClient.Delete(ctx, folderID)
}

type internalFolderImpl struct {
	c service.Connector
}

func (i *internalFolderImpl) Gets(ctx context.Context, options *model.FolderOptionsScheme, cursor string, limit int) (*model.FolderChunkScheme, *model.ResponseScheme, error) {

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	if options != nil {

		if options.Sort != "" {
			query.Add("sort", options.Sort)
		}

		if options.ParentID != "" {
			query.Add("parent-id", options.ParentID)
		}

		if len(options.SpaceIDs) > 0 {

			var spaceIDs = make([]string, 0, len(options.SpaceIDs))
			for _, spaceIDAsInt := range options.SpaceIDs {
				spaceIDs = append(spaceIDs, strconv.Itoa(spaceIDAsInt))
			}

			query.Add("space-id", strings.Join(spaceIDs, ","))
		}
	}

	endpoint := fmt.Sprintf("wiki/api/v2/folders?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	chunk := new(model.FolderChunkScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	return chunk, response, nil
}

func (i *internalFolderImpl) Get(ctx context.Context, folderID string) (*model.FolderScheme, *model.ResponseScheme, error) {

	if folderID == "" {
		return nil, nil, fmt.Errorf("confluence: %w", model.ErrNoFolderID)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/folders/%v", folderID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	folder := new(model.FolderScheme)
	response, err := i.c.Call(request, folder)
	if err != nil {
		return nil, response, err
	}

	return folder, response, nil
}

func (i *internalFolderImpl) GetsBySpace(ctx context.Context, spaceID int, cursor string, limit int) (*model.FolderChunkScheme, *model.ResponseScheme, error) {

	if spaceID == 0 {
		return nil, nil, fmt.Errorf("confluence: %w", model.ErrNoSpaceID)
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/spaces/%v/folders?%v", spaceID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	chunk := new(model.FolderChunkScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	return chunk, response, nil
}

func (i *internalFolderImpl) GetsByParent(ctx context.Context, parentID string, cursor string, limit int) (*model.FolderChunkScheme, *model.ResponseScheme, error) {

	if parentID == "" {
		return nil, nil, fmt.Errorf("confluence: %w", model.ErrNoFolderID)
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/folders/%v/children?%v", parentID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	chunk := new(model.FolderChunkScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	return chunk, response, nil
}

func (i *internalFolderImpl) Create(ctx context.Context, payload *model.FolderCreatePayloadScheme) (*model.FolderScheme, *model.ResponseScheme, error) {

	endpoint := "wiki/api/v2/folders"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	folder := new(model.FolderScheme)
	response, err := i.c.Call(request, folder)
	if err != nil {
		return nil, response, err
	}

	return folder, response, nil
}

func (i *internalFolderImpl) Update(ctx context.Context, folderID string, payload *model.FolderUpdatePayloadScheme) (*model.FolderScheme, *model.ResponseScheme, error) {

	if folderID == "" {
		return nil, nil, fmt.Errorf("confluence: %w", model.ErrNoFolderID)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/folders/%v", folderID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	folder := new(model.FolderScheme)
	response, err := i.c.Call(request, folder)
	if err != nil {
		return nil, response, err
	}

	return folder, response, nil
}

func (i *internalFolderImpl) Delete(ctx context.Context, folderID string) (*model.ResponseScheme, error) {

	if folderID == "" {
		return nil, fmt.Errorf("confluence: %w", model.ErrNoFolderID)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/folders/%v", folderID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

