// Package models provides the data structures used in the AQL search parameters.
package models

// AQLSearchParamsScheme represents the parameters for an AQL search.
type AQLSearchParamsScheme struct {

	// Query is the AQL query to determine which objects should be fetched.
	Query string `query:"qlQuery,omitempty"`

	// Page is the page number to fetch when paginating through the response.
	Page int `query:"page,omitempty"`

	// ResultPerPage is the number of objects returned per page.
	ResultPerPage int `query:"resultPerPage,omitempty"`

	// IncludeAttributes determines if the objects' attributes should be included in the response.
	IncludeAttributes bool `query:"includeAttributes,omitempty"`

	// IncludeAttributesDeep determines how many levels of attributes should be included.
	IncludeAttributesDeep int `query:"includeAttributesDeep,omitempty"`

	// IncludeTypeAttributes determines if the response should include the object type attribute definition for each attribute returned with the objects.
	IncludeTypeAttributes bool `query:"includeTypeAttributes,omitempty"`

	// IncludeExtendedInfo determines if each object should have information if open tickets are connected to the object.
	IncludeExtendedInfo bool `query:"includeExtendedInfo,omitempty"`
}

// ObjectPageScheme represents a page of objects in an AQL search.
type ObjectPageScheme struct {
	// ObjectEntries is a slice of the objects in the page.
	ObjectEntries []*ObjectScheme `json:"objectEntries"`

	// ObjectTypeAttributes is a slice of the object type attributes in the page.
	ObjectTypeAttributes []*ObjectTypeAttributeScheme `json:"objectTypeAttributes"`

	// ObjectTypeID is the ID of the object type.
	ObjectTypeID string `json:"objectTypeId"`

	// ObjectTypeIsInherited determines if the object type is inherited.
	ObjectTypeIsInherited bool `json:"objectTypeIsInherited"`

	// AbstractObjectType determines if the object type is abstract.
	AbstractObjectType bool `json:"abstractObjectType"`

	// TotalFilterCount is the total number of filters applied.
	TotalFilterCount int `json:"totalFilterCount"`

	// StartIndex is the starting index of the page.
	StartIndex int `json:"startIndex"`

	// ToIndex is the ending index of the page.
	ToIndex int `json:"toIndex"`

	// PageObjectSize is the number of objects in the page.
	PageObjectSize int `json:"pageObjectSize"`

	// PageNumber is the number of the page.
	PageNumber int `json:"pageNumber"`

	// OrderByTypeAttrID is the ID of the attribute used for ordering.
	OrderByTypeAttrID int `json:"orderByTypeAttrId"`

	// OrderWay is the way of ordering (ascending or descending).
	OrderWay string `json:"orderWay"`

	// QlQuery is the AQL query used for the search.
	QlQuery string `json:"qlQuery"`

	// QlQuerySearchResult determines if the result is from an AQL query search.
	QlQuerySearchResult bool `json:"qlQuerySearchResult"`

	// Iql is the IQL query used for the search.
	Iql string `json:"iql"`

	// IqlSearchResult determines if the result is from an IQL query search.
	IqlSearchResult bool `json:"iqlSearchResult"`

	// ConversionPossible determines if a conversion is possible.
	ConversionPossible bool `json:"conversionPossible"`
}
