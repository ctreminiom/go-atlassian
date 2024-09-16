package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

// ParseTempoAccountCustomField parses the Jira Tempo account type elements from the given buffer
// data associated with the specified custom field ID and returns a struct CustomFieldTempoAccountScheme
//
// Parameters:
//   - customfieldID: A string representing the unique identifier of the custom field.
//   - buffer: A bytes.Buffer containing the serialized data to be parsed.
//
// Returns:
//   - *CustomFieldTempoAccountScheme: the customfield value as CustomFieldTempoAccountScheme type
//
// Example usage:
//
//	customfieldID := "customfield_10038"
//	buffer := bytes.NewBuffer([]byte{ /* Serialized data */ })
//	tempoAccount, err := ParseTempoAccountCustomField(customfieldID, buffer)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(tempoAccount)
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-tempoaccount-customfield
func ParseTempoAccountCustomField(buffer bytes.Buffer, customField string) (*CustomFieldTempoAccountScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())
	path := fmt.Sprintf("fields.%v", customField)

	// Check if the buffer contains the "fields" object
	if !raw.Get("fields").Exists() {
		return nil, ErrNoFieldInformation
	}

	// Check if the issue iteration contains information on the customfield selected,
	// if not, continue
	if raw.Get(path).Type == gjson.Null {
		return nil, ErrNoTempoAccountType
	}

	var tempoAccount *CustomFieldTempoAccountScheme
	if err := json.Unmarshal([]byte(raw.Get(path).String()), &tempoAccount); err != nil {
		return nil, ErrNoTempoAccountType
	}

	return tempoAccount, nil
}

// ParseTempoAccountCustomFields extracts and parses jira tempo account type customfield data from
// a given bytes.Buffer from multiple issues
//
// This function takes the name of the custom field to parse and a bytes.Buffer containing
// JSON data representing the custom field values associated with different issues. It returns
// a map where the key is the issue key and the value is a slice of CustomFieldTempoAccountScheme
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
//	customFieldName := "customfield_10038"
//	buffer := // Populate the buffer with JSON data
//	customFields, err := ParseTempoAccountCustomFields(customFieldName, buffer)
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
//   - map[string]*ParseTempoAccountCustomFields: A map where the key is the issue key and the
//     value is a ParseTempoAccountCustomFields struct representing the parsed
//     jira tempo account type values.
//   - error: An error if there was a problem parsing the custom field data or if the JSON data
//     did not conform to the expected structure.
//
// Docs: https://docs.go-atlassian.io/cookbooks/extract-customfields-from-issue-s#parse-requesttype-customfields
func ParseTempoAccountCustomFields(buffer bytes.Buffer, customField string) (map[string]*CustomFieldTempoAccountScheme, error) {

	raw := gjson.ParseBytes(buffer.Bytes())

	// Check if the buffer contains the "issues" object
	if !raw.Get("issues").Exists() {
		return nil, ErrNoIssuesSlice
	}

	// Loop through each custom field, extract the information and stores the data on a map
	customfieldsAsMap := make(map[string]*CustomFieldTempoAccountScheme)
	raw.Get("issues").ForEach(func(key, value gjson.Result) bool {

		path, issueKey := fmt.Sprintf("fields.%v", customField), value.Get("key").String()

		// Check if the issue iteration contains information on the customfield selected,
		// if not, continue
		if value.Get(path).Type == gjson.Null {
			return true
		}

		var customField *CustomFieldTempoAccountScheme
		if err := json.Unmarshal([]byte(value.Get(path).String()), &customField); err != nil {
			return true
		}

		customfieldsAsMap[issueKey] = customField
		return true
	})

	// Check if the map processed contains elements
	// if so, return an error interface
	if len(customfieldsAsMap) == 0 {
		return nil, ErrNoMapValues
	}

	return customfieldsAsMap, nil
}
