// Package models provides the data structures used in the admin package.
package models

// SCIMUserScheme represents a SCIM user.
type SCIMUserScheme struct {
	ID                string                        `json:"id"`                                                                   // The ID of the user.
	ExternalID        string                        `json:"externalId"`                                                           // The external ID of the user.
	Meta              *SCIMUserMetaScheme           `json:"meta,omitempty"`                                                       // The metadata of the user.
	Groups            []*SCIMUserGroupScheme        `json:"groups,omitempty"`                                                     // The groups the user belongs to.
	UserName          string                        `json:"userName,omitempty"`                                                   // The username of the user.
	Emails            []*SCIMUserEmailScheme        `json:"emails,omitempty"`                                                     // The emails of the user.
	Name              *SCIMUserNameScheme           `json:"name,omitempty"`                                                       // The name of the user.
	DisplayName       string                        `json:"displayName,omitempty"`                                                // The display name of the user.
	NickName          string                        `json:"nickName,omitempty"`                                                   // The nickname of the user.
	Title             string                        `json:"title,omitempty"`                                                      // The title of the user.
	PreferredLanguage string                        `json:"preferredLanguage,omitempty"`                                          // The preferred language of the user.
	Department        string                        `json:"department,omitempty"`                                                 // The department of the user.
	Organization      string                        `json:"organization,omitempty"`                                               // The organization of the user.
	Timezone          string                        `json:"timezone,omitempty"`                                                   // The timezone of the user.
	PhoneNumbers      []*SCIMUserPhoneNumberScheme  `json:"phoneNumbers,omitempty"`                                               // The phone numbers of the user.
	Active            bool                          `json:"active,omitempty"`                                                     // Whether the user is active.
	EnterpriseInfo    *SCIMEnterpriseUserInfoScheme `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.1:User,omitempty"` // The enterprise user info of the user.
	SCIMExtension     *SCIMExtensionScheme          `json:"urn:scim:schemas:extension:atlassian-external:1.1,omitempty"`          // The SCIM extension of the user.
}

// SCIMUserEmailScheme represents an email of a SCIM user.
type SCIMUserEmailScheme struct {
	Value   string `json:"value,omitempty"`   // The value of the email.
	Type    string `json:"type,omitempty"`    // The type of the email.
	Primary bool   `json:"primary,omitempty"` // Whether the email is primary.
}

// SCIMUserNameScheme represents the name of a SCIM user.
type SCIMUserNameScheme struct {
	Formatted       string `json:"formatted,omitempty"`       // The formatted name.
	FamilyName      string `json:"familyName,omitempty"`      // The family name.
	GivenName       string `json:"givenName,omitempty"`       // The given name.
	MiddleName      string `json:"middleName,omitempty"`      // The middle name.
	HonorificPrefix string `json:"honorificPrefix,omitempty"` // The honorific prefix.
	HonorificSuffix string `json:"honorificSuffix,omitempty"` // The honorific suffix.
}

// SCIMUserPhoneNumberScheme represents a phone number of a SCIM user.
type SCIMUserPhoneNumberScheme struct {
	Value   string `json:"value,omitempty"`   // The value of the phone number.
	Type    string `json:"type,omitempty"`    // The type of the phone number.
	Primary bool   `json:"primary,omitempty"` // Whether the phone number is primary.
}

// SCIMUserMetaScheme represents the metadata of a SCIM user.
type SCIMUserMetaScheme struct {
	ResourceType string `json:"resourceType,omitempty"` // The resource type.
	Location     string `json:"location,omitempty"`     // The location.
	LastModified string `json:"lastModified,omitempty"` // The last modified time.
	Created      string `json:"created,omitempty"`      // The creation time.
}

// SCIMUserGroupScheme represents a group of a SCIM user.
type SCIMUserGroupScheme struct {
	Type    string `json:"type,omitempty"`    // The type of the group.
	Value   string `json:"value,omitempty"`   // The value of the group.
	Display string `json:"display,omitempty"` // The display of the group.
	Ref     string `json:"$ref,omitempty"`    // The reference of the group.
}

// SCIMEnterpriseUserInfoScheme represents the enterprise user info of a SCIM user.
type SCIMEnterpriseUserInfoScheme struct {
	Organization string `json:"organization,omitempty"` // The organization.
	Department   string `json:"department,omitempty"`   // The department.
}

// SCIMExtensionScheme represents the SCIM extension of a SCIM user.
type SCIMExtensionScheme struct {
	AtlassianAccountID string `json:"atlassianAccountId,omitempty"` // The Atlassian account ID.
}

// SCIMUserGetsOptionsScheme represents the options for getting SCIM users.
type SCIMUserGetsOptionsScheme struct {
	Attributes         []string // The attributes to get.
	ExcludedAttributes []string // The attributes to exclude.
	Filter             string   // The filter.
}

