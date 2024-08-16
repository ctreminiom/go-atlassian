package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

// NewScreenTabFieldService creates a new instance of ScreenTabFieldService.
func NewScreenTabFieldService(client service.Connector, version string) (*ScreenTabFieldService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ScreenTabFieldService{
		internalClient: &internalScreenTabFieldImpl{c: client, version: version},
	}, nil
}

// ScreenTabFieldService provides methods to manage screen tab fields in Jira Service Management.
type ScreenTabFieldService struct {
	// internalClient is the connector interface for screen tab field operations.
	internalClient jira.ScreenTabFieldConnector
}

// Gets returns all fields for a screen tab.
//
// GET /rest/api/{2-3}/screens/{screenId}/tabs/{tabId}/fields
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#get-all-screen-tab-fields
func (s *ScreenTabFieldService) Gets(ctx context.Context, screenId, tabId int) ([]*model.ScreenTabFieldScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, screenId, tabId)
}

// Add adds a field to a screen tab.
//
// POST /rest/api/{2-3}/screens/{screenId}/tabs/{tabId}/fields
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#add-screen-tab-field
func (s *ScreenTabFieldService) Add(ctx context.Context, screenId, tabId int, fieldId string) (*model.ScreenTabFieldScheme, *model.ResponseScheme, error) {
	return s.internalClient.Add(ctx, screenId, tabId, fieldId)
}

// Remove removes a field from a screen tab.
//
// DELETE /rest/api/{2-3}/screens/{screenId}/tabs/{tabId}/fields/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#remove-screen-tab-field
func (s *ScreenTabFieldService) Remove(ctx context.Context, screenId, tabId int, fieldId string) (*model.ResponseScheme, error) {
	return s.internalClient.Remove(ctx, screenId, tabId, fieldId)
}

// Move moves a screen tab field.
//
// If after and position are provided in the request, position is ignored.
//
// POST /rest/api/{2-3}/screens/{screenId}/tabs/{tabId}/fields/{id}/move
//
// TODO: Add documentation
func (s *ScreenTabFieldService) Move(ctx context.Context, screenId, tabId int, fieldId, after, position string) (*model.ResponseScheme, error) {
	return s.internalClient.Move(ctx, screenId, tabId, fieldId, after, position)
}

type internalScreenTabFieldImpl struct {
	c       service.Connector
	version string
}

func (i *internalScreenTabFieldImpl) Gets(ctx context.Context, screenId, tabId int) ([]*model.ScreenTabFieldScheme, *model.ResponseScheme, error) {

	if screenId == 0 {
		return nil, nil, model.ErrNoScreenIDError
	}

	if tabId == 0 {
		return nil, nil, model.ErrNoScreenTabIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs/%v/fields", i.version, screenId, tabId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var fields []*model.ScreenTabFieldScheme
	response, err := i.c.Call(request, &fields)
	if err != nil {
		return nil, response, err
	}

	return fields, response, nil
}

func (i *internalScreenTabFieldImpl) Add(ctx context.Context, screenId, tabId int, fieldId string) (*model.ScreenTabFieldScheme, *model.ResponseScheme, error) {

	if screenId == 0 {
		return nil, nil, model.ErrNoScreenIDError
	}

	if tabId == 0 {
		return nil, nil, model.ErrNoScreenTabIDError
	}

	if fieldId == "" {
		return nil, nil, model.ErrNoFieldIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs/%v/fields", i.version, screenId, tabId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"fieldId": fieldId})
	if err != nil {
		return nil, nil, err
	}

	field := new(model.ScreenTabFieldScheme)
	response, err := i.c.Call(request, field)
	if err != nil {
		return nil, response, err
	}

	return field, response, nil
}

func (i *internalScreenTabFieldImpl) Remove(ctx context.Context, screenId, tabId int, fieldId string) (*model.ResponseScheme, error) {

	if screenId == 0 {
		return nil, model.ErrNoScreenIDError
	}

	if tabId == 0 {
		return nil, model.ErrNoScreenTabIDError
	}

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs/%v/fields/%v", i.version, screenId, tabId, fieldId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalScreenTabFieldImpl) Move(ctx context.Context, screenId, tabId int, fieldId, after, position string) (*model.ResponseScheme, error) {

	if screenId == 0 {
		return nil, model.ErrNoScreenIDError
	}

	if tabId == 0 {
		return nil, model.ErrNoScreenTabIDError
	}

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs/%v/fields/%v/move", i.version, screenId, tabId, fieldId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"after": after, "position": position})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
