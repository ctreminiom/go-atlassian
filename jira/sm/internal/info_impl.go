package internal

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/sm"
	"net/http"
)

func NewInfoService(client service.Client, version string) (*InfoService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &InfoService{
		internalClient: &internalInfoImpl{c: client, version: version},
	}, nil
}

type InfoService struct {
	internalClient sm.InfoConnector
}

// Get retrieves information about the Jira Service Management instance such as software version, builds, and related links.
//
// GET /rest/servicedeskapi/info
//
// https://docs.go-atlassian.io/jira-service-management-cloud/info#get-info
func (i *InfoService) Get(ctx context.Context) (*model.InfoScheme, *model.ResponseScheme, error) {
	return i.internalClient.Get(ctx)
}

type internalInfoImpl struct {
	c       service.Client
	version string
}

func (i *internalInfoImpl) Get(ctx context.Context) (*model.InfoScheme, *model.ResponseScheme, error) {

	endpoint := "rest/servicedeskapi/info"

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	info := new(model.InfoScheme)
	response, err := i.c.Call(request, info)
	if err != nil {
		return nil, response, err
	}

	return info, response, nil
}
