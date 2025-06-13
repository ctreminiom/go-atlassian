package models

// UpdateTemplateScheme represents a Confluence template update.
type UpdateTemplateScheme struct {
	// TemplateID the ID of the template being updated.
	TemplateID string `json:"templateId"`
	// Name of the template. Set to the current name if this field is not being updated.
	Name string `json:"name"`
	// TemplateType of the template.
	TemplateType string `json:"templateType"`
	// Body of the new content.
	Body *ContentTemplateBodyCreateScheme `json:"body"`
	// Description of the template.
	// Max length: 100
	Description string `json:"description"`
	// Labels to add to the template.
	Labels []LabelValueScheme `json:"labels,omitempty"`
	// Space the template is in.
	// Only the Space.Key is required.
	Space *SpaceScheme `json:"space"`
}

// CreateTemplateScheme represents a Confluence template creation.
type CreateTemplateScheme struct {
	// Name of the template. Set to the current name if this field is not being updated.
	Name string `json:"name"`
	// TemplateType of the template.
	TemplateType string `json:"templateType"`
	// Body of the new content.
	Body *ContentTemplateBodyCreateScheme `json:"body"`
	// Description of the template.
	// Max length: 255
	Description string `json:"description"`
	// Labels to add to the template.
	Labels []LabelValueScheme `json:"labels,omitempty"`
	// Space the template is in.
	// Only the Space.Key is required.
	Space *SpaceScheme `json:"space"`
}

// ContentTemplateBodySchema is the body of the template.
type ContentTemplateBodySchema struct {
	View                *ContentBodyCreateScheme `json:"view,omitempty"`
	ExportView          *ContentBodyCreateScheme `json:"export_view,omitempty"`
	StyledView          *ContentBodyCreateScheme `json:"styled_view,omitempty"`
	Storage             *ContentBodyCreateScheme `json:"storage,omitempty"`
	Editor              *ContentBodyCreateScheme `json:"editor,omitempty"`
	Editor2             *ContentBodyCreateScheme `json:"editor2,omitempty"`
	Wiki                *ContentBodyCreateScheme `json:"wiki,omitempty"`
	AtlasDocFormat      *ContentBodyCreateScheme `json:"atlas_doc_format,omitempty"`
	AnonymousExportView *ContentBodyCreateScheme `json:"anonymous_export_view,omitempty"`
}

// ContentBodyScheme is used when creating or updating content.
type ContentBodyScheme struct {
	// The body of the content in the relevant format.
	Value string `json:"value"`
	// The content format type. Set the value of this property to the name of the format being used, e.g. 'storage'.
	// Valid values: view, export_view, styled_view, storage, editor, editor2, anonymous_export_view, wiki, atlas_doc_format, plain, raw
	Representation string `json:"representation"`
}

// ContentTemplateBodyCreateScheme is the body of the template for creating.
// Only one body format should be specified as the property for this object, e.g. storage.
// Note, editor2 format is used by Atlassian only. anonymous_export_view is the same as export_view format but only content viewable by an anonymous user is included.
type ContentTemplateBodyCreateScheme struct {
	View                *ContentBodyCreateScheme `json:"view,omitempty"`
	ExportView          *ContentBodyCreateScheme `json:"export_view,omitempty"`
	StyledView          *ContentBodyCreateScheme `json:"styled_view,omitempty"`
	Storage             *ContentBodyCreateScheme `json:"storage,omitempty"`
	Editor              *ContentBodyCreateScheme `json:"editor,omitempty"`
	Editor2             *ContentBodyCreateScheme `json:"editor2,omitempty"`
	Wiki                *ContentBodyCreateScheme `json:"wiki,omitempty"`
	AtlasDocFormat      *ContentBodyCreateScheme `json:"atlas_doc_format,omitempty"`
	AnonymousExportView *ContentBodyCreateScheme `json:"anonymous_export_view,omitempty"`
}

// ContentBodyCreateScheme is used when creating or updating content.
type ContentBodyCreateScheme struct {
	// The body of the content in the relevant format.
	Value string `json:"value"`
	// The content format type. Set the value of this property to the name of the format being used, e.g. 'storage'.
	// Valid values: view, export_view, styled_view, storage, editor, editor2, anonymous_export_view, wiki, atlas_doc_format, plain, raw
	Representation string `json:"representation"`
}

// ContentTemplateScheme represents a Confluence template.
type ContentTemplateScheme struct {
	TemplateID           string                  `json:"templateId"`
	OriginalTemplate     *OriginalTemplateScheme `json:"originalTemplate,omitempty"`
	ReferencingBlueprint string                  `json:"referencingBlueprint"`
	Name                 string                  `json:"name"`
	Description          string                  `json:"description"`
	// Labels to add to the template.
	Labels []LabelValueScheme `json:"labels"`
	// TemplateType of the template.
	TemplateType  string                           `json:"templateType"`
	EditorVersion string                           `json:"editorVersion"`
	Body          *ContentTemplateBodySchema       `json:"body"`
	Expandable    *ContentTemplateExpandableScheme `json:"_expandable,omitempty"`
}

// OriginalTemplateScheme contains the original template reference.
type OriginalTemplateScheme struct {
	PluginKey string `json:"pluginKey"`
	ModuleKey string `json:"moduleKey"`
}

// ContentTemplateExpandableScheme represents the expandable properties of a template.
type ContentTemplateExpandableScheme struct {
	Body string `json:"body"`
}
