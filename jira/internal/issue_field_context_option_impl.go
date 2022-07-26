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
)

func NewIssueFieldContextOptionService(client service.Client, version string) (*IssueFieldContextOptionService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssueFieldContextOptionService{
		internalClient: &internalIssueFieldContextOptionServiceImpl{c: client, version: version},
	}, nil
}

type IssueFieldContextOptionService struct {
	internalClient jira.FieldContextOptionConnector
}

// Gets returns a paginated list of all custom field option for a context.
//
// Options are returned first then cascading options, in the order they display in Jira.
// GET /rest/api/3/field/{fieldId}/context/{contextId}/option
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#get-custom-field-options
func (i *IssueFieldContextOptionService) Gets(ctx context.Context, fieldId string, contextId int, options *model.FieldOptionContextParams, startAt, maxResults int) (*model.CustomFieldContextOptionPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, fieldId, contextId, options, startAt, maxResults)
}

// Create creates options and, where the custom select field is of the type Select List (cascading), cascading options for a custom select field.
// The options are added to a context of the field.
//
// The maximum number of options that can be created per request is 1000 and each field can have a maximum of 10000 options.
// POST /rest/api/3/field/{fieldId}/context/{contextId}/option
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#create-custom-field-options
func (i *IssueFieldContextOptionService) Create(ctx context.Context, fieldId string, contextId int, payload *model.FieldContextOptionListScheme) (*model.FieldContextOptionListScheme, *model.ResponseScheme, error) {
	return i.internalClient.Create(ctx, fieldId, contextId, payload)
}

/*
Update updates the options of a custom field.

If any of the options are not found, no options are updated.
Options where the values in the request match the current values aren't updated and aren't reported in the response.
PUT /rest/api/3/field/{fieldId}/context/{contextId}/option
Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#update-custom-field-options
*/
func (i *IssueFieldContextOptionService) Update(ctx context.Context, fieldId string, contextId int, payload *model.FieldContextOptionListScheme) (*model.FieldContextOptionListScheme, *model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, fieldId, contextId, payload)
}

// Delete deletes a custom field option.
//
// Options with cascading options cannot be deleted without deleting the cascading options first.
// DELETE /rest/api/3/field/{fieldId}/context/{contextId}/option/{optionId}
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#delete-custom-field-options
func (i *IssueFieldContextOptionService) Delete(ctx context.Context, fieldId string, contextId, optionId int) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, fieldId, contextId, optionId)
}

// Order changes the order of custom field options or cascading options in a context.
//
// PUT /rest/api/3/field/{fieldId}/context/{contextId}/option/move
// Docs: TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
func (i *IssueFieldContextOptionService) Order(ctx context.Context, fieldId string, contextId int, payload *model.OrderFieldOptionPayloadScheme) (*model.ResponseScheme, error) {
	return i.internalClient.Order(ctx, fieldId, contextId, payload)
}

type internalIssueFieldContextOptionServiceImpl struct {
	c       service.Client
	version string
}

func (i *internalIssueFieldContextOptionServiceImpl) Gets(ctx context.Context, fieldId string, contextId int, options *model.FieldOptionContextParams, startAt, maxResults int) (*model.CustomFieldContextOptionPageScheme, *model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, nil, model.ErrNoFieldIDError
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

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/option?%v", i.version, fieldId, contextId, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalIssueFieldContextOptionServiceImpl) Create(ctx context.Context, fieldId string, contextId int, payload *model.FieldContextOptionListScheme) (*model.FieldContextOptionListScheme, *model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, nil, model.ErrNoFieldIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/option", i.version, fieldId, contextId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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

func (i *internalIssueFieldContextOptionServiceImpl) Update(ctx context.Context, fieldId string, contextId int, payload *model.FieldContextOptionListScheme) (*model.FieldContextOptionListScheme, *model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, nil, model.ErrNoFieldIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/option", i.version, fieldId, contextId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
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

func (i *internalIssueFieldContextOptionServiceImpl) Delete(ctx context.Context, fieldId string, contextId, optionId int) (*model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	if contextId == 0 {
		return nil, model.ErrNoFieldContextIDError
	}

	if optionId == 0 {
		return nil, model.ErrNoContextOptionIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/option/%v", i.version, fieldId, contextId, optionId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextOptionServiceImpl) Order(ctx context.Context, fieldId string, contextId int, payload *model.OrderFieldOptionPayloadScheme) (*model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	if contextId == 0 {
		return nil, model.ErrNoFieldContextIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/option/move", i.version, fieldId, contextId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
