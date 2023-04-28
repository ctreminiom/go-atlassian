package admin

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// SCIMSchemaConnector represents the cloud admin SCIM schema actions.
// Use it to search, get, create, delete, and change schemas.
type SCIMSchemaConnector interface {

	// Gets get all SCIM features metadata.
	//
	// Filtering, pagination and sorting are not supported.
	//
	// GET /scim/directory/{directoryId}/Schemas
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemas#get-all-schemas
	Gets(ctx context.Context, directoryID string) (*model.SCIMSchemasScheme, *model.ResponseScheme, error)

	// Group get the group schemas from the SCIM provider.
	//
	// Filtering, pagination and sorting are not supported.
	//
	// GET /scim/directory/{directoryId}/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemas#get-group-schemas
	Group(ctx context.Context, directoryID string) (*model.SCIMSchemaScheme, *model.ResponseScheme, error)

	// User get the user schemas from the SCIM provider.
	//
	// Filtering, pagination and sorting are not supported.
	//
	// GET /scim/directory/{directoryId}/Schemas/urn:ietf:params:scim:schemas:core:2.0:User
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemas#get-user-schemas
	User(ctx context.Context, directoryID string) (*model.SCIMSchemaScheme, *model.ResponseScheme, error)

	// Enterprise get the user enterprise extension schemas from the SCIM provider.
	//
	// Filtering, pagination and sorting are not supported.
	//
	// GET /scim/directory/{directoryId}/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemas#get-user-enterprise-extension-schemas
	Enterprise(ctx context.Context, directoryID string) (*model.SCIMSchemaScheme, *model.ResponseScheme, error)

	// Feature get metadata about the supported SCIM features.
	//
	// This is a service provider configuration endpoint providing supported SCIM features.
	//
	// Filtering, pagination and sorting are not supported.
	//
	// GET /scim/directory/{directoryId}/ServiceProviderConfig
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemas#get-feature-metadata
	Feature(ctx context.Context, directoryID string) (*model.ServiceProviderConfigScheme, *model.ResponseScheme, error)
}
