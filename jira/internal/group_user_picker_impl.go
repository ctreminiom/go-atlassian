package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewGroupUserPickerService creates a new GroupUserPickerService instance.
func NewGroupUserPickerService(client service.Connector, version string) (*GroupUserPickerService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &GroupUserPickerService{
		internalClient: &internalGroupUserPickerServiceImpl{c: client, version: version},
	}, nil
}

// GroupUserPickerService handles communication with the GroupUserPicker related methods of the Jira API.
type GroupUserPickerService struct {
	internalClient jira.GroupUserPickerConnector
}

// Find returns a list of users and groups matching a string.
//
// GET /rest/api/{2-3}/groupuserpicker
//
// https://docs.go-atlassian.io/jira-software-cloud/groupuserpicker#find-users-and-groups
func (g *GroupUserPickerService) Find(ctx context.Context, options *model.GroupUserPickerFindOptionScheme) (*model.GroupUserPickerFindScheme, *model.ResponseScheme, error) {
	return g.internalClient.Find(ctx, options)
}

type internalGroupUserPickerServiceImpl struct {
	c       service.Connector
	version string
}

// Find returns a list of users and groups matching a string.
func (i internalGroupUserPickerServiceImpl) Find(ctx context.Context, options *model.GroupUserPickerFindOptionScheme) (*model.GroupUserPickerFindScheme, *model.ResponseScheme, error) {

	if options == nil || options.Query == "" {
		return nil, nil, model.ErrNoQuery
	}

	endpoint := fmt.Sprintf("rest/api/%v/groupuserpicker", i.version)

	q, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	endpoint += "?" + q.Encode()

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	find := new(model.GroupUserPickerFindScheme)
	response, err := i.c.Call(request, find)
	if err != nil {
		return nil, response, err
	}

	return find, response, nil
}
