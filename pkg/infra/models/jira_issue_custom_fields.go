package models

import (
	"bytes"
	"github.com/perimeterx/marshmallow"
)

func ParseMultiSelectField(buffer bytes.Buffer, customField string) ([]*CustomFieldContextOptionScheme, bool, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, false, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, false, ErrNoFieldInformationError
	}

	customFields := fields.(map[string]interface{})
	var records []*CustomFieldContextOptionScheme

	switch options := customFields[customField].(type) {
	case []interface{}:

		for _, option := range options {

			record := &CustomFieldContextOptionScheme{
				ID:    option.(map[string]interface{})["id"].(string),
				Value: option.(map[string]interface{})["value"].(string),
			}

			isDisabled, wasFound := option.(map[string]interface{})["disabled"].(bool)
			if wasFound {
				record.Disabled = isDisabled
			}

			optionID, wasFound := option.(map[string]interface{})["optionId"].(string)
			if wasFound {
				record.OptionID = optionID
			}

			records = append(records, record)
		}

	case nil:
		return nil, false, nil
	default:
		return nil, false, ErrNoMultiSelectTypeError
	}

	return records, true, nil
}

func ParseMultiGroupPickerField(buffer bytes.Buffer, customField string) ([]*GroupDetailScheme, bool, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, false, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, false, ErrNoFieldInformationError
	}

	customFields := fields.(map[string]interface{})
	var groups []*GroupDetailScheme

	switch options := customFields[customField].(type) {
	case []interface{}:

		for _, option := range options {

			group := &GroupDetailScheme{
				Self:    option.(map[string]interface{})["self"].(string),
				Name:    option.(map[string]interface{})["name"].(string),
				GroupID: option.(map[string]interface{})["groupId"].(string),
			}

			groups = append(groups, group)
		}

	case nil:
		return nil, false, nil
	default:
		return nil, false, ErrNoMultiSelectTypeError
	}

	return groups, true, nil
}

func ParseMultiUserPickerField(buffer bytes.Buffer, customField string) ([]*UserDetailScheme, bool, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, false, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, false, ErrNoFieldInformationError
	}

	customFields := fields.(map[string]interface{})
	var users []*UserDetailScheme

	switch options := customFields[customField].(type) {
	case []interface{}:

		for _, option := range options {

			user := &UserDetailScheme{
				Self:        option.(map[string]interface{})["self"].(string),
				AccountID:   option.(map[string]interface{})["accountId"].(string),
				AccountType: option.(map[string]interface{})["accountType"].(string),
				DisplayName: option.(map[string]interface{})["displayName"].(string),
				Active:      option.(map[string]interface{})["active"].(bool),
				TimeZone:    option.(map[string]interface{})["timeZone"].(string),
			}

			email, wasFound := option.(map[string]interface{})["emailAddress"].(string)
			if wasFound {
				user.EmailAddress = email
			}

			users = append(users, user)
		}

	case nil:
		return nil, false, nil
	default:
		return nil, false, ErrNoMultiSelectTypeError
	}

	return users, true, nil
}

func ParseCascadingSelectField(buffer bytes.Buffer, customField string) (*CascadingSelectScheme, bool, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, false, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, false, ErrNoFieldInformationError
	}

	customFields := fields.(map[string]interface{})
	cascading := &CascadingSelectScheme{}

	switch option := customFields[customField].(type) {
	case map[string]interface{}:

		cascading = &CascadingSelectScheme{
			Self:  option["self"].(string),
			Value: option["value"].(string),
			Id:    option["id"].(string),
			Child: &CascadingSelectChildScheme{
				Self:  option["child"].(map[string]interface{})["self"].(string),
				Value: option["child"].(map[string]interface{})["value"].(string),
				Id:    option["child"].(map[string]interface{})["id"].(string),
			},
		}

	case nil:
		return nil, false, nil
	default:
		return nil, false, ErrNoMultiSelectTypeError
	}

	return cascading, false, nil
}
