package models

type ParsedQueryPageScheme struct {
	Queries []*ParseQueryScheme `json:"queries"`
}

type ParseQueryScheme struct {
	Query     string `json:"query"`
	Structure struct {
		Where struct {
		} `json:"where"`
		OrderBy *QueryStructureOrderScheme `json:"orderBy"`
	} `json:"structure"`
	Errors []string `json:"errors"`
}

type QueryStructureScheme struct {
	OrderBy *QueryStructureOrderScheme `json:"orderBy"`
}

type QueryStructureOrderScheme struct {
	Fields []*QueryStructureOrderFieldScheme `json:"fields"`
}

type QueryStructureOrderFieldScheme struct {
	Field     *QueryStructureOrderFieldNodeScheme `json:"field"`
	Direction string                              `json:"direction"`
}

type QueryStructureOrderFieldNodeScheme struct {
	Name     string                 `json:"name"`
	Property []*QueryPropertyScheme `json:"property"`
}

type QueryPropertyScheme struct {
	Entity string `json:"entity"`
	Key    string `json:"key"`
	Path   string `json:"path"`
	Type   string `json:"type"`
}
