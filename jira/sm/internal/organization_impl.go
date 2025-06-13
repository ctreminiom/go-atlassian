package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/sm"
	"net/http"
	"net/url"
	"strconv"
)

// NewOrganizationService creates a new instance of OrganizationService.
// It takes a service.Connector and a version string as input and returns a pointer to OrganizationService.
func NewOrganizationService(client service.Connector, version string) *OrganizationService {
	return &OrganizationService{
		internalClient: &internalOrganizationImpl{c: client, version: version},
	}
}

// OrganizationService provides methods to interact with organization operations in Jira Service Management.
type OrganizationService struct {
	// internalClient is the connector interface for organization operations.
	internalClient sm.OrganizationConnector
}

// Gets returns a list of organizations in the Jira Service Management instance.
//
// Use this method when you want to present a list of organizations or want to locate an organization by name.
//
// GET /rest/servicedeskapi/organization
//
// https://docs.go-atlassian.io/jira-service-management-cloud/organization#get-organizations
func (o *OrganizationService) Gets(ctx context.Context, accountID string, start, limit int) (*model.OrganizationPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Gets(ctx, accountID, start, limit)
}

// Get returns details of an organization.
//
// # Use this method to get organization details whenever your application component is passed an organization ID
//
// but needs to display other organization details.
//
// GET /rest/servicedeskapi/organization/{organizationId}
//
// https://docs.go-atlassian.io/jira-service-management-cloud/organization#get-organization
func (o *OrganizationService) Get(ctx context.Context, organizationID int) (*model.OrganizationScheme, *model.ResponseScheme, error) {
	return o.internalClient.Get(ctx, organizationID)
}

// Delete deletes an organization.
//
// Note that the organization is deleted regardless of other associations it may have.
//
// For example, associations with service desks.
//
// DELETE /rest/servicedeskapi/organization/{organizationId}
//
// https://docs.go-atlassian.io/jira-service-management/organization#delete-organization
func (o *OrganizationService) Delete(ctx context.Context, organizationID int) (*model.ResponseScheme, error) {
	return o.internalClient.Delete(ctx, organizationID)
}

// Create creates an organization by passing the name of the organization.
//
// POST /rest/servicedeskapi/organization
//
// https://docs.go-atlassian.io/jira-service-management-cloud/organization#create-organization
func (o *OrganizationService) Create(ctx context.Context, name string) (*model.OrganizationScheme, *model.ResponseScheme, error) {
	return o.internalClient.Create(ctx, name)
}

// Users returns all the users associated with an organization.
//
// # Use this method where you want to provide a list of users for an
//
// organization or determine if a user is associated with an organization.
//
// GET /rest/servicedeskapi/organization/{organizationId}/user
//
// https://docs.go-atlassian.io/jira-service-management-cloud/organization#get-users-in-organization
func (o *OrganizationService) Users(ctx context.Context, organizationID, start, limit int) (*model.OrganizationUsersPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Users(ctx, organizationID, start, limit)
}

// Add adds users to an organization.
//
// POST /rest/servicedeskapi/organization/{organizationId}/user
//
// https://docs.go-atlassian.io/jira-service-management-cloud/organization#add-users-to-organization
func (o *OrganizationService) Add(ctx context.Context, organizationID int, accountIDs []string) (*model.ResponseScheme, error) {
	return o.internalClient.Add(ctx, organizationID, accountIDs)
}

// Remove removes users from an organization.
//
// DELETE /rest/servicedeskapi/organization/{organizationId}/user
//
// https://docs.go-atlassian.io/jira-service-management-cloud/organization#remove-users-from-organization
func (o *OrganizationService) Remove(ctx context.Context, organizationID int, accountIDs []string) (*model.ResponseScheme, error) {
	return o.internalClient.Remove(ctx, organizationID, accountIDs)
}

// Project returns a list of all organizations associated with a service desk.
//
// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/organization
//
// https://docs.go-atlassian.io/jira-service-management/organization#get-project-organizations
func (o *OrganizationService) Project(ctx context.Context, accountID string, serviceDeskID, start, limit int) (*model.OrganizationPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Project(ctx, accountID, serviceDeskID, start, limit)
}

// Associate adds an organization to a service desk.
//
// If the organization ID is already associated with the service desk,
//
// no change is made and the resource returns a 204 success code.
//
// POST /rest/servicedeskapi/servicedesk/{serviceDeskId}/organization
//
// https://docs.go-atlassian.io/jira-service-management-cloud/organization#associate-organization
func (o *OrganizationService) Associate(ctx context.Context, serviceDeskID, organizationID int) (*model.ResponseScheme, error) {
	return o.internalClient.Associate(ctx, serviceDeskID, organizationID)
}

