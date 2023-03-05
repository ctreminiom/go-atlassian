package models

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

type SCIMUserPageScheme struct {
	Schemas      []string          `json:"schemas,omitempty"`
	TotalResults int               `json:"totalResults,omitempty"`
	StartIndex   int               `json:"startIndex,omitempty"`
	ItemsPerPage int               `json:"itemsPerPage,omitempty"`
	Resources    []*SCIMUserScheme `json:"Resources,omitempty"`
}

type SCIMUserToPathScheme struct {
	Schemas    []string                         `json:"schemas,omitempty"`
	Operations []*SCIMUserToPathOperationScheme `json:"operations,omitempty"`
}

func (s *SCIMUserToPathScheme) AddStringOperation(operation, path, value string) error {

	if operation == "" {
		return ErrNoSCIMOperationError
	}

	if path == "" {
		return ErrNoSCIMPathError
	}

	if value == "" {
		return ErrNoSCIMValueError
	}

	s.Operations = append(s.Operations, &SCIMUserToPathOperationScheme{
		Op:    operation,
		Path:  path,
		Value: value,
	})

	return nil
}

func (s *SCIMUserToPathScheme) AddBoolOperation(operation, path string, value bool) error {

	if operation == "" {
		return ErrNoSCIMOperationError
	}

	if path == "" {
		return ErrNoSCIMPathError
	}

	s.Operations = append(s.Operations, &SCIMUserToPathOperationScheme{
		Op:    operation,
		Path:  path,
		Value: value,
	})

	return nil
}

func (s *SCIMUserToPathScheme) AddComplexOperation(operation, path string, values []*SCIMUserComplexOperationScheme) error {

	if operation == "" {
		return ErrNoSCIMOperationError
	}

	if path == "" {
		return ErrNoSCIMPathError
	}

	if values == nil {
		return ErrNoSCIMComplexValueError
	}

	s.Operations = append(s.Operations, &SCIMUserToPathOperationScheme{
		Op:    operation,
		Path:  path,
		Value: values,
	})

	return nil
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
