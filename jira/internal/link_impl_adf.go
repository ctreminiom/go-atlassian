package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

type LinkADFService struct {
	internalClient jira.LinkAdfIssueConnector
	Type           *LinkTypeService
	Remote         *RemoteLinkService
}

type internalLinkADFServiceImpl struct {
	c       service.Connector
	version string
}

// Get returns an issue link.
//
// GET /rest/api/{2-3}/issueLink/{linkId}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-link
func (l *LinkADFService) Get(ctx context.Context, linkId string) (*model.IssueLinkScheme, *model.ResponseScheme, error) {
	return l.internalClient.Get(ctx, linkId)
}

// Gets get the issue links ID's associated with a Jira Issue
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-links
func (l *LinkADFService) Gets(ctx context.Context, issueKeyOrId string) (*model.IssueLinkPageScheme, *model.ResponseScheme, error) {
	return l.internalClient.Gets(ctx, issueKeyOrId)
}

// Delete deletes an issue link.
//
// DELETE /rest/api/{2-3}/issueLink/{linkId}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link#delete-issue-link
func (l *LinkADFService) Delete(ctx context.Context, linkId string) (*model.ResponseScheme, error) {
	return l.internalClient.Delete(ctx, linkId)
}

// Create creates a link between two issues. Use this operation to indicate a relationship between two issues
//
// and optionally add a comment to the form (outward) issue.
//
// To use this resource the site must have Issue Linking enabled.
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/link#create-issue-link
func (l *LinkADFService) Create(ctx context.Context, payload *model.LinkPayloadSchemeV3) (*model.ResponseScheme, error) {
	return l.internalClient.Create(ctx, payload)
}

func (i *internalLinkADFServiceImpl) Get(ctx context.Context, linkId string) (*model.IssueLinkScheme, *model.ResponseScheme, error) {

	if linkId == "" {
		return nil, nil, model.ErrNoTypeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issueLink/%v", i.version, linkId)

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

func (i *internalLinkADFServiceImpl) Gets(ctx context.Context, issueKeyOrId string) (*model.IssueLinkPageScheme, *model.ResponseScheme, error) {

	if len(issueKeyOrId) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v?fields=issuelinks", i.version, issueKeyOrId)

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

func (i *internalLinkADFServiceImpl) Delete(ctx context.Context, linkId string) (*model.ResponseScheme, error) {

	if linkId == "" {
		return nil, model.ErrNoTypeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issueLink/%v", i.version, linkId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalLinkADFServiceImpl) Create(ctx context.Context, payload *model.LinkPayloadSchemeV3) (*model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/issueLink", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
