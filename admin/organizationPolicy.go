package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type OrganizationPolicyService struct {
	client   *Client
	Resource *OrganizationPolicyResourceService
}

// Returns information about org policies, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. policyType = Sets the type for the page of policies to return.
// 3. cursor = the next pagination result, The cursor is not a number that you can increment through predictably.
// Atlassian Docs: Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-policies-get
// Library Docs: N/A
func (o *OrganizationPolicyService) Gets(ctx context.Context, organizationID, policyType, cursor string) (result *OrganizationPolicyPageScheme, response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	params := url.Values{}
	if len(policyType) != 0 {
		params.Add("type", policyType)
	}

	if len(cursor) != 0 {
		params.Add("cursor", cursor)
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("/admin/v1/orgs/%v/policies?%v", organizationID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("/admin/v1/orgs/%v/policies", organizationID)
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

	result = new(OrganizationPolicyPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type OrganizationPolicyPageScheme struct {
	Data []*OrganizationPolicyData `json:"data"`
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

// Returns information about a single policy by ID, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 2. policyID = ID of the policy to query (REQUIRED)
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-policies-policyid-get
// Example Library: N/A
func (o *OrganizationPolicyService) Get(ctx context.Context, organizationID, policyID string) (result *OrganizationPolicyScheme, response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	if len(policyID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid policyID value")
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/policies/%v", organizationID, policyID)

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationPolicyScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type OrganizationPolicyScheme struct {
	Data OrganizationPolicyData `json:"data"`
}
type OrganizationPolicyResource struct {
	ID                string `json:"id,omitempty"`
	ApplicationStatus string `json:"applicationStatus,omitempty"`
}
type OrganizationPolicyAttributes struct {
	Type      string                        `json:"type,omitempty"`
	Name      string                        `json:"name,omitempty"`
	Status    string                        `json:"status,omitempty"`
	Resources []*OrganizationPolicyResource `json:"resources,omitempty"`
	CreatedAt time.Time                     `json:"createdAt,omitempty"`
	UpdatedAt time.Time                     `json:"updatedAt,omitempty"`
}
type OrganizationPolicyData struct {
	ID         string                       `json:"id,omitempty"`
	Type       string                       `json:"type,omitempty"`
	Attributes OrganizationPolicyAttributes `json:"attributes,omitempty"`
}

// Create a policy for an org, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 2. payload = PolicyCreateModel (REQUIRED)
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-policies-post
// Library Example: N/A
func (o *OrganizationPolicyService) Create(ctx context.Context, organizationID string, payload *OrganizationPolicyData) (result *OrganizationPolicyScheme, response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	if payload == nil {
		return nil, nil, fmt.Errorf("error!, please provide a valid OrganizationPolicyScheme pointer")
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/policies", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationPolicyScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Update a policy for an org, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. payload = PolicyCreateModel (REQUIRED)
// 4. policyID = ID of the policy to update (REQUIRED)
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-policies-policyid-put
func (o *OrganizationPolicyService) Update(ctx context.Context, organizationID, policyID string, payload *OrganizationPolicyData) (result *OrganizationPolicyScheme, response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	if len(policyID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid policyID value")
	}

	if payload == nil {
		return nil, nil, fmt.Errorf("error!, please provide a valid OrganizationPolicyData pointer")
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/policies/%v", organizationID, policyID)

	request, err := o.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationPolicyScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Delete a policy for an org, this func needs the following parameters:
// 1. ctx = it's the context.context value
// 2. organizationID = ID of the organization to return (REQUIRED)
// 3. policyID = ID of the policy to update (REQUIRED)
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-policies-policyid-delete
func (o *OrganizationPolicyService) Delete(ctx context.Context, organizationID, policyID string) (response *Response, err error) {

	if len(organizationID) == 0 {
		return nil, fmt.Errorf("error!, please provide a valid organizationID value")
	}

	if len(policyID) == 0 {
		return nil, fmt.Errorf("error!, please provide a valid policyID value")
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/policies/%v", organizationID, policyID)

	request, err := o.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	return
}
