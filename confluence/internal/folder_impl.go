package internal

import (
	"context"
	"fmt"
	"net/http"

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

// Create creates a folder in the space.
//
// Folders are created as published by default unless specified as a draft in the status field.
//
// If creating a published folder, the title must be specified.
//
// POST /wiki/api/v2/folders
//
// https://docs.go-atlassian.io/confluence-cloud/v2/folder#create-folder
func (p *FolderService) Create(ctx context.Context, payload *model.FolderCreatePayloadScheme) (*model.FolderScheme, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, payload)
}

// Get returns a specific folder.
//
// GET /wiki/api/v2/folders/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/folder#get-folder-by-id
func (p *FolderService) Get(ctx context.Context, folderID int) (*model.FolderScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, folderID)
}

// Delete deletes a folder by id.
//
// DELETE /wiki/api/v2/folders/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/folder#delete-folder
func (p *FolderService) Delete(ctx context.Context, folderID int) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, folderID)
}

type internalFolderImpl struct {
	c service.Connector
}

func (i *internalFolderImpl) Create(ctx context.Context, payload *model.FolderCreatePayloadScheme) (*model.FolderScheme, *model.ResponseScheme, error) {

	endpoint := "wiki/api/v2/folders"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, fmt.Errorf("%w, payload: %+v", err, payload)
	}

	folder := new(model.FolderScheme)
	response, err := i.c.Call(request, folder)
	if err != nil {
		return nil, response, fmt.Errorf("%w, payload: %+v", err, payload)
	}

	return folder, response, nil
}

func (i *internalFolderImpl) Get(ctx context.Context, folderID int) (*model.FolderScheme, *model.ResponseScheme, error) {

	if folderID == 0 {
		return nil, nil, model.ErrNoFolderID
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

func (i *internalFolderImpl) Delete(ctx context.Context, folderID int) (*model.ResponseScheme, error) {

	if folderID == 0 {
		return nil, model.ErrNoFolderID
	}

	endpoint := fmt.Sprintf("wiki/api/v2/folders/%v", folderID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
