package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

// ParseMultiSelectCustomField parses a multi-select custom field from the given buffer data
// associated with the specified custom field ID and returns a slice of pointers to
// CustomFieldContextOptionSchema structs.
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - []*CustomFieldContextOptionSchema: A slice of pointers to CustomFieldContextOptionSchema
//     structs representing the parsed options associated with the custom field.
//
// The ParseMultiSelectCustomField method is responsible for extracting and parsing the
// serialized data from the provided buffer, which is expected to be in a specific format.
// It then constructs and returns a slice of pointers to CustomFieldContextOptionSchema
// structs that represent the parsed options for the given custom field.
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	options, err := ParseMultiSelectCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, option := range options {
//	    fmt.Printf("Option ID: %s, Option Name: %s\n", option.ID, option.Name)
//	}
//
// NOTE: This method can be used to extract check-box customfield values
func ParseMultiSelectCustomField(buffer bytes.Buffer, customField string) ([]*CustomFieldContextOptionScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "issues" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectTypeError
	}

	var options []*CustomFieldContextOptionScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &options); err != nil {
		return nil, ErrNoMultiSelectTypeError
	}

	return options, nil
}

// ParseMultiSelectCustomFields extracts and parses multi-select custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a slice of CustomFieldContextOptionScheme
// structs, representing the parsed multi-select custom field values.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseMultiSelectCustomFields(customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    for _, value := range customFieldValues {
//	        fmt.Printf("Custom Field Value: %+v\n", value)
//	    }
//	}
//
// Parameters:
//   - customField: The name of the multi-select custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string][]*CustomFieldContextOptionScheme: A map where the key is the issue key and the
//     value is a slice of CustomFieldContextOptionScheme structs representing the parsed
//     multi-select custom field values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
func ParseMultiSelectCustomFields(buffer bytes.Buffer, customField string) (map[string][]*CustomFieldContextOptionScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string][]*CustomFieldContextOptionScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields []*CustomFieldContextOptionScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customfieldsAsMap, nil
}

// ParseMultiGroupPickerCustomField parses a group-picker custom field from the given buffer data
// associated with the specified custom field ID and returns a slice of pointers to
// GroupDetailScheme structs.
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - []*GroupDetailScheme: A slice of pointers to GroupDetailScheme
//     structs representing the parsed group picker associated with the custom field.
//
// The GroupDetailScheme method is responsible for extracting and parsing the
// serialized data from the provided buffer, which is expected to be in a specific format.
// It then constructs and returns a slice of pointers to GroupDetailScheme
// structs that represent the parsed groups for the given custom field.
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	groups, err := ParseMultiGroupPickerCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, group := range groups {
//	    fmt.Printf("Group ID: %s, Group Name: %s\n", group.GroupID, group.Name)
//	}
func ParseMultiGroupPickerCustomField(buffer bytes.Buffer, customField string) ([]*GroupDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "issues" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectTypeError
	}

	var options []*GroupDetailScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &options); err != nil {
		return nil, ErrNoMultiSelectTypeError
	}

	return options, nil
}

// ParseMultiGroupPickerCustomFields extracts and parses a group picker custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a slice of GroupDetailScheme
// structs, representing the parsed multi-select custom field values.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseMultiGroupPickerCustomFields(customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    for _, group := range customFieldValues {
//	        fmt.Printf("Custom Field Value: %+v\n", group)
//	    }
//	}
//
// Parameters:
//   - customField: The name of the multi-select custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string][]*GroupDetailScheme: A map where the key is the issue key and the
//     value is a slice of GroupDetailScheme structs representing the parsed
//     multi-select custom field values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
func ParseMultiGroupPickerCustomFields(buffer bytes.Buffer, customField string) (map[string][]*GroupDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string][]*GroupDetailScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields []*GroupDetailScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customfieldsAsMap, nil
}

