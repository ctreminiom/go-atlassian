package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/admin"
	"net/http"
)

// NewSCIMSchemaService creates a new instance of SCIMSchemaService.
// It takes a service.Connector as input and returns a pointer to SCIMSchemaService.
func NewSCIMSchemaService(client service.Connector) *SCIMSchemaService {
	return &SCIMSchemaService{internalClient: &internalSCIMSchemaImpl{c: client}}
}

// SCIMSchemaService provides methods to interact with SCIM schemas in Atlassian Administration.
type SCIMSchemaService struct {
	// internalClient is the connector interface for SCIM schema operations.
	internalClient admin.SCIMSchemaConnector
}

// Gets get all SCIM features metadata.
//
// Filtering, pagination and sorting are not supported.
//
// GET /scim/directory/{directoryId}/Schemas
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemas#get-all-schemas
func (s *SCIMSchemaService) Gets(ctx context.Context, directoryID string) (*model.SCIMSchemasScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, directoryID)
}

// Group get the group schemas from the SCIM provider.
//
// Filtering, pagination and sorting are not supported.
//
// GET /scim/directory/{directoryId}/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemas#get-group-schemas
func (s *SCIMSchemaService) Group(ctx context.Context, directoryID string) (*model.SCIMSchemaScheme, *model.ResponseScheme, error) {
	return s.internalClient.Group(ctx, directoryID)
}

// User get the user schemas from the SCIM provider.
//
// Filtering, pagination and sorting are not supported.
//
// GET /scim/directory/{directoryId}/Schemas/urn:ietf:params:scim:schemas:core:2.0:User
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemas#get-user-schemas
func (s *SCIMSchemaService) User(ctx context.Context, directoryID string) (*model.SCIMSchemaScheme, *model.ResponseScheme, error) {
	return s.internalClient.User(ctx, directoryID)
}

// Enterprise get the user enterprise extension schemas from the SCIM provider.
//
// Filtering, pagination and sorting are not supported.
//
// GET /scim/directory/{directoryId}/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemas#get-user-enterprise-extension-schemas
func (s *SCIMSchemaService) Enterprise(ctx context.Context, directoryID string) (*model.SCIMSchemaScheme, *model.ResponseScheme, error) {
	return s.internalClient.Enterprise(ctx, directoryID)
}

// Feature get metadata about the supported SCIM features.
//
// This is a service provider configuration endpoint providing supported SCIM features.
//
// Filtering, pagination and sorting are not supported.
//
// GET /scim/directory/{directoryId}/ServiceProviderConfig
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemas#get-feature-metadata
func (s *SCIMSchemaService) Feature(ctx context.Context, directoryID string) (*model.ServiceProviderConfigScheme, *model.ResponseScheme, error) {
	return s.internalClient.Feature(ctx, directoryID)
}

type internalSCIMSchemaImpl struct {
	c service.Connector
}

func (i *internalSCIMSchemaImpl) Gets(ctx context.Context, directoryID string) (*model.SCIMSchemasScheme, *model.ResponseScheme, error) {

	if directoryID == "" {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	endpoint := fmt.Sprintf("scim/directory/%v/Schemas", directoryID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	schemas := new(model.SCIMSchemasScheme)
	response, err := i.c.Call(request, schemas)
	if err != nil {
		return nil, response, err
	}

	return schemas, response, nil
}

func (i *internalSCIMSchemaImpl) Group(ctx context.Context, directoryID string) (*model.SCIMSchemaScheme, *model.ResponseScheme, error) {

	if directoryID == "" {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	endpoint := fmt.Sprintf("scim/directory/%v/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group", directoryID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(model.SCIMSchemaScheme)
	response, err := i.c.Call(request, group)
	if err != nil {
		return nil, response, err
	}

	return group, response, nil
}

func (i *internalSCIMSchemaImpl) User(ctx context.Context, directoryID string) (*model.SCIMSchemaScheme, *model.ResponseScheme, error) {

	if directoryID == "" {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	endpoint := fmt.Sprintf("scim/directory/%v/Schemas/urn:ietf:params:scim:schemas:core:2.0:User", directoryID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(model.SCIMSchemaScheme)
	response, err := i.c.Call(request, user)
	if err != nil {
		return nil, response, err
	}

	return user, response, nil
}

func (i *internalSCIMSchemaImpl) Enterprise(ctx context.Context, directoryID string) (*model.SCIMSchemaScheme, *model.ResponseScheme, error) {

	if directoryID == "" {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	endpoint := fmt.Sprintf("scim/directory/%v/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User", directoryID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	enterprise := new(model.SCIMSchemaScheme)
	response, err := i.c.Call(request, enterprise)
	if err != nil {
		return nil, response, err
	}

	return enterprise, response, nil
}

func (i *internalSCIMSchemaImpl) Feature(ctx context.Context, directoryID string) (*model.ServiceProviderConfigScheme, *model.ResponseScheme, error) {

	if directoryID == "" {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	endpoint := fmt.Sprintf("scim/directory/%v/ServiceProviderConfig", directoryID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	provider := new(model.ServiceProviderConfigScheme)
	response, err := i.c.Call(request, provider)
	if err != nil {
		return nil, response, err
	}

	return provider, response, nil
}
