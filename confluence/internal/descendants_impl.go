package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
)

// NewDescendantsService creates a new instance of DescendantsService.
// It takes a service.Connector as input and returns a pointer to DescendantsService.
func NewDescendantsService(client service.Connector) *DescendantsService {
	return &DescendantsService{internalClient: &internalDescendantsImpl{c: client}}
}

// DescendantsService provides methods to interact with descendants operations in Confluence.
type DescendantsService struct {
	// internalClient is the connector interface for descendants operations.
	internalClient confluence.DescendantsConnector
}

// Get descendants of a whiteboard.
//
// GET /wiki/api/v2/whiteboards/{id}/descendants
//
// https://docs.go-atlassian.io/confluence-cloud/v2/descendants#get-descendants-of-a-whiteboard
func (p *DescendantsService) GetForWhiteboard(
	ctx context.Context,
	whiteboardID int,
	limit int,
	depth int,
	cursor string,
) (*model.DescendantsScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetForWhiteboard(ctx, whiteboardID, limit, depth, cursor)
}

// Get descendants of a database.
//
// GET /wiki/api/v2/databases/{id}/descendants
//
// https://docs.go-atlassian.io/confluence-cloud/v2/descendants#get-descendants-of-a-database
func (p *DescendantsService) GetForDatabase(
	ctx context.Context,
	databaseID int,
	limit int,
	depth int,
	cursor string,
) (*model.DescendantsScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetForDatabase(ctx, databaseID, limit, depth, cursor)
}

// Get descendants of a smart link.
//
// GET /wiki/api/v2/embeds/{id}/descendants
//
// https://docs.go-atlassian.io/confluence-cloud/v2/descendants#get-descendants-of-a-smart-link
func (p *DescendantsService) GetForSmartLink(
	ctx context.Context,
	embedID int,
	limit int,
	depth int,
	cursor string,
) (*model.DescendantsScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetForSmartLink(ctx, embedID, limit, depth, cursor)
}

// Get descendants of a folder.
//
// GET /wiki/api/v2/folders/{id}/descendants
//
// https://docs.go-atlassian.io/confluence-cloud/v2/descendants#get-descendants-of-a-folder
func (p *DescendantsService) GetForFolder(
	ctx context.Context,
	folderID int,
	limit int,
	depth int,
	cursor string,
) (*model.DescendantsScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetForFolder(ctx, folderID, limit, depth, cursor)
}

// Get descendants of a page.
//
// GET /wiki/api/v2/pages/{id}/descendants
//
// https://docs.go-atlassian.io/confluence-cloud/v2/descendants#get-descendants-of-a-page
func (p *DescendantsService) GetForPage(
	ctx context.Context,
	pageID int,
	limit int,
	depth int,
	cursor string,
) (*model.DescendantsScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetForPage(ctx, pageID, limit, depth, cursor)
}

type internalDescendantsImpl struct {
	c service.Connector
}

func (i *internalDescendantsImpl) GetForWhiteboard(
	ctx context.Context,
	whiteboardID int,
	limit int,
	depth int,
	cursor string,
) (*model.DescendantsScheme, *model.ResponseScheme, error) {

	if whiteboardID == 0 {
		return nil, nil, model.ErrNoWhiteboardID
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	query.Add("depth", strconv.Itoa(depth))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/whiteboards/%v/descendants?%v", whiteboardID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	descendants := new(model.DescendantsScheme)
	response, err := i.c.Call(request, descendants)
	if err != nil {
		return nil, response, err
	}

	return descendants, response, nil
}

func (i *internalDescendantsImpl) GetForDatabase(
	ctx context.Context,
	databaseID int,
	limit int,
	depth int,
	cursor string,
) (*model.DescendantsScheme, *model.ResponseScheme, error) {

	if databaseID == 0 {
		return nil, nil, model.ErrNoDatabaseID
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	query.Add("depth", strconv.Itoa(depth))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/databases/%v/descendants?%v", databaseID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	descendants := new(model.DescendantsScheme)
	response, err := i.c.Call(request, descendants)
	if err != nil {
		return nil, response, err
	}

	return descendants, response, nil
}

func (i *internalDescendantsImpl) GetForSmartLink(
	ctx context.Context,
	embedID int,
	limit int,
	depth int,
	cursor string,
) (*model.DescendantsScheme, *model.ResponseScheme, error) {

	if embedID == 0 {
		return nil, nil, model.ErrNoEmbedID
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	query.Add("depth", strconv.Itoa(depth))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/embeds/%v/descendants?%v", embedID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	descendants := new(model.DescendantsScheme)
	response, err := i.c.Call(request, descendants)
	if err != nil {
		return nil, response, err
	}

	return descendants, response, nil
}

func (i *internalDescendantsImpl) GetForFolder(
	ctx context.Context,
	folderID int,
	limit int,
	depth int,
	cursor string,
) (*model.DescendantsScheme, *model.ResponseScheme, error) {

	if folderID == 0 {
		return nil, nil, model.ErrNoFolderID
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	query.Add("depth", strconv.Itoa(depth))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/folders/%v/descendants?%v", folderID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	descendants := new(model.DescendantsScheme)
	response, err := i.c.Call(request, descendants)
	if err != nil {
		return nil, response, err
	}

	return descendants, response, nil
}

func (i *internalDescendantsImpl) GetForPage(
	ctx context.Context,
	pageID int,
	limit int,
	depth int,
	cursor string,
) (*model.DescendantsScheme, *model.ResponseScheme, error) {

	if pageID == 0 {
		return nil, nil, model.ErrNoPageID
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	query.Add("depth", strconv.Itoa(depth))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/pages/%v/descendants?%v", pageID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	descendants := new(model.DescendantsScheme)
	response, err := i.c.Call(request, descendants)
	if err != nil {
		return nil, response, err
	}

	return descendants, response, nil
}
