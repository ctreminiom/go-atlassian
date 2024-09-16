package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

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
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-multiselect-customfield
func ParseMultiSelectCustomField(buffer bytes.Buffer, customField string) ([]*CustomFieldContextOptionScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "issues" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectType
	}

	var options []*CustomFieldContextOptionScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &options); err != nil {
		return nil, ErrNoMultiSelectType
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
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-multiselect-customfields
func ParseMultiSelectCustomFields(buffer bytes.Buffer, customField string) (map[string][]*CustomFieldContextOptionScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
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
		return nil, ErrNoMapValues
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
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-grouppicker-customfield
func ParseMultiGroupPickerCustomField(buffer bytes.Buffer, customField string) ([]*GroupDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "issues" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectType
	}

	var options []*GroupDetailScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &options); err != nil {
		return nil, ErrNoMultiSelectType
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
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-grouppicker-customfields
func ParseMultiGroupPickerCustomFields(buffer bytes.Buffer, customField string) (map[string][]*GroupDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
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
		return nil, ErrNoMapValues
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
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-multi-userpicker-customfield
func ParseMultiUserPickerCustomField(buffer bytes.Buffer, customField string) ([]*UserDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "issues" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiSelectType
	}

	var users []*UserDetailScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &users); err != nil {
		return nil, ErrNoMultiSelectType
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
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-multi-userpicker-customfields
func ParseMultiUserPickerCustomFields(buffer bytes.Buffer, customField string) (map[string][]*UserDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
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
		return nil, ErrNoMapValues
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
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-cascading-customfield
func ParseCascadingSelectCustomField(buffer bytes.Buffer, customField string) (*CascadingSelectScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "issues" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoCascadingParent
	}

	var cascading CascadingSelectScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &cascading); err != nil {
		return nil, ErrNoCascadingParent
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
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-cascading-customfields
func ParseCascadingCustomFields(buffer bytes.Buffer, customField string) (map[string]*CascadingSelectScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
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
		return nil, ErrNoMapValues
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
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-versionpicker-customfield
func ParseMultiVersionCustomField(buffer bytes.Buffer, customField string) ([]*VersionDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoMultiVersionType
	}

	var versions []*VersionDetailScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &versions); err != nil {
		return nil, ErrNoMultiSelectType
	}

	return versions, nil
}

// ParseMultiVersionCustomFields extracts and parses a version picker custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a slice of VersionDetailScheme
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
//   - map[string][]*VersionDetailScheme: A map where the key is the issue key and the
//     value is a slice of VersionDetailScheme structs representing the parsed
//     multi-select custom field values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-versionpicker-customfields
func ParseMultiVersionCustomFields(buffer bytes.Buffer, customField string) (map[string][]*VersionDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
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
		return nil, ErrNoMapValues
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
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-userpicker-customfield
func ParseUserPickerCustomField(buffer bytes.Buffer, customField string) (*UserDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoUserType
	}

	var user UserDetailScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &user); err != nil {
		return nil, ErrNoMultiSelectType
	}

	return &user, nil
}

// ParseUserPickerCustomFields extracts and parses a user custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a UserDetailScheme struct pointer
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
//   - map[string]*UserDetailScheme: A map where the key is the issue key and the
//     value is a UserDetailScheme struct pointer representing the parsed
//     cascading custom field value.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-userpicker-customfields
func ParseUserPickerCustomFields(buffer bytes.Buffer, customField string) (map[string]*UserDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
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
		return nil, ErrNoMapValues
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
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-textfield-customfield
func ParseStringCustomField(buffer bytes.Buffer, customField string) (string, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return "", ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return "", ErrNoTextType
	}

	return raw.Get(path).String(), nil
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
//   - map[string]string: A map where the key is the issue key and the
//     value is a string with the customfield textfield value
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-textfield-customfields
func ParseStringCustomFields(buffer bytes.Buffer, customField string) (map[string]string, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
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

		// Check if the customfield value is a string type
		// if not, don't append the value on the map
		if reflect.TypeOf(value.Get(path).Value()).Kind() == reflect.String {
			customfieldsAsMap[issueKey] = value.Get(path).String()
		}
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValues
	}

	return customfieldsAsMap, nil
}

// ParseFloatCustomField parses a float custom field from the given buffer data
// associated with the specified custom field ID and returns float64 value
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - float64: the customfield value as float64 type
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	floatField, err := ParseFloatCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(floatField)
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-float-customfield
func ParseFloatCustomField(buffer bytes.Buffer, customField string) (float64, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return 0, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return 0, ErrNoFloatType
	}

	var textFloat float64
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &textFloat); err != nil {
		return 0, ErrNoFloatType
	}

	return textFloat, nil
}

// ParseFloatCustomFields extracts and parses the float customfield information from multiple issues using a bytes.Buffer.
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a float64 representing the parsed float customfield value.
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
//	customFields, err := ParseFloatCustomFields(customFieldName, buffer)
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
//   - map[string]float: A map where the key is the issue key and the
//     value is a float64 value.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-float-customfields
func ParseFloatCustomFields(buffer bytes.Buffer, customField string) (map[string]float64, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
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
		return nil, ErrNoMapValues
	}

	return customfieldsAsMap, nil
}

