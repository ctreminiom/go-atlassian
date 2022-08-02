package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewPriorityService(client service.Client, version string) (*PriorityService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &PriorityService{
		internalClient: &internalPriorityImpl{c: client, version: version},
	}, nil
}

type PriorityService struct {
	internalClient jira.PriorityConnector
}

// Gets returns the list of all issue priorities.
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/priorities#get-priorities
func (p *PriorityService) Gets(ctx context.Context) ([]*model.PriorityScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx)
}

// Get returns an issue priority.
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/priorities#get-priority
func (p *PriorityService) Get(ctx context.Context, priorityId string) (*model.PriorityScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, priorityId)
}

type internalPriorityImpl struct {
	c       service.Client
	version string
}

func (i *internalPriorityImpl) Gets(ctx context.Context) ([]*model.PriorityScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/priority", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalPriorityImpl) Get(ctx context.Context, priorityId string) (*model.PriorityScheme, *model.ResponseScheme, error) {

	if priorityId == "" {
		return nil, nil, model.ErrNoPriorityIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/priority/%v", i.version, priorityId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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
