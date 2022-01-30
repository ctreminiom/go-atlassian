package models

type ProjectFeaturesScheme struct {
	Features []*ProjectFeatureScheme `json:"features,omitempty"`
}

type ProjectFeatureScheme struct {
	ProjectID            int      `json:"projectId,omitempty"`
	State                string   `json:"state,omitempty"`
	ToggleLocked         bool     `json:"toggleLocked,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	Prerequisites        []string `json:"prerequisites,omitempty"`
	LocalisedName        string   `json:"localisedName,omitempty"`
	LocalisedDescription string   `json:"localisedDescription,omitempty"`
	ImageURI             string   `json:"imageUri,omitempty"`
}