// Detach removes an organization from a service desk.
//
// If the organization ID does not match an organization associated with the service desk,
//
// no change is made and the resource returns a 204 success code.
//
// DELETE /rest/servicedeskapi/servicedesk/{serviceDeskId}/organization
//
// https://docs.go-atlassian.io/jira-service-management-cloud/organization#detach-organization
func (o *OrganizationService) Detach(ctx context.Context, serviceDeskID, organizationID int) (*model.ResponseScheme, error) {
	return o.internalClient.Detach(ctx, serviceDeskID, organizationID)
}

type internalOrganizationImpl struct {
	c       service.Connector
	version string
}

func (i *internalOrganizationImpl) Gets(ctx context.Context, accountID string, start, limit int) (*model.OrganizationPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if accountID != "" {
		params.Add("accountId", accountID)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/organization?%v", params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.OrganizationPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalOrganizationImpl) Get(ctx context.Context, organizationID int) (*model.OrganizationScheme, *model.ResponseScheme, error) {

	if organizationID == 0 {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoOrganizationID)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/organization/%v", organizationID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	organization := new(model.OrganizationScheme)
	res, err := i.c.Call(req, organization)
	if err != nil {
		return nil, res, err
	}

	return organization, res, nil
}

func (i *internalOrganizationImpl) Delete(ctx context.Context, organizationID int) (*model.ResponseScheme, error) {

	if organizationID == 0 {
		return nil, fmt.Errorf("sm: %w", model.ErrNoOrganizationID)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/organization/%v", organizationID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalOrganizationImpl) Create(ctx context.Context, name string) (*model.OrganizationScheme, *model.ResponseScheme, error) {

	if name == "" {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoOrganizationName)
	}

	endpoint := "rest/servicedeskapi/organization"

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"name": name})
	if err != nil {
		return nil, nil, err
	}

	organization := new(model.OrganizationScheme)
	res, err := i.c.Call(req, organization)
	if err != nil {
		return nil, res, err
	}

	return organization, res, nil
}

func (i *internalOrganizationImpl) Users(ctx context.Context, organizationID, start, limit int) (*model.OrganizationUsersPageScheme, *model.ResponseScheme, error) {

	if organizationID == 0 {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoOrganizationID)
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("rest/servicedeskapi/organization/%v/user?%v", organizationID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.OrganizationUsersPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalOrganizationImpl) Add(ctx context.Context, organizationID int, accountIDs []string) (*model.ResponseScheme, error) {

	if organizationID == 0 {
		return nil, fmt.Errorf("sm: %w", model.ErrNoOrganizationID)
	}

	if len(accountIDs) == 0 {
		return nil, fmt.Errorf("sm: %w", model.ErrNoAccountSlice)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/organization/%v/user", organizationID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"accountIds": accountIDs})
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalOrganizationImpl) Remove(ctx context.Context, organizationID int, accountIDs []string) (*model.ResponseScheme, error) {

	if organizationID == 0 {
		return nil, fmt.Errorf("sm: %w", model.ErrNoOrganizationID)
	}

	if len(accountIDs) == 0 {
		return nil, fmt.Errorf("sm: %w", model.ErrNoAccountSlice)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/organization/%v/user", organizationID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", map[string]interface{}{"accountIds": accountIDs})
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalOrganizationImpl) Project(ctx context.Context, accountID string, serviceDeskID, start, limit int) (*model.OrganizationPageScheme, *model.ResponseScheme, error) {

	if serviceDeskID == 0 {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoServiceDeskID)
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if len(accountID) != 0 {
		params.Add("accountId", accountID)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization?%v", serviceDeskID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.OrganizationPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalOrganizationImpl) Associate(ctx context.Context, serviceDeskID, organizationID int) (*model.ResponseScheme, error) {

	if serviceDeskID == 0 {
		return nil, fmt.Errorf("sm: %w", model.ErrNoServiceDeskID)
	}

	if organizationID == 0 {
		return nil, fmt.Errorf("sm: %w", model.ErrNoOrganizationID)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization", serviceDeskID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"organizationId": organizationID})
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalOrganizationImpl) Detach(ctx context.Context, serviceDeskID, organizationID int) (*model.ResponseScheme, error) {

	if serviceDeskID == 0 {
		return nil, fmt.Errorf("sm: %w", model.ErrNoServiceDeskID)
	}

	if organizationID == 0 {
		return nil, fmt.Errorf("sm: %w", model.ErrNoOrganizationID)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization", serviceDeskID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", map[string]interface{}{"organizationId": organizationID})
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}
