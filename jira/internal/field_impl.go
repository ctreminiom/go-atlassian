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
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewIssueFieldService creates a new instance of IssueFieldService.
// It takes a service.Connector, a version string, an IssueFieldConfigService, an IssueFieldContextService, and an IssueFieldTrashService as input.
// Returns a pointer to IssueFieldService and an error if the version is not provided.
func NewIssueFieldService(client service.Connector, version string, configuration *IssueFieldConfigService, context *IssueFieldContextService,
	trash *IssueFieldTrashService) (*IssueFieldService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssueFieldService{
		internalClient: &internalIssueFieldServiceImpl{c: client, version: version},
		Configuration:  configuration,
		Context:        context,
		Trash:          trash,
	}, nil
}

// IssueFieldService provides methods to manage issue fields in Jira Service Management.
type IssueFieldService struct {
	// internalClient is the connector interface for issue field operations.
	internalClient jira.FieldConnector
	// Configuration is the service for managing field configurations.
	Configuration *IssueFieldConfigService
	// Context is the service for managing field contexts.
	Context *IssueFieldContextService
	// Trash is the service for managing trashed fields.
	Trash *IssueFieldTrashService
}

// Gets returns system and custom issue fields according to the following rules:
//
// 1. Fields that cannot be added to the issue navigator are always returned.
//
// 2. Fields that cannot be placed on an issue screen are always returned.
//
// 3. Fields that depend on global Jira settings are only returned if the setting is enabled.
// That is, timetracking fields, subtasks, votes, and watches.
//
// 4. For all other fields, this operation only returns the fields that the user has permission to view
//
// GET /rest/api/{2-3}/field
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields#get-fields
func (i *IssueFieldService) Gets(ctx context.Context) ([]*model.IssueFieldScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx)
}

// Create creates a custom field.
//
// POST /rest/api/{2-3}/field
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields#create-custom-field
func (i *IssueFieldService) Create(ctx context.Context, payload *model.CustomFieldScheme) (*model.IssueFieldScheme, *model.ResponseScheme, error) {
	return i.internalClient.Create(ctx, payload)
}

// Search returns a paginated list of fields for Classic Jira projects.
//
// GET /rest/api/{2-3}/field/search
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields#get-fields-paginated
func (i *IssueFieldService) Search(ctx context.Context, options *model.FieldSearchOptionsScheme, startAt, maxResults int) (*model.FieldSearchPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Search(ctx, options, startAt, maxResults)
}

// Delete deletes a custom field. The custom field is deleted whether it is in the trash or not.
//
// See Edit or delete a custom field for more information on trashing and deleting custom fields.
//
// DELETE /rest/api/{2-3}/field/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields#delete-field
func (i *IssueFieldService) Delete(ctx context.Context, fieldID string) (*model.TaskScheme, *model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, fieldID)
}

type internalIssueFieldServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueFieldServiceImpl) Gets(ctx context.Context) ([]*model.IssueFieldScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/field", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var fields []*model.IssueFieldScheme
	response, err := i.c.Call(request, &fields)
	if err != nil {
		return nil, response, err
	}

	return fields, response, nil
}

func (i *internalIssueFieldServiceImpl) Create(ctx context.Context, payload *model.CustomFieldScheme) (*model.IssueFieldScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/field", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	field := new(model.IssueFieldScheme)
	response, err := i.c.Call(request, field)
	if err != nil {
		return nil, response, err
	}

	return field, response, nil
}

func (i *internalIssueFieldServiceImpl) Search(ctx context.Context, options *model.FieldSearchOptionsScheme, startAt, maxResults int) (*model.FieldSearchPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

		if len(options.Types) != 0 {
			params.Add("type", strings.Join(options.Types, ","))
		}

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

	endpoint := fmt.Sprintf("rest/api/%v/field/search?%v", i.version, params.Encode())

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

func (i *internalIssueFieldServiceImpl) Delete(ctx context.Context, fieldID string) (*model.TaskScheme, *model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, nil, model.ErrNoFieldID
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v", i.version, fieldID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	task := new(model.TaskScheme)
	response, err := i.c.Call(request, task)
	if err != nil {
		return nil, response, err
	}

	return task, response, nil
}
