// Package models provides the data structures used in the admin user management.
package models

// AdminUserScheme represents an admin user account.
type AdminUserScheme struct {
	Account *AdminUserAccountScheme `json:"account,omitempty"` // The account details of the admin user.
}

// AdminUserAccountScheme represents the account details of an admin user.
type AdminUserAccountScheme struct {
	AccountID       string                          `json:"account_id,omitempty"`     // The account ID of the admin user.
	Name            string                          `json:"name,omitempty"`           // The name of the admin user.
	Nickname        string                          `json:"nickname,omitempty"`       // The nickname of the admin user.
	ZoneInfo        string                          `json:"zoneinfo,omitempty"`       // The timezone information of the admin user.
	Locale          string                          `json:"locale,omitempty"`         // The locale of the admin user.
	Email           string                          `json:"email,omitempty"`          // The email of the admin user.
	Picture         string                          `json:"picture,omitempty"`        // The picture of the admin user.
	AccountType     string                          `json:"account_type,omitempty"`   // The account type of the admin user.
	AccountStatus   string                          `json:"account_status,omitempty"` // The account status of the admin user.
	EmailVerified   bool                            `json:"email_verified,omitempty"` // Whether the email of the admin user is verified.
	ExtendedProfile *AdminUserExtendedProfileScheme `json:"extended_profile"`         // The extended profile of the admin user.
	PrivacySettings *AdminUserPrivacySettingsScheme `json:"privacy_settings"`         // The privacy settings of the admin user.
}

// AdminUserExtendedProfileScheme represents the extended profile of an admin user.
type AdminUserExtendedProfileScheme struct {
	JobTitle string `json:"job_title,omitempty"` // The job title of the admin user.
	TeamType string `json:"team_type,omitempty"` // The team type of the admin user.
}

// AdminUserPrivacySettingsScheme represents the privacy settings of an admin user.
type AdminUserPrivacySettingsScheme struct {
	Name                        string `json:"name,omitempty"`                          // The name privacy setting of the admin user.
	Nickname                    string `json:"nickname,omitempty"`                      // The nickname privacy setting of the admin user.
	Picture                     string `json:"picture,omitempty"`                       // The picture privacy setting of the admin user.
	ExtendedProfileJobTitle     string `json:"extended_profile.job_title,omitempty"`    // The job title privacy setting of the admin user.
	ExtendedProfileDepartment   string `json:"extended_profile.department,omitempty"`   // The department privacy setting of the admin user.
	ExtendedProfileOrganization string `json:"extended_profile.organization,omitempty"` // The organization privacy setting of the admin user.
	ExtendedProfileLocation     string `json:"extended_profile.location,omitempty"`     // The location privacy setting of the admin user.
	ZoneInfo                    string `json:"zoneinfo,omitempty"`                      // The timezone information privacy setting of the admin user.
	Email                       string `json:"email,omitempty"`                         // The email privacy setting of the admin user.
	ExtendedProfilePhoneNumber  string `json:"extended_profile.phone_number,omitempty"` // The phone number privacy setting of the admin user.
	ExtendedProfileTeamType     string `json:"extended_profile.team_type,omitempty"`    // The team type privacy setting of the admin user.
}

// AdminUserPermissionGrantScheme represents a permission grant of an admin user.
type AdminUserPermissionGrantScheme struct {
	Allowed bool                                  `json:"allowed,omitempty"` // Whether the permission is allowed.
	Reason  *AdminUserPermissionGrantReasonScheme `json:"reason,omitempty"`  // The reason for the permission grant.
}

// AdminUserPermissionGrantReasonScheme represents the reason for a permission grant of an admin user.
type AdminUserPermissionGrantReasonScheme struct {
	Key string `json:"key,omitempty"` // The key of the reason.
}

// AdminUserPermissionScheme represents the permissions of an admin user.
type AdminUserPermissionScheme struct {
	EmailSet            *AdminUserPermissionGrantScheme   `json:"email.set,omitempty"`            // The email set permission of the admin user.
	LifecycleEnablement *AdminUserPermissionGrantScheme   `json:"lifecycle.enablement,omitempty"` // The lifecycle enablement permission of the admin user.
	Profile             *AdminUserPermissionProfileScheme `json:"profile,omitempty"`              // The profile permission of the admin user.
	ProfileWrite        *AdminUserPermissionProfileScheme `json:"profile.write,omitempty"`        // The profile write permission of the admin user.
	ProfileRead         *AdminUserPermissionGrantScheme   `json:"profile.read,omitempty"`         // The profile read permission of the admin user.
	LinkedAccountsRead  *AdminUserPermissionGrantScheme   `json:"linkedAccounts.read,omitempty"`  // The linked accounts read permission of the admin user.
	APITokenRead        *AdminUserPermissionGrantScheme   `json:"apiToken.read,omitempty"`        // The API token read permission of the admin user.
	APITokenDelete      *AdminUserPermissionGrantScheme   `json:"apiToken.delete,omitempty"`      // The API token delete permission of the admin user.
	Avatar              *AdminUserPermissionGrantScheme   `json:"avatar,omitempty"`               // The avatar permission of the admin user.
	PrivacySet          *AdminUserPermissionGrantScheme   `json:"privacy.set,omitempty"`          // The privacy set permission of the admin user.
	SessionRead         *AdminUserPermissionGrantScheme   `json:"session.read,omitempty"`         // The session read permission of the admin user.
}

// AdminUserPermissionProfileScheme represents the profile permissions of an admin user.
type AdminUserPermissionProfileScheme struct {
	Name                        *AdminUserPermissionGrantScheme `json:"name,omitempty"`                          // The name permission of the admin user.
	Nickname                    *AdminUserPermissionGrantScheme `json:"nickname,omitempty"`                      // The nickname permission of the admin user.
	Zoneinfo                    *AdminUserPermissionGrantScheme `json:"zoneinfo,omitempty"`                      // The timezone information permission of the admin user.
	Locale                      *AdminUserPermissionGrantScheme `json:"locale,omitempty"`                        // The locale permission of the admin user.
	ExtendedProfilePhoneNumber  *AdminUserPermissionGrantScheme `json:"extended_profile.phone_number,omitempty"` // The phone number permission of the admin user.
	ExtendedProfileJobTitle     *AdminUserPermissionGrantScheme `json:"extended_profile.job_title,omitempty"`    // The job title permission of the admin user.
	ExtendedProfileOrganization *AdminUserPermissionGrantScheme `json:"extended_profile.organization,omitempty"` // The organization permission of the admin user.
	ExtendedProfileDepartment   *AdminUserPermissionGrantScheme `json:"extended_profile.department,omitempty"`   // The department permission of the admin user.
	ExtendedProfileLocation     *AdminUserPermissionGrantScheme `json:"extended_profile.location,omitempty"`     // The location permission of the admin user.
	ExtendedProfileTeamType     *AdminUserPermissionGrantScheme `json:"extended_profile.team_type,omitempty"`    // The team type permission of the admin user.
}
