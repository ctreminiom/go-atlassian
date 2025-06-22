package internal

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/bitbucket"
)

// NewWorkspaceHookService creates a new WorkspaceHookService.
func NewWorkspaceHookService(client service.Connector) *WorkspaceHookService {
	return &WorkspaceHookService{
		internalClient: &internalWorkspaceHookServiceImpl{c: client},
	}
}

// WorkspaceHookService handles communication with the workspace hook related methods of the Bitbucket API.
type WorkspaceHookService struct {
	internalClient bitbucket.WorkspaceHookConnector
}

// Gets returns a paginated list of webhooks installed on this workspace.
//
// GET /2.0/workspaces/{workspace}/hooks
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#list-webhooks-for-a-workspace
func (w *WorkspaceHookService) Gets(ctx context.Context, workspace string) (*model.WebhookSubscriptionPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkspaceHookService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	return w.internalClient.Gets(ctx, workspace)
}

// Create creates a new webhook on the specified workspace.
//
// Workspace webhooks are fired for events from all repositories contained by that workspace.
//
// POST /2.0/workspaces/{workspace}/hooks
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#create-webhook-for-a-workspace
func (w *WorkspaceHookService) Create(ctx context.Context, workspace string, payload *model.WebhookSubscriptionPayloadScheme) (*model.WebhookSubscriptionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkspaceHookService).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create"))

	return w.internalClient.Create(ctx, workspace, payload)
}

// Get returns the webhook with the specified id installed on the given workspace.
//
// GET /2.0/workspaces/{workspace}/hooks/{uid}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#get-webhook-for-a-workspace
func (w *WorkspaceHookService) Get(ctx context.Context, workspace, webhookID string) (*model.WebhookSubscriptionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkspaceHookService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	return w.internalClient.Get(ctx, workspace, webhookID)
}

// Update updates the specified webhook subscription.
//
// PUT /2.0/workspaces/{workspace}/hooks/{uid}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#update-webhook-for-a-workspace
func (w *WorkspaceHookService) Update(ctx context.Context, workspace, webhookID string, payload *model.WebhookSubscriptionPayloadScheme) (*model.WebhookSubscriptionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkspaceHookService).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update"))

	return w.internalClient.Update(ctx, workspace, webhookID, payload)
}

// Delete deletes the specified webhook subscription from the given workspace.
//
// DELETE /2.0/workspaces/{workspace}/hooks/{uid}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#delete-webhook-for-a-workspace
func (w *WorkspaceHookService) Delete(ctx context.Context, workspace, webhookID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkspaceHookService).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete"))

	return w.internalClient.Delete(ctx, workspace, webhookID)
}

type internalWorkspaceHookServiceImpl struct {
	c service.Connector
}

func (i *internalWorkspaceHookServiceImpl) Gets(ctx context.Context, workspace string) (*model.WebhookSubscriptionPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkspaceHookServiceImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	if workspace == "" {

			return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/hooks", workspace)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	page := new(model.WebhookSubscriptionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalWorkspaceHookServiceImpl) Create(ctx context.Context, workspace string, payload *model.WebhookSubscriptionPayloadScheme) (*model.WebhookSubscriptionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkspaceHookServiceImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create"))

	if workspace == "" {

			return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/hooks", workspace)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	webhook := new(model.WebhookSubscriptionScheme)
	response, err := i.c.Call(request, webhook)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return webhook, response, nil
}

func (i *internalWorkspaceHookServiceImpl) Get(ctx context.Context, workspace, webhookID string) (*model.WebhookSubscriptionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkspaceHookServiceImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	if workspace == "" {

			return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
	}

	if webhookID == "" {

			return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWebhookID)
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/hooks/%v", workspace, webhookID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	webhook := new(model.WebhookSubscriptionScheme)
	response, err := i.c.Call(request, webhook)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return webhook, response, nil
}

func (i *internalWorkspaceHookServiceImpl) Update(ctx context.Context, workspace, webhookID string, payload *model.WebhookSubscriptionPayloadScheme) (*model.WebhookSubscriptionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkspaceHookServiceImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update"))

	if workspace == "" {

			return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
	}

	if webhookID == "" {

			return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWebhookID)
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/hooks/%v", workspace, webhookID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	webhook := new(model.WebhookSubscriptionScheme)
	response, err := i.c.Call(request, webhook)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return webhook, response, nil
}

func (i *internalWorkspaceHookServiceImpl) Delete(ctx context.Context, workspace, webhookID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkspaceHookServiceImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete"))

	if workspace == "" {
		return nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
	}

	if webhookID == "" {
		return nil, fmt.Errorf("bitbucket: %w", model.ErrNoWebhookID)
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/hooks/%v", workspace, webhookID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	return i.c.Call(request, nil)
}
