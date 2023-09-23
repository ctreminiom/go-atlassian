package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/bitbucket"
	"net/http"
)

func NewWorkspaceHookService(client service.Connector) *WorkspaceHookService {

	return &WorkspaceHookService{
		internalClient: &internalWorkspaceHookServiceImpl{c: client},
	}
}

type WorkspaceHookService struct {
	internalClient bitbucket.WorkspaceHookConnector
}

// Gets returns a paginated list of webhooks installed on this workspace.
//
// GET /2.0/workspaces/{workspace}/hooks
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#list-webhooks-for-a-workspace
func (w *WorkspaceHookService) Gets(ctx context.Context, workspace string) (*model.WebhookSubscriptionPageScheme, *model.ResponseScheme, error) {
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
	return w.internalClient.Create(ctx, workspace, payload)
}

// Get returns the webhook with the specified id installed on the given workspace.
//
// GET /2.0/workspaces/{workspace}/hooks/{uid}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#get-webhook-for-a-workspace
func (w *WorkspaceHookService) Get(ctx context.Context, workspace, webhookId string) (*model.WebhookSubscriptionScheme, *model.ResponseScheme, error) {
	return w.internalClient.Get(ctx, workspace, webhookId)
}

// Update updates the specified webhook subscription.
//
// PUT /2.0/workspaces/{workspace}/hooks/{uid}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#update-webhook-for-a-workspace
func (w *WorkspaceHookService) Update(ctx context.Context, workspace, webhookId string, payload *model.WebhookSubscriptionPayloadScheme) (*model.WebhookSubscriptionScheme, *model.ResponseScheme, error) {
	return w.internalClient.Update(ctx, workspace, webhookId, payload)
}

// Delete deletes the specified webhook subscription from the given workspace.
//
// DELETE /2.0/workspaces/{workspace}/hooks/{uid}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#delete-webhook-for-a-workspace
func (w *WorkspaceHookService) Delete(ctx context.Context, workspace, webhookId string) (*model.ResponseScheme, error) {
	return w.internalClient.Delete(ctx, workspace, webhookId)
}

type internalWorkspaceHookServiceImpl struct {
	c service.Connector
}

func (i *internalWorkspaceHookServiceImpl) Gets(ctx context.Context, workspace string) (*model.WebhookSubscriptionPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspaceError
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/hooks", workspace)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.WebhookSubscriptionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalWorkspaceHookServiceImpl) Create(ctx context.Context, workspace string, payload *model.WebhookSubscriptionPayloadScheme) (*model.WebhookSubscriptionScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspaceError
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/hooks", workspace)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	webhook := new(model.WebhookSubscriptionScheme)
	response, err := i.c.Call(request, webhook)
	if err != nil {
		return nil, response, err
	}

	return webhook, response, nil
}

func (i *internalWorkspaceHookServiceImpl) Get(ctx context.Context, workspace, webhookId string) (*model.WebhookSubscriptionScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspaceError
	}

	if webhookId == "" {
		return nil, nil, model.ErrNoWebhookIDError
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/hooks/%v", workspace, webhookId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	webhook := new(model.WebhookSubscriptionScheme)
	response, err := i.c.Call(request, webhook)
	if err != nil {
		return nil, response, err
	}

	return webhook, response, nil
}

func (i *internalWorkspaceHookServiceImpl) Update(ctx context.Context, workspace, webhookId string, payload *model.WebhookSubscriptionPayloadScheme) (*model.WebhookSubscriptionScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspaceError
	}

	if webhookId == "" {
		return nil, nil, model.ErrNoWebhookIDError
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/hooks/%v", workspace, webhookId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	webhook := new(model.WebhookSubscriptionScheme)
	response, err := i.c.Call(request, webhook)
	if err != nil {
		return nil, response, err
	}

	return webhook, response, nil
}

func (i *internalWorkspaceHookServiceImpl) Delete(ctx context.Context, workspace, webhookId string) (*model.ResponseScheme, error) {

	if workspace == "" {
		return nil, model.ErrNoWorkspaceError
	}

	if webhookId == "" {
		return nil, model.ErrNoWebhookIDError
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/hooks/%v", workspace, webhookId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
