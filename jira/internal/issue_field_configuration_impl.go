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

func NewIssueFieldConfigurationService(client service.Client, version string, item jira.FieldConfigItemConnector,
	scheme jira.FieldConfigSchemeConnector) (*IssueFieldConfigService, error) {

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
	Item           jira.FieldConfigItemConnector
	Scheme         jira.FieldConfigSchemeConnector
}

func (i *IssueFieldConfigService) Gets(ctx context.Context, ids []int, isDefault bool, startAt, maxResults int) (*model.FieldConfigurationPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, ids, isDefault, startAt, maxResults)
}

func (i *IssueFieldConfigService) Create(ctx context.Context, name, description string) (*model.FieldConfigurationScheme, *model.ResponseScheme, error) {
	return i.internalClient.Create(ctx, name, description)
}

func (i *IssueFieldConfigService) Update(ctx context.Context, id int, name, description string) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, id, name, description)
}

func (i *IssueFieldConfigService) Delete(ctx context.Context, id int) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, id)
}

type internalIssueFieldConfigServiceImpl struct {
	c       service.Client
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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration/%v", i.version, id)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
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

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
