package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewIssuePropertyService creates a new instance of the IssuePropertyService.
func NewIssuePropertyService(client service.Connector, version string) (*IssuePropertyService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssuePropertyService{
		internalClient: &internalIssuePropertyImpl{c: client, version: version},
	}, nil
}

// IssuePropertyService handles the issue property methods for the Jira Cloud REST API.
type IssuePropertyService struct {
	internalClient jira.IssuePropertyConnector
}

/*
Gets returns the URLs and keys of an issue's properties.
  - This operation can be accessed anonymously.

Permissions required:
  - Browse projects project permission for the project containing the issue.
  - If issue-level security is configured, issue-level security permission to view the issue.

Endpoint: GET /rest/api/{apiVersion}/issue/{issueKeyOrID}/properties

You can refer to the documentation: [Get issue property keys]

[Get issue property keys]: https://docs.go-atlassian.io/jira-software-cloud/issues/properties#get-issue-property-keys
*/
func (i *IssuePropertyService) Gets(ctx context.Context, issueKeyOrID string) (*model.PropertyPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, issueKeyOrID)
}

/*
Get returns the key and value of an issue's property.
  - This operation can be accessed anonymously.

Permissions required:
  - Browse projects project permission for the project containing the issue.
  - If issue-level security is configured, issue-level security permission to view the issue.

Endpoint: GET /rest/api/{apiVersion}/issue/{issueKeyOrID}/properties/{propertyKey}

You can refer to the documentation: [Get issue property]

[Get issue property]: https://docs.go-atlassian.io/jira-software-cloud/issues/properties#get-issue-property
*/
func (i *IssuePropertyService) Get(ctx context.Context, issueKeyOrID, propertyKey string) (*model.EntityPropertyScheme, *model.ResponseScheme, error) {
	return i.internalClient.Get(ctx, issueKeyOrID, propertyKey)
}

/*
Set sets the value of an issue's property. Use this resource to store custom data against an issue.
  - The value of the request body must be a valid, non-empty JSON blob. The maximum length is 32768 characters.
  - This operation can be accessed anonymously.

Permissions required:
  - Browse projects and Edit issues project permissions for the project containing the issue.
  - If issue-level security is configured, issue-level security permission to view the issue.

Endpoint: PUT /rest/api/{apiVersion}/issue/{issueKeyOrID}/properties/{propertyKey}

You can refer to the documentation: [Set issue property]

[Set issue property]: https://docs.go-atlassian.io/jira-software-cloud/issues/properties#set-issue-property
*/
func (i *IssuePropertyService) Set(ctx context.Context, issueKeyOrID, propertyKey string, payload interface{}) (*model.ResponseScheme, error) {
	return i.internalClient.Set(ctx, issueKeyOrID, propertyKey, payload)
}

/*
Delete deletes an issue's property.
  - This operation can be accessed anonymously.

Permissions required:
  - Browse projects and Edit issues project permissions for the project containing the issue.
  - If issue-level security is configured, issue-level security permission to view the issue.

Endpoint: DELETE /rest/api/{apiVersion}/issue/{issueKeyOrID}/properties/{propertyKey}

You can refer to the documentation: [Delete issue property]

[Delete issue property]: https://docs.go-atlassian.io/jira-software-cloud/issues/properties#delete-issue-property
*/
func (i *IssuePropertyService) Delete(ctx context.Context, issueKeyOrID, propertyKey string) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, issueKeyOrID, propertyKey)
}

type internalIssuePropertyImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssuePropertyImpl) Gets(ctx context.Context, issueKeyOrID string) (*model.PropertyPageScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/properties", i.version, issueKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	properties := new(model.PropertyPageScheme)
	response, err := i.c.Call(request, properties)
	if err != nil {
		return nil, response, err
	}

	return properties, response, nil

}

func (i *internalIssuePropertyImpl) Get(ctx context.Context, issueKey, propertyKey string) (*model.EntityPropertyScheme, *model.ResponseScheme, error) {

	if issueKey == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	if propertyKey == "" {
		return nil, nil, model.ErrNoPropertyKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/properties/%v", i.version, issueKey, propertyKey)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	property := new(model.EntityPropertyScheme)
	response, err := i.c.Call(request, property)
	if err != nil {
		return nil, response, err
	}

	return property, response, nil
}

func (i *internalIssuePropertyImpl) Set(ctx context.Context, issueKey, propertyKey string, payload interface{}) (*model.ResponseScheme, error) {

	if issueKey == "" {
		return nil, model.ErrNoIssueKeyOrID
	}

	if propertyKey == "" {
		return nil, model.ErrNoPropertyKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/properties/%v", i.version, issueKey, propertyKey)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssuePropertyImpl) Delete(ctx context.Context, issueKey, propertyKey string) (*model.ResponseScheme, error) {

	if issueKey == "" {
		return nil, model.ErrNoIssueKeyOrID
	}

	if propertyKey == "" {
		return nil, model.ErrNoPropertyKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/properties/%v", i.version, issueKey, propertyKey)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