// SCIMUserPageScheme represents a page of SCIM users.
type SCIMUserPageScheme struct {
	Schemas      []string          `json:"schemas,omitempty"`      // The schemas.
	TotalResults int               `json:"totalResults,omitempty"` // The total results.
	StartIndex   int               `json:"startIndex,omitempty"`   // The start index.
	ItemsPerPage int               `json:"itemsPerPage,omitempty"` // The items per page.
	Resources    []*SCIMUserScheme `json:"Resources,omitempty"`    // The resources.
}

// SCIMUserToPathScheme represents the path scheme for a SCIM user.
type SCIMUserToPathScheme struct {
	Schemas    []string                         `json:"schemas,omitempty"`    // The schemas.
	Operations []*SCIMUserToPathOperationScheme `json:"operations,omitempty"` // The operations.
}

// AddStringOperation adds a string operation to the SCIM user path scheme.
func (s *SCIMUserToPathScheme) AddStringOperation(operation, path, value string) error {
	if operation == "" {
		return ErrNoSCIMOperation
	}

	if path == "" {
		return ErrNoSCIMPath
	}

	if value == "" {
		return ErrNoSCIMValue
	}

	s.Operations = append(s.Operations, &SCIMUserToPathOperationScheme{
		Op:    operation,
		Path:  path,
		Value: value,
	})

	return nil
}

// AddBoolOperation adds a boolean operation to the SCIM user path scheme.
func (s *SCIMUserToPathScheme) AddBoolOperation(operation, path string, value bool) error {
	if operation == "" {
		return ErrNoSCIMOperation
	}

	if path == "" {
		return ErrNoSCIMPath
	}

	s.Operations = append(s.Operations, &SCIMUserToPathOperationScheme{
		Op:    operation,
		Path:  path,
		Value: value,
	})

	return nil
}

// AddComplexOperation adds a complex operation to the SCIM user path scheme.
func (s *SCIMUserToPathScheme) AddComplexOperation(operation, path string, values []*SCIMUserComplexOperationScheme) error {
	if operation == "" {
		return ErrNoSCIMOperation
	}

	if path == "" {
		return ErrNoSCIMPath
	}

	if values == nil {
		return ErrNoSCIMComplexValue
	}

	if len(values) == 0 {
		return ErrNoSCIMComplexValue
	}

	s.Operations = append(s.Operations, &SCIMUserToPathOperationScheme{
		Op:    operation,
		Path:  path,
		Value: values,
	})

	return nil
}

// SCIMUserComplexOperationScheme represents a complex operation of a SCIM user.
type SCIMUserComplexOperationScheme struct {
	Value     string `json:"value,omitempty"`   // The value of the operation.
	ValueType string `json:"type,omitempty"`    // Available values (work, home, other)
	Primary   bool   `json:"primary,omitempty"` // Whether the operation is primary.
}

// SCIMUserToPathValueScheme represents the value scheme for a path of a SCIM user.
type SCIMUserToPathValueScheme struct {
	Array               bool   `json:"array,omitempty"`               // Whether the value is an array.
	Null                bool   `json:"null,omitempty"`                // Whether the value is null.
	ValueNode           bool   `json:"valueNode,omitempty"`           // Whether the value is a node.
	ContainerNode       bool   `json:"containerNode,omitempty"`       // Whether the value is a container node.
	MissingNode         bool   `json:"missingNode,omitempty"`         // Whether the value is a missing node.
	Object              bool   `json:"object,omitempty"`              // Whether the value is an object.
	NodeType            string `json:"nodeType,omitempty"`            // The node type.
	Pojo                bool   `json:"pojo,omitempty"`                // Whether the value is a POJO.
	Number              bool   `json:"number,omitempty"`              // Whether the value is a number.
	IntegralNumber      bool   `json:"integralNumber,omitempty"`      // Whether the value is an integral number.
	FloatingPointNumber bool   `json:"floatingPointNumber,omitempty"` // Whether the value is a floating point number.
	Short               bool   `json:"short,omitempty"`               // Whether the value is short.
	Int                 bool   `json:"int,omitempty"`                 // Whether the value is an integer.
	Long                bool   `json:"long,omitempty"`                // Whether the value is long.
	Double              bool   `json:"double,omitempty"`              // Whether the value is double.
	BigDecimal          bool   `json:"bigDecimal,omitempty"`          // Whether the value is a big decimal.
	BigInteger          bool   `json:"bigInteger,omitempty"`          // Whether the value is a big integer.
	Textual             bool   `json:"textual,omitempty"`             // Whether the value is textual.
	Boolean             bool   `json:"boolean,omitempty"`             // Whether the value is boolean.
	Binary              bool   `json:"binary,omitempty"`              // Whether the value is binary.
	Float               bool   `json:"float,omitempty"`               // Whether the value is float.
}

// SCIMUserToPathOperationScheme represents the operation scheme for a path of a SCIM user.
type SCIMUserToPathOperationScheme struct {
	Op    string      `json:"op,omitempty"`    // The operation.
	Path  string      `json:"path,omitempty"`  // The path.
	Value interface{} `json:"value,omitempty"` // The value.
}
