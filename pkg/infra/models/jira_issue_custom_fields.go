package models

import (
	"bytes"
	"github.com/perimeterx/marshmallow"
)

func ExtractMultiSelectField(buffer bytes.Buffer, customField string) ([]*CustomFieldContextOptionScheme, bool, error) {

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
