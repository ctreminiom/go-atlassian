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

func NewIssueFieldConfigurationService(client service.Connector, version string, item *IssueFieldConfigItemService,
	scheme *IssueFieldConfigSchemeService) (*IssueFieldConfigService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssueFieldConfigService{
		internalClient: &internalIssueFieldConfigServiceImpl{c: client, version: version},
		Item:           item,
		Scheme:         scheme,
	}, nil
}

type IssueFieldConfigService struct {
	internalClient jira.FieldConfigConnector
	Item           *IssueFieldConfigItemService
	Scheme         *IssueFieldConfigSchemeService
}

// Gets Returns a paginated list of all field configurations.
//
// GET /rest/api/{2-3}/fieldconfiguration
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configurations
func (i *IssueFieldConfigService) Gets(ctx context.Context, ids []int, isDefault bool, startAt, maxResults int) (*model.FieldConfigurationPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, ids, isDefault, startAt, maxResults)
}

// Create creates a field configuration. The field configuration is created with the same field properties as the
// default configuration, with all the fields being optional.
//
// This operation can only create configurations for use in company-managed (classic) projects.
//
// POST /rest/api/{2-3}/fieldconfiguration
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#create-field-configuration
func (i *IssueFieldConfigService) Create(ctx context.Context, name, description string) (*model.FieldConfigurationScheme, *model.ResponseScheme, error) {
	return i.internalClient.Create(ctx, name, description)
}

// Update updates a field configuration. The name and the description provided in the request override the existing values.
//
// This operation can only update configurations used in company-managed (classic) projects.
//
// PUT /rest/api/{2-3}/fieldconfiguration/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#update-field-configuration
func (i *IssueFieldConfigService) Update(ctx context.Context, id int, name, description string) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, id, name, description)
}

// Delete deletes a field configuration.
//
// This operation can only delete configurations used in company-managed (classic) projects.
//
// DELETE /rest/api/{2-3}/fieldconfiguration/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#delete-field-configuration
func (i *IssueFieldConfigService) Delete(ctx context.Context, id int) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, id)
}

type internalIssueFieldConfigServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueFieldConfigServiceImpl) Gets(ctx context.Context, ids []int, isDefault bool, startAt, maxResults int) (*model.FieldConfigurationPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("isDefault", fmt.Sprintf("%v", isDefault))

	for _, id := range ids {
		params.Add("id", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.FieldConfigurationPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalIssueFieldConfigServiceImpl) Create(ctx context.Context, name, description string) (*model.FieldConfigurationScheme, *model.ResponseScheme, error) {

	if name == "" {
		return nil, nil, model.ErrNoFieldConfigurationNameError
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration", i.version)

	payload := map[string]interface{}{"name": name}

	if description != "" {
		payload["description"] = description
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	issueConfig := new(model.FieldConfigurationScheme)
	response, err := i.c.Call(request, issueConfig)
	if err != nil {
		return nil, response, err
	}

	return issueConfig, response, nil
}

func (i *internalIssueFieldConfigServiceImpl) Update(ctx context.Context, id int, name, description string) (*model.ResponseScheme, error) {

	if id == 0 {
		return nil, model.ErrNoFieldConfigurationIDError
	}

	if name == "" {
		return nil, model.ErrNoFieldConfigurationNameError
	}

	payload := map[string]interface{}{"name": name}

	if description != "" {
		payload["description"] = description
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration/%v", i.version, id)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldConfigServiceImpl) Delete(ctx context.Context, id int) (*model.ResponseScheme, error) {

	if id == 0 {
		return nil, model.ErrNoFieldConfigurationIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration/%v", i.version, id)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
