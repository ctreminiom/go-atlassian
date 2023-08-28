package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strings"
)

func NewMySelfService(client service.Connector, version string) (*MySelfService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &MySelfService{
		internalClient: &internalMySelfImpl{c: client, version: version},
	}, nil
}

type MySelfService struct {
	internalClient jira.MySelfConnector
}

// Details returns details for the current user.
//
// GET /rest/api/{2-3}/myself
//
// https://docs.go-atlassian.io/jira-software-cloud/myself#get-current-user
func (m *MySelfService) Details(ctx context.Context, expand []string) (*model.UserScheme, *model.ResponseScheme, error) {
	return m.internalClient.Details(ctx, expand)
}

type internalMySelfImpl struct {
	c       service.Connector
	version string
}

func (i *internalMySelfImpl) Details(ctx context.Context, expand []string) (*model.UserScheme, *model.ResponseScheme, error) {

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/myself", i.version))

	if expand != nil {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	my := new(model.UserScheme)
	response, err := i.c.Call(request, my)
	if err != nil {
		return nil, response, err
	}

	return my, response, nil
}
