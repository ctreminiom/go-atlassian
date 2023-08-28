package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewIssueFieldTrashService(client service.Connector, version string) (*IssueFieldTrashService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssueFieldTrashService{
		internalClient: &internalFieldTrashServiceImpl{c: client, version: version},
	}, nil
}

type IssueFieldTrashService struct {
	internalClient jira.FieldTrashConnector
}

// Search returns a paginated list of fields in the trash.
//
// The list may be restricted to field whose field name or description partially match a string.
//
// Only custom fields can be queried, type must be set to custom.
//
// GET /rest/api/{2-3}/field/search/trashed
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/trash#search-fields-in-trash
func (i *IssueFieldTrashService) Search(ctx context.Context, options *model.FieldSearchOptionsScheme, startAt, maxResults int) (*model.FieldSearchPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Search(ctx, options, startAt, maxResults)
}

// Move moves a custom field to trash.
//
// See Edit or delete a custom field for more information on trashing and deleting custom fields.
//
// POST /rest/api/{2-3}/field/{id}/trash
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/trash#move-field-to-trash
func (i *IssueFieldTrashService) Move(ctx context.Context, id string) (*model.ResponseScheme, error) {
	return i.internalClient.Move(ctx, id)
}

// Restore restores a custom field from trash.
//
// See Edit or delete a custom field for more information on trashing and deleting custom fields.
//
// POST /rest/api/{2-3}/field/{id}/restore
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/trash#move-field-to-trash
func (i *IssueFieldTrashService) Restore(ctx context.Context, id string) (*model.ResponseScheme, error) {
	return i.internalClient.Restore(ctx, id)
}

type internalFieldTrashServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalFieldTrashServiceImpl) Search(ctx context.Context, options *model.FieldSearchOptionsScheme, startAt, maxResults int) (*model.FieldSearchPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.IDs) != 0 {
			params.Add("id", strings.Join(options.IDs, ","))
		}

		if len(options.OrderBy) != 0 {
			params.Add("orderBy", options.OrderBy)
		}

		if len(options.Query) != 0 {
			params.Add("query", options.Query)
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/search/trashed?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.FieldSearchPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalFieldTrashServiceImpl) Move(ctx context.Context, id string) (*model.ResponseScheme, error) {

	if id == "" {
		return nil, model.ErrNoFieldIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/trash", i.version, id)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalFieldTrashServiceImpl) Restore(ctx context.Context, id string) (*model.ResponseScheme, error) {

	if id == "" {
		return nil, model.ErrNoFieldIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/restore", i.version, id)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