// ParseMultiUserPickerCustomField parses a group-picker custom field from the given buffer data
// associated with the specified custom field ID and returns a slice of pointers to
// UserDetailScheme structs.
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - []*UserDetailScheme: A slice of pointers to UserDetailScheme
//     structs representing the parsed user picker associated with the custom field.
//
// The UserDetailScheme method is responsible for extracting and parsing the
// serialized data from the provided buffer, which is expected to be in a specific format.
// It then constructs and returns a slice of pointers to UserDetailScheme
// structs that represent the parsed groups for the given custom field.
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	users, err := ParseMultiUserPickerCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, user := range users {
//	    fmt.Printf("User ID: %s, User Name: %s\n", user.AccountID, user.DisplayName)
//	}
func ParseMultiUserPickerCustomField(buffer bytes.Buffer, customField string) ([]*UserDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "issues" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectTypeError
	}

	var users []*UserDetailScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &users); err != nil {
		return nil, ErrNoMultiSelectTypeError
	}

	return users, nil
}

// ParseMultiUserPickerCustomFields extracts and parses a user picker custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a slice of UserDetailScheme
// structs, representing the parsed multi-select custom field values.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseMultiUserPickerCustomFields(customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    for _, user := range customFieldValues {
//	        fmt.Printf("Custom Field Value: %+v\n", user)
//	    }
//	}
//
// Parameters:
//   - customField: The name of the multi-select custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string][]*UserDetailScheme: A map where the key is the issue key and the
//     value is a slice of UserDetailScheme structs representing the parsed
//     multi-select custom field values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
func ParseMultiUserPickerCustomFields(buffer bytes.Buffer, customField string) (map[string][]*UserDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string][]*UserDetailScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields []*UserDetailScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customfieldsAsMap, nil
}

// ParseCascadingSelectCustomField parses a cascading custom field from the given buffer data
// associated with the specified custom field ID and returns a CascadingSelectScheme struct pointer.
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - *CascadingSelectScheme: A pointer of the struct CascadingSelectScheme
//     representing the parsed cascading data associated with the custom field.
//
// The CascadingSelectScheme method is responsible for extracting and parsing the
// serialized data from the provided buffer, which is expected to be in a specific format.
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	field, err := ParseCascadingSelectCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, user := range users {
//	    fmt.Printf("Parent Value: %s, Child Value: %s\n", field.Value, field.Child.Value)
//	}
func ParseCascadingSelectCustomField(buffer bytes.Buffer, customField string) (*CascadingSelectScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "issues" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectTypeError
	}

	var cascading CascadingSelectScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &cascading); err != nil {
		return nil, ErrNoMultiSelectTypeError
	}

	return &cascading, nil
}

// ParseCascadingCustomFields extracts and parses a cascading custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a CascadingSelectScheme struct pointer
// ,representing the parsed cascading custom field value.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseCascadingCustomFields (customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    for _, data := range customFieldValues {
//	        fmt.Printf("Custom Field Value: %+v\n", data)
//	    }
//	}
//
// Parameters:
//   - customField: The name of the multi-select custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string]*CascadingSelectScheme: A map where the key is the issue key and the
//     value is a CascadingSelectScheme struct pointer representing the parsed
//     cascading custom field value.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
func ParseCascadingCustomFields(buffer bytes.Buffer, customField string) (map[string]*CascadingSelectScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string]*CascadingSelectScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields *CascadingSelectScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customfieldsAsMap, nil
}

// ParseMultiVersionCustomField parses a version-picker custom field from the given buffer data
// associated with the specified custom field ID and returns a slice of pointers to
// VersionDetailScheme structs.
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - []*VersionDetailScheme: A slice of pointers to VersionDetailScheme
//     structs representing the parsed group picker associated with the custom field.
//
// The VersionDetailScheme method is responsible for extracting and parsing the
// serialized data from the provided buffer, which is expected to be in a specific format.
// It then constructs and returns a slice of pointers to VersionDetailScheme
// structs that represent the parsed groups for the given custom field.
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	versions, err := ParseMultiVersionCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, version := range versions {
//	    fmt.Printf("Version ID: %s, Version Name: %s\n", version.ID, version.Name)
//	}
func ParseMultiVersionCustomField(buffer bytes.Buffer, customField string) ([]*VersionDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectTypeError
	}

	var versions []*VersionDetailScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &versions); err != nil {
		return nil, ErrNoMultiSelectTypeError
	}

	return versions, nil
}

