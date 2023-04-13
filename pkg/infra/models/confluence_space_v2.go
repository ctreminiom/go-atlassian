package models

type SpaceChunkV2Scheme struct {
	Results []*SpaceSchemeV2       `json:"results,omitempty"`
	Links   *SpacePageLinkSchemeV2 `json:"_links,omitempty"`
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
	Id          string                    `json:"id,omitempty"`
	Key         string                    `json:"key,omitempty"`
	Name        string                    `json:"name,omitempty"`
	Type        string                    `json:"type,omitempty"`
	Status      string                    `json:"status,omitempty"`
	HomepageId  string                    `json:"homepageId,omitempty"`
	Description *SpaceDescriptionSchemeV2 `json:"description,omitempty"`
}

type SpaceDescriptionSchemeV2 struct {
	Plain *PageBodyRepresentationScheme `json:"plain,omitempty"`
	View  *PageBodyRepresentationScheme `json:"view,omitempty"`
}
