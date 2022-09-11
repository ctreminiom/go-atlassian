package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strconv"
)

func NewScreenService(client service.Client, version string, scheme *ScreenSchemeService, tab *ScreenTabService) (*ScreenService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ScreenService{
		internalClient: &internalScreenImpl{c: client, version: version},
		Scheme:         scheme,
		Tab:            tab,
	}, nil
}

type ScreenService struct {
	internalClient jira.ScreenConnector
	Scheme         *ScreenSchemeService
	Tab            *ScreenTabService
}

// Fields returns a paginated list of the screens a field is used in.
//
// GET /rest/api/{2-3}/field/{fieldId}/screens
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#get-screens-for-a-field
func (s *ScreenService) Fields(ctx context.Context, fieldId string, startAt, maxResults int) (*model.ScreenFieldPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Fields(ctx, fieldId, startAt, maxResults)
}

// Gets returns a paginated list of all screens or those specified by one or more screen IDs.
//
// GET /rest/api/{2-3}/screens
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#get-screens
func (s *ScreenService) Gets(ctx context.Context, screenIds []int, startAt, maxResults int) (*model.ScreenSearchPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, screenIds, startAt, maxResults)
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
// POST /rest/api/{2-3}/screens/addToDefault/{fieldId}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#add-field-to-default-screen
func (s *ScreenService) AddToDefault(ctx context.Context, fieldId string) (*model.ResponseScheme, error) {
	return s.internalClient.AddToDefault(ctx, fieldId)
}

// Update updates a screen. Only screens used in classic projects can be updated.
//
// PUT /rest/api/{2-3}/screens/{screenId}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#update-screen
func (s *ScreenService) Update(ctx context.Context, screenId int, name, description string) (*model.ScreenScheme, *model.ResponseScheme, error) {
	return s.internalClient.Update(ctx, screenId, name, description)
}

// Delete deletes a screen.
// A screen cannot be deleted if it is used in a screen scheme,
//
// workflow, or workflow draft. Only screens used in classic projects can be deleted.
//
// DELETE /rest/api/{2-3}/screens/{screenId}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#delete-screen
func (s *ScreenService) Delete(ctx context.Context, screenId int) (*model.ResponseScheme, error) {
	return s.internalClient.Delete(ctx, screenId)
}

// Available returns the fields that can be added to a tab on a screen.
//
// GET /rest/api/{2-3}/screens/{screenId}/availableFields
//
// https://docs.go-atlassian.io/jira-software-cloud/screens#get-available-screen-fields
func (s *ScreenService) Available(ctx context.Context, screenId int) ([]*model.AvailableScreenFieldScheme, *model.ResponseScheme, error) {
	return s.internalClient.Available(ctx, screenId)
}

type internalScreenImpl struct {
	c       service.Client
	version string
}

func (i *internalScreenImpl) Fields(ctx context.Context, fieldId string, startAt, maxResults int) (*model.ScreenFieldPageScheme, *model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, nil, model.ErrNoFieldIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/screens?%v", i.version, fieldId, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalScreenImpl) Gets(ctx context.Context, screenIds []int, startAt, maxResults int) (*model.ScreenSearchPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, screenID := range screenIds {
		params.Add("id", strconv.Itoa(screenID))
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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
		return nil, nil, model.ErrNoScreenNameError
	}

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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

func (i *internalScreenImpl) AddToDefault(ctx context.Context, fieldId string) (*model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/addToDefault/%v", i.version, fieldId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalScreenImpl) Update(ctx context.Context, screenId int, name, description string) (*model.ScreenScheme, *model.ResponseScheme, error) {

	if screenId == 0 {
		return nil, nil, model.ErrNoScreenIDError
	}

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v", i.version, screenId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
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

func (i *internalScreenImpl) Delete(ctx context.Context, screenId int) (*model.ResponseScheme, error) {

	if screenId == 0 {
		return nil, model.ErrNoScreenIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v", i.version, screenId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalScreenImpl) Available(ctx context.Context, screenId int) ([]*model.AvailableScreenFieldScheme, *model.ResponseScheme, error) {

	if screenId == 0 {
		return nil, nil, model.ErrNoScreenIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/availableFields", i.version, screenId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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
