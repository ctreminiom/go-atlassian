// Package models provides the data structures used in the admin package.
package models

import (
	"time"
)

// CustomFields represents a collection of custom fields.
type CustomFields struct{ Fields []map[string]interface{} }

// Groups adds a group custom field to the collection.
func (c *CustomFields) Groups(customFieldID string, groups []string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if len(groups) == 0 {
		return ErrNoGroupsNameError
	}

	var groupsNode []map[string]interface{}
	for _, group := range groups {

		var groupNode = map[string]interface{}{}
		groupNode["name"] = group

		groupsNode = append(groupsNode, groupNode)
	}

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = groupsNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// Group adds a single group custom field to the collection.
func (c *CustomFields) Group(customFieldID, group string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if len(group) == 0 {
		return ErrNoGroupNameError
	}

	var groupNode = map[string]interface{}{}
	groupNode["name"] = group

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = groupNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// URL adds a URL custom field to the collection.
func (c *CustomFields) URL(customFieldID, URL string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if len(URL) == 0 {
		return ErrNoURLTypeError
	}

	var urlNode = map[string]interface{}{}
	urlNode[customFieldID] = URL

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = urlNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// Text adds a text custom field to the collection.
func (c *CustomFields) Text(customFieldID, textValue string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if len(textValue) == 0 {
		return ErrNoTextTypeError
	}

	var textNode = map[string]interface{}{}
	textNode[customFieldID] = textValue

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = textNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// DateTime adds a datetime custom field to the collection.
func (c *CustomFields) DateTime(customFieldID string, dateValue time.Time) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if dateValue.IsZero() {
		return ErrNoDatePickerTypeError
	}

	var dateNode = map[string]interface{}{}
	dateNode[customFieldID] = dateValue.Format(time.RFC3339)

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = dateNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// Date adds a date custom field to the collection.
func (c *CustomFields) Date(customFieldID string, dateTimeValue time.Time) (err error) {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if dateTimeValue.IsZero() {
		return ErrNoDateTimeTypeError
	}

	var dateTimeNode = map[string]interface{}{}
	dateTimeNode[customFieldID] = dateTimeValue.Format("2006-01-02")

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = dateTimeNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

// MultiSelect adds a multi-select custom field to the collection.
func (c *CustomFields) MultiSelect(customFieldID string, options []string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if len(options) == 0 {
		return ErrNoMultiSelectTypeError
	}

	var groupsNode []map[string]interface{}
	for _, group := range options {

		var groupNode = map[string]interface{}{}
		groupNode["value"] = group

		groupsNode = append(groupsNode, groupNode)
	}

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = groupsNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// Select adds a select custom field to the collection.
func (c *CustomFields) Select(customFieldID string, option string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if len(option) == 0 {
		return ErrNoSelectTypeError
	}

	var selectNode = map[string]interface{}{}
	selectNode["value"] = option

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = selectNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// RadioButton adds a radio button custom field to the collection.
func (c *CustomFields) RadioButton(customFieldID, button string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if len(button) == 0 {
		return ErrNoButtonTypeError
	}

	var selectNode = map[string]interface{}{}
	selectNode["value"] = button

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = selectNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// User adds a user custom field to the collection.
func (c *CustomFields) User(customFieldID string, accountID string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if len(accountID) == 0 {
		return ErrNoUserTypeError
	}

	var userNode = map[string]interface{}{}
	userNode["accountId"] = accountID

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = userNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// Users adds a multi-user custom field to the collection.
func (c *CustomFields) Users(customFieldID string, accountIDs []string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if len(accountIDs) == 0 {
		return ErrNoMultiUserTypeError
	}

	var accountsNode []map[string]interface{}
	for _, accountID := range accountIDs {

		var groupNode = map[string]interface{}{}
		groupNode["accountId"] = accountID

		accountsNode = append(accountsNode, groupNode)
	}

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = accountsNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// Number adds a number custom field to the collection.
func (c *CustomFields) Number(customFieldID string, numberValue float64) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	var numberNode = map[string]interface{}{}
	numberNode[customFieldID] = numberValue

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = numberNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// CheckBox adds a checkbox custom field to the collection.
func (c *CustomFields) CheckBox(customFieldID string, options []string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if len(options) == 0 {
		return ErrNoCheckBoxTypeError
	}

	var groupsNode []map[string]interface{}
	for _, group := range options {

		var groupNode = map[string]interface{}{}
		groupNode["value"] = group

		groupsNode = append(groupsNode, groupNode)
	}

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = groupsNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// Cascading adds a cascading custom field to the collection.
func (c *CustomFields) Cascading(customFieldID, parent, child string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if parent == "" {
		return ErrNoCascadingParentError
	}

	if child == "" {
		return ErrNoCascadingChildError
	}

	var childNode = map[string]interface{}{}
	childNode["value"] = child

	var parentNode = map[string]interface{}{}
	parentNode["value"] = parent
	parentNode["child"] = childNode

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = parentNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}

// Raw adds an untyped field to the collection.
func (c *CustomFields) Raw(customFieldID string, value interface{}) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if value == nil {
		return ErrNoValueTypeError
	}

	var valueNode = map[string]interface{}{}
	valueNode[customFieldID] = value

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = valueNode

	c.Fields = append(c.Fields, fieldsNode)
	return nil
}
