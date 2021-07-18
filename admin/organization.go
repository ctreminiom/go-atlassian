package admin

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type OrganizationService struct {
	client *Client
	Policy *OrganizationPolicyService
}

// Gets returns a list of your organizations
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-organizations
func (o *OrganizationService) Gets(ctx context.Context, cursor string) (result *OrganizationPageScheme,
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
func (o *OrganizationService) Get(ctx context.Context, organizationID string) (result *OrganizationScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, notOrganizationError
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

type OrganizationPageScheme struct {
	Data  []*OrganizationModelScheme `json:"data,omitempty"`
	Links *LinkPageModelScheme       `json:"links,omitempty"`
}

type LinkPageModelScheme struct {
	Self string `json:"self,omitempty"`
	Prev string `json:"prev,omitempty"`
	Next string `json:"next,omitempty"`
}

type OrganizationModelScheme struct {
	ID            string                          `json:"id,omitempty"`
	Type          string                          `json:"type,omitempty"`
	Attributes    *OrganizationModelAttribute     `json:"attributes,omitempty"`
	Relationships *OrganizationModelRelationships `json:"relationships,omitempty"`
	Links         *LinkSelfModelScheme            `json:"links,omitempty"`
}

type OrganizationModelAttribute struct {
	Name string `json:"name,omitempty"`
}

type OrganizationModelRelationships struct {
	Domains *OrganizationModelSchemes `json:"domains,omitempty"`
	Users   *OrganizationModelSchemes `json:"users,omitempty"`
}

type OrganizationModelSchemes struct {
	Links struct {
		Related string `json:"related,omitempty"`
	} `json:"links,omitempty"`
}

type LinkSelfModelScheme struct {
	Self string `json:"self,omitempty"`
}

type OrganizationScheme struct {
	Data *OrganizationModelScheme `json:"data,omitempty"`
}

// Users returns a list of users in an organization
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-users-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-users-in-an-organization
func (o *OrganizationService) Users(ctx context.Context, organizationID, cursor string) (result *OrganizationUserPageScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, notOrganizationError
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

type OrganizationUserPageScheme struct {
	Data  []*OrganizationUserScheme `json:"data,omitempty"`
	Links *LinkPageModelScheme      `json:"links,omitempty"`
	Meta  struct {
		Total int `json:"total,omitempty"`
	} `json:"meta,omitempty"`
}

type OrganizationUserScheme struct {
	AccountID      string                           `json:"account_id,omitempty"`
	AccountType    string                           `json:"account_type,omitempty"`
	AccountStatus  string                           `json:"account_status,omitempty"`
	Name           string                           `json:"name,omitempty"`
	Picture        string                           `json:"picture,omitempty"`
	Email          string                           `json:"email,omitempty"`
	AccessBillable bool                             `json:"access_billable,omitempty"`
	LastActive     string                           `json:"last_active,omitempty"`
	ProductAccess  []*OrganizationUserProductScheme `json:"product_access,omitempty"`
	Links          *LinkSelfModelScheme             `json:"links,omitempty"`
}

type OrganizationUserProductScheme struct {
	Key        string `json:"key,omitempty"`
	Name       string `json:"name,omitempty"`
	URL        string `json:"url,omitempty"`
	LastActive string `json:"last_active,omitempty"`
}

// Domains returns a list of domains in an organization one page at a time
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-domains-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domains-in-an-organization
func (o *OrganizationService) Domains(ctx context.Context, organizationID, cursor string) (result *OrganizationDomainPageScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, notOrganizationError
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

type OrganizationDomainPageScheme struct {
	Data  []*OrganizationDomainModelScheme `json:"data,omitempty"`
	Links *LinkPageModelScheme             `json:"links,omitempty"`
}

type OrganizationDomainModelScheme struct {
	ID         string `json:"id,omitempty"`
	Type       string `json:"type,omitempty"`
	Attributes struct {
		Name  string `json:"name,omitempty"`
		Claim struct {
			Type   string `json:"type,omitempty"`
			Status string `json:"status,omitempty"`
		} `json:"claim,omitempty"`
	} `json:"attributes,omitempty"`
	Links *LinkSelfModelScheme `json:"links,omitempty"`
}

// Domain returns information about a single verified domain by ID
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-domains-domainid-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domain-by-id
func (o *OrganizationService) Domain(ctx context.Context, organizationID, domainID string) (result *OrganizationDomainScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, notOrganizationError
	}

	if len(domainID) == 0 {
		return nil, nil, notDomainError
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

type OrganizationDomainScheme struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name  string `json:"name"`
			Claim struct {
				Type   string `json:"type"`
				Status string `json:"status"`
			} `json:"claim"`
		} `json:"attributes"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
}

type OrganizationEventOptScheme struct {
	Q      string    //Single query term for searching events.
	From   time.Time //The earliest date and time of the event represented as a UNIX epoch time.
	To     time.Time //The latest date and time of the event represented as a UNIX epoch time.
	Action string    //A query filter that returns events of a specific action type.
}

// Events returns an audit log of events from an organization one page at a time
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-events-get
// Library Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-audit-log-of-events
func (o *OrganizationService) Events(ctx context.Context, organizationID string, options *OrganizationEventOptScheme,
	cursor string) (result *OrganizationEventPageScheme, response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, notOrganizationError
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

type OrganizationEventPageScheme struct {
	Data  []*OrganizationEventModelScheme `json:"data,omitempty"`
	Links *LinkPageModelScheme            `json:"links,omitempty"`
	Meta  struct {
		Next     string `json:"next,omitempty"`
		PageSize int    `json:"page_size,omitempty"`
	} `json:"meta,omitempty"`
}

type OrganizationEventModelScheme struct {
	ID         string                                  `json:"id,omitempty"`
	Type       string                                  `json:"type,omitempty"`
	Attributes *OrganizationEventModelAttributesScheme `json:"attributes,omitempty"`
	Links      *LinkSelfModelScheme                    `json:"links,omitempty"`
}

type OrganizationEventModelAttributesScheme struct {
	Time      string                          `json:"time,omitempty"`
	Action    string                          `json:"action,omitempty"`
	Actor     *OrganizationEventActorModel    `json:"actor,omitempty"`
	Context   []*OrganizationEventObjectModel `json:"context,omitempty"`
	Container []*OrganizationEventObjectModel `json:"container,omitempty"`
	Location  *OrganizationEventLocationModel `json:"location,omitempty"`
}

type OrganizationEventActorModel struct {
	ID    string               `json:"id,omitempty"`
	Name  string               `json:"name,omitempty"`
	Links *LinkSelfModelScheme `json:"links,omitempty"`
}

type OrganizationEventObjectModel struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
		Alt  string `json:"alt,omitempty"`
	} `json:"links,omitempty"`
}

type OrganizationEventLocationModel struct {
	IP  string `json:"ip,omitempty"`
	Geo string `json:"geo,omitempty"`
}

// Event returns information about a single event by ID.
// Library Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-event-by-id
func (o *OrganizationService) Event(ctx context.Context, organizationID, eventID string) (result *OrganizationEventScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, notOrganizationError
	}

	if len(eventID) == 0 {
		return nil, nil, notEventError
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

type OrganizationEventScheme struct {
	Data *OrganizationEventModelScheme `json:"data,omitempty"`
}

// Actions returns information localized event actions
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-event-actions-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-list-of-event-actions
func (o *OrganizationService) Actions(ctx context.Context, organizationID string) (result *OrganizationEventActionScheme,
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, notOrganizationError
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

type OrganizationEventActionScheme struct {
	Data []*OrganizationEventActionModelScheme `json:"data,omitempty"`
}

type OrganizationEventActionModelScheme struct {
	ID         string                                        `json:"id,omitempty"`
	Type       string                                        `json:"type,omitempty"`
	Attributes *OrganizationEventActionModelAttributesScheme `json:"attributes,omitempty"`
}

type OrganizationEventActionModelAttributesScheme struct {
	DisplayName      string `json:"displayName,omitempty"`
	GroupDisplayName string `json:"groupDisplayName,omitempty"`
}

var (
	notOrganizationError = fmt.Errorf("error!, please provide a valid organizationID value")
	notDomainError       = fmt.Errorf("error!, please provide a valid domainID value")
	notEventError        = fmt.Errorf("error!, please provide a valid eventID value")
)
