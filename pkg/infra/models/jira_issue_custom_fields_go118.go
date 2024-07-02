//go:build go1.18
// +build go1.18

package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

// ParseCustomField parses a generic custom field from the given buffer data associated
// with the specified custom field ID and returns a pointers to the generic T.
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - *T: A pointer to CustomFieldContextOptionSchema
//     structs representing the parsed generic custom field values.
//
// The ParseCustomField method is responsible for extracting and parsing the serialized
// data from the provided buffer, which is expected to be in a specific format.
// It then constructs and returns a pointer to T that represent the parsed generic
// custom field.
//
// Example usage:
//
//	type MyType struct {
//	    FieldInsideCustomField string
//	}
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	options, err := ParseCustomField[MyType](customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Custom Field Value: %+v\n", customFieldValue.FieldInsideCustomField)
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-customfield
func ParseCustomField[T any](buffer bytes.Buffer, customField string) (*T, error) {
	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "issues" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoCustomTypeError
	}

	var value *T
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &value); err != nil {
		return nil, ErrNoCustomTypeError
	}

	return value, nil
}

// ParseCustomFields extracts and parses generic custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a pointer of T a generic type,
// representing the parsed custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	type MyType struct {
//	    FieldInsideCustomField string
//	}
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseCustomFields[MyType](customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValue := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    fmt.Printf("Custom Field Value: %+v\n", customFieldValue.FieldInsideCustomField)
//	}
//
// Parameters:
//   - customField: The name of the generic custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string]*T: A map where the key is the issue key and the
//     value is a pointer of T representing the parsed generic custom field values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-customfields
func ParseCustomFields[T any](buffer bytes.Buffer, customField string) (map[string]*T, error) {
	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customFieldsAsMap := make(map[string]*T)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields *T
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customFieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customFieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customFieldsAsMap, nil
}
