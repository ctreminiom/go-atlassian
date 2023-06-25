package models

type AQLSearchParamsScheme struct {

	// The query to determine which objects that should be fetched.
	//
	// E.g. objectType = "Computer". The empty AQL means all objects
	Query string `json:"qlQuery,omitempty"`

	// Which page to fetch when paginating through the response
	Page int `json:"page,omitempty"`

	// The amount of objects returned per page
	ResultPerPage int `json:"resultPerPage,omitempty"`

	// Should the objects attributes be included in the response.
	//
	// If this parameter is false only the information on the object will
	//
	// be returned and the object attributes will not be present
	IncludeAttributes bool `json:"includeAttributes,omitempty"`

	// How many levels of attributes should be included.
	//
	// E.g. consider an object A that has a reference to object B that has a reference to object C.
	//
	// If object A is included in the response and includeAttributesDeep=1 object A's reference to
	//
	// object B will be included in the attributes of object A but object B's reference to
	//
	// object C will not be included. However, if the includeAttributesDeep=2 then object B's
	//
	// reference to object C will be included in object B's attributes
	IncludeAttributesDeep bool `json:"includeAttributesDeep,omitempty"`

	// Should the response include the object type attribute definition
	//
	// for each attribute that is returned with the objects
	IncludeTypeAttributes bool `json:"includeTypeAttributes,omitempty"`

	// Include information about open Jira issues.
	//
	// Should each object have information if open tickets are connected to the object?
	IncludeExtendedInfo bool `json:"includeExtendedInfo,omitempty"`
}

type ObjectTypeAttributesScheme struct {
	AttributesToDisplayIds []int `json:"attributesToDisplayIds,omitempty"`
}

type ObjectPageScheme struct {
	ObjectEntries         []*ObjectScheme              `json:"objectEntries"`
	ObjectTypeAttributes  []*ObjectTypeAttributeScheme `json:"objectTypeAttributes"`
	ObjectTypeID          string                       `json:"objectTypeId"`
	ObjectTypeIsInherited bool                         `json:"objectTypeIsInherited"`
	AbstractObjectType    bool                         `json:"abstractObjectType"`
	TotalFilterCount      int                          `json:"totalFilterCount"`
	StartIndex            int                          `json:"startIndex"`
	ToIndex               int                          `json:"toIndex"`
	PageObjectSize        int                          `json:"pageObjectSize"`
	PageNumber            int                          `json:"pageNumber"`
	OrderByTypeAttrId     int                          `json:"orderByTypeAttrId"`
	OrderWay              string                       `json:"orderWay"`
	QlQuery               string                       `json:"qlQuery"`
	QlQuerySearchResult   bool                         `json:"qlQuerySearchResult"`
	Iql                   string                       `json:"iql"`
	IqlSearchResult       bool                         `json:"iqlSearchResult"`
	ConversionPossible    bool                         `json:"conversionPossible"`
}
