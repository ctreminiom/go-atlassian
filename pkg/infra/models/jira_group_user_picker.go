package models

// GroupUserPickerFindOptionScheme represents the options for finding users and groups in Jira.
type GroupUserPickerFindOptionScheme struct {
	// Query is the search string.
	Query string `url:"query"`
	// MaxResults is the maximum number of items to return in each list.
	MaxResults int `url:"maxResults,omitempty"`
	// ShowAvatar indicates whether the user avatar should be returned.
	ShowAvatar bool `url:"showAvatar,omitempty"`
	// FieldID is the custom field ID of the field this request is for.
	FieldID string `url:"fieldId,omitempty"`
	// ProjectIDs is the ID of a project that returned users and groups must have permission to view.
	// This parameter is only used when FieldID is present.
	ProjectIDs []string `url:"projectId,omitempty"`
	// IssueTypeIDs is the ID of an issue type that returned users and groups must have permission to view.
	// Special values, such as -1 (all standard issue types) and -2 (all subtask issue types), are supported.
	// This parameter is only used when FieldID is present.
	IssueTypeIDs []string `url:"issueTypeId,omitempty"`
	// AvatarSize is the size of the avatar to return.
	// Possible values are:
	// xsmall, xsmall@2x, xsmall@3x, small, small@2x, small@3x, medium, medium@2x, medium@3x, large,
	// large@2x, large@3x, xlarge, xlarge@2x, xlarge@3x, xxlarge, xxlarge@2x, xxlarge@3x, xxxlarge,
	// xxxlarge@2x, xxxlarge@3x
	AvatarSize string `url:"avatarSize,omitempty"`
	// CaseInsensitive indicates whether the search for groups should be case-insensitive.
	CaseInsensitive bool `url:"caseInsensitive,omitempty"`
	// ExcludeConnectAddons indicates whether Connect app users and groups should be excluded from the search results.
	ExcludeConnectAddons bool `url:"excludeConnectAddons,omitempty"`
}

// GroupUserPickerFindScheme represents the result of finding users and groups in Jira.
type GroupUserPickerFindScheme struct {
	// Groups is the list of groups found in a search, including header text (Showing X of Y matching groups)
	// and total of matched groups.
	Groups *GroupUserPickerFoundGroupsScheme `json:"groups"`
	// Users is the list of users found in a search.
	Users *GroupUserPickerFoundUsersScheme `json:"users"`
}

// GroupUserPickerFoundGroupsScheme represents the groups found in a search.
type GroupUserPickerFoundGroupsScheme struct {
	// Groups is the list of groups found in a search.
	Groups []*GroupUserPickerFoundGroupScheme `json:"groups"`
	// Header is the header text indicating the number of groups in the response and the total number of groups found in the search.
	Header string `json:"header"`
	// Total is the total number of groups found in the search.
	Total int `json:"total"`
}

// GroupUserPickerFoundGroupScheme represents a group found in a search.
type GroupUserPickerFoundGroupScheme struct {
	// GroupID is the ID of the group, which uniquely identifies the group across all Atlassian products.
	// For example, 952d12c3-5b5b-4d04-bb32-44d383afc4b2.
	GroupID string `json:"groupId"`
	// HTML is the group name with the matched query string highlighted with the HTML bold tag.
	HTML string `json:"html"`
	// Labels is the list of labels associated with the group.
	Labels []*GroupUserPickerFoundGroupLabelScheme `json:"labels"`
	// Name is the name of the group.
	// The name of a group is mutable, to reliably identify a group use GroupID.
	Name string `json:"name"`
}

// GroupUserPickerFoundGroupLabelScheme represents a label associated with a group.
type GroupUserPickerFoundGroupLabelScheme struct {
	// Text is the group label name.
	Text string `json:"text"`
	// Title is the title of the group label.
	Title string `json:"title"`
	// Type is the type of the group label.
	// Valid values: ADMIN, SINGLE, MULTIPLE
	Type string `json:"type"`
}

// GroupUserPickerFoundUsersScheme represents the users found in a search.
type GroupUserPickerFoundUsersScheme struct {
	// Users is the list of users found in a search.
	Users []*GroupUserPickerFoundUserScheme `json:"users"`
	// Header is the header text indicating the number of groups in the response and the total number of groups found in the search.
	Header string `json:"header"`
	// Total is the total number of groups found in the search.
	Total int `json:"total"`
}

// GroupUserPickerFoundUserScheme represents a user found in a search.
type GroupUserPickerFoundUserScheme struct {
	// AccountID is the account ID of the user, which uniquely identifies the user across all Atlassian products.
	// For example, 5b10ac8d82e05b22cc7d4ef5.
	AccountID string `json:"accountId"`
	// AvatarURL is the avatar URL of the user.
	AvatarURL string `json:"avatarUrl"`
	// DisplayName is the display name of the user. Depending on the user’s privacy setting, this may be returned as null.
	DisplayName string `json:"displayName"`
	// HTML is the display name, email address, and key of the user with the matched query string highlighted with the HTML bold tag.
	HTML string `json:"html"`
}
