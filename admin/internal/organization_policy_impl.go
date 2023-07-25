package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/admin"
	"net/http"
	"net/url"
	"strings"
)

func NewOrganizationPolicyService(client service.Connector) *OrganizationPolicyService {
	return &OrganizationPolicyService{internalClient: &internalOrganizationPolicyImpl{c: client}}
}

type OrganizationPolicyService struct {
	internalClient admin.OrganizationPolicyConnector
}

// Gets returns information about org policies.
//
// GET /admin/v1/orgs/{orgId}/policies
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#get-list-of-policies
func (o *OrganizationPolicyService) Gets(ctx context.Context, organizationID, policyType, cursor string) (*model.OrganizationPolicyPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Gets(ctx, organizationID, policyType, cursor)
}

// Get returns information about a single policy by ID.
//
// GET /admin/v1/orgs/{orgId}/policies/{policyId}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#get-a-policy-by-id
func (o *OrganizationPolicyService) Get(ctx context.Context, organizationID, policyID string) (*model.OrganizationPolicyScheme, *model.ResponseScheme, error) {
	return o.internalClient.Get(ctx, organizationID, policyID)
}

// Create creates a policy for an org
//
// POST /admin/v1/orgs/{orgId}/policies
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#create-a-policy
func (o *OrganizationPolicyService) Create(ctx context.Context, organizationID string, payload *model.OrganizationPolicyData) (*model.OrganizationPolicyScheme, *model.ResponseScheme, error) {
	return o.internalClient.Create(ctx, organizationID, payload)
}

// Update updates a policy for an org
//
// PUT /admin/v1/orgs/{orgId}/policies/{policyId}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#update-a-policy
func (o *OrganizationPolicyService) Update(ctx context.Context, organizationID, policyID string, payload *model.OrganizationPolicyData) (*model.OrganizationPolicyScheme, *model.ResponseScheme, error) {
	return o.internalClient.Update(ctx, organizationID, policyID, payload)
}

// Delete deletes a policy for an org.
//
// DELETE /admin/v1/orgs/{orgId}/policies/{policyId}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#delete-a-policy
func (o *OrganizationPolicyService) Delete(ctx context.Context, organizationID, policyID string) (*model.ResponseScheme, error) {
	return o.internalClient.Delete(ctx, organizationID, policyID)
}

type internalOrganizationPolicyImpl struct {
	c service.Connector
}

func (i *internalOrganizationPolicyImpl) Gets(ctx context.Context, organizationID, policyType, cursor string) (*model.OrganizationPolicyPageScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	params := url.Values{}
	if policyType != "" {
		params.Add("type", policyType)
	}

	if cursor != "" {
		params.Add("cursor", cursor)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("admin/v1/orgs/%v/policies", organizationID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	policies := new(model.OrganizationPolicyPageScheme)
	res, err := i.c.Call(req, policies)
	if err != nil {
		return nil, res, err
	}

	return policies, res, nil
}

func (i *internalOrganizationPolicyImpl) Get(ctx context.Context, organizationID, policyID string) (*model.OrganizationPolicyScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	if policyID == "" {
		return nil, nil, model.ErrNoAdminPolicyError
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/policies/%v", organizationID, policyID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	policy := new(model.OrganizationPolicyScheme)
	res, err := i.c.Call(req, policy)
	if err != nil {
		return nil, res, err
	}

	return policy, res, nil
}

func (i *internalOrganizationPolicyImpl) Create(ctx context.Context, organizationID string, payload *model.OrganizationPolicyData) (*model.OrganizationPolicyScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/policies", organizationID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	policy := new(model.OrganizationPolicyScheme)
	response, err := i.c.Call(request, policy)
	if err != nil {
		return nil, response, err
	}

	return policy, response, nil
}

func (i *internalOrganizationPolicyImpl) Update(ctx context.Context, organizationID, policyID string, payload *model.OrganizationPolicyData) (*model.OrganizationPolicyScheme, *model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, nil, model.ErrNoAdminOrganizationError
	}

	if policyID == "" {
		return nil, nil, model.ErrNoAdminPolicyError
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/policies/%v", organizationID, policyID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	policy := new(model.OrganizationPolicyScheme)
	response, err := i.c.Call(request, policy)
	if err != nil {
		return nil, response, err
	}

	return policy, response, nil
}

func (i *internalOrganizationPolicyImpl) Delete(ctx context.Context, organizationID, policyID string) (*model.ResponseScheme, error) {

	if organizationID == "" {
		return nil, model.ErrNoAdminOrganizationError
	}

	if policyID == "" {
		return nil, model.ErrNoAdminPolicyError
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/policies/%v", organizationID, policyID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
