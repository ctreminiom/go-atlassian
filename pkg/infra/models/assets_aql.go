package models

type AQLAssetSearchParamsScheme struct {
	Query               string                           `json:"qlQuery,omitempty"`
	ObjectTypeID        string                           `json:"objectTypeId,omitempty"`
	Page                int                              `json:"page,omitempty"`
	ResultsPerPage      int                              `json:"resultsPerPage,omitempty"`
	OrderByTypeAttrID   int                              `json:"orderByTypeAttrId,omitempty"`
	Asc                 int                              `json:"asc,omitempty"`
	ObjectID            string                           `json:"objectId,omitempty"`
	ObjectSchemaID      string                           `json:"objectSchemaId,omitempty"`
	IncludeAttributes   bool                             `json:"includeAttributes,omitempty"`
	AttributesToDisplay *ObjectAssetTypeAttributesScheme `json:"attributesToDisplay,omitempty"`
}

type ObjectAssetTypeAttributesScheme struct {
	AttributesToDisplayIds []int `json:"attributesToDisplayIds,omitempty"`
}

type ObjectAssetPageScheme struct {
	ObjectEntries         []*ObjectAssetScheme              `json:"objectEntries"`
	ObjectTypeAttributes  []*ObjectTypeAssetAttributeScheme `json:"objectTypeAttributes"`
	ObjectTypeId          int                               `json:"objectTypeId"`
	ObjectTypeIsInherited bool                              `json:"objectTypeIsInherited"`
	AbstractObjectType    bool                              `json:"abstractObjectType"`
	TotalFilterCount      int                               `json:"totalFilterCount"`
	StartIndex            int                               `json:"startIndex"`
	ToIndex               int                               `json:"toIndex"`
	PageObjectSize        int                               `json:"pageObjectSize"`
	PageNumber            int                               `json:"pageNumber"`
	OrderByTypeAttrId     int                               `json:"orderByTypeAttrId"`
	OrderWay              string                            `json:"orderWay"`
	QlQuery               string                            `json:"qlQuery"`
	QlQuerySearchResult   bool                              `json:"qlQuerySearchResult"`
	Iql                   string                            `json:"iql"`
	IqlSearchResult       bool                              `json:"iqlSearchResult"`
	ConversionPossible    bool                              `json:"conversionPossible"`
}
