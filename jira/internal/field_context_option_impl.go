package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewIssueFieldContextOptionService creates a new instance of IssueFieldContextOptionService.
// It takes a service.Connector and a version string as input.
// Returns a pointer to IssueFieldContextOptionService and an error if the version is not provided.
func NewIssueFieldContextOptionService(client service.Connector, version string) (*IssueFieldContextOptionService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssueFieldContextOptionService{
		internalClient: &internalIssueFieldContextOptionServiceImpl{c: client, version: version},
	}, nil
}

// IssueFieldContextOptionService provides methods to manage field context options in Jira Service Management.
type IssueFieldContextOptionService struct {
	// internalClient is the connector interface for field context option operations.
	internalClient jira.FieldContextOptionConnector
}

// Gets returns a paginated list of all custom field option for a context.
//
// Options are returned first then cascading options, in the order they display in Jira.
//
// GET /rest/api/{2-3}/field/{fieldID}/context/{contextID}/option
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#get-custom-field-options
func (i *IssueFieldContextOptionService) Gets(ctx context.Context, fieldID string, contextID int, options *model.FieldOptionContextParams, startAt, maxResults int) (*model.CustomFieldContextOptionPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, fieldID, contextID, options, startAt, maxResults)
}

// Create creates options and, where the custom select field is of the type Select List (cascading), cascading options for a custom select field.
//
// 1. The options are added to a context of the field.
//
// 2. The maximum number of options that can be created per request is 1000 and each field can have a maximum of 10000 options.
//
// POST /rest/api/{2-3}/field/{fieldID}/context/{contextID}/option
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#create-custom-field-options
func (i *IssueFieldContextOptionService) Create(ctx context.Context, fieldID string, contextID int, payload *model.FieldContextOptionListScheme) (*model.FieldContextOptionListScheme, *model.ResponseScheme, error) {
	return i.internalClient.Create(ctx, fieldID, contextID, payload)
}

// Update updates the options of a custom field.
//
// 1. If any of the options are not found, no options are updated.
//
// 2. Options where the values in the request match the current values aren't updated and aren't reported in the response.
//
// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}/option
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#update-custom-field-options
func (i *IssueFieldContextOptionService) Update(ctx context.Context, fieldID string, contextID int, payload *model.FieldContextOptionListScheme) (*model.FieldContextOptionListScheme, *model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, fieldID, contextID, payload)
}

// Delete deletes a custom field option.
//
// 1. Options with cascading options cannot be deleted without deleting the cascading options first.
//
// DELETE /rest/api/{2-3}/field/{fieldID}/context/{contextID}/option/{optionID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#delete-custom-field-options
func (i *IssueFieldContextOptionService) Delete(ctx context.Context, fieldID string, contextID, optionID int) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, fieldID, contextID, optionID)
}

// Order changes the order of custom field options or cascading options in a context.
//
// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}/option/move
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#reorder-custom-field-options
func (i *IssueFieldContextOptionService) Order(ctx context.Context, fieldID string, contextID int, payload *model.OrderFieldOptionPayloadScheme) (*model.ResponseScheme, error) {
	return i.internalClient.Order(ctx, fieldID, contextID, payload)
}

type internalIssueFieldContextOptionServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueFieldContextOptionServiceImpl) Gets(ctx context.Context, fieldID string, contextID int, options *model.FieldOptionContextParams, startAt, maxResults int) (*model.CustomFieldContextOptionPageScheme, *model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, nil, model.ErrNoFieldID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {
		params.Add("onlyOptions", fmt.Sprintf("%v", options.OnlyOptions))

		if options.OptionID != 0 {
			params.Add("optionId", strconv.Itoa(options.OptionID))
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/option?%v", i.version, fieldID, contextID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	pagination := new(model.CustomFieldContextOptionPageScheme)
	response, err := i.c.Call(request, pagination)
	if err != nil {
		return nil, response, err
	}

	return pagination, response, nil
}

func (i *internalIssueFieldContextOptionServiceImpl) Create(ctx context.Context, fieldID string, contextID int, payload *model.FieldContextOptionListScheme) (*model.FieldContextOptionListScheme, *model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, nil, model.ErrNoFieldID
	}

	if contextID == 0 {
		return nil, nil, model.ErrNoFieldContextID
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/option", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	options := new(model.FieldContextOptionListScheme)
	response, err := i.c.Call(request, options)
	if err != nil {
		return nil, response, err
	}

	return options, response, nil
}

func (i *internalIssueFieldContextOptionServiceImpl) Update(ctx context.Context, fieldID string, contextID int, payload *model.FieldContextOptionListScheme) (*model.FieldContextOptionListScheme, *model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, nil, model.ErrNoFieldID
	}

	if contextID == 0 {
		return nil, nil, model.ErrNoFieldContextID
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/option", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	options := new(model.FieldContextOptionListScheme)
	response, err := i.c.Call(request, options)
	if err != nil {
		return nil, response, err
	}

	return options, response, nil
}

func (i *internalIssueFieldContextOptionServiceImpl) Delete(ctx context.Context, fieldID string, contextID, optionID int) (*model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	if contextID == 0 {
		return nil, model.ErrNoFieldContextID
	}

	if optionID == 0 {
		return nil, model.ErrNoContextOptionID
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/option/%v", i.version, fieldID, contextID, optionID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextOptionServiceImpl) Order(ctx context.Context, fieldID string, contextID int, payload *model.OrderFieldOptionPayloadScheme) (*model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	if contextID == 0 {
		return nil, model.ErrNoFieldContextID
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/option/move", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
