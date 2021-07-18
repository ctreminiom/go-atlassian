package admin

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type OrganizationPolicyService struct {
	client *Client
}

// Gets returns information about org policies
// Atlassian Docs: Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-policies-get
// Library Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#get-list-of-policies
func (o *OrganizationPolicyService) Gets(ctx context.Context, organizationID, policyType, cursor string) (
	result *OrganizationPolicyPageScheme, response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, notOrganizationError
	}

	params := url.Values{}

	if len(policyType) != 0 {
		params.Add("type", policyType)
	}

	if len(cursor) != 0 {
		params.Add("cursor", cursor)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/admin/v1/orgs/%v/policies", organizationID))

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

type OrganizationPolicyPageScheme struct {
	Data  []*OrganizationPolicyData `json:"data"`
	Links *LinkPageModelScheme      `json:"links"`
	Meta  struct {
		Next     string `json:"next"`
		PageSize int    `json:"page_size"`
	} `json:"meta"`
}

// Get information about a single policy by ID
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-policies-policyid-get
// Example Library: https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#get-a-policy-by-id
func (o *OrganizationPolicyService) Get(ctx context.Context, organizationID, policyID string) (
	result *OrganizationPolicyScheme, response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, notOrganizationError
	}

	if len(policyID) == 0 {
		return nil, nil, notPolicyError
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/policies/%v", organizationID, policyID)

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
	ID         string                        `json:"id,omitempty"`
	Type       string                        `json:"type,omitempty"`
	Attributes *OrganizationPolicyAttributes `json:"attributes,omitempty"`
}

// Create a policy for an org
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-policies-post
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#create-a-policy
func (o *OrganizationPolicyService) Create(ctx context.Context, organizationID string, payload *OrganizationPolicyData) (
	result *OrganizationPolicyScheme, response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, notOrganizationError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/policies", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update a policy for an org
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-policies-policyid-put
// Library Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#update-a-policy
func (o *OrganizationPolicyService) Update(ctx context.Context, organizationID, policyID string,
	payload *OrganizationPolicyData) (result *OrganizationPolicyScheme, response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, nil, notOrganizationError
	}

	if len(policyID) == 0 {
		return nil, nil, notPolicyError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/policies/%v", organizationID, policyID)

	request, err := o.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete a policy for an org
// Official Docs: https://developer.atlassian.com/cloud/admin/organization/rest/api-group-orgs/#api-orgs-orgid-policies-policyid-delete
// Library Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#delete-a-policy
func (o *OrganizationPolicyService) Delete(ctx context.Context, organizationID, policyID string) (
	response *ResponseScheme, err error) {

	if len(organizationID) == 0 {
		return nil, notOrganizationError
	}

	if len(policyID) == 0 {
		return nil, notPolicyError
	}

	var endpoint = fmt.Sprintf("/admin/v1/orgs/%v/policies/%v", organizationID, policyID)

	request, err := o.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

var (
	notPolicyError = fmt.Errorf("error!, please provide a valid policyID value")
)