// ParseMultiVersionCustomFields extracts and parses a version picker custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a slice of ParseMultiVersionCustomFields
// structs, representing the parsed multi-version custom field values.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseMultiVersionCustomFields(customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    for _, version := range customFieldValues {
//	        fmt.Printf("Custom Field Value: %+v\n", version)
//	    }
//	}
//
// Parameters:
//   - customField: The name of the multi-select custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string][]*ParseMultiVersionCustomFields: A map where the key is the issue key and the
//     value is a slice of ParseMultiVersionCustomFields structs representing the parsed
//     multi-select custom field values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
func ParseMultiVersionCustomFields(buffer bytes.Buffer, customField string) (map[string][]*VersionDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string][]*VersionDetailScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields []*VersionDetailScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customfieldsAsMap, nil
}

// ParseUserPickerCustomField parses a user custom field from the given buffer data
// associated with the specified custom field ID and pointer of the UserDetailScheme struct
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - *UserDetailScheme: A pointer of the UserDetailScheme struct
//     representing the parsed group picker associated with the custom field.
//
// The UserDetailScheme is responsible for extracting and parsing the
// serialized data from the provided buffer, which is expected to be in a specific format.
// It then constructs and returns a slice of pointers to UserDetailScheme
// structs that represent the parsed groups for the given custom field.
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	user, err := ParseUserPickerCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Printf("User ID: %s, User Name: %s\n", version.AccountID, version. DisplayName)
func ParseUserPickerCustomField(buffer bytes.Buffer, customField string) (*UserDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectTypeError
	}

	var user UserDetailScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &user); err != nil {
		return nil, ErrNoMultiSelectTypeError
	}

	return &user, nil
}

// ParseUserPickerCustomFields extracts and parses a user custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a ParseUserPickerCustomFields struct pointer
// ,representing the parsed cascading custom field value.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseUserPickerCustomFields(customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    for _, data := range customFieldValues {
//	        fmt.Printf("Custom Field Value: %+v\n", data)
//	    }
//	}
//
// Parameters:
//   - customField: The name of the multi-select custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string]*ParseUserPickerCustomFields: A map where the key is the issue key and the
//     value is a ParseUserPickerCustomFields struct pointer representing the parsed
//     cascading custom field value.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
func ParseUserPickerCustomFields(buffer bytes.Buffer, customField string) (map[string]*UserDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string]*UserDetailScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields *UserDetailScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customfieldsAsMap, nil
}

// ParseStringCustomField parses a textfield custom field from the given buffer data
// associated with the specified custom field ID and returns string
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - string: the customfield value as string type
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	textField, err := ParseStringCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(textField)
func ParseStringCustomField(buffer bytes.Buffer, customField string) (string, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return "", ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return "", ErrNoMultiSelectTypeError
	}

	var textField string
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &textField); err != nil {
		return "", ErrNoMultiSelectTypeError
	}

	return textField, nil
}

// ParseStringCustomFields extracts and parses the textfield customfield information from multiple issues using a bytes.Buffer.
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a string representing the parsed textfield customfield value.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseStringCustomFields(customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    for _, data := range customFieldValues {
//	        fmt.Printf("Custom Field Value: %+v\n", data)
//	    }
//	}
//
// Parameters:
//   - customField: The name of the multi-select custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string]*ParseUserPickerCustomFields: A map where the key is the issue key and the
//     value is a ParseUserPickerCustomFields struct pointer representing the parsed
//     cascading custom field value.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
func ParseStringCustomFields(buffer bytes.Buffer, customField string) (map[string]string, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string]string)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields string
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customfieldsAsMap, nil
}

// ParseStringCustomField parses a textfield custom field from the given buffer data
// associated with the specified custom field ID and returns string
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - string: the customfield value as string type
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	textField, err := ParseStringCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(textField)
func ParseFloatCustomField(buffer bytes.Buffer, customField string) (float64, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return 0, ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return 0, ErrNoMultiSelectTypeError
	}

	var textFloat float64
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &textFloat); err != nil {
		return 0, ErrNoMultiSelectTypeError
	}

	return textFloat, nil
}

