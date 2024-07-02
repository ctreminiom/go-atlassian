package models

type GroupUserPickerFindOptionScheme struct {
	// The search string.
	Query string `url:"query"`
	// The maximum number of items to return in each list.
	MaxResults int `url:"maxResults,omitempty"`
	// Whether the user avatar should be returned.
	ShowAvatar bool `url:"showAvatar,omitempty"`
	// The custom field ID of the field this request is for.
	FieldID string `url:"fieldId,omitempty"`
	// The ID of a project that returned users and groups must have permission to view.
	// This parameter is only used when FieldID is present.
	ProjectIDs []string `url:"projectId,omitempty"`
	// The ID of an issue type that returned users and groups must have permission to view.
	// Special values, such as -1 (all standard issue types) and -2 (all subtask issue types), are supported.
	// This parameter is only used when FieldID is present.
	IssueTypeIDs []string `url:"issueTypeId,omitempty"`
	// The size of the avatar to return.
	// Possible values are:
	// xsmall, xsmall@2x, xsmall@3x, small, small@2x, small@3x, medium, medium@2x, medium@3x, large,
	// large@2x, large@3x, xlarge, xlarge@2x, xlarge@3x, xxlarge, xxlarge@2x, xxlarge@3x, xxxlarge,
	// xxxlarge@2x, xxxlarge@3x
	AvatarSize string `url:"avatarSize,omitempty"`
	// Whether the search for groups should be case-insensitive.
	CaseInsensitive bool `url:"caseInsensitive,omitempty"`
	// Whether Connect app users and groups should be excluded from the search results.
	ExcludeConnectAddons bool `url:"excludeConnectAddons,omitempty"`
}

type GroupUserPickerFindScheme struct {
	// The list of groups found in a search, including header text (Showing X of Y matching groups)
	// and total of matched groups.
	Groups *GroupUserPickerFoundGroupsScheme `json:"groups"`
	Users  *GroupUserPickerFoundUsersScheme  `json:"users"`
}

type GroupUserPickerFoundGroupsScheme struct {
	Groups []*GroupUserPickerFoundGroupScheme `json:"groups"`
	// Header text indicating the number of groups in the response and the total number of groups found in the search.
	Header string `json:"header"`
	// The total number of groups found in the search.
	Total int `json:"total"`
}

type GroupUserPickerFoundGroupScheme struct {
	// The ID of the group, which uniquely identifies the group across all Atlassian products.
	// For example, 952d12c3-5b5b-4d04-bb32-44d383afc4b2.
	GroupID string `json:"groupId"`
	// The group name with the matched query string highlighted with the HTML bold tag.
	HTML   string                                  `json:"html"`
	Labels []*GroupUserPickerFoundGroupLabelScheme `json:"labels"`
	// The name of the group.
	// The name of a group is mutable, to reliably identify a group use GroupID.
	Name string `json:"name"`
}

type GroupUserPickerFoundGroupLabelScheme struct {
	// The group label name.
	Text string `json:"text"`
	// The title of the group label.
	Title string `json:"title"`
	// The type of the group label.
	// Valid values: ADMIN, SINGLE, MULTIPLE
	Type string `json:"type"`
}

type GroupUserPickerFoundUsersScheme struct {
	Users []*GroupUserPickerFoundUserScheme `json:"users"`
	// Header text indicating the number of groups in the response and the total number of groups found in the search.
	Header string `json:"header"`
	// The total number of groups found in the search.
	Total int `json:"total"`
}

type GroupUserPickerFoundUserScheme struct {
	// The account ID of the user, which uniquely identifies the user across all Atlassian products.
	// For example, 5b10ac8d82e05b22cc7d4ef5.
	AccountID string `json:"accountId"`
	// The avatar URL of the user.
	AvatarURL string `json:"avatarUrl"`
	//  The display name of the user. Depending on the user’s privacy setting, this may be returned as null.
	DisplayName string `json:"displayName"`
	// The display name, email address, and key of the user with the matched query string highlighted with the HTML bold tag.
	HTML string `json:"html"`
}
