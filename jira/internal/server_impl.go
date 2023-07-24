package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewServerService(client service.Connector, version string) (*ServerService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ServerService{
		internalClient: &internalServerServiceImpl{c: client, version: version},
	}, nil
}

type ServerService struct {
	internalClient jira.ServerConnector
}

// Info returns information about the Jira instance
//
// GET /rest/api/{2-3}/serverInfo
//
// https://docs.go-atlassian.io/jira-software-cloud/server#get-jira-instance-info
func (s *ServerService) Info(ctx context.Context) (*model.ServerInformationScheme, *model.ResponseScheme, error) {
	return s.internalClient.Info(ctx)
}

type internalServerServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalServerServiceImpl) Info(ctx context.Context) (*model.ServerInformationScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/serverInfo", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	server := new(model.ServerInformationScheme)
	response, err := i.c.Call(request, server)
	if err != nil {
		return nil, response, err
	}

	return server, response, nil
}
