package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewWorkflowValidatorService(client service.Connector, version string) (*WorkflowValidatorService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &WorkflowValidatorService{
		internalClient: &internalWorkflowValidatorImpl{c: client, version: version},
	}, nil
}

type WorkflowValidatorService struct {
	internalClient jira.WorkflowValidatorConnector
}

// Creation validates the payload for bulk create workflows.
//
// POST /rest/api/{2-3}/workflows/create/validation
func (w *WorkflowValidatorService) Creation(ctx context.Context, payload *model.WorkflowCreateValidatorPayloadScheme) (*model.WorkflowValidationErrorListScheme, *model.ResponseScheme, error) {
	return w.internalClient.Creation(ctx, payload)
}

// Modification validates the payload for bulk update workflows.
//
// POST /rest/api/{2-3}/workflows/update/validation
func (w *WorkflowValidatorService) Modification(ctx context.Context, payload *model.WorkflowUpdateValidatorPayloadScheme) (*model.WorkflowValidationErrorListScheme, *model.ResponseScheme, error) {
	return w.internalClient.Modification(ctx, payload)
}

type internalWorkflowValidatorImpl struct {
	c       service.Connector
	version string
}

func (i *internalWorkflowValidatorImpl) Creation(ctx context.Context, payload *model.WorkflowCreateValidatorPayloadScheme) (*model.WorkflowValidationErrorListScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/workflows/create/validation", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	errorList := new(model.WorkflowValidationErrorListScheme)
	response, err := i.c.Call(request, errorList)
	if err != nil {
		return nil, response, err
	}

	return errorList, response, nil
}

func (i *internalWorkflowValidatorImpl) Modification(ctx context.Context, payload *model.WorkflowUpdateValidatorPayloadScheme) (*model.WorkflowValidationErrorListScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/workflows/update/validation", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	errorList := new(model.WorkflowValidationErrorListScheme)
	response, err := i.c.Call(request, errorList)
	if err != nil {
		return nil, response, err
	}

	return errorList, response, nil
}
