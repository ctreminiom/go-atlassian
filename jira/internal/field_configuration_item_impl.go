package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
	"net/http"
	"net/url"
	"strconv"
)

// NewIssueFieldConfigurationItemService creates a new instance of IssueFieldConfigItemService.
// It takes a service.Connector and a version string as input.
// Returns a pointer to IssueFieldConfigItemService and an error if the version is not provided.
func NewIssueFieldConfigurationItemService(client service.Connector, version string) (*IssueFieldConfigItemService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssueFieldConfigItemService{
		internalClient: &internalIssueFieldConfigItemServiceImpl{c: client, version: version},
	}, nil
}

// IssueFieldConfigItemService provides methods to manage field configuration items in Jira Service Management.
type IssueFieldConfigItemService struct {
	// internalClient is the connector interface for field configuration item operations.
	internalClient jira.FieldConfigItemConnector
}

// Gets Returns a paginated list of all fields for a configuration.
//
// GET /rest/api/{2-3}/fieldconfiguration/{id}/fields
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/items#get-field-configuration-items
func (i *IssueFieldConfigItemService) Gets(ctx context.Context, id, startAt, maxResults int) (*model.FieldConfigurationItemPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, id, startAt, maxResults)
}

// Update updates fields in a field configuration. The properties of the field configuration fields provided
// override the existing values.
//
// 1. This operation can only update field configurations used in company-managed (classic) projects.
//
// PUT /rest/api/{2-3}/fieldconfiguration/{id}/fields
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/items#update-field-configuration-items
func (i *IssueFieldConfigItemService) Update(ctx context.Context, id int, payload *model.UpdateFieldConfigurationItemPayloadScheme) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, id, payload)
}

type internalIssueFieldConfigItemServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueFieldConfigItemServiceImpl) Gets(ctx context.Context, id, startAt, maxResults int) (*model.FieldConfigurationItemPageScheme, *model.ResponseScheme, error) {

	if id == 0 {
		return nil, nil, model.ErrNoFieldConfigurationID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration/%v/fields?%v", i.version, id, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.FieldConfigurationItemPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalIssueFieldConfigItemServiceImpl) Update(ctx context.Context, id int, payload *model.UpdateFieldConfigurationItemPayloadScheme) (*model.ResponseScheme, error) {

	if id == 0 {
		return nil, model.ErrNoFieldConfigurationID
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration/%v/fields", i.version, id)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