// ParseFloatCustomFields extracts and parses the textfield customfield information from multiple issues using a bytes.Buffer.
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a string representing the parsed textfield customfield value.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseStringCustomFields(customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    for _, data := range customFieldValues {
//	        fmt.Printf("Custom Field Value: %+v\n", data)
//	    }
//	}
//
// Parameters:
//   - customField: The name of the multi-select custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string]*ParseUserPickerCustomFields: A map where the key is the issue key and the
//     value is a ParseUserPickerCustomFields struct pointer representing the parsed
//     cascading custom field value.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
func ParseFloatCustomFields(buffer bytes.Buffer, customField string) (map[string]float64, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string]float64)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields float64
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customfieldsAsMap, nil
}

// ParseLabelCustomField parses a textfield custom field from the given buffer data
// associated with the specified custom field ID and returns string
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - string: the customfield value as string type
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	textField, err := ParseStringCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(textField)
func ParseLabelCustomField(buffer bytes.Buffer, customField string) ([]string, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectTypeError
	}

	var labels []string
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &labels); err != nil {
		return nil, ErrNoMultiSelectTypeError
	}

	return labels, nil
}

// ParseLabelCustomFields extracts and parses the textfield customfield information from multiple issues using a bytes.Buffer.
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a string representing the parsed textfield customfield value.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseStringCustomFields(customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    for _, data := range customFieldValues {
//	        fmt.Printf("Custom Field Value: %+v\n", data)
//	    }
//	}
//
// Parameters:
//   - customField: The name of the multi-select custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string]*ParseUserPickerCustomFields: A map where the key is the issue key and the
//     value is a ParseUserPickerCustomFields struct pointer representing the parsed
//     cascading custom field value.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
func ParseLabelCustomFields(buffer bytes.Buffer, customField string) (map[string][]string, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string][]string)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields []string
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customfieldsAsMap, nil
}

// ParseSprintCustomField parses a sprints custom field from the given buffer data
// associated with the specified custom field ID and returns string
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - string: the customfield value as string type
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	textField, err := ParseStringCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(textField)
func ParseSprintCustomField(buffer bytes.Buffer, customField string) ([]*SprintDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectTypeError
	}

	var sprints []*SprintDetailScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &sprints); err != nil {
		return nil, ErrNoMultiSelectTypeError
	}

	return sprints, nil
}

// ParseSprintCustomFields extracts and parses multi-select custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a slice of CustomFieldContextOptionScheme
// structs, representing the parsed multi-select custom field values.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseMultiSelectCustomFields(customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    for _, value := range customFieldValues {
//	        fmt.Printf("Custom Field Value: %+v\n", value)
//	    }
//	}
//
// Parameters:
//   - customField: The name of the multi-select custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string][]*CustomFieldContextOptionScheme: A map where the key is the issue key and the
//     value is a slice of CustomFieldContextOptionScheme structs representing the parsed
//     multi-select custom field values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
func ParseSprintCustomFields(buffer bytes.Buffer, customField string) (map[string][]*SprintDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string][]*SprintDetailScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields []*SprintDetailScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customfieldsAsMap, nil
}

// ParseSelectCustomField parses a sprints custom field from the given buffer data
// associated with the specified custom field ID and returns string
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - string: the customfield value as string type
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	textField, err := ParseStringCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(textField)
func ParseSelectCustomField(buffer bytes.Buffer, customField string) (*CustomFieldContextOptionScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformationError
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectTypeError
	}

	var sprints *CustomFieldContextOptionScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &sprints); err != nil {
		return nil, ErrNoMultiSelectTypeError
	}

	return sprints, nil
}

// ParseSprintCustomFields extracts and parses multi-select custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a slice of CustomFieldContextOptionScheme
// structs, representing the parsed multi-select custom field values.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10001"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseMultiSelectCustomFields(customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    for _, value := range customFieldValues {
//	        fmt.Printf("Custom Field Value: %+v\n", value)
//	    }
//	}
//
// Parameters:
//   - customField: The name of the multi-select custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string][]*CustomFieldContextOptionScheme: A map where the key is the issue key and the
//     value is a slice of CustomFieldContextOptionScheme structs representing the parsed
//     multi-select custom field values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
func ParseSelectCustomFields(buffer bytes.Buffer, customField string) (map[string][]*CustomFieldContextOptionScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSliceError
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string][]*CustomFieldContextOptionScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields []*CustomFieldContextOptionScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValuesError
	}

	return customfieldsAsMap, nil
}
