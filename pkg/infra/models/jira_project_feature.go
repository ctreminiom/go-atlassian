package models

// ProjectFeaturesScheme represents the features of a project in Jira.
type ProjectFeaturesScheme struct {
	Features []*ProjectFeatureScheme `json:"features,omitempty"` // The features of the project.
}

// ProjectFeatureScheme represents a feature of a project in Jira.
type ProjectFeatureScheme struct {
	ProjectID            int      `json:"projectId,omitempty"`            // The ID of the project.
	State                string   `json:"state,omitempty"`                // The state of the feature.
	ToggleLocked         bool     `json:"toggleLocked,omitempty"`         // Indicates if the feature is locked.
	Feature              string   `json:"feature,omitempty"`              // The name of the feature.
	Prerequisites        []string `json:"prerequisites,omitempty"`        // The prerequisites of the feature.
	LocalisedName        string   `json:"localisedName,omitempty"`        // The localized name of the feature.
	LocalisedDescription string   `json:"localisedDescription,omitempty"` // The localized description of the feature.
	ImageURI             string   `json:"imageUri,omitempty"`             // The URI of the feature image.
}
