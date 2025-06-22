package internal

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/admin"
)

// NewOrganizationDirectoryService creates a new instance of OrganizationDirectoryService.
// It takes a service.Connector as input and returns a pointer to OrganizationDirectoryService.
func NewOrganizationDirectoryService(client service.Connector) *OrganizationDirectoryService {
	return &OrganizationDirectoryService{internalClient: &internalOrganizationDirectoryServiceImpl{c: client}}
}

// OrganizationDirectoryService provides methods to interact with the organization directory in Atlassian Administration.
type OrganizationDirectoryService struct {
	// internalClient is the connector interface for organization directory operations.
	internalClient admin.OrganizationDirectoryConnector
}

// Activity returns a user’s last active date for each product listed in Atlassian Administration.
//
// Active is defined as viewing a product's page for a minimum of 2 seconds.
//
// Last activity data can be delayed by up to 4 hours.
//
// If the user has not accessed a product, the product_access response field will be empty.
//
// The added_to_org date field is available only to customers using the new user management experience.
//
// GET /admin/v1/orgs/{orgId}/directory/users/{accountID}/last-active-dates
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/directory#users-last-active-dates
func (o *OrganizationDirectoryService) Activity(ctx context.Context, organizationID, accountID string) (*model.UserProductAccessScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*OrganizationDirectoryService).Activity", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "activity"))

	return o.internalClient.Activity(ctx, organizationID, accountID)
}

// Remove removes user access to products listed in Atlassian Administration.
//
// -- The API is available for customers using the new user management experience only. --
//
// Note: Users with emails whose domain is claimed can still be found in Managed accounts in Directory.
//
// DELETE /admin/v1/orgs/{orgId}/directory/users/{accountId}
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/directory#remove-user-access
func (o *OrganizationDirectoryService) Remove(ctx context.Context, organizationID, accountID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*OrganizationDirectoryService).Remove", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "remove"))

	return o.internalClient.Remove(ctx, organizationID, accountID)
}

// Suspend suspends user access to products listed in Atlassian Administration.
//
// -- The API is available for customers using the new user management experience only. --
//
// Note: Users with emails whose domain is claimed can still be found in Managed accounts in Directory.
//
// POST /admin/v1/orgs/{orgId}/directory/users/{accountId}/suspend-access
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/directory#suspend-user-access
func (o *OrganizationDirectoryService) Suspend(ctx context.Context, organizationID, accountID string) (*model.GenericActionSuccessScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*OrganizationDirectoryService).Suspend", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "suspend"))

	return o.internalClient.Suspend(ctx, organizationID, accountID)
}

// Restore restores user access to products listed in Atlassian Administration.
//
// -- The API is available for customers using the new user management experience only. --
//
// Note: Users with emails whose domain is claimed can still be found in Managed accounts in Directory.
//
// POST /admin/v1/orgs/{orgId}/directory/users/{accountId}/restore-access
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/directory#restore-user-access
func (o *OrganizationDirectoryService) Restore(ctx context.Context, organizationID, accountID string) (*model.GenericActionSuccessScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*OrganizationDirectoryService).Restore", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "restore"))

	return o.internalClient.Restore(ctx, organizationID, accountID)
}

type internalOrganizationDirectoryServiceImpl struct {
	c service.Connector
}

func (i *internalOrganizationDirectoryServiceImpl) Activity(ctx context.Context, organizationID, accountID string) (*model.UserProductAccessScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalOrganizationDirectoryServiceImpl).Activity", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "activity"))

	if organizationID == "" {

			return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminOrganization)
	}

	if accountID == "" {

			return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminAccountID)
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/directory/users/%v/last-active-dates", organizationID, accountID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	activity := new(model.UserProductAccessScheme)
	res, err := i.c.Call(req, activity)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return activity, res, nil
}

func (i *internalOrganizationDirectoryServiceImpl) Remove(ctx context.Context, organizationID, accountID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalOrganizationDirectoryServiceImpl).Remove", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "remove"))

	if organizationID == "" {
		return nil, fmt.Errorf("admin: %w", model.ErrNoAdminOrganization)
	}

	if accountID == "" {
		return nil, fmt.Errorf("admin: %w", model.ErrNoAdminAccountID)
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/directory/users/%v", organizationID, accountID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalOrganizationDirectoryServiceImpl) Suspend(ctx context.Context, organizationID, accountID string) (*model.GenericActionSuccessScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalOrganizationDirectoryServiceImpl).Suspend", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "suspend"))

	if organizationID == "" {

			return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminOrganization)
	}

	if accountID == "" {

			return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminAccountID)
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/directory/users/%v/suspend-access", organizationID, accountID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	message := new(model.GenericActionSuccessScheme)
	res, err := i.c.Call(req, message)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return message, res, nil
}

func (i *internalOrganizationDirectoryServiceImpl) Restore(ctx context.Context, organizationID, accountID string) (*model.GenericActionSuccessScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalOrganizationDirectoryServiceImpl).Restore", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "restore"))

	if organizationID == "" {

			return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminOrganization)
	}

	if accountID == "" {

			return nil, nil, fmt.Errorf("admin: %w", model.ErrNoAdminAccountID)
	}

	endpoint := fmt.Sprintf("admin/v1/orgs/%v/directory/users/%v/restore-access", organizationID, accountID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	message := new(model.GenericActionSuccessScheme)
	res, err := i.c.Call(req, message)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return message, res, nil
}
