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
// Library Example: N/A
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
// Library Example: N/A
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
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name string `json:"name"`
		} `json:"attributes"`
		Relationships struct {
			Domains struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"domains"`
			Users struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"users"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
}

type OrganizationScheme struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name string `json:"name"`
		} `json:"attributes"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Relationships struct {
			Domains struct {
				Related string `json:"related"`
			} `json:"domains"`
			Users struct {
				Related string `json:"related"`
			} `json:"users"`
		} `json:"relationships"`
	} `json:"data"`
}

// Returns a list of users in an organization, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. cursor = the next pagination result, The cursor is not a number that you can increment through predictably.
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-users-get
// Library Example: N/A
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
	Data []struct {
		AccountID      string `json:"account_id"`
		AccountType    string `json:"account_type"`
		AccountStatus  string `json:"account_status"`
		Name           string `json:"name"`
		Picture        string `json:"picture"`
		Email          string `json:"email"`
		AccessBillable bool   `json:"access_billable"`
		LastActive     string `json:"last_active"`
		ProductAccess  []struct {
			Key        string `json:"key"`
			Name       string `json:"name"`
			URL        string `json:"url"`
			LastActive string `json:"last_active"`
		} `json:"product_access"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Meta struct {
		Total int `json:"total"`
	} `json:"meta"`
	Links struct {
		Self string `json:"self"`
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
}

// Returns a list of domains in an organization one page at a time, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. cursor = the next pagination result, The cursor is not a number that you can increment through predictably.
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-domains-get
// Library Example: N/A
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
	Data []struct {
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
	Links struct {
		Self string `json:"self"`
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
}

// Returns information about a single verified domain by ID, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. domainID = ID of the domain to return (REQUIRED)
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-domains-domainid-get
// Library Example: N/A
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
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Time   string `json:"time"`
			Action string `json:"action"`
			Actor  struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Links struct {
					Self string `json:"self"`
				} `json:"links"`
			} `json:"actor"`
			Context []struct {
				ID         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
				} `json:"attributes"`
				Links struct {
					Self string `json:"self"`
					Alt  string `json:"alt"`
				} `json:"links"`
			} `json:"context"`
			Container []struct {
				ID         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
				} `json:"attributes"`
				Links struct {
					Self string `json:"self"`
					Alt  string `json:"alt"`
				} `json:"links"`
			} `json:"container"`
			Location struct {
				IP  string `json:"ip"`
				Geo string `json:"geo"`
			} `json:"location"`
		} `json:"attributes"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Meta struct {
		Next     string `json:"next"`
		PageSize int    `json:"page_size"`
	} `json:"meta"`
	Links struct {
		Self string `json:"self"`
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
}

// Returns information about a single event by ID.
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. eventID = ID of the event to return (REQUIRED)
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
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Time   string `json:"time"`
			Action string `json:"action"`
			Actor  struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Links struct {
					Self string `json:"self"`
				} `json:"links"`
			} `json:"actor"`
			Context []struct {
				ID         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
				} `json:"attributes"`
				Links struct {
					Self string `json:"self"`
					Alt  string `json:"alt"`
				} `json:"links"`
			} `json:"context"`
			Container []struct {
				ID         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
				} `json:"attributes"`
				Links struct {
					Self string `json:"self"`
					Alt  string `json:"alt"`
				} `json:"links"`
			} `json:"container"`
			Location struct {
				IP  string `json:"ip"`
				Geo string `json:"geo"`
			} `json:"location"`
		} `json:"attributes"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
}

// Returns information localized event actions
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-event-actions-get
// Library Example: N/A
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
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			DisplayName      string `json:"displayName"`
			GroupDisplayName string `json:"groupDisplayName"`
		} `json:"attributes"`
	} `json:"data"`
}
