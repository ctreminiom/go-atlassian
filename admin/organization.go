package admin

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type OrganizationService struct {
	client *Client
	Policy *OrganizationPolicyService
}

// Gets returns a list of your organizations
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-organizations
func (o *OrganizationService) Gets(ctx context.Context, cursor string) (result *model.AdminOrganizationPageScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	if cursor != "" {
		params.Add("cursor", cursor)
	}

	var endpoint strings.Builder
	endpoint.WriteString("/admin/v1/orgs")

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns information about a single organization by ID
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-organization-by-id
func (o *OrganizationService) Get(ctx context.Context, organizationID string) (result *model.AdminOrganizationScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Users returns a list of users in an organization
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-users-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-users-in-an-organization
func (o *OrganizationService) Users(ctx context.Context, organizationID, cursor string) (result *model.OrganizationUserPageScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	params := url.Values{}
	if cursor != "" {
		params.Add("cursor", cursor)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/admin/v1/orgs/%v/users", organizationID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Domains returns a list of domains in an organization one page at a time
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-domains-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domains-in-an-organization
func (o *OrganizationService) Domains(ctx context.Context, organizationID, cursor string) (result *model.OrganizationDomainPageScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	params := url.Values{}
	if cursor != "" {
		params.Add("cursor", cursor)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/admin/v1/orgs/%v/domains", organizationID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Domain returns information about a single verified domain by ID
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-domains-domainid-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domain-by-id
func (o *OrganizationService) Domain(ctx context.Context, organizationID, domainID string) (result *model.OrganizationDomainScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	if len(domainID) == 0 {
		return nil, nil, model.ErrNoAdminDomainIDError
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/domains/%v", organizationID, domainID)

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Events returns an audit log of events from an organization one page at a time
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-events-get
// Library Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-audit-log-of-events
func (o *OrganizationService) Events(ctx context.Context, organizationID string, options *model.OrganizationEventOptScheme,
	cursor string) (result *model.OrganizationEventPageScheme, response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	params := url.Values{}
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

	if len(cursor) != 0 {
		params.Add("cursor", cursor)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/admin/v1/orgs/%v/events", organizationID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Event returns information about a single event by ID.
// Library Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-event-by-id
func (o *OrganizationService) Event(ctx context.Context, organizationID, eventID string) (result *model.OrganizationEventScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	if len(eventID) == 0 {
		return nil, nil, model.ErrNoEventIDError
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/events/%v", organizationID, eventID)

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Actions returns information localized event actions
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-event-actions-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-list-of-event-actions
func (o *OrganizationService) Actions(ctx context.Context, organizationID string) (result *model.OrganizationEventActionScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/event-actions", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