// ParseLabelCustomField parses a textfield slice custom field from the given buffer data
// associated with the specified custom field ID and returns string slice
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - []string: the customfield value as string slice type
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	labels, err := ParseLabelCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(labels)
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-labels-customfield
func ParseLabelCustomField(buffer bytes.Buffer, customField string) ([]string, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoLabelsType
	}

	var labels []string
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &labels); err != nil {
		return nil, ErrNoLabelsType
	}

	return labels, nil
}

// ParseLabelCustomFields extracts and parses the textfield customfield information from multiple issues using a bytes.Buffer.
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a string slice representing the parsed labels customfield value.
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
//	customFields, err := ParseLabelCustomFields(customFieldName, buffer)
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
//   - map[string][]string: A map where the key is the issue key and the
//     value is a string slice with the labels.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-labels-customfields
func ParseLabelCustomFields(buffer bytes.Buffer, customField string) (map[string][]string, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
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
		return nil, ErrNoMapValues
	}

	return customfieldsAsMap, nil
}

// ParseSprintCustomField parses a sprints custom field from the given buffer data
// associated with the specified custom field ID and returns the issue sprints
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
//	sprints, err := ParseSprintCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(sprints)
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-sprints-customfield
func ParseSprintCustomField(buffer bytes.Buffer, customField string) ([]*SprintDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoSprintType
	}

	var sprints []*SprintDetailScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &sprints); err != nil {
		return nil, ErrNoSprintType
	}

	return sprints, nil
}

// ParseSprintCustomFields extracts and parses sprint custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a slice of SprintDetailScheme
// structs, representing the parsed sprint custom field values.
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
//	customFields, err := ParseSprintCustomFields(customFieldName, buffer)
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
//   - customField: The name of the sprint custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string][]*SprintDetailScheme: A map where the key is the issue key and the
//     value is a slice of SprintDetailScheme structs representing the parsed
//     multi-select custom field values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-sprints-customfields
func ParseSprintCustomFields(buffer bytes.Buffer, customField string) (map[string][]*SprintDetailScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
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
		return nil, ErrNoMapValues
	}

	return customfieldsAsMap, nil
}

// ParseSelectCustomField parses a select custom field from the given buffer data
// associated with the specified custom field ID and returns a CustomFieldContextOptionScheme struct
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - CustomFieldContextOptionScheme: the customfield value as CustomFieldContextOptionScheme type
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	option, err := ParseSelectCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(option.ID, option.Value)
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-select-customfield
func ParseSelectCustomField(buffer bytes.Buffer, customField string) (*CustomFieldContextOptionScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoSelectType
	}

	var sprints *CustomFieldContextOptionScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &sprints); err != nil {
		return nil, ErrNoSelectType
	}

	return sprints, nil
}

// ParseSelectCustomFields extracts and parses select custom field data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a CustomFieldContextOptionScheme
// struct, representing the parsed select custom field value.
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
//	customFields, err := ParseSelectCustomFields(customFieldName, buffer)
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
//   - map[string]*CustomFieldContextOptionScheme: A map where the key is the issue key and the
//     value is a CustomFieldContextOptionScheme struct representing the parsed
//     multi-select custom field values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-select-customfields
func ParseSelectCustomFields(buffer bytes.Buffer, customField string) (map[string]*CustomFieldContextOptionScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string]*CustomFieldContextOptionScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields *CustomFieldContextOptionScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValues
	}

	return customfieldsAsMap, nil
}

// ParseAssetCustomField parses the Jira assets elements from the given buffer data
// associated with the specified custom field ID and returns a struct CustomFieldAssetScheme slice
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - []*CustomFieldAssetScheme: the customfield value as CustomFieldAssetScheme slice type
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	assets, err := ParseAssetCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(assets)
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-assets-customfield
func ParseAssetCustomField(buffer bytes.Buffer, customField string) ([]*CustomFieldAssetScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoAssetType
	}

	var assets []*CustomFieldAssetScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &assets); err != nil {
		return nil, ErrNoAssetType
	}

	return assets, nil
}

// ParseAssetCustomFields extracts and parses jira assets customfield data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a slice of CustomFieldAssetScheme
// structs, representing the parsed assets associated with a Jira issues.
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
//	customFields, err := ParseAssetCustomFields(customFieldName, buffer)
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
//   - customField: The name of the asset custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string][]*CustomFieldAssetScheme: A map where the key is the issue key and the
//     value is a CustomFieldAssetScheme struct representing the parsed
//     jira assets values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-assets-customfields
func ParseAssetCustomFields(buffer bytes.Buffer, customField string) (map[string][]*CustomFieldAssetScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string][]*CustomFieldAssetScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields []*CustomFieldAssetScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValues
	}

	return customfieldsAsMap, nil
}

