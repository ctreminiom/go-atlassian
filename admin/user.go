package admin

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type UserService struct {
	client *Client
	Token  *UserTokenService
}

// Permissions returns the set of permissions you have for managing the specified Atlassian account, this func needs the following parameters:
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-account-id-manage-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/user#get-user-management-permissions
func (u *UserService) Permissions(ctx context.Context, accountID string, privileges []string) (result *UserPermissionScheme,
	response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, nil, notAccountError
	}

	params := url.Values{}
	if len(privileges) != 0 {
		params.Add("privileges", strings.Join(privileges, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/users/%v/manage", accountID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns information about a single Atlassian account by ID, this func needs the following parameters:
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-account-id-manage-profile-get
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/user#get-profile
func (u *UserService) Get(ctx context.Context, accountID string) (result *UserScheme, response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, nil, notAccountError
	}

	var endpoint = fmt.Sprintf("/users/%v/manage/profile", accountID)

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type UserScheme struct {
	Account struct {
		AccountID       string `json:"account_id"`
		Name            string `json:"name"`
		Nickname        string `json:"nickname"`
		Zoneinfo        string `json:"zoneinfo"`
		Locale          string `json:"locale"`
		Email           string `json:"email"`
		Picture         string `json:"picture"`
		ExtendedProfile struct {
			JobTitle string `json:"job_title"`
			TeamType string `json:"team_type"`
		} `json:"extended_profile"`
		AccountType     string `json:"account_type"`
		AccountStatus   string `json:"account_status"`
		EmailVerified   bool   `json:"email_verified"`
		PrivacySettings struct {
			Name                        string `json:"name"`
			Nickname                    string `json:"nickname"`
			Picture                     string `json:"picture"`
			ExtendedProfileJobTitle     string `json:"extended_profile.job_title"`
			ExtendedProfileDepartment   string `json:"extended_profile.department"`
			ExtendedProfileOrganization string `json:"extended_profile.organization"`
			ExtendedProfileLocation     string `json:"extended_profile.location"`
			ZoneInfo                    string `json:"zoneinfo"`
			Email                       string `json:"email"`
			ExtendedProfilePhoneNumber  string `json:"extended_profile.phone_number"`
			ExtendedProfileTeamType     string `json:"extended_profile.team_type"`
		} `json:"privacy_settings"`
	} `json:"account"`
}

type UserPermissionScheme struct {
	EmailSet struct {
		Allowed bool `json:"allowed"`
		Reason  struct {
			Key string `json:"key"`
		} `json:"reason"`
	} `json:"email.set"`
	LifecycleEnablement struct {
		Allowed bool `json:"allowed"`
		Reason  struct {
			Key string `json:"key"`
		} `json:"reason"`
	} `json:"lifecycle.enablement"`
	Profile struct {
		Name struct {
			Allowed bool `json:"allowed"`
			Reason  struct {
				Key string `json:"key"`
			} `json:"reason"`
		} `json:"name"`
		Nickname struct {
			Allowed bool `json:"allowed"`
		} `json:"nickname"`
		Zoneinfo struct {
			Allowed bool `json:"allowed"`
		} `json:"zoneinfo"`
		Locale struct {
			Allowed bool `json:"allowed"`
		} `json:"locale"`
		ExtendedProfilePhoneNumber struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.phone_number"`
		ExtendedProfileJobTitle struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.job_title"`
		ExtendedProfileOrganization struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.organization"`
		ExtendedProfileDepartment struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.department"`
		ExtendedProfileLocation struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.location"`
		ExtendedProfileTeamType struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.team_type"`
	} `json:"profile"`
	ProfileWrite struct {
		Name struct {
			Allowed bool `json:"allowed"`
			Reason  struct {
				Key string `json:"key"`
			} `json:"reason"`
		} `json:"name"`
		Nickname struct {
			Allowed bool `json:"allowed"`
		} `json:"nickname"`
		Zoneinfo struct {
			Allowed bool `json:"allowed"`
		} `json:"zoneinfo"`
		Locale struct {
			Allowed bool `json:"allowed"`
		} `json:"locale"`
		ExtendedProfilePhoneNumber struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.phone_number"`
		ExtendedProfileJobTitle struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.job_title"`
		ExtendedProfileOrganization struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.organization"`
		ExtendedProfileDepartment struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.department"`
		ExtendedProfileLocation struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.location"`
		ExtendedProfileTeamType struct {
			Allowed bool `json:"allowed"`
		} `json:"extended_profile.team_type"`
	} `json:"profile.write"`
	ProfileRead struct {
		Allowed bool `json:"allowed"`
	} `json:"profile.read"`
	LinkedAccountsRead struct {
		Allowed bool `json:"allowed"`
	} `json:"linkedAccounts.read"`
	APITokenRead struct {
		Allowed bool `json:"allowed"`
	} `json:"apiToken.read"`
	APITokenDelete struct {
		Allowed bool `json:"allowed"`
	} `json:"apiToken.delete"`
	Avatar struct {
		Allowed bool `json:"allowed"`
	} `json:"avatar"`
	PrivacySet struct {
		Allowed bool `json:"allowed"`
		Reason  struct {
			Key string `json:"key"`
		} `json:"reason"`
	} `json:"privacy.set"`
	SessionRead struct {
		Allowed bool `json:"allowed"`
	} `json:"session.read"`
}

// Update updates fields in a user account. The profile.write privilege details which fields you can change
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-account-id-manage-profile-patch
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/user#update-profile
func (u *UserService) Update(ctx context.Context, accountID string, payload map[string]interface{}) (
	result *UserScheme, response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, nil, notAccountError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	if len(payload) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid payload map with keys")
	}

	var endpoint = fmt.Sprintf("/users/%v/manage/profile", accountID)

	request, err := u.client.newRequest(ctx, http.MethodPatch, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Disable disables the specified user account.
// The permission to make use of this resource is exposed by the lifecycle.enablement privilege
// You can optionally set a message associated with the block that will be shown to the user on attempted authentication.
// If none is supplied, a default message will be used.
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-account-id-manage-lifecycle-disable-post
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/user#disable-a-user
func (u *UserService) Disable(ctx context.Context, accountID, message string) (response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, notAccountError
	}

	var (
		endpoint = fmt.Sprintf("/users/%v/manage/lifecycle/disable", accountID)
		request  *http.Request
	)

	if len(message) != 0 {

		payload := struct {
			Message string `json:"message"`
		}{
			Message: message,
		}

		payloadAsReader, _ := transformStructToReader(&payload)
		request, err = u.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
		if err != nil {
			return
		}

		request.Header.Set("Accept", "application/json")
		request.Header.Set("Content-Type", "application/json")

	} else {
		request, err = u.client.newRequest(ctx, http.MethodPost, endpoint, nil)
		if err != nil {
			return
		}

		request.Header.Set("Accept", "application/json")
	}

	response, err = u.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Enable enables the specified user account.
// The permission to make use of this resource is exposed by the lifecycle.enablement privilege.
// This func needs the following parameters:
// Atlassian Docs: https://developer.atlassian.com/cloud/admin/user-management/rest/api-group-users/#api-users-account-id-manage-lifecycle-enable-post
// Library Example: https://docs.go-atlassian.io/atlassian-admin-cloud/user#enable-a-user
func (u *UserService) Enable(ctx context.Context, accountID string) (response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, notAccountError
	}

	var endpoint = fmt.Sprintf("/users/%v/manage/lifecycle/enable", accountID)

	request, err := u.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	response, err = u.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

var (
	notAccountError = fmt.Errorf("error!, please provide a valid accountID value")
)
