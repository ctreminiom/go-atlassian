package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/admin"
)

// NewOrganizationService creates a new instance of the OrganizationService.
func NewOrganizationService(client service.Connector, policy *OrganizationPolicyService, directory *OrganizationDirectoryService) *OrganizationService {
	return &OrganizationService{internalClient: &internalOrganizationImpl{c: client}, Policy: policy, Directory: directory}
}

// OrganizationService handles communication with the organization related methods of the Atlassian Admin API.
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
// GET /admin/v1/orgs/{organizationID}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-organization-by-id
func (o *OrganizationService) Get(ctx context.Context, organizationID string) (*model.AdminOrganizationScheme, *model.ResponseScheme, error) {
	return o.internalClient.Get(ctx, organizationID)
}

// Users returns a list of users in an organization.
//
// GET /admin/v1/orgs/{organizationID}/users
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-users-in-an-organization
func (o *OrganizationService) Users(ctx context.Context, organizationID, cursor string) (*model.OrganizationUserPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Users(ctx, organizationID, cursor)
}

// Domains returns a list of domains in an organization one page at a time.
//
// GET /admin/v1/orgs/{organizationID}/domains
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domains-in-an-organization
func (o *OrganizationService) Domains(ctx context.Context, organizationID, cursor string) (*model.OrganizationDomainPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Domains(ctx, organizationID, cursor)
}

// Domain returns information about a single verified domain by ID.
//
// GET /admin/v1/orgs/{organizationID}/domains/{domainID}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domain-by-id
func (o *OrganizationService) Domain(ctx context.Context, organizationID, domainID string) (*model.OrganizationDomainScheme, *model.ResponseScheme, error) {
	return o.internalClient.Domain(ctx, organizationID, domainID)
}

// Events returns an audit log of events from an organization one page at a time.
//
// GET /admin/v1/orgs/{organizationID}/events
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-audit-log-of-events
func (o *OrganizationService) Events(ctx context.Context, organizationID string, options *model.OrganizationEventOptScheme, cursor string) (*model.OrganizationEventPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Events(ctx, organizationID, options, cursor)
}

// Event returns information about a single event by ID.
//
// GET /admin/v1/orgs/{organizationID}/events/{eventID}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-event-by-id
func (o *OrganizationService) Event(ctx context.Context, organizationID, eventID string) (*model.OrganizationEventScheme, *model.ResponseScheme, error) {
	return o.internalClient.Event(ctx, organizationID, eventID)
}

// Actions returns information localized event actions.
//
// GET /admin/v1/orgs/{organizationID}/event-actions
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-list-of-event-actions
func (o *OrganizationService) Actions(ctx context.Context, organizationID string) (*model.OrganizationEventActionScheme, *model.ResponseScheme, error) {
	return o.internalClient.Actions(ctx, organizationID)
}

// EventsStream returns a paginated list of audit log events from an organization using the /events-stream endpoint.
//
// GET /admin/v1/orgs/{organizationID}/events-stream
//
// https://developer.atlassian.com/cloud/admin/organization/rest/api-group-events/#api-v1-orgs-orgid-events-stream-get
func (o *OrganizationService) EventsStream(ctx context.Context, organizationID string, options *model.OrganizationEventStreamOptScheme) (*model.OrganizationEventStreamPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.EventsStream(ctx, organizationID, options)
}

type internalOrganizationImpl struct {
	c service.Connector
}

