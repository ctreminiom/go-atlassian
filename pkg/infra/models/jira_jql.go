package models

// ParsedQueryPageScheme represents a page of parsed queries in Jira.
// Queries is a slice of pointers to ParseQueryScheme which represents the parsed queries in the page.
type ParsedQueryPageScheme struct {
	Queries []*ParseQueryScheme `json:"queries"` // The parsed queries in the page.
}

// ParseQueryScheme represents a parsed query in Jira.
type ParseQueryScheme struct {
	Query     string `json:"query"` // The query.
	Structure struct {
		Where struct {
		} `json:"where"` // The where clause of the query.
		OrderBy *QueryStructureOrderScheme `json:"orderBy"` // The order by clause of the query.
	} `json:"structure"` // The structure of the query.
	Errors []string `json:"errors"` // The errors occurred during parsing the query.
}

// QueryStructureScheme represents the structure of a query in Jira.
type QueryStructureScheme struct {
	OrderBy *QueryStructureOrderScheme `json:"orderBy"` // The order by clause of the query.
}

// QueryStructureOrderScheme represents the order by clause of a query in Jira.
type QueryStructureOrderScheme struct {
	Fields []*QueryStructureOrderFieldScheme `json:"fields"` // The fields in the order by clause.
}

// QueryStructureOrderFieldScheme represents a field in the order by clause of a query in Jira.
type QueryStructureOrderFieldScheme struct {
	Field     *QueryStructureOrderFieldNodeScheme `json:"field"`     // The field node.
	Direction string                              `json:"direction"` // The direction of the order.
}

// QueryStructureOrderFieldNodeScheme represents a field node in the order by clause of a query in Jira.
type QueryStructureOrderFieldNodeScheme struct {
	Name     string                 `json:"name"`     // The name of the field.
	Property []*QueryPropertyScheme `json:"property"` // The properties of the field.
}

// QueryPropertyScheme represents a property of a field in the order by clause of a query in Jira.
type QueryPropertyScheme struct {
	Entity string `json:"entity"` // The entity of the property.
	Key    string `json:"key"`    // The key of the property.
	Path   string `json:"path"`   // The path of the property.
	Type   string `json:"type"`   // The type of the property.
}
