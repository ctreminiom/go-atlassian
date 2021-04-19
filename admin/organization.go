package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type OrganizationService struct {
	client *Client
	Policy *OrganizationPolicyService
}

// Returns a list of your organizations, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. cursor = the next pagination result, The cursor is not a number that you can increment through predictably.
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-organizations
func (o *OrganizationService) Gets(ctx context.Context, cursor string) (result *OrganizationPageScheme, response *Response, err error) {

	var endpoint string
	if len(cursor) != 0 {
		endpoint = fmt.Sprintf("/admin/v1/orgs?cursor=%v", cursor)
	} else {
		endpoint = "/admin/v1/orgs"
	}

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns information about a single organization by ID, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-organization-by-id
func (o *OrganizationService) Get(ctx context.Context, organizationID string) (result *OrganizationScheme, response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
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

// Returns a list of users in an organization, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. cursor = the next pagination result, The cursor is not a number that you can increment through predictably.
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-users-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-users-in-an-organization
func (o *OrganizationService) Users(ctx context.Context, organizationID, cursor string) (result *OrganizationUserPageScheme, response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	var endpoint string
	if len(cursor) != 0 {
		endpoint = fmt.Sprintf("/admin/v1/orgs/%v/users?cursor=%v", organizationID, cursor)
	} else {
		endpoint = fmt.Sprintf("/admin/v1/orgs/%v/users", organizationID)
	}

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationUserPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
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

// Returns a list of domains in an organization one page at a time, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. cursor = the next pagination result, The cursor is not a number that you can increment through predictably.
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-domains-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domains-in-an-organization
func (o *OrganizationService) Domains(ctx context.Context, organizationID, cursor string) (result *OrganizationDomainPageScheme, response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	var endpoint string
	if len(cursor) != 0 {
		endpoint = fmt.Sprintf("/admin/v1/orgs/%v/domains?cursor=%v", organizationID, cursor)
	} else {
		endpoint = fmt.Sprintf("/admin/v1/orgs/%v/domains", organizationID)
	}

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationDomainPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
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

// Returns information about a single verified domain by ID, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. domainID = ID of the domain to return (REQUIRED)
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-domains-domainid-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domain-by-id
func (o *OrganizationService) Domain(ctx context.Context, organizationID, domainID string) (result *OrganizationDomainScheme, response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	if len(domainID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid domainID value")
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/domains/%v", organizationID, domainID)

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationDomainScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
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

// Returns an audit log of events from an organization one page at a time, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. opts = the Event options
// 4. cursor = the next pagination result, The cursor is not a number that you can increment through predictably.
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-events-get
// Library Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-audit-log-of-events
func (o *OrganizationService) Events(ctx context.Context, organizationID string, opts *OrganizationEventOptScheme, cursor string) (result *OrganizationEventPageScheme, response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	params := url.Values{}
	if opts != nil {

		if !opts.To.IsZero() {
			timeAsEpoch := int(opts.To.Unix())
			params.Add("to", strconv.Itoa(timeAsEpoch))
		}

		if !opts.From.IsZero() {
			timeAsEpoch := int(opts.From.Unix())
			params.Add("from", strconv.Itoa(timeAsEpoch))
		}

		if len(opts.Q) != 0 {
			params.Add("q", opts.Q)
		}

		if len(opts.Action) != 0 {
			params.Add("action", opts.Action)
		}

	}

	if len(cursor) != 0 {
		params.Add("cursor", cursor)
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("/admin/v1/orgs/%v/events?%v", organizationID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("/admin/v1/orgs/%v/events", organizationID)
	}

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationEventPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
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

// Returns information about a single event by ID.
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. eventID = ID of the event to return (REQUIRED)
// Library Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-event-by-id
func (o *OrganizationService) Event(ctx context.Context, organizationID, eventID string) (result *OrganizationEventScheme, response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	if len(eventID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid eventID value")
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/events/%v", organizationID, eventID)

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationEventScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type OrganizationEventScheme struct {
	Data *OrganizationEventModelScheme `json:"data,omitempty"`
}

// Returns information localized event actions
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-event-actions-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-list-of-event-actions
func (o *OrganizationService) Actions(ctx context.Context, organizationID string) (result *OrganizationEventActionScheme, response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/event-actions", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationEventActionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
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
