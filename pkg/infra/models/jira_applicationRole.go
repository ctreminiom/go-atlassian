package models

// ApplicationRoleScheme represents an application role in Jira.
type ApplicationRoleScheme struct {
	Key                  string   `json:"key,omitempty"`                  // The key of the application role.
	Groups               []string `json:"groups,omitempty"`               // The groups associated with the application role.
	Name                 string   `json:"name,omitempty"`                 // The name of the application role.
	DefaultGroups        []string `json:"defaultGroups,omitempty"`        // The default groups of the application role.
	SelectedByDefault    bool     `json:"selectedByDefault,omitempty"`    // Indicates if the application role is selected by default.
	Defined              bool     `json:"defined,omitempty"`              // Indicates if the application role is defined.
	NumberOfSeats        int      `json:"numberOfSeats,omitempty"`        // The number of seats for the application role.
	RemainingSeats       int      `json:"remainingSeats,omitempty"`       // The remaining seats for the application role.
	UserCount            int      `json:"userCount,omitempty"`            // The user count for the application role.
	UserCountDescription string   `json:"userCountDescription,omitempty"` // The user count description for the application role.
	HasUnlimitedSeats    bool     `json:"hasUnlimitedSeats,omitempty"`    // Indicates if the application role has unlimited seats.
	Platform             bool     `json:"platform,omitempty"`             // Indicates if the application role is a platform role.
}
