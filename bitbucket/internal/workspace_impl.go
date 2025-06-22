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

// NewWorkspaceService handles communication with the workspace related methods of the Bitbucket API.
func NewWorkspaceService(client service.Connector, webhook *WorkspaceHookService, permission *WorkspacePermissionService) *WorkspaceService {

	return &WorkspaceService{
		internalClient: &internalWorkspaceServiceImpl{c: client},
		Hook:           webhook,
		Permission:     permission,
	}
}

// WorkspaceService handles communication with the workspace related methods of the Bitbucket API.
type WorkspaceService struct {
	internalClient bitbucket.WorkspaceConnector
	Hook           *WorkspaceHookService
	Permission     *WorkspacePermissionService
}

// Get returns the requested workspace.
//
// GET /2.0/workspaces/{workspace}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-a-workspace
func (w *WorkspaceService) Get(ctx context.Context, workspace string) (*model.WorkspaceScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkspaceService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	return w.internalClient.Get(ctx, workspace)
}

// Members returns all members of the requested workspace.
//
// GET /2.0/workspaces/{workspace}/members
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-a-workspace
func (w *WorkspaceService) Members(ctx context.Context, workspace string) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkspaceService).Members", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "members"))

	return w.internalClient.Members(ctx, workspace)
}

// Membership returns the workspace membership.
//
// which includes a User object for the member and a Workspace object for the requested workspace.
//
// GET /2.0/workspaces/{workspace}/members/{memberID}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-member-in-a-workspace
func (w *WorkspaceService) Membership(ctx context.Context, workspace, memberID string) (*model.WorkspaceMembershipScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkspaceService).Membership", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "membership"))

	return w.internalClient.Membership(ctx, workspace, memberID)
}

// Projects returns the list of projects in this workspace.
//
// GET /2.0/workspaces/{workspace}/projects
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-projects-in-a-workspace
func (w *WorkspaceService) Projects(ctx context.Context, workspace string) (*model.BitbucketProjectPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkspaceService).Projects", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "projects"))

	return w.internalClient.Projects(ctx, workspace)
}

type internalWorkspaceServiceImpl struct {
	c service.Connector
}

// Get returns the requested workspace.
func (i *internalWorkspaceServiceImpl) Get(ctx context.Context, workspace string) (*model.WorkspaceScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkspaceServiceImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	if workspace == "" {

			return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v", workspace)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	result := new(model.WorkspaceScheme)
	response, err := i.c.Call(request, result)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Members returns all members of the requested workspace.
func (i *internalWorkspaceServiceImpl) Members(ctx context.Context, workspace string) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkspaceServiceImpl).Members", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "members"))

	if workspace == "" {

			return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/members", workspace)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	page := new(model.WorkspaceMembershipPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

// Membership returns the workspace membership.
func (i *internalWorkspaceServiceImpl) Membership(ctx context.Context, workspace, memberID string) (*model.WorkspaceMembershipScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkspaceServiceImpl).Membership", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "membership"))

	if workspace == "" {

			return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
	}

	if memberID == "" {

			return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoMemberID)
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/members/%v", workspace, memberID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	member := new(model.WorkspaceMembershipScheme)
	response, err := i.c.Call(request, member)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return member, response, nil
}

// Projects returns the list of projects in this workspace.
func (i *internalWorkspaceServiceImpl) Projects(ctx context.Context, workspace string) (*model.BitbucketProjectPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkspaceServiceImpl).Projects", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "projects"))

	if workspace == "" {

			return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/projects", workspace)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	page := new(model.BitbucketProjectPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}
