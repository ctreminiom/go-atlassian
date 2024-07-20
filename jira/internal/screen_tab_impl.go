package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
)

func NewScreenTabService(client service.Connector, version string, field *ScreenTabFieldService) (*ScreenTabService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ScreenTabService{
		internalClient: &internalScreenTabImpl{c: client, version: version},
		Field:          field,
	}, nil
}

type ScreenTabService struct {
	internalClient jira.ScreenTabConnector
	Field          *ScreenTabFieldService
}

// Gets returns the list of tabs for a screen.
//
// GET /rest/api/{2-3}/screens/{screenID}/tabs
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#get-all-screen-tabs
func (s *ScreenTabService) Gets(ctx context.Context, screenID int, projectKey string) ([]*model.ScreenTabScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, screenID, projectKey)
}

// Create creates a tab for a screen.
//
// POST /rest/api/{2-3}/screens/{screenID}/tabs
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#create-screen-tab
func (s *ScreenTabService) Create(ctx context.Context, screenID int, tabName string) (*model.ScreenTabScheme, *model.ResponseScheme, error) {
	return s.internalClient.Create(ctx, screenID, tabName)
}

// Update updates the name of a screen tab.
//
// PUT /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#update-screen-tab
func (s *ScreenTabService) Update(ctx context.Context, screenID, tabID int, newTabName string) (*model.ScreenTabScheme, *model.ResponseScheme, error) {
	return s.internalClient.Update(ctx, screenID, tabID, newTabName)
}

// Delete deletes a screen tab.
//
// DELETE /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#delete-screen-tab
func (s *ScreenTabService) Delete(ctx context.Context, screenID, tabID int) (*model.ResponseScheme, error) {
	return s.internalClient.Delete(ctx, screenID, tabID)
}

// Move moves a screen tab.
//
// POST /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}/move/{pos}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#move-screen-tab
func (s *ScreenTabService) Move(ctx context.Context, screenID, tabID, position int) (*model.ResponseScheme, error) {
	return s.internalClient.Move(ctx, screenID, tabID, position)
}

type internalScreenTabImpl struct {
	c       service.Connector
	version string
}

func (i *internalScreenTabImpl) Gets(ctx context.Context, screenID int, projectKey string) ([]*model.ScreenTabScheme, *model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, nil, model.ErrNoScreenIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/screens/%v/tabs", i.version, screenID))

	if projectKey != "" {

		params := url.Values{}
		params.Add("projectKey", projectKey)

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	var tabs []*model.ScreenTabScheme
	response, err := i.c.Call(request, &tabs)
	if err != nil {
		return nil, response, err
	}

	return tabs, response, nil
}

func (i *internalScreenTabImpl) Create(ctx context.Context, screenID int, tabName string) (*model.ScreenTabScheme, *model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, nil, model.ErrNoScreenIDError
	}

	if tabName == "" {
		return nil, nil, model.ErrNoScreenTabNameError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs", i.version, screenID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"name": tabName})
	if err != nil {
		return nil, nil, err
	}

	tab := new(model.ScreenTabScheme)
	response, err := i.c.Call(request, tab)
	if err != nil {
		return nil, response, err
	}

	return tab, response, nil
}

func (i *internalScreenTabImpl) Update(ctx context.Context, screenID, tabID int, newTabName string) (*model.ScreenTabScheme, *model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, nil, model.ErrNoScreenIDError
	}

	if tabID == 0 {
		return nil, nil, model.ErrNoScreenTabIDError
	}

	if newTabName == "" {
		return nil, nil, model.ErrNoScreenTabNameError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs/%v", i.version, screenID, tabID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"name": newTabName})
	if err != nil {
		return nil, nil, err
	}

	tab := new(model.ScreenTabScheme)
	response, err := i.c.Call(request, tab)
	if err != nil {
		return nil, response, err
	}

	return tab, response, nil
}

func (i *internalScreenTabImpl) Delete(ctx context.Context, screenID, tabID int) (*model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, model.ErrNoScreenIDError
	}

	if tabID == 0 {
		return nil, model.ErrNoScreenTabIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs/%v", i.version, screenID, tabID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalScreenTabImpl) Move(ctx context.Context, screenID, tabID, position int) (*model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, model.ErrNoScreenIDError
	}

	if tabID == 0 {
		return nil, model.ErrNoScreenTabIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs/%v/move/%v", i.version, screenID, tabID, position)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
