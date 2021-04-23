package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SCIMSchemeService struct{ client *Client }

// Get all SCIM features metadata. Filtering, pagination and sorting are not supported.
// --- This func needs the following parameters: ---
// 1. ctx = it's the context.context value (REQUIRED)
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-provisioning/rest/api-group-schemas/#api-scim-directory-directoryid-schemas-get
// Library Docs: N/A
func (s *SCIMSchemeService) Gets(ctx context.Context, directoryID string) (result *SCIMSchemasScheme, response *Response, err error) {

	if len(directoryID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Schemas", directoryID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SCIMSchemasScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type SCIMSchemasScheme struct {
	TotalResults int               `json:"totalResults,omitempty"`
	ItemsPerPage int               `json:"itemsPerPage,omitempty"`
	StartIndex   int               `json:"startIndex,omitempty"`
	Schemas      []string          `json:"schemas,omitempty"`
	Resources    []*ResourceScheme `json:"Resources,omitempty"`
}

type ResourceScheme struct {
	ID          string              `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Attributes  []*AttributeScheme  `json:"attributes,omitempty"`
	Meta        *ResourceMetaScheme `json:"meta,omitempty"`
}

type ResourceMetaScheme struct {
	ResourceType string `json:"resourceType,omitempty"`
	Location     string `json:"location,omitempty"`
}

type AttributeScheme struct {
	Name          string                `json:"name,omitempty"`
	Type          string                `json:"type,omitempty"`
	MultiValued   bool                  `json:"multiValued,omitempty"`
	Description   string                `json:"description,omitempty"`
	Required      bool                  `json:"required,omitempty"`
	CaseExact     bool                  `json:"caseExact,omitempty"`
	Mutability    string                `json:"mutability,omitempty"`
	Returned      string                `json:"returned,omitempty"`
	Uniqueness    string                `json:"uniqueness,omitempty"`
	SubAttributes []*SubAttributeScheme `json:"subAttributes,omitempty"`
}

type SubAttributeScheme struct {
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	MultiValued bool   `json:"multiValued,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	CaseExact   bool   `json:"caseExact,omitempty"`
	Mutability  string `json:"mutability,omitempty"`
	Returned    string `json:"returned,omitempty"`
	Uniqueness  string `json:"uniqueness,omitempty"`
}

// Get the group schemas from the SCIM provider. Filtering, pagination and sorting are not supported.
// --- This func needs the following parameters: ---
// 1. ctx = it's the context.context value (REQUIRED)
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-provisioning/rest/api-group-schemas/#api-scim-directory-directoryid-schemas-urn-ietf-params-scim-schemas-core-2-0-group-get
// Library Docs: N/A
func (s *SCIMSchemeService) Group(ctx context.Context, directoryID string) (result *SCIMSchemaScheme, response *Response, err error) {

	if len(directoryID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group", directoryID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SCIMSchemaScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Get the user schemas from the SCIM provider. Filtering, pagination and sorting are not supported.
// --- This func needs the following parameters: ---
// 1. ctx = it's the context.context value (REQUIRED)
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-provisioning/rest/api-group-schemas/#api-scim-directory-directoryid-schemas-urn-ietf-params-scim-schemas-core-2-0-user-get
// Library Docs: N/A
func (s *SCIMSchemeService) User(ctx context.Context, directoryID string) (result *SCIMSchemaScheme, response *Response, err error) {

	if len(directoryID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Schemas/urn:ietf:params:scim:schemas:core:2.0:User", directoryID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SCIMSchemaScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Get the user enterprise extension schemas from the SCIM provider. Filtering, pagination and sorting are not supported.
// --- This func needs the following parameters: ---
// 1. ctx = it's the context.context value (REQUIRED)
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-provisioning/rest/api-group-schemas/#api-scim-directory-directoryid-schemas-urn-ietf-params-scim-schemas-extension-enterprise-2-0-user-get
// Library Docs: N/A
func (s *SCIMSchemeService) Enterprise(ctx context.Context, directoryID string) (result *SCIMSchemaScheme, response *Response, err error) {

	if len(directoryID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User", directoryID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SCIMSchemaScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type SCIMSchemaScheme struct {
	ID          string              `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Attributes  []*AttributeScheme  `json:"attributes,omitempty"`
	Meta        *ResourceMetaScheme `json:"meta,omitempty"`
}

// Get metadata about the supported SCIM features.
// This is a service provider configuration endpoint providing supported SCIM features. Filtering, pagination and sorting are not supported.
// --- This func needs the following parameters: ---
// 1. ctx = it's the context.context value (REQUIRED)
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-provisioning/rest/api-group-schemas/#api-scim-directory-directoryid-serviceproviderconfig-get
// Library Docs: N/A
func (s *SCIMSchemeService) Feature(ctx context.Context, directoryID string) (result *ServiceProviderConfigScheme, response *Response, err error) {

	if len(directoryID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/ServiceProviderConfig", directoryID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ServiceProviderConfigScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ServiceProviderConfigScheme struct {
	Schemas []string `json:"schemas"`
	Patch   struct {
		Supported bool `json:"supported"`
	} `json:"patch"`
	Bulk struct {
		Supported      bool `json:"supported"`
		MaxOperations  int  `json:"maxOperations"`
		MaxPayloadSize int  `json:"maxPayloadSize"`
	} `json:"bulk"`
	Filter struct {
		MaxResults int  `json:"maxResults"`
		Supported  bool `json:"supported"`
	} `json:"filter"`
	ChangePassword struct {
		Supported bool `json:"supported"`
	} `json:"changePassword"`
	Sort struct {
		Supported bool `json:"supported"`
	} `json:"sort"`
	Etag struct {
		Supported bool `json:"supported"`
	} `json:"etag"`
	AuthenticationSchemes []struct {
		Type        string `json:"type"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"authenticationSchemes"`
	Meta struct {
		Location     string    `json:"location"`
		ResourceType string    `json:"resourceType"`
		LastModified time.Time `json:"lastModified"`
		Created      time.Time `json:"created"`
	} `json:"meta"`
}
