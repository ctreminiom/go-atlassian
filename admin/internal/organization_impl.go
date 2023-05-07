package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/admin"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewOrganizationService(client service.Client, policy *OrganizationPolicyService, directory *OrganizationDirectoryService) *OrganizationService {
	return &OrganizationService{internalClient: &internalOrganizationImpl{c: client}, Policy: policy, Directory: directory}
}

type OrganizationService struct {
	internalClient admin.OrganizationConnector
	Policy         *OrganizationPolicyService
	Directory      *OrganizationDirectoryService
}

// Gets returns a list of your organizations (based on your API key).
//
// GET /admin/v1/orgs
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-organizations
func (o *OrganizationService) Gets(ctx context.Context, cursor string) (*model.AdminOrganizationPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Gets(ctx, cursor)
}

// Get returns information about a single organization by ID
//
// GET /admin/v1/orgs/{orgId}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-organization-by-id
func (o *OrganizationService) Get(ctx context.Context, organizationID string) (*model.AdminOrganizationScheme, *model.ResponseScheme, error) {
	return o.internalClient.Get(ctx, organizationID)
}

// Users returns a list of users in an organization.
//
// GET /admin/v1/orgs/{orgId}/users
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-users-in-an-organization
func (o *OrganizationService) Users(ctx context.Context, organizationID, cursor string) (*model.OrganizationUserPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Users(ctx, organizationID, cursor)
}

// Domains returns a list of domains in an organization one page at a time.
//
// GET /admin/v1/orgs/{orgId}/domains
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domains-in-an-organization
func (o *OrganizationService) Domains(ctx context.Context, organizationID, cursor string) (*model.OrganizationDomainPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Domains(ctx, organizationID, cursor)
}

// Domain returns information about a single verified domain by ID.
//
// GET /admin/v1/orgs/{orgId}/domains/{domainId}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domain-by-id
func (o *OrganizationService) Domain(ctx context.Context, organizationID, domainID string) (*model.OrganizationDomainScheme, *model.ResponseScheme, error) {
	return o.internalClient.Domain(ctx, organizationID, domainID)
}

// Events returns an audit log of events from an organization one page at a time.
//
// GET /admin/v1/orgs/{orgId}/events
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-audit-log-of-events
func (o *OrganizationService) Events(ctx context.Context, organizationID string, options *model.OrganizationEventOptScheme, cursor string) (*model.OrganizationEventPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Events(ctx, organizationID, options, cursor)
}

// Event returns information about a single event by ID.
//
// GET /admin/v1/orgs/{orgId}/events/{eventId}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-event-by-id
func (o *OrganizationService) Event(ctx context.Context, organizationID, eventID string) (*model.OrganizationEventScheme, *model.ResponseScheme, error) {
	return o.internalClient.Event(ctx, organizationID, eventID)
}

// Actions returns information localized event actions.
//
// GET /admin/v1/orgs/{orgId}/event-actions
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-list-of-event-actions
func (o *OrganizationService) Actions(ctx context.Context, organizationID string) (*model.OrganizationEventActionScheme, *model.ResponseScheme, error) {
	return o.internalClient.Actions(ctx, organizationID)
}

type internalOrganizationImpl struct {
	c service.Client
}

func (i *internalOrganizationImpl) Gets(ctx context.Context, cursor string) (*model.AdminOrganizationPageScheme, *model.ResponseScheme, error) {

	var endpoint strings.Builder
	endpoint.WriteString("admin/v1/orgs")

	if cursor != "" {
		params := url.Values{}
		params.Add("cursor", cursor)

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	organizations := new(model.AdminOrganizationPageScheme)
	response, err := i.c.Call(request, organizations)
	if err != nil {
		return nil, response, err
	}

	return organizations, response, nil
}

func (i *internalOrganizationImpl) Get(ctx context.Context, organizationID string) (*model.AdminOrganizationScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v", organizationID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	organization := new(model.AdminOrganizationScheme)
	response, err := i.c.Call(request, organization)
	if err != nil {
		return nil, response, err
	}

	return organization, response, nil
}

func (i *internalOrganizationImpl) Users(ctx context.Context, organizationID, cursor string) (*model.OrganizationUserPageScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("admin/v1/orgs/%v/users", organizationID))

	if cursor != "" {
		params := url.Values{}
		params.Add("cursor", cursor)

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	users := new(model.OrganizationUserPageScheme)
	response, err := i.c.Call(request, users)
	if err != nil {
		return nil, response, err
	}

	return users, response, nil
}

func (i *internalOrganizationImpl) Domains(ctx context.Context, organizationID, cursor string) (*model.OrganizationDomainPageScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("admin/v1/orgs/%v/domains", organizationID))

	if cursor != "" {
		params := url.Values{}
		params.Add("cursor", cursor)

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	domains := new(model.OrganizationDomainPageScheme)
	response, err := i.c.Call(request, domains)
	if err != nil {
		return nil, response, err
	}

	return domains, response, nil
}

func (i *internalOrganizationImpl) Domain(ctx context.Context, organizationID, domainID string) (*model.OrganizationDomainScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	if domainID == "" {
		return nil, nil, model.ErrNoAdminDomainIDError
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/domains/%v", organizationID, domainID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	domain := new(model.OrganizationDomainScheme)
	response, err := i.c.Call(request, domain)
	if err != nil {
		return nil, response, err
	}

	return domain, response, nil
}

func (i *internalOrganizationImpl) Events(ctx context.Context, organizationID string, options *model.OrganizationEventOptScheme, cursor string) (*model.OrganizationEventPageScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	params := url.Values{}
	if cursor != "" {
		params.Add("cursor", cursor)
	}

	if options != nil {

		if !options.To.IsZero() {
			timeAsEpoch := int(options.To.Unix())
			params.Add("to", strconv.Itoa(timeAsEpoch))
		}

		if !options.From.IsZero() {
			timeAsEpoch := int(options.From.Unix())
			params.Add("from", strconv.Itoa(timeAsEpoch))
		}

		if len(options.Q) != 0 {
			params.Add("q", options.Q)
		}

		if len(options.Action) != 0 {
			params.Add("action", options.Action)
		}
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("admin/v1/orgs/%v/events", organizationID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	events := new(model.OrganizationEventPageScheme)
	response, err := i.c.Call(request, events)
	if err != nil {
		return nil, response, err
	}

	return events, response, nil
}

func (i *internalOrganizationImpl) Event(ctx context.Context, organizationID, eventID string) (*model.OrganizationEventScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	if eventID == "" {
		return nil, nil, model.ErrNoEventIDError
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/events/%v", organizationID, eventID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	event := new(model.OrganizationEventScheme)
	response, err := i.c.Call(request, event)
	if err != nil {
		return nil, response, err
	}

	return event, response, nil
}

func (i *internalOrganizationImpl) Actions(ctx context.Context, organizationID string) (*model.OrganizationEventActionScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/event-actions", organizationID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	event := new(model.OrganizationEventActionScheme)
	response, err := i.c.Call(request, event)
	if err != nil {
		return nil, response, err
	}

	return event, response, nil
}
