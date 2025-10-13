package models

// DescendantsScheme represents a chunk of descendants in Confluence.
type DescendantsScheme struct {
	Results []*DescendantScheme          `json:"results,omitempty"` // The descendants in the chunk.
	Links   *DescendantsChunkLinksScheme `json:"_links,omitempty"`  // The links of the chunk.
}

// DescendantsScheme represents a descendant in Confluence.
type DescendantScheme struct {
	ID            string `json:"id,omitempty"`            // The ID of the descendant.
	Status        string `json:"status,omitempty"`        // The status of the descendant.
	Title         string `json:"title,omitempty"`         // The title of the descendant.
	Type          string `json:"type,omitempty"`          // The type of the descendant.
	ChildPosition int    `json:"childPosition,omitempty"` // The position of the descendant.
}

// DescendantsChunkLinksScheme represents the links of a chunk of descendants in Confluence.
type DescendantsChunkLinksScheme struct {
	Next string `json:"next,omitempty"` // The link to the next chunk of descendants.
	Base string `json:"base,omitempty"` // The link to the base chunk of descendants.
}
