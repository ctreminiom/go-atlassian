package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewMySelfService creates a new instance of MySelfService.
func NewMySelfService(client service.Connector, version string) (*MySelfService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &MySelfService{
		internalClient: &internalMySelfImpl{c: client, version: version},
	}, nil
}

// MySelfService provides methods to manage the current user's details in Jira Service Management.
type MySelfService struct {
	// internalClient is the connector interface for current user operations.
	internalClient jira.MySelfConnector
}

// Details returns details for the current user.
//
// GET /rest/api/{2-3}/myself
//
// https://docs.go-atlassian.io/jira-software-cloud/myself#get-current-user
func (m *MySelfService) Details(ctx context.Context, expand []string) (*model.UserScheme, *model.ResponseScheme, error) {
	return m.internalClient.Details(ctx, expand)
}

// Get returns the values of the user's preferences.
//
// GET /rest/api/{2-3}/mypreferences
//
// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-myself/#api-rest-api-3-mypreferences-get
func (m *MySelfService) Get(ctx context.Context, key string) (map[string]interface{}, *model.ResponseScheme, error) {
	return m.internalClient.Get(ctx, key)
}

// Set sets the value of the user's preference.
//
// PUT /rest/api/{2-3}/mypreferences
//
// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-myself/#api-rest-api-3-mypreferences-put
func (m *MySelfService) Set(ctx context.Context, key string, value string) (map[string]interface{}, *model.ResponseScheme, error) {
	return m.internalClient.Set(ctx, key, value)
}

// Delete deletes the user's preference.
//
// DELETE /rest/api/{2-3}/mypreferences
//
// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-myself/#api-rest-api-3-mypreferences-delete
func (m *MySelfService) Delete(ctx context.Context, key string) (*model.ResponseScheme, error) {
	return m.internalClient.Delete(ctx, key)
}

type internalMySelfImpl struct {
	c       service.Connector
	version string
}

func (i *internalMySelfImpl) Details(ctx context.Context, expand []string) (*model.UserScheme, *model.ResponseScheme, error) {

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/myself", i.version))

	if expand != nil {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	my := new(model.UserScheme)
	response, err := i.c.Call(request, my)
	if err != nil {
		return nil, response, err
	}

	return my, response, nil
}

func (i *internalMySelfImpl) Get(ctx context.Context, key string) (map[string]interface{}, *model.ResponseScheme, error) {

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/mypreferences", i.version))

	if key != "" {
		params := url.Values{}
		params.Add("key", key)
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	preferences := make(map[string]interface{})
	response, err := i.c.Call(request, &preferences)
	if err != nil {
		return nil, response, err
	}

	return preferences, response, nil
}

func (i *internalMySelfImpl) Set(ctx context.Context, key string, value string) (map[string]interface{}, *model.ResponseScheme, error) {

	if key == "" {
		return nil, nil, model.ErrNoKeyError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/mypreferences", i.version))

	params := url.Values{}
	params.Add("key", key)
	endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))

	payload := struct {
		Value string `json:"value"`
	}{
		Value: value,
	}

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	preferences := make(map[string]interface{})
	response, err := i.c.Call(request, &preferences)
	if err != nil {
		return nil, response, err
	}

	return preferences, response, nil
}

func (i *internalMySelfImpl) Delete(ctx context.Context, key string) (*model.ResponseScheme, error) {

	if key == "" {
		return nil, model.ErrNoKeyError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/mypreferences", i.version))

	params := url.Values{}
	params.Add("key", key)
	endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint.String(), "", nil)
	if err != nil {
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		return response, err
	}

	return response, nil
}
