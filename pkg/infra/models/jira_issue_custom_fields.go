package models

import (
	"bytes"
	"github.com/perimeterx/marshmallow"
)

func ParseMultiSelectCustomField(buffer bytes.Buffer, customField string) ([]*CustomFieldContextOptionScheme, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, ErrNoFieldInformationError
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
		return nil, nil
	default:
		return nil, ErrNoMultiSelectTypeError
	}

	return records, nil
}

func ParseMultiGroupPickerCustomField(buffer bytes.Buffer, customField string) ([]*GroupDetailScheme, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, ErrNoFieldInformationError
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
		return nil, nil
	default:
		return nil, ErrNoMultiSelectTypeError
	}

	return groups, nil
}

func ParseMultiUserPickerCustomField(buffer bytes.Buffer, customField string) ([]*UserDetailScheme, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, ErrNoFieldInformationError
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
		return nil, nil
	default:
		return nil, ErrNoMultiSelectTypeError
	}

	return users, nil
}

func ParseCascadingSelectCustomField(buffer bytes.Buffer, customField string) (*CascadingSelectScheme, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, ErrNoFieldInformationError
	}

	customFields := fields.(map[string]interface{})
	var cascading *CascadingSelectScheme

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
		return nil, nil
	default:
		return nil, ErrNoMultiSelectTypeError
	}

	return cascading, nil
}

func ParseMultiCheckboxesCustomField(buffer bytes.Buffer, customField string) ([]*CustomFieldContextOptionScheme, error) {
	return ParseMultiSelectCustomField(buffer, customField)
}

func ParseMultiVersionCustomField(buffer bytes.Buffer, customField string) ([]*VersionDetailScheme, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, ErrNoFieldInformationError
	}

	customFields := fields.(map[string]interface{})
	var records []*VersionDetailScheme

	switch options := customFields[customField].(type) {
	case []interface{}:

		for _, option := range options {

			record := &VersionDetailScheme{
				Self:        option.(map[string]interface{})["self"].(string),
				ID:          option.(map[string]interface{})["id"].(string),
				Description: option.(map[string]interface{})["description"].(string),
				Name:        option.(map[string]interface{})["name"].(string),
				Archived:    option.(map[string]interface{})["archived"].(bool),
				Released:    option.(map[string]interface{})["released"].(bool),
				ReleaseDate: option.(map[string]interface{})["releaseDate"].(string),
			}

			records = append(records, record)
		}

	case nil:
		return nil, nil
	default:
		return nil, ErrNoMultiSelectTypeError
	}

	return records, nil
}

func ParseUserPickerCustomField(buffer bytes.Buffer, customField string) (*UserDetailScheme, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, ErrNoFieldInformationError
	}

	customFields := fields.(map[string]interface{})
	var user *UserDetailScheme

	switch option := customFields[customField].(type) {
	case map[string]interface{}:

		user = &UserDetailScheme{
			Self:        option["self"].(string),
			AccountID:   option["accountId"].(string),
			DisplayName: option["displayName"].(string),
			Active:      option["active"].(bool),
			TimeZone:    option["timeZone"].(string),
			AccountType: option["accountType"].(string),
		}

		email, wasFound := option["emailAddress"].(string)
		if wasFound {
			user.EmailAddress = email
		}

	case nil:
		return nil, nil
	default:
		return nil, ErrNoMultiSelectTypeError
	}

	return user, nil
}

func ParseFloatCustomField(buffer bytes.Buffer, customField string) (float64, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return 0, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return 0, ErrNoFieldInformationError
	}
	var number float64
	customFields := fields.(map[string]interface{})

	switch value := customFields[customField].(type) {

	case float64:
		number = value
	case nil:
		return 0, nil
	default:
		return 0, ErrNoMultiSelectTypeError
	}

	return number, err
}

func ParseLabelCustomField(buffer bytes.Buffer, customField string) ([]string, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, ErrNoFieldInformationError
	}

	customFields := fields.(map[string]interface{})
	var labels []string
	switch value := customFields[customField].(type) {
	case []interface{}:
		for _, label := range value {
			labels = append(labels, label.(string))
		}
	case nil:
		return nil, nil
	default:
		return nil, ErrNoMultiSelectTypeError
	}

	return labels, err
}

func ParseSprintCustomField(buffer bytes.Buffer, customField string) ([]*SprintDetailScheme, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, ErrNoFieldInformationError
	}

	customFields := fields.(map[string]interface{})
	var records []*SprintDetailScheme

	switch options := customFields[customField].(type) {
	case []interface{}:

		for _, option := range options {

			record := &SprintDetailScheme{
				Name:          option.(map[string]interface{})["name"].(string),
				State:         option.(map[string]interface{})["state"].(string),
				ID:            int(option.(map[string]interface{})["id"].(float64)),
				StartDate:     option.(map[string]interface{})["startDate"].(string),
				EndDate:       option.(map[string]interface{})["endDate"].(string),
				OriginBoardID: int(option.(map[string]interface{})["boardId"].(float64)),
				Goal:          option.(map[string]interface{})["goal"].(string),
			}

			completeDate, wasFound := option.(map[string]interface{})["completeDate"].(string)
			if wasFound {
				record.CompleteDate = completeDate
			}
			records = append(records, record)
		}

	case nil:
		return nil, nil
	default:
		return nil, ErrNoMultiSelectTypeError
	}

	return records, nil
}

func ParseSelectCustomField(buffer bytes.Buffer, customField string) (*CustomFieldContextOptionScheme, error) {

	raw, err := marshmallow.Unmarshal(buffer.Bytes(), &struct{}{})
	if err != nil {
		return nil, ErrNoCustomFieldUnmarshalError
	}

	fields, containsFields := raw["fields"]
	if !containsFields {
		return nil, ErrNoFieldInformationError
	}

	customFields := fields.(map[string]interface{})
	var cascading *CustomFieldContextOptionScheme

	switch option := customFields[customField].(type) {
	case map[string]interface{}:

		cascading = &CustomFieldContextOptionScheme{
			ID:    option["id"].(string),
			Value: option["value"].(string),
		}

	case nil:
		return nil, nil
	default:
		return nil, ErrNoMultiSelectTypeError
	}

	return cascading, nil
}
