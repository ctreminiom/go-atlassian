package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
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
// GET /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}/fields
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#get-all-screen-tab-fields
func (s *ScreenTabFieldService) Gets(ctx context.Context, screenID, tabID int) ([]*model.ScreenTabFieldScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, screenID, tabID)
}

// Add adds a field to a screen tab.
//
// POST /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}/fields
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#add-screen-tab-field
func (s *ScreenTabFieldService) Add(ctx context.Context, screenID, tabID int, fieldID string) (*model.ScreenTabFieldScheme, *model.ResponseScheme, error) {
	return s.internalClient.Add(ctx, screenID, tabID, fieldID)
}

// Remove removes a field from a screen tab.
//
// DELETE /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}/fields/{fieldID}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#remove-screen-tab-field
func (s *ScreenTabFieldService) Remove(ctx context.Context, screenID, tabID int, fieldID string) (*model.ResponseScheme, error) {
	return s.internalClient.Remove(ctx, screenID, tabID, fieldID)
}

// Move moves a screen tab field.
//
// If after and position are provided in the request, position is ignored.
//
// POST /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}/fields/{fieldID}/move
//
// TODO: Add documentation
func (s *ScreenTabFieldService) Move(ctx context.Context, screenID, tabID int, fieldID, after, position string) (*model.ResponseScheme, error) {
	return s.internalClient.Move(ctx, screenID, tabID, fieldID, after, position)
}

type internalScreenTabFieldImpl struct {
	c       service.Connector
	version string
}

func (i *internalScreenTabFieldImpl) Gets(ctx context.Context, screenID, tabID int) ([]*model.ScreenTabFieldScheme, *model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, nil, model.ErrNoScreenID
	}

	if tabID == 0 {
		return nil, nil, model.ErrNoScreenTabID
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs/%v/fields", i.version, screenID, tabID)

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

func (i *internalScreenTabFieldImpl) Add(ctx context.Context, screenID, tabID int, fieldID string) (*model.ScreenTabFieldScheme, *model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, nil, model.ErrNoScreenID
	}

	if tabID == 0 {
		return nil, nil, model.ErrNoScreenTabID
	}

	if fieldID == "" {
		return nil, nil, model.ErrNoFieldID
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs/%v/fields", i.version, screenID, tabID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"fieldId": fieldID})
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

func (i *internalScreenTabFieldImpl) Remove(ctx context.Context, screenID, tabID int, fieldID string) (*model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, model.ErrNoScreenID
	}

	if tabID == 0 {
		return nil, model.ErrNoScreenTabID
	}

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs/%v/fields/%v", i.version, screenID, tabID, fieldID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalScreenTabFieldImpl) Move(ctx context.Context, screenID, tabID int, fieldID, after, position string) (*model.ResponseScheme, error) {

	if screenID == 0 {
		return nil, model.ErrNoScreenID
	}

	if tabID == 0 {
		return nil, model.ErrNoScreenTabID
	}

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	endpoint := fmt.Sprintf("rest/api/%v/screens/%v/tabs/%v/fields/%v/move", i.version, screenID, tabID, fieldID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"after": after, "position": position})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
