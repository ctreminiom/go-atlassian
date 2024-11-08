package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewScreenService creates a new instance of ScreenService.
func NewScreenService(client service.Connector, version string, scheme *ScreenSchemeService, tab *ScreenTabService) (*ScreenService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ScreenService{
		internalClient: &internalScreenImpl{c: client, version: version},
		Scheme:         scheme,
		Tab:            tab,
	}, nil
}

// ScreenService provides methods to manage screens in Jira Service Management.
type ScreenService struct {
	// internalClient is the connector interface for screen operations.
	internalClient jira.ScreenConnector
	// Scheme is the service for managing screen schemes.
	Scheme *ScreenSchemeService
	// Tab is the service for managing screen tabs.
	Tab *ScreenTabService
}

// Fields returns a paginated list of the screens a field is used in.
//
// GET /rest/api/{2-3}/field/{fieldID}/screens
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#get-screens-for-a-field
func (s *ScreenService) Fields(ctx context.Context, fieldID string, startAt, maxResults int) (*model.ScreenFieldPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Fields(ctx, fieldID, startAt, maxResults)
}

// Gets returns a paginated list of all screens or those specified by one or more screen IDs.
//
// GET /rest/api/{2-3}/screens
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#get-screens
func (s *ScreenService) Gets(ctx context.Context, options *model.ScreenParamsScheme, startAt, maxResults int) (*model.ScreenSearchPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, options, startAt, maxResults)
}

// Create creates a screen with a default field tab
//
// POST /rest/api/{2-3}/screens
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#create-screen
func (s *ScreenService) Create(ctx context.Context, name, description string) (*model.ScreenScheme, *model.ResponseScheme, error) {
	return s.internalClient.Create(ctx, name, description)
}

// AddToDefault adds a field to the default tab of the default screen.
//
// POST /rest/api/{2-3}/screens/addToDefault/{fieldID}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#add-field-to-default-screen
func (s *ScreenService) AddToDefault(ctx context.Context, fieldID string) (*model.ResponseScheme, error) {
	return s.internalClient.AddToDefault(ctx, fieldID)
}

// Update updates a screen. Only screens used in classic projects can be updated.
//
// PUT /rest/api/{2-3}/screens/{screenID}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#update-screen
func (s *ScreenService) Update(ctx context.Context, screenID int, name, description string) (*model.ScreenScheme, *model.ResponseScheme, error) {
	return s.internalClient.Update(ctx, screenID, name, description)
}

// Delete deletes a screen.
// A screen cannot be deleted if it is used in a screen scheme,
//
// workflow, or workflow draft. Only screens used in classic projects can be deleted.
//
// DELETE /rest/api/{2-3}/screens/{screenID}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#delete-screen
func (s *ScreenService) Delete(ctx context.Context, screenID int) (*model.ResponseScheme, error) {
	return s.internalClient.Delete(ctx, screenID)
}

// Available returns the fields that can be added to a tab on a screen.
//
// GET /rest/api/{2-3}/screens/{screenID}/availableFields
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#get-available-screen-fields
func (s *ScreenService) Available(ctx context.Context, screenID int) ([]*model.AvailableScreenFieldScheme, *model.ResponseScheme, error) {
	return s.internalClient.Available(ctx, screenID)
}

type internalScreenImpl struct {
	c       service.Connector
	version string
}

func (i *internalScreenImpl) Fields(ctx context.Context, fieldID string, startAt, maxResults int) (*model.ScreenFieldPageScheme, *model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, nil, model.ErrNoFieldID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/screens?%v", i.version, fieldID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ScreenFieldPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalScreenImpl) Gets(ctx context.Context, options *model.ScreenParamsScheme, startAt, maxResults int) (*model.ScreenSearchPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		for _, id := range options.IDs {
			params.Add("id", strconv.Itoa(id))
		}

		if options.QueryString != "" {
			params.Add("queryString", options.QueryString)
		}

		for _, scope := range options.Scope {
			params.Add("scope", scope)
		}

		if options.OrderBy != "" {
			params.Add("orderBy", options.OrderBy)
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ScreenSearchPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalScreenImpl) Create(ctx context.Context, name, description string) (*model.ScreenScheme, *model.ResponseScheme, error) {

	if name == "" {
		return nil, nil, model.ErrNoScreenName
	}

	payload := map[string]interface{}{"name": name}

	if description != "" {
		payload["description"] = description
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	screen := new(model.ScreenScheme)
	response, err := i.c.Call(request, screen)
	if err != nil {
		return nil, response, err
	}

	return screen, response, nil
}

func (i *internalScreenImpl) AddToDefault(ctx context.Context, fieldID string) (*model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/addToDefault/%v", i.version, fieldID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalScreenImpl) Update(ctx context.Context, screenID int, name, description string) (*model.ScreenScheme, *model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, nil, model.ErrNoScreenID
	}

	payload := map[string]interface{}{"name": name}

	if description != "" {
		payload["description"] = description
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v", i.version, screenID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	screen := new(model.ScreenScheme)
	response, err := i.c.Call(request, screen)
	if err != nil {
		return nil, response, err
	}

	return screen, response, nil
}

func (i *internalScreenImpl) Delete(ctx context.Context, screenID int) (*model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, model.ErrNoScreenID
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v", i.version, screenID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalScreenImpl) Available(ctx context.Context, screenID int) ([]*model.AvailableScreenFieldScheme, *model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, nil, model.ErrNoScreenID
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/availableFields", i.version, screenID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var fields []*model.AvailableScreenFieldScheme
	response, err := i.c.Call(request, &fields)
	if err != nil {
		return nil, response, err
	}

	return fields, response, nil
}