func (i *internalOrganizationImpl) Gets(ctx context.Context, cursor string) (*model.AdminOrganizationPageScheme, *model.ResponseScheme, error) {

	var endpoint strings.Builder
	endpoint.WriteString("admin/v1/orgs")

	if cursor != "" {
		params := url.Values{}
		params.Add("cursor", cursor)

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	organizations := new(model.AdminOrganizationPageScheme)
	res, err := i.c.Call(req, organizations)
	if err != nil {
		return nil, res, err
	}

	return organizations, res, nil
}

func (i *internalOrganizationImpl) Get(ctx context.Context, organizationID string) (*model.AdminOrganizationScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganization
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v", organizationID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	organization := new(model.AdminOrganizationScheme)
	res, err := i.c.Call(req, organization)
	if err != nil {
		return nil, res, err
	}

	return organization, res, nil
}

func (i *internalOrganizationImpl) Users(ctx context.Context, organizationID, cursor string) (*model.OrganizationUserPageScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganization
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("admin/v1/orgs/%v/users", organizationID))

	if cursor != "" {
		params := url.Values{}
		params.Add("cursor", cursor)

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	users := new(model.OrganizationUserPageScheme)
	res, err := i.c.Call(req, users)
	if err != nil {
		return nil, res, err
	}

	return users, res, nil
}

func (i *internalOrganizationImpl) Domains(ctx context.Context, organizationID, cursor string) (*model.OrganizationDomainPageScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganization
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("admin/v1/orgs/%v/domains", organizationID))

	if cursor != "" {
		params := url.Values{}
		params.Add("cursor", cursor)

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	domains := new(model.OrganizationDomainPageScheme)
	res, err := i.c.Call(req, domains)
	if err != nil {
		return nil, res, err
	}

	return domains, res, nil
}

func (i *internalOrganizationImpl) Domain(ctx context.Context, organizationID, domainID string) (*model.OrganizationDomainScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganization
	}

	if domainID == "" {
		return nil, nil, model.ErrNoAdminDomainID
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/domains/%v", organizationID, domainID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	domain := new(model.OrganizationDomainScheme)
	res, err := i.c.Call(req, domain)
	if err != nil {
		return nil, res, err
	}

	return domain, res, nil
}

func (i *internalOrganizationImpl) Events(ctx context.Context, organizationID string, options *model.OrganizationEventOptScheme, cursor string) (*model.OrganizationEventPageScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganization
	}

	params := url.Values{}
	if cursor != "" {
		params.Add("cursor", cursor)
	}

	if options != nil {

		if !options.To.IsZero() {
			timeAsEpoch := int(options.To.UnixMilli())
			params.Add("to", strconv.Itoa(timeAsEpoch))
		}

		if !options.From.IsZero() {
			timeAsEpoch := int(options.From.UnixMilli())
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

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	events := new(model.OrganizationEventPageScheme)
	res, err := i.c.Call(req, events)
	if err != nil {
		return nil, res, err
	}

	return events, res, nil
}

func (i *internalOrganizationImpl) Event(ctx context.Context, organizationID, eventID string) (*model.OrganizationEventScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganization
	}

	if eventID == "" {
		return nil, nil, model.ErrNoEventID
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/events/%v", organizationID, eventID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	event := new(model.OrganizationEventScheme)
	res, err := i.c.Call(req, event)
	if err != nil {
		return nil, res, err
	}

	return event, res, nil
}

func (i *internalOrganizationImpl) Actions(ctx context.Context, organizationID string) (*model.OrganizationEventActionScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganization
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/event-actions", organizationID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	event := new(model.OrganizationEventActionScheme)
	res, err := i.c.Call(req, event)
	if err != nil {
		return nil, res, err
	}

	return event, res, nil
}

func (i *internalOrganizationImpl) EventsStream(ctx context.Context, organizationID string, options *model.OrganizationEventStreamOptScheme) (*model.OrganizationEventStreamPageScheme, *model.ResponseScheme, error) {
	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganization
	}

	params := url.Values{}
	if options != nil {
		if !options.From.IsZero() {
			params.Add("from", strconv.FormatInt(options.From.UnixMilli(), 10))
		}
		if !options.To.IsZero() {
			params.Add("to", strconv.FormatInt(options.To.UnixMilli(), 10))
		}
		if options.Cursor != "" {
			params.Add("cursor", options.Cursor)
		}
		if options.SortOrder != "" {
			params.Add("sortOrder", options.SortOrder)
		}
		if options.Limit > 0 {
			params.Add("limit", strconv.Itoa(options.Limit))
		}
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("admin/v1/orgs/%v/events-stream", organizationID))
	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.OrganizationEventStreamPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}
