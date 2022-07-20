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

func NewIssueFieldConfigurationItemService(client service.Client, version string) (*IssueFieldConfigItemService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssueFieldConfigItemService{
		internalClient: &internalIssueFieldConfigItemServiceImpl{c: client, version: version},
	}, nil
}

type IssueFieldConfigItemService struct {
	internalClient jira.FieldConfigurationItem
}

func (i *IssueFieldConfigItemService) Gets(ctx context.Context, id, startAt, maxResults int) (*model.FieldConfigurationItemPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, id, startAt, maxResults)
}

func (i *IssueFieldConfigItemService) Update(ctx context.Context, id int, payload *model.UpdateFieldConfigurationItemPayloadScheme) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, id, payload)
}

type internalIssueFieldConfigItemServiceImpl struct {
	c       service.Client
	version string
}

func (i *internalIssueFieldConfigItemServiceImpl) Gets(ctx context.Context, id, startAt, maxResults int) (*model.FieldConfigurationItemPageScheme, *model.ResponseScheme, error) {

	if id == 0 {
		return nil, nil, model.ErrNoFieldConfigurationIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration/%v/fields?%v", i.version, id, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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
		return nil, model.ErrNoFieldConfigurationIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration/%v/fields", i.version, id)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
