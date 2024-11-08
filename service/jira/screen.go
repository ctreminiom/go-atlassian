package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type ScreenConnector interface {

	// Fields returns a paginated list of the screens a field is used in.
	//
	// GET /rest/api/{2-3}/field/{fieldID}/screens
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens#get-screens-for-a-field
	Fields(ctx context.Context, fieldID string, startAt, maxResults int) (*model.ScreenFieldPageScheme, *model.ResponseScheme, error)

	// Gets returns a paginated list of all screens or those specified by one or more screen IDs.
	//
	// GET /rest/api/{2-3}/screens
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens#get-screens
	Gets(ctx context.Context, options *model.ScreenParamsScheme, startAt, maxResults int) (*model.ScreenSearchPageScheme, *model.ResponseScheme, error)

	// Create creates a screen with a default field tab
	//
	// POST /rest/api/{2-3}/screens
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens#create-screen
	Create(ctx context.Context, name, description string) (*model.ScreenScheme, *model.ResponseScheme, error)

	// AddToDefault adds a field to the default tab of the default screen.
	//
	// POST /rest/api/{2-3}/screens/addToDefault/{fieldID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens#add-field-to-default-screen
	AddToDefault(ctx context.Context, fieldID string) (*model.ResponseScheme, error)

	// Update updates a screen. Only screens used in classic projects can be updated.
	//
	// PUT /rest/api/{2-3}/screens/{screenID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens#update-screen
	Update(ctx context.Context, screenID int, name, description string) (*model.ScreenScheme, *model.ResponseScheme, error)

	// Delete deletes a screen.
	// A screen cannot be deleted if it is used in a screen scheme,
	//
	// workflow, or workflow draft. Only screens used in classic projects can be deleted.
	//
	// DELETE /rest/api/{2-3}/screens/{screenID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens#delete-screen
	Delete(ctx context.Context, screenID int) (*model.ResponseScheme, error)

	// Available returns the fields that can be added to a tab on a screen.
	//
	// GET /rest/api/{2-3}/screens/{screenID}/availableFields
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens#get-available-screen-fields
	Available(ctx context.Context, screenID int) ([]*model.AvailableScreenFieldScheme, *model.ResponseScheme, error)
}

type ScreenSchemeConnector interface {

	// Gets returns a paginated list of screen schemes.
	//
	// Only screen schemes used in classic projects are returned.
	//
	// GET /rest/api/{2-3}/screenscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#get-screen-schemes
	Gets(ctx context.Context, options *model.ScreenSchemeParamsScheme, startAt, maxResults int) (*model.ScreenSchemePageScheme, *model.ResponseScheme, error)

	// Create creates a screen scheme.
	//
	// POST /rest/api/{2-3}/screenscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#create-screen-scheme
	Create(ctx context.Context, payload *model.ScreenSchemePayloadScheme) (*model.ScreenSchemeScheme, *model.ResponseScheme, error)

	// Update updates a screen scheme. Only screen schemes used in classic projects can be updated.
	//
	// PUT /rest/api/{2-3}/screenscheme/{screenSchemeID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#update-screen-scheme
	Update(ctx context.Context, screenSchemeID string, payload *model.ScreenSchemePayloadScheme) (*model.ResponseScheme, error)

	// Delete deletes a screen scheme. A screen scheme cannot be deleted if it is used in an issue type screen scheme.
	//
	// Only screens schemes used in classic projects can be deleted.
	//
	// DELETE /rest/api/{2-3}/screenscheme/{screenSchemeID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#delete-screen-scheme
	Delete(ctx context.Context, screenSchemeID string) (*model.ResponseScheme, error)
}

type ScreenTabConnector interface {

	// Gets returns the list of tabs for a screen.
	//
	// GET /rest/api/{2-3}/screens/{screenID}/tabs
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#get-all-screen-tabs
	Gets(ctx context.Context, screenID int, projectKey string) ([]*model.ScreenTabScheme, *model.ResponseScheme, error)

	// Create creates a tab for a screen.
	//
	// POST /rest/api/{2-3}/screens/{screenID}/tabs
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#create-screen-tab
	Create(ctx context.Context, screenID int, tabName string) (*model.ScreenTabScheme, *model.ResponseScheme, error)

	// Update updates the name of a screen tab.
	//
	// PUT /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#update-screen-tab
	Update(ctx context.Context, screenID, tabID int, newTabName string) (*model.ScreenTabScheme, *model.ResponseScheme, error)

	// Delete deletes a screen tab.
	//
	// DELETE /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#delete-screen-tab
	Delete(ctx context.Context, screenID, tabID int) (*model.ResponseScheme, error)

	// Move moves a screen tab.
	//
	// POST /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}/move/{pos}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#move-screen-tab
	Move(ctx context.Context, screenID, tabID, position int) (*model.ResponseScheme, error)
}

type ScreenTabFieldConnector interface {

	// Gets returns all fields for a screen tab.
	//
	// GET /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}/fields
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#get-all-screen-tab-fields
	Gets(ctx context.Context, screenID, tabID int) ([]*model.ScreenTabFieldScheme, *model.ResponseScheme, error)

	// Add adds a field to a screen tab.
	//
	// POST /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}/fields
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#add-screen-tab-field
	Add(ctx context.Context, screenID, tabID int, fieldID string) (*model.ScreenTabFieldScheme, *model.ResponseScheme, error)

	// Remove removes a field from a screen tab.
	//
	// DELETE /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}/fields/{fieldID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#remove-screen-tab-field
	Remove(ctx context.Context, screenID, tabID int, fieldID string) (*model.ResponseScheme, error)

	// Move moves a screen tab field.
	//
	// If after and position are provided in the request, position is ignored.
	//
	// POST /rest/api/{2-3}/screens/{screenID}/tabs/{tabID}/fields/{fieldID}/move
	//
	// TODO: Add documentation
	Move(ctx context.Context, screenID, tabID int, fieldID, after, position string) (*model.ResponseScheme, error)
}
