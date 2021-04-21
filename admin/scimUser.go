package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type SCIMUserService struct{ client *Client }

// Create a user in a directory.
// An attempt to create an existing user fails with a 409 (Conflict) error.
// A user account can only be created if it has an email address on a verified domain.
// If a managed Atlassian account already exists on the Atlassian platform for the specified email address,
// the user in your identity provider is linked to the user in your Atlassian organization.
// --- This func needs the following parameters: ---
// 1. ctx = it's the context.context value
// 2. directoryId = Directory Id (REQUIRED)
// 3. payload = The information of the new user to create (REQUIRED)
// 4. attributes = Resource attributes to be included in response
// 5. excludedAttributes = Resource attributes to be excluded from response
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-provisioning/rest/api-group-users/#api-scim-directory-directoryid-users-post
// Library Docs: N/A
func (s *SCIMUserService) Create(ctx context.Context, directoryID string, payload *SCIMUserScheme, attributes, excludedAttributes []string) (result *SCIMUserScheme, response *Response, err error) {

	if len(directoryID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	if payload == nil {
		return nil, nil, fmt.Errorf("error!, please provide a valid SCIMUserScheme pointer value")
	}

	params := url.Values{}

	var attributesAsString string
	for pos, attribute := range attributes {

		if pos == 0 {
			attributesAsString = attribute
			continue
		}

		attributesAsString += "," + attribute
	}

	if len(attributesAsString) != 0 {
		params.Add("attributes", attributesAsString)
	}

	var excludedAttributesAsString string
	for pos, attribute := range excludedAttributes {

		if pos == 0 {
			excludedAttributesAsString = attribute
			continue
		}

		excludedAttributesAsString += "," + attribute
	}

	if len(excludedAttributesAsString) != 0 {
		params.Add("excludedAttributes", excludedAttributesAsString)
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("/scim/directory/%v/Users?%v", directoryID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("/scim/directory/%v/Users", directoryID)
	}

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SCIMUserScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type SCIMUserScheme struct {
	ID                string                        `json:"id"`
	ExternalID        string                        `json:"externalId"`
	Meta              *SCIMUserMetaScheme           `json:"meta,omitempty"`
	Groups            []*SCIMUserGroupScheme        `json:"groups,omitempty"`
	UserName          string                        `json:"userName,omitempty"`
	Emails            []*SCIMUserEmailScheme        `json:"emails,omitempty"`
	Name              *SCIMUserNameScheme           `json:"name,omitempty"`
	DisplayName       string                        `json:"displayName,omitempty"`
	NickName          string                        `json:"nickName,omitempty"`
	Title             string                        `json:"title,omitempty"`
	PreferredLanguage string                        `json:"preferredLanguage,omitempty"`
	Department        string                        `json:"department,omitempty"`
	Organization      string                        `json:"organization,omitempty"`
	Timezone          string                        `json:"timezone,omitempty"`
	PhoneNumbers      []*SCIMUserPhoneNumberScheme  `json:"phoneNumbers,omitempty"`
	Active            bool                          `json:"active,omitempty"`
	EnterpriseInfo    *SCIMEnterpriseUserInfoScheme `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.1:User,omitempty"`
	SCIMExtension     *SCIMExtensionScheme          `json:"urn:scim:schemas:extension:atlassian-external:1.1,omitempty"`
}
type SCIMUserEmailScheme struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type SCIMUserNameScheme struct {
	Formatted       string `json:"formatted,omitempty"`
	FamilyName      string `json:"familyName,omitempty"`
	GivenName       string `json:"givenName,omitempty"`
	MiddleName      string `json:"middleName,omitempty"`
	HonorificPrefix string `json:"honorificPrefix,omitempty"`
	HonorificSuffix string `json:"honorificSuffix,omitempty"`
}

type SCIMUserPhoneNumberScheme struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type SCIMUserMetaScheme struct {
	ResourceType string `json:"resourceType,omitempty"`
	Location     string `json:"location,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
	Created      string `json:"created,omitempty"`
}

type SCIMUserGroupScheme struct {
	Type    string `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type SCIMEnterpriseUserInfoScheme struct {
	Organization string `json:"organization,omitempty"`
	Department   string `json:"department,omitempty"`
}

type SCIMExtensionScheme struct {
	AtlassianAccountID string `json:"atlassianAccountId,omitempty"`
}

type SCIMUserGetsOptionsScheme struct {
	Attributes         []string
	ExcludedAttributes []string
	Filter             string
}

// Get users from the specified directory
// --- This func needs the following parameters: ---
// 1. ctx = it's the context.context value
// 2. directoryId = Directory Id (REQUIRED)
// 3. opts = More Filtering queries
// 4. startIndex = A 1-based index of the first query result.
// 5. count = Desired maximum number of query results in the list response page.
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-provisioning/rest/api-group-users/#api-scim-directory-directoryid-users-get
// Library Docs: N/A
func (s *SCIMUserService) Gets(ctx context.Context, directoryID string, opts *SCIMUserGetsOptionsScheme, startIndex, count int) (result *SCIMUserPageScheme, response *Response, err error) {

	if len(directoryID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	params := url.Values{}
	params.Add("startIndex", strconv.Itoa(startIndex))
	params.Add("count", strconv.Itoa(count))

	if opts != nil {

		var attributesAsString string
		for pos, attribute := range opts.Attributes {

			if pos == 0 {
				attributesAsString = attribute
				continue
			}

			attributesAsString += "," + attribute
		}

		if len(attributesAsString) != 0 {
			params.Add("attributes", attributesAsString)
		}

		var excludedAttributesAsString string
		for pos, attribute := range opts.ExcludedAttributes {

			if pos == 0 {
				excludedAttributesAsString = attribute
				continue
			}

			excludedAttributesAsString += "," + attribute
		}

		if len(excludedAttributesAsString) != 0 {
			params.Add("excludedAttributes", excludedAttributesAsString)
		}

		if len(opts.Filter) != 0 {
			params.Add("filter", opts.Filter)
		}

	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Users?%v", directoryID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SCIMUserPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type SCIMUserPageScheme struct {
	Schemas      []string          `json:"schemas,omitempty"`
	TotalResults int               `json:"totalResults,omitempty"`
	StartIndex   int               `json:"startIndex,omitempty"`
	ItemsPerPage int               `json:"itemsPerPage,omitempty"`
	Resources    []*SCIMUserScheme `json:"Resources,omitempty"`
}

// Get a user from a directory by userId.
// --- This func needs the following parameters: ---
// 1. ctx = it's the context.context value
// 2. directoryId = Directory Id (REQUIRED)
// 3. userId = The user ID
// 4. attributes = Resource attributes to be included in response
// 5. excludedAttributes = Resource attributes to be excluded from response
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-provisioning/rest/api-group-users/#api-scim-directory-directoryid-users-userid-get
// Library Docs: N/A
func (s *SCIMUserService) Get(ctx context.Context, directoryID, userID string, attributes, excludedAttributes []string) (result *SCIMUserScheme, response *Response, err error) {

	if len(directoryID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	if len(userID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid userID value")
	}

	params := url.Values{}

	var attributesAsString string
	for pos, attribute := range attributes {

		if pos == 0 {
			attributesAsString = attribute
			continue
		}

		attributesAsString += "," + attribute
	}

	if len(attributesAsString) != 0 {
		params.Add("attributes", attributesAsString)
	}

	var excludedAttributesAsString string
	for pos, attribute := range excludedAttributes {

		if pos == 0 {
			excludedAttributesAsString = attribute
			continue
		}

		excludedAttributesAsString += "," + attribute
	}

	if len(excludedAttributesAsString) != 0 {
		params.Add("excludedAttributes", excludedAttributesAsString)
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("/scim/directory/%v/Users/%v?%v", directoryID, userID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("/scim/directory/%v/Users/%v", directoryID, userID)
	}

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SCIMUserScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deactivate a user by userId.
// The user is not available for future requests until activated again.
// Any future operation for the deactivated user returns the 404 (resource not found) error.
// --- This func needs the following parameters: ---
// 1. ctx = it's the context.context value
// 2. directoryId = Directory Id (REQUIRED)
// 3. userId = The user ID (REQUIRED)
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-provisioning/rest/api-group-users/#api-scim-directory-directoryid-users-userid-delete
// Library Docs: N/A
func (s *SCIMUserService) Deactivate(ctx context.Context, directoryID, userID string) (response *Response, err error) {

	if len(directoryID) == 0 {
		return nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	if len(userID) == 0 {
		return nil, fmt.Errorf("error!, please provide a valid userID value")
	}

	var endpoint = fmt.Sprintf("/scim/directory/%v/Users/%v", directoryID, userID)

	request, err := s.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	return
}

// This operation updates a user's information in a directory by userId via PATCH.
// Refer to GET /ServiceProviderConfig for details on the supported operations.
// --- This func needs the following parameters: ---
// 1. ctx = it's the context.context value
// 2. directoryId = Directory Id (REQUIRED)
// 3. userId = The user ID (REQUIRED)
// 4. payload = the user update payload (REQUIRED)
// 5. attributes = Resource attributes to be included in response.
// 6. excludedAttributes = Resource attributes to be excluded from response.
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-provisioning/rest/api-group-users/#api-scim-directory-directoryid-users-userid-patch
// Library Docs: N/A
func (s *SCIMUserService) Path(ctx context.Context, directoryID, userID string, payload *SCIMUserToPathScheme, attributes, excludedAttributes []string) (result *SCIMUserScheme, response *Response, err error) {

	if len(directoryID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	if len(userID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid userID value")
	}

	if payload == nil {
		return nil, nil, fmt.Errorf("error!, please provide a valid SCIMUserToPathScheme pointer value")
	}

	params := url.Values{}

	var attributesAsString string
	for pos, attribute := range attributes {

		if pos == 0 {
			attributesAsString = attribute
			continue
		}

		attributesAsString += "," + attribute
	}

	if len(attributesAsString) != 0 {
		params.Add("attributes", attributesAsString)
	}

	var excludedAttributesAsString string
	for pos, attribute := range excludedAttributes {

		if pos == 0 {
			excludedAttributesAsString = attribute
			continue
		}

		excludedAttributesAsString += "," + attribute
	}

	if len(excludedAttributesAsString) != 0 {
		params.Add("excludedAttributes", excludedAttributesAsString)
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("/scim/directory/%v/Users/%v?%v", directoryID, userID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("/scim/directory/%v/Users/%v", directoryID, userID)
	}

	request, err := s.client.newRequest(ctx, http.MethodPatch, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SCIMUserScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type SCIMUserToPathScheme struct {
	Schemas    []string                         `json:"schemas,omitempty"`
	Operations []*SCIMUserToPathOperationScheme `json:"operations,omitempty"`
}

func (s *SCIMUserToPathScheme) AddStringOperation(operation, path, value string) (err error) {

	if len(operation) == 0 {
		return fmt.Errorf("error!, please provide a valid operation value, you can check the availables values calling the user schemas")
	}

	if len(path) == 0 {
		return fmt.Errorf("error!, please provide a valid path value")
	}

	if len(value) == 0 {
		return fmt.Errorf("error!, please provide a valid value value")
	}

	s.Operations = append(s.Operations, &SCIMUserToPathOperationScheme{
		Op:    operation,
		Path:  path,
		Value: value,
	})

	return
}

func (s *SCIMUserToPathScheme) AddBoolOperation(operation, path string, value bool) (err error) {

	if len(operation) == 0 {
		return fmt.Errorf("error!, please provide a valid operation value, you can check the availables values calling the user schemas")
	}

	if len(path) == 0 {
		return fmt.Errorf("error!, please provide a valid path value")
	}

	s.Operations = append(s.Operations, &SCIMUserToPathOperationScheme{
		Op:    operation,
		Path:  path,
		Value: value,
	})

	return
}

func (s *SCIMUserToPathScheme) AddComplexOperation(operation, path string, values []*SCIMUserComplexOperationScheme) (err error) {

	if len(operation) == 0 {
		return fmt.Errorf("error!, please provide a valid operation value, you can check the availables values calling the user schemas")
	}

	if len(path) == 0 {
		return fmt.Errorf("error!, please provide a valid path value")
	}

	if values == nil {
		return fmt.Errorf("error!, please provide a valid SCIMUserComplexOperationScheme slice pointer")
	}

	if len(values) == 0 {
		return fmt.Errorf("error!, the values variable must contains SCIMUserComplexOperationScheme nodes")
	}

	s.Operations = append(s.Operations, &SCIMUserToPathOperationScheme{
		Op:    operation,
		Path:  path,
		Value: values,
	})

	return
}

type SCIMUserComplexOperationScheme struct {
	Value     string `json:"value,omitempty"`
	ValueType string `json:"type,omitempty"` // Available values (work, home, other)
	Primary   bool   `json:"primary,omitempty"`
}

type SCIMUserToPathValueScheme struct {
	Array               bool   `json:"array,omitempty"`
	Null                bool   `json:"null,omitempty"`
	ValueNode           bool   `json:"valueNode,omitempty"`
	ContainerNode       bool   `json:"containerNode,omitempty"`
	MissingNode         bool   `json:"missingNode,omitempty"`
	Object              bool   `json:"object,omitempty"`
	NodeType            string `json:"nodeType,omitempty"`
	Pojo                bool   `json:"pojo,omitempty"`
	Number              bool   `json:"number,omitempty"`
	IntegralNumber      bool   `json:"integralNumber,omitempty"`
	FloatingPointNumber bool   `json:"floatingPointNumber,omitempty"`
	Short               bool   `json:"short,omitempty"`
	Int                 bool   `json:"int,omitempty"`
	Long                bool   `json:"long,omitempty"`
	Double              bool   `json:"double,omitempty"`
	BigDecimal          bool   `json:"bigDecimal,omitempty"`
	BigInteger          bool   `json:"bigInteger,omitempty"`
	Textual             bool   `json:"textual,omitempty"`
	Boolean             bool   `json:"boolean,omitempty"`
	Binary              bool   `json:"binary,omitempty"`
	Float               bool   `json:"float,omitempty"`
}

type SCIMUserToPathOperationScheme struct {
	Op    string      `json:"op,omitempty"`
	Path  string      `json:"path,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// Updates a user's information in a directory by userId via user attributes.
// User information is replaced attribute-by-attribute, with the exception of immutable and read-only attributes.
// Existing values of unspecified attributes are cleaned.
// --- This func needs the following parameters: ---
// 1. ctx = it's the context.context value
// 2. directoryId = Directory Id (REQUIRED)
// 3. userId = The user ID (REQUIRED)
// 4. payload = the user update payload (REQUIRED)
// 5. attributes = Resource attributes to be included in response.
// 6. excludedAttributes = Resource attributes to be excluded from response.
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-provisioning/rest/api-group-users/#api-scim-directory-directoryid-users-userid-put
// Library Docs: N/A
func (s *SCIMUserService) Update(ctx context.Context, directoryID, userID string, payload *SCIMUserScheme, attributes, excludedAttributes []string) (result *SCIMUserScheme, response *Response, err error) {

	if len(directoryID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid directoryID value")
	}

	if len(userID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid userID value")
	}

	if payload == nil {
		return nil, nil, fmt.Errorf("error!, please provide a valid SCIMUserToPathScheme pointer value")
	}

	params := url.Values{}

	var attributesAsString string
	for pos, attribute := range attributes {

		if pos == 0 {
			attributesAsString = attribute
			continue
		}

		attributesAsString += "," + attribute
	}

	if len(attributesAsString) != 0 {
		params.Add("attributes", attributesAsString)
	}

	var excludedAttributesAsString string
	for pos, attribute := range excludedAttributes {

		if pos == 0 {
			excludedAttributesAsString = attribute
			continue
		}

		excludedAttributesAsString += "," + attribute
	}

	if len(excludedAttributesAsString) != 0 {
		params.Add("excludedAttributes", excludedAttributesAsString)
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("/scim/directory/%v/Users/%v?%v", directoryID, userID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("/scim/directory/%v/Users/%v", directoryID, userID)
	}

	request, err := s.client.newRequest(ctx, http.MethodPut, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/scim+json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SCIMUserScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
