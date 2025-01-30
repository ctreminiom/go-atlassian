package admin

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// OrganizationConnector represents the cloud admin organization endpoints.
// Use it to search, get, create, delete, and change organizations.
type OrganizationConnector interface {

	// Gets returns a list of your organizations (based on your API key).
	//
	// GET /admin/v1/orgs
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-organizations
	Gets(ctx context.Context, cursor string) (*model.AdminOrganizationPageScheme, *model.ResponseScheme, error)

	// Get returns information about a single organization by ID
	//
	// GET /admin/v1/orgs/{organizationID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-organization-by-id
	Get(ctx context.Context, organizationID string) (*model.AdminOrganizationScheme, *model.ResponseScheme, error)

	// Users returns a list of users in an organization.
	//
	// GET /admin/v1/orgs/{organizationID}/users
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-users-in-an-organization
	Users(ctx context.Context, organizationID, cursor string) (*model.OrganizationUserPageScheme, *model.ResponseScheme, error)

	// Domains returns a list of domains in an organization one page at a time.
	//
	// GET /admin/v1/orgs/{organizationID}/domains
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domains-in-an-organization
	Domains(ctx context.Context, organizationID, cursor string) (*model.OrganizationDomainPageScheme, *model.ResponseScheme, error)

	// Domain returns information about a single verified domain by ID.
	//
	// GET /admin/v1/orgs/{organizationID}/domains/{domainID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-domain-by-id
	Domain(ctx context.Context, organizationID, domainID string) (*model.OrganizationDomainScheme, *model.ResponseScheme, error)

	// Events returns an audit log of events from an organization one page at a time.
	//
	// GET /admin/v1/orgs/{organizationID}/events
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-audit-log-of-events
	Events(ctx context.Context, organizationID string, options *model.OrganizationEventOptScheme, cursor string) (*model.OrganizationEventPageScheme, *model.ResponseScheme, error)

	// Event returns information about a single event by ID.
	//
	// GET /admin/v1/orgs/{organizationID}/events/{eventID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-an-event-by-id
	Event(ctx context.Context, organizationID, eventID string) (*model.OrganizationEventScheme, *model.ResponseScheme, error)

	// Actions returns information localized event actions.
	//
	// GET /admin/v1/orgs/{organizationID}/event-actions
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization#get-list-of-event-actions
	Actions(ctx context.Context, organizationID string) (*model.OrganizationEventActionScheme, *model.ResponseScheme, error)

	// SearchUsers searches for users within an organization with the specified filters
	//
	// POST /admin/v1/orgs/{organizationID}/users/search
	//
	// https://developer.atlassian.com/cloud/admin/organization/rest/api-group-users/#api-v1-orgs-orgid-users-search-post
	SearchUsers(ctx context.Context, organizationID string, payload *model.OrganizationUserSearchParams) (*model.OrganizationUserSearchPage, *model.ResponseScheme, error)

	// SearchGroups searches for groups within an organization with the specified filters
	//
	// POST /admin/v1/orgs/{organizationID}/groups/search
	//
	// https://developer.atlassian.com/cloud/admin/organization/rest/api-group-groups/#api-v1-orgs-orgid-groups-search-post
	SearchGroups(ctx context.Context, organizationID string, payload *model.OrganizationGroupSearchParams) (*model.OrganizationGroupSearchPage, *model.ResponseScheme, error)

	// SearchWorkspaces searches for workspaces within an organization with the specified filters
	//
	// POST /v2/orgs/{organizationID}/workspaces
	//
	// https://developer.atlassian.com/cloud/admin/organization/rest/api-group-workspaces/#api-v2-orgs-orgid-workspaces-post
	SearchWorkspaces(ctx context.Context, organizationID string, payload *model.WorkspaceSearchParams) (*model.WorkspaceSearchPage, *model.ResponseScheme, error)
}
