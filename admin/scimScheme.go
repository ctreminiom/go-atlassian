package admin

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type SCIMSchemeService struct{ client *Client }

// Gets all SCIM features metadata. Filtering, pagination and sorting are not supported.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemes#get-all-schemas
func (s *SCIMSchemeService) Gets(ctx context.Context, directoryID string) (result *model.SCIMSchemasScheme,
	response *ResponseScheme, err error) {

	if len(directoryID) == 0 {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Schemas", directoryID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Group get the group schemas from the SCIM provider. Filtering, pagination and sorting are not supported.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemes#get-group-schemas
func (s *SCIMSchemeService) Group(ctx context.Context, directoryID string) (result *model.SCIMSchemaScheme,
	response *ResponseScheme, err error) {

	if len(directoryID) == 0 {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group", directoryID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// User get the user schemas from the SCIM provider. Filtering, pagination and sorting are not supported.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemes#get-user-schemas
func (s *SCIMSchemeService) User(ctx context.Context, directoryID string) (result *model.SCIMSchemaScheme,
	response *ResponseScheme, err error) {

	if len(directoryID) == 0 {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Schemas/urn:ietf:params:scim:schemas:core:2.0:User", directoryID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Enterprise get the user enterprise extension schemas from the SCIM provider.
// Filtering, pagination and sorting are not supported.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemes#get-user-enterprise-extension-schemas
func (s *SCIMSchemeService) Enterprise(ctx context.Context, directoryID string) (result *model.SCIMSchemaScheme,
	response *ResponseScheme, err error) {

	if len(directoryID) == 0 {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User", directoryID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Feature get metadata about the supported SCIM features.
// This is a service provider configuration endpoint providing supported SCIM features.
// Filtering, pagination and sorting are not supported.
// Docs: https://docs.go-atlassian.io/atlassian-admin-cloud/scim/schemes#get-feature-metadata
func (s *SCIMSchemeService) Feature(ctx context.Context, directoryID string) (result *model.ServiceProviderConfigScheme,
	response *ResponseScheme, err error) {

	if len(directoryID) == 0 {
		return nil, nil, model.ErrNoAdminDirectoryIDError
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/ServiceProviderConfig", directoryID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
