package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
	"net/http"
	"net/url"
)

// NewProjectValidatorService creates a new instance of ProjectValidatorService.
func NewProjectValidatorService(client service.Connector, version string) (*ProjectValidatorService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectValidatorService{
		internalClient: &internalProjectValidatorImpl{c: client, version: version},
	}, nil
}

// ProjectValidatorService provides methods to validate project keys and names in Jira Service Management.
type ProjectValidatorService struct {
	// internalClient is the connector interface for project validation operations.
	internalClient jira.ProjectValidatorConnector
}

// Validate validates a project key by confirming the key is a valid string and not in use.
//
// GET /rest/api/{2-3}/projectvalidate/key
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/validation#validate-project-key
func (p *ProjectValidatorService) Validate(ctx context.Context, key string) (*model.ProjectValidationMessageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Validate(ctx, key)
}

// Key validates a project key and, if the key is invalid or in use,
//
// generates a valid random string for the project key.
//
// GET /rest/api/{2-3}/projectvalidate/validProjectKey
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/validation#get-valid-project-key
func (p *ProjectValidatorService) Key(ctx context.Context, key string) (string, *model.ResponseScheme, error) {
	return p.internalClient.Key(ctx, key)
}

// Name checks that a project name isn't in use.
//
// If the name isn't in use, the passed string is returned.
//
// If the name is in use, this operation attempts to generate a valid project name based on the one supplied,
//
// usually by adding a sequence number. If a valid project name cannot be generated, a 404 response is returned.
//
// GET /rest/api/{2-3}/projectvalidate/validProjectName
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/validation#get-valid-project-name
func (p *ProjectValidatorService) Name(ctx context.Context, name string) (string, *model.ResponseScheme, error) {
	return p.internalClient.Name(ctx, name)
}

type internalProjectValidatorImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectValidatorImpl) Validate(ctx context.Context, key string) (*model.ProjectValidationMessageScheme, *model.ResponseScheme, error) {

	if key == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	params := url.Values{}
	params.Add("key", key)

	endpoint := fmt.Sprintf("rest/api/%v/projectvalidate/key?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	message := new(model.ProjectValidationMessageScheme)
	response, err := i.c.Call(request, message)
	if err != nil {
		return nil, response, err
	}

	return message, response, nil
}

func (i *internalProjectValidatorImpl) Key(ctx context.Context, key string) (string, *model.ResponseScheme, error) {

	if key == "" {
		return "", nil, model.ErrNoProjectIDOrKey
	}

	params := url.Values{}
	params.Add("key", key)

	endpoint := fmt.Sprintf("rest/api/%v/projectvalidate/validProjectKey?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return "", nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		return "", response, err
	}

	return response.Bytes.String(), response, nil
}

func (i *internalProjectValidatorImpl) Name(ctx context.Context, name string) (string, *model.ResponseScheme, error) {

	if name == "" {
		return "", nil, model.ErrNoProjectName
	}

	params := url.Values{}
	params.Add("name", name)

	endpoint := fmt.Sprintf("rest/api/%v/projectvalidate/validProjectName?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return "", nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		return "", response, err
	}

	return response.Bytes.String(), response, nil
}