// ParseDatePickerCustomField parses the datepicker customfield from the given buffer data
// associated with the specified custom field ID and returns a struct time.Time value
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - time.Time: the customfield value as time.Time type
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	datepicker, err := ParseDatePickerCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(datepicker)
//
// https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-datepicker-customfield
func ParseDatePickerCustomField(buffer bytes.Buffer, customField string) (time.Time, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "issues" object
	if !raw.Get("fields").Exists() {
		return time.Time{}, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return time.Time{}, ErrNoDatePickerType
	}

	value, err := time.Parse("2006-01-02", raw.Get(path).String())
	if err != nil {
		return time.Time{}, ErrNoDatePickerType
	}

	return value, nil
}

// ParseDatePickerCustomFields extracts and parses the datepicker customfield data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a time.Time value,
// representing the parsed datepicker values with the Jira issues.
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
//	customFields, err := ParseDatePickerCustomFields(customFieldName, buffer)
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
//   - map[string]time.Time: A map where the key is the issue key and the
//     value is a time.Time representing the parsed jira datepicker values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-datepicker-customfields
func ParseDatePickerCustomFields(buffer bytes.Buffer, customField string) (map[string]time.Time, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string]time.Time)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		datepicker, err := time.Parse("2006-01-02", value.Get(path).String())
		if err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = datepicker
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValues
	}

	return customfieldsAsMap, nil
}

// ParseDateTimeCustomField parses the datetime customfield from the given buffer data
// associated with the specified custom field ID and returns a struct time.Time value
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - time.Time: the customfield value as time.Time type
//
// Example usage:
//
//	customfieldID := "customfield_10001"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	datetime, err := ParseDateTimeCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(datetime)
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-datetime-customfield
func ParseDateTimeCustomField(buffer bytes.Buffer, customField string) (time.Time, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "issues" object
	if !raw.Get("fields").Exists() {
		return time.Time{}, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return time.Time{}, ErrNoDateTimeType
	}

	value, err := time.Parse("2006-01-02T15:04:05.000-0700", raw.Get(path).String())
	if err != nil {
		return time.Time{}, ErrNoDateTimeType
	}

	return value, nil
}

// ParseDateTimeCustomFields extracts and parses the datetime customfield data from a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a time.Time value,
// representing the parsed datetime values with the Jira issues.
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
//	customFields, err := ParseDateTimeCustomFields(customFieldName, buffer)
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
//   - map[string]time.Time: A map where the key is the issue key and the
//     value is a time.Time representing the parsed jira datetime values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-datetime-customfields
func ParseDateTimeCustomFields(buffer bytes.Buffer, customField string) (map[string]time.Time, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string]time.Time)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		datepicker, err := time.Parse("2006-01-02T15:04:05.000-0700", value.Get(path).String())
		if err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = datepicker
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValues
	}

	return customfieldsAsMap, nil
}

// ParseRequestTypeCustomField parses the Jira service desk request type elements from the given buffer
// data associated with the specified custom field ID and returns a struct CustomFieldRequestTypeScheme
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - *CustomFieldRequestTypeScheme: the customfield value as CustomFieldRequestTypeScheme type
//
// Example usage:
//
//	customfieldID := "customfield_10010"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	requestType, err := ParseRequestTypeCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(requestType)
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-requesttype-customfield
func ParseRequestTypeCustomField(buffer bytes.Buffer, customField string) (*CustomFieldRequestTypeScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoRequestType
	}

	var requestType *CustomFieldRequestTypeScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &requestType); err != nil {
		return nil, ErrNoRequestType
	}

	return requestType, nil
}

// ParseRequestTypeCustomFields extracts and parses jira service desk reqeust type customfield data from
// a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a slice of CustomFieldRequestTypeScheme
// structs, representing the parsed assets associated with a Jira issues.
//
// The JSON data within the buffer is expected to have a specific structure where the custom field
// values are organized by issue keys and options are represented within a context. The function
// parses this structure to extract and organize the custom field values.
//
// If the custom field data cannot be parsed successfully, an error is returned.
//
// Example Usage:
//
//	customFieldName := "customfield_10010"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseRequestTypeCustomFields(customFieldName, buffer)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Iterate through the parsed custom fields
//	for issueKey, customFieldValues := range customFields {
//	    fmt.Printf("Issue Key: %s\n", issueKey)
//	    fmt.Printf("Custom Field Value: %+v\n", customFieldValues)
//	}
//
// Parameters:
//   - customField: The name of the request type custom field to parse.
//   - buffer: A bytes.Buffer containing JSON data representing custom field values.
//
// Returns:
//   - map[string]*CustomFieldRequestTypeScheme: A map where the key is the issue key and the
//     value is a CustomFieldRequestTypeScheme struct representing the parsed
//     jira service desk request type values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-requesttype-customfields
func ParseRequestTypeCustomFields(buffer bytes.Buffer, customField string) (map[string]*CustomFieldRequestTypeScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string]*CustomFieldRequestTypeScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customFields *CustomFieldRequestTypeScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customFields); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customFields
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValues
	}

	return customfieldsAsMap, nil
}
