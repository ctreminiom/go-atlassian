package models

type SpaceChunkV2Scheme struct {
	Results []*SpaceSchemeV2 `json:"results,omitempty"`
	Links   struct {
		Next string `json:"next"`
	} `json:"_links"`
}

type SpacePageLinkSchemeV2 struct {
	Next string `json:"next,omitempty"`
}

type GetSpacesOptionSchemeV2 struct {
	IDs               []int
	Keys              []string
	Type              string
	Status            string
	Labels            []string
	Sort              string
	DescriptionFormat string
	SerializeIDs      bool
}

type SpaceSchemeV2 struct {
	ID          int                       `json:"id,omitempty"`
	Key         string                    `json:"key,omitempty"`
	Name        string                    `json:"name,omitempty"`
	Type        string                    `json:"type,omitempty"`
	Status      string                    `json:"status,omitempty"`
	HomepageId  int                       `json:"homepageId,omitempty"`
	Description *SpaceDescriptionSchemeV2 `json:"description,omitempty"`
}

type SpaceDescriptionSchemeV2 struct {
	Plain *PageBodyRepresentationScheme `json:"plain,omitempty"`
	View  *PageBodyRepresentationScheme `json:"view,omitempty"`
}

type SpacePermissionPageScheme struct {
	Results []*SpacePermissionsV2Scheme    `json:"results,omitempty"`
	Links   *SpacePermissionPageLinkScheme `json:"_links,omitempty"`
}

type SpacePermissionsV2Scheme struct {
	ID        string                           `json:"id"`
	Principal *SpacePermissionsPrincipalScheme `json:"principal"`
	Operation *SpacePermissionsOperationScheme `json:"operation"`
}

type SpacePermissionsPrincipalScheme struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type SpacePermissionsOperationScheme struct {
	Key        string `json:"key,omitempty"`
	TargetType string `json:"targetType,omitempty"`
}

type SpacePermissionPageLinkScheme struct {
	Next string `json:"next,omitempty"`
}
