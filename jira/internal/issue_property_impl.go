package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewIssuePropertyService(client service.Connector, version string) (*IssuePropertyService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssuePropertyService{
		internalClient: &internalIssuePropertyImpl{c: client, version: version},
	}, nil
}

type IssuePropertyService struct {
	internalClient jira.IssuePropertyConnector
}

/*
Gets returns the URLs and keys of an issue's properties.
  - This operation can be accessed anonymously.

Permissions required:
  - Browse projects project permission for the project containing the issue.
  - If issue-level security is configured, issue-level security permission to view the issue.

Endpoint: GET /rest/api/{apiVersion}/issue/{issueIdOrKey}/properties

You can refer to the documentation: [Get issue property keys]

[Get issue property keys]: https://docs.go-atlassian.io/jira-software-cloud/issues/properties#get-issue-property-keys
*/
func (i *IssuePropertyService) Gets(ctx context.Context, issueIdOrKey string) (*model.PropertyPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, issueIdOrKey)
}

/*
Get returns the key and value of an issue's property.
  - This operation can be accessed anonymously.

Permissions required:
  - Browse projects project permission for the project containing the issue.
  - If issue-level security is configured, issue-level security permission to view the issue.

Endpoint: GET /rest/api/{apiVersion}/issue/{issueIdOrKey}/properties/{propertyKey}

You can refer to the documentation: [Get issue property]

[Get issue property]: https://docs.go-atlassian.io/jira-software-cloud/issues/properties#get-issue-property
*/
func (i *IssuePropertyService) Get(ctx context.Context, issueKey, propertyKey string) (*model.EntityPropertyScheme, *model.ResponseScheme, error) {
	return i.internalClient.Get(ctx, issueKey, propertyKey)
}

/*
Set sets the value of an issue's property. Use this resource to store custom data against an issue.
  - The value of the request body must be a valid, non-empty JSON blob. The maximum length is 32768 characters.
  - This operation can be accessed anonymously.

Permissions required:
  - Browse projects and Edit issues project permissions for the project containing the issue.
  - If issue-level security is configured, issue-level security permission to view the issue.

Endpoint: PUT /rest/api/{apiVersion}/issue/{issueIdOrKey}/properties/{propertyKey}

You can refer to the documentation: [Set issue property]

[Set issue property]: https://docs.go-atlassian.io/jira-software-cloud/issues/properties#set-issue-property
*/
func (i *IssuePropertyService) Set(ctx context.Context, issueKey, propertyKey string, payload interface{}) (*model.ResponseScheme, error) {
	return i.internalClient.Set(ctx, issueKey, propertyKey, payload)
}

/*
Delete deletes an issue's property.
  - This operation can be accessed anonymously.

Permissions required:
  - Browse projects and Edit issues project permissions for the project containing the issue.
  - If issue-level security is configured, issue-level security permission to view the issue.

Endpoint: DELETE /rest/api/{apiVersion}/issue/{issueIdOrKey}/properties/{propertyKey}

You can refer to the documentation: [Delete issue property]

[Delete issue property]: https://docs.go-atlassian.io/jira-software-cloud/issues/properties#delete-issue-property
*/
func (i *IssuePropertyService) Delete(ctx context.Context, issueKey, propertyKey string) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, issueKey, propertyKey)
}

type internalIssuePropertyImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssuePropertyImpl) Gets(ctx context.Context, issueIdOrKey string) (*model.PropertyPageScheme, *model.ResponseScheme, error) {

	if issueIdOrKey == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/properties", i.version, issueIdOrKey)

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
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if propertyKey == "" {
		return nil, nil, model.ErrNoPropertyKeyError
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
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if propertyKey == "" {
		return nil, model.ErrNoPropertyKeyError
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
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if propertyKey == "" {
		return nil, model.ErrNoPropertyKeyError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/properties/%v", i.version, issueKey, propertyKey)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
