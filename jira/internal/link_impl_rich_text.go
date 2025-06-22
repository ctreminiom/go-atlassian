package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// LinkRichTextService provides methods to manage issue links in Jira Service Management using the rich text format.
type LinkRichTextService struct {
	// internalClient is the connector interface for rich text issue link operations.
	internalClient jira.LinkRichTextConnector
	// Type is the service for managing link types.
	Type *LinkTypeService
	// Remote is the service for managing remote links.
	Remote *RemoteLinkService
}

type internalLinkRichTextServiceImpl struct {
	c       service.Connector
	version string
}

// Get returns an issue link.
//
// GET /rest/api/{2-3}/issueLink/{linkID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-link
func (l *LinkRichTextService) Get(ctx context.Context, linkID string) (*model.IssueLinkScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*LinkRichTextService).Get")
	defer span.End()

	return l.internalClient.Get(ctx, linkID)
}

// Gets get the issue links ID's associated with a Jira Issue
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-links
func (l *LinkRichTextService) Gets(ctx context.Context, issueKeyOrID string) (*model.IssueLinkPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*LinkRichTextService).Gets")
	defer span.End()

	return l.internalClient.Gets(ctx, issueKeyOrID)
}

// Delete deletes an issue link.
//
// DELETE /rest/api/{2-3}/issueLink/{linkID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link#delete-issue-link
func (l *LinkRichTextService) Delete(ctx context.Context, linkID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*LinkRichTextService).Delete")
	defer span.End()

	return l.internalClient.Delete(ctx, linkID)
}

// Create creates a link between two issues. Use this operation to indicate a relationship between two issues
//
// and optionally add a comment to the from (outward) issue.
//
// To use this resource the site must have Issue Linking enabled.
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link#create-issue-link
func (l *LinkRichTextService) Create(ctx context.Context, payload *model.LinkPayloadSchemeV2) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*LinkRichTextService).Create")
	defer span.End()

	return l.internalClient.Create(ctx, payload)
}

func (i *internalLinkRichTextServiceImpl) Get(ctx context.Context, linkID string) (*model.IssueLinkScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalLinkRichTextServiceImpl).Get")
	defer span.End()

	if linkID == "" {
		return nil, nil, fmt.Errorf("jira: %w", model.ErrNoTypeID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/issueLink/%v", i.version, linkID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	link := new(model.IssueLinkScheme)
	response, err := i.c.Call(request, link)
	if err != nil {
		return nil, response, err
	}

	return link, response, nil
}

func (i *internalLinkRichTextServiceImpl) Gets(ctx context.Context, issueKeyOrID string) (*model.IssueLinkPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalLinkRichTextServiceImpl).Gets")
	defer span.End()

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("jira: %w", model.ErrNoIssueKeyOrID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v?fields=issuelinks", i.version, issueKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	links := new(model.IssueLinkPageScheme)
	response, err := i.c.Call(request, links)
	if err != nil {
		return nil, response, err
	}

	return links, response, nil
}

func (i *internalLinkRichTextServiceImpl) Delete(ctx context.Context, linkID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalLinkRichTextServiceImpl).Delete")
	defer span.End()

	if linkID == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoTypeID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/issueLink/%v", i.version, linkID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalLinkRichTextServiceImpl) Create(ctx context.Context, payload *model.LinkPayloadSchemeV2) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalLinkRichTextServiceImpl).Create")
	defer span.End()

	endpoint := fmt.Sprintf("rest/api/%v/issueLink", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
