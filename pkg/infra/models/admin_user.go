package models

type AdminUserScheme struct {
	Account *AdminUserAccountScheme `json:"account,omitempty"`
}

type AdminUserAccountScheme struct {
	AccountID       string                          `json:"account_id,omitempty"`
	Name            string                          `json:"name,omitempty"`
	Nickname        string                          `json:"nickname,omitempty"`
	ZoneInfo        string                          `json:"zoneinfo,omitempty"`
	Locale          string                          `json:"locale,omitempty"`
	Email           string                          `json:"email,omitempty"`
	Picture         string                          `json:"picture,omitempty"`
	AccountType     string                          `json:"account_type,omitempty"`
	AccountStatus   string                          `json:"account_status,omitempty"`
	EmailVerified   bool                            `json:"email_verified,omitempty"`
	ExtendedProfile *AdminUserExtendedProfileScheme `json:"extended_profile"`
	PrivacySettings *AdminUserPrivacySettingsScheme `json:"privacy_settings"`
}

type AdminUserExtendedProfileScheme struct {
	JobTitle string `json:"job_title,omitempty"`
	TeamType string `json:"team_type,omitempty"`
}

type AdminUserPrivacySettingsScheme struct {
	Name                        string `json:"name,omitempty"`
	Nickname                    string `json:"nickname,omitempty"`
	Picture                     string `json:"picture,omitempty"`
	ExtendedProfileJobTitle     string `json:"extended_profile.job_title,omitempty"`
	ExtendedProfileDepartment   string `json:"extended_profile.department,omitempty"`
	ExtendedProfileOrganization string `json:"extended_profile.organization,omitempty"`
	ExtendedProfileLocation     string `json:"extended_profile.location,omitempty"`
	ZoneInfo                    string `json:"zoneinfo,omitempty"`
	Email                       string `json:"email,omitempty"`
	ExtendedProfilePhoneNumber  string `json:"extended_profile.phone_number,omitempty"`
	ExtendedProfileTeamType     string `json:"extended_profile.team_type,omitempty"`
}

type AdminUserPermissionGrantScheme struct {
	Allowed bool                                  `json:"allowed,omitempty"`
	Reason  *AdminUserPermissionGrantReasonScheme `json:"reason,omitempty"`
}

type AdminUserPermissionGrantReasonScheme struct {
	Key string `json:"key,omitempty"`
}

type AdminUserPermissionScheme struct {
	EmailSet            *AdminUserPermissionGrantScheme   `json:"email.set,omitempty"`
	LifecycleEnablement *AdminUserPermissionGrantScheme   `json:"lifecycle.enablement,omitempty"`
	Profile             *AdminUserPermissionProfileScheme `json:"profile,omitempty"`
	ProfileWrite        *AdminUserPermissionProfileScheme `json:"profile.write,omitempty"`
	ProfileRead         *AdminUserPermissionGrantScheme   `json:"profile.read,omitempty"`
	LinkedAccountsRead  *AdminUserPermissionGrantScheme   `json:"linkedAccounts.read,omitempty"`
	APITokenRead        *AdminUserPermissionGrantScheme   `json:"apiToken.read,omitempty"`
	APITokenDelete      *AdminUserPermissionGrantScheme   `json:"apiToken.delete,omitempty"`
	Avatar              *AdminUserPermissionGrantScheme   `json:"avatar,omitempty"`
	PrivacySet          *AdminUserPermissionGrantScheme   `json:"privacy.set,omitempty"`
	SessionRead         *AdminUserPermissionGrantScheme   `json:"session.read,omitempty"`
}

type AdminUserPermissionProfileScheme struct {
	Name                        *AdminUserPermissionGrantScheme `json:"name,omitempty"`
	Nickname                    *AdminUserPermissionGrantScheme `json:"nickname,omitempty"`
	Zoneinfo                    *AdminUserPermissionGrantScheme `json:"zoneinfo,omitempty"`
	Locale                      *AdminUserPermissionGrantScheme `json:"locale,omitempty"`
	ExtendedProfilePhoneNumber  *AdminUserPermissionGrantScheme `json:"extended_profile.phone_number,omitempty"`
	ExtendedProfileJobTitle     *AdminUserPermissionGrantScheme `json:"extended_profile.job_title,omitempty"`
	ExtendedProfileOrganization *AdminUserPermissionGrantScheme `json:"extended_profile.organization,omitempty"`
	ExtendedProfileDepartment   *AdminUserPermissionGrantScheme `json:"extended_profile.department,omitempty"`
	ExtendedProfileLocation     *AdminUserPermissionGrantScheme `json:"extended_profile.location,omitempty"`
	ExtendedProfileTeamType     *AdminUserPermissionGrantScheme `json:"extended_profile.team_type,omitempty"`
}
