package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewPriorityService creates a new instance of PriorityService.
func NewPriorityService(client service.Connector, version string) (*PriorityService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &PriorityService{
		internalClient: &internalPriorityImpl{c: client, version: version},
	}, nil
}

// PriorityService provides methods to manage issue priorities in Jira Service Management.
type PriorityService struct {
	// internalClient is the connector interface for priority operations.
	internalClient jira.PriorityConnector
}

// Gets returns the list of all issue priorities.
//
// GET /rest/api/{2-3}/priority
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/priorities#get-priorities
func (p *PriorityService) Gets(ctx context.Context) ([]*model.PriorityScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx)
}

// Get returns an issue priority.
//
// GET /rest/api/{2-3}/priority/{priorityID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/priorities#get-priority
func (p *PriorityService) Get(ctx context.Context, priorityID string) (*model.PriorityScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, priorityID)
}

type internalPriorityImpl struct {
	c       service.Connector
	version string
}

func (i *internalPriorityImpl) Gets(ctx context.Context) ([]*model.PriorityScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/priority", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var priorities []*model.PriorityScheme
	response, err := i.c.Call(request, &priorities)
	if err != nil {
		return nil, response, err
	}

	return priorities, response, nil
}

func (i *internalPriorityImpl) Get(ctx context.Context, priorityID string) (*model.PriorityScheme, *model.ResponseScheme, error) {

	if priorityID == "" {
		return nil, nil, model.ErrNoPriorityID
	}

	endpoint := fmt.Sprintf("rest/api/%v/priority/%v", i.version, priorityID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	priority := new(model.PriorityScheme)
	response, err := i.c.Call(request, priority)
	if err != nil {
		return nil, response, err
	}

	return priority, response, nil
}
