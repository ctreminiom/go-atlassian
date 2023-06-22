package models

import (
	"encoding/json"
	"github.com/imdario/mergo"
	"time"
)

type CreateCustomerRequestPayloadScheme struct {
	RequestParticipants []string `json:"requestParticipants,omitempty"`
	ServiceDeskID       string   `json:"serviceDeskId,omitempty"`
	RequestTypeID       string   `json:"requestTypeId,omitempty"`
}

func (c *CreateCustomerRequestPayloadScheme) MergeFields(fields *CustomerRequestFields) (map[string]interface{}, error) {

	if fields == nil || len(fields.Fields) == 0 {
		return nil, ErrNoCustomFieldError
	}

	//Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	issueSchemeAsMap := make(map[string]interface{})
	if err := json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap); err != nil {
		return nil, err
	}

	// Create the requestFieldValues object with the custom fields
	fieldsAsMap := make(map[string]interface{})
	for key, value := range fields.Fields {
		fieldsAsMap[key] = value
	}

	requestFieldValues := map[string]interface{}{"requestFieldValues": fieldsAsMap}

	// Merge the requestFieldValues map struct with the service desk request struct information
	if err := mergo.Merge(&issueSchemeAsMap, requestFieldValues, mergo.WithOverride); err != nil {
		return nil, err
	}

	return issueSchemeAsMap, nil
}

type CustomerRequestFields struct{ Fields map[string]interface{} }

func (c *CustomerRequestFields) Attachments(attachments []string) error {

	if len(attachments) == 0 {
		return ErrNoAttachmentIdsError
	}

	c.Fields["attachments"] = attachments
	return nil
}

func (c *CustomerRequestFields) Labels(labels []string) error {

	if len(labels) == 0 {
		return ErrNoLabelsError
	}

	c.Fields["labels"] = labels
	return nil
}

func (c *CustomerRequestFields) Components(components []string) error {

	if len(components) == 0 {
		return ErrNoComponentsError
	}

	var componentNode []map[string]interface{}
	for _, component := range components {

		var groupNode = map[string]interface{}{}
		groupNode["name"] = component

		componentNode = append(componentNode, groupNode)
	}

	c.Fields["components"] = componentNode
	return nil
}

func (c *CustomerRequestFields) Groups(customFieldID string, groups []string) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
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

	c.Fields["components"] = fieldNode
	return nil
}

func (c *CustomerRequestFields) Group(customFieldID, group string) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
	}

	if len(group) == 0 {
		return ErrNoGroupNameError
	}

	var groupNode = map[string]interface{}{}
	groupNode["name"] = group

	c.Fields[customFieldID] = groupNode
	return nil
}

func (c *CustomerRequestFields) URL(customFieldID, URL string) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
	}

	if len(URL) == 0 {
		return ErrNoUrlTypeError
	}

	c.Fields[customFieldID] = URL
	return nil
}

func (c *CustomerRequestFields) Text(customFieldID, textValue string) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
	}

	if len(textValue) == 0 {
		return ErrNoTextTypeError
	}

	c.Fields[customFieldID] = textValue
	return nil
}

func (c *CustomerRequestFields) DateTime(customFieldID string, dateValue time.Time) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
	}

	if dateValue.IsZero() {
		return ErrNoDateTimeTypeError
	}

	c.Fields[customFieldID] = dateValue.Format(time.RFC3339)
	return nil
}

func (c *CustomerRequestFields) Date(customFieldID string, dateTimeValue time.Time) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
	}

	if dateTimeValue.IsZero() {
		return ErrNoDateTypeError
	}

	c.Fields[customFieldID] = dateTimeValue.Format("2006-01-02")
	return nil
}

func (c *CustomerRequestFields) MultiSelect(customFieldID string, options []string) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
	}

	if len(options) == 0 {
		return ErrNoMultiSelectTypeError
	}

	var multiSelectOptions []map[string]interface{}
	for _, group := range options {

		var groupNode = map[string]interface{}{}
		groupNode["value"] = group

		multiSelectOptions = append(multiSelectOptions, groupNode)
	}

	c.Fields[customFieldID] = multiSelectOptions
	return nil
}

func (c *CustomerRequestFields) Select(customFieldID string, option string) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
	}

	if len(option) == 0 {
		return ErrNoSelectTypeError
	}

	var selectNode = map[string]interface{}{}
	selectNode["value"] = option

	c.Fields[customFieldID] = selectNode
	return nil
}

func (c *CustomerRequestFields) RadioButton(customFieldID, button string) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
	}

	if len(button) == 0 {
		return ErrNoButtonTypeError
	}

	var selectNode = map[string]interface{}{}
	selectNode["value"] = button

	c.Fields[customFieldID] = selectNode
	return nil
}

func (c *CustomerRequestFields) User(customFieldID string, accountID string) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
	}

	if len(accountID) == 0 {
		return ErrNoUserTypeError
	}

	var userNode = map[string]interface{}{}
	userNode["accountId"] = accountID

	c.Fields[customFieldID] = userNode
	return nil
}

func (c *CustomerRequestFields) Users(customFieldID string, accountIDs []string) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
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

	c.Fields[customFieldID] = accountsNode
	return nil
}

func (c *CustomerRequestFields) Number(customFieldID string, numberValue float64) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
	}

	c.Fields[customFieldID] = numberValue
	return nil
}

func (c *CustomerRequestFields) CheckBox(customFieldID string, options []string) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
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

	c.Fields[customFieldID] = groupsNode
	return nil
}

func (c *CustomerRequestFields) Cascading(customFieldID, parent, child string) error {

	if len(customFieldID) == 0 {
		return ErrNoCustomFieldIDError
	}

	if len(parent) == 0 {
		return ErrNoCascadingParentError
	}

	if len(child) == 0 {
		return ErrNoCascadingChildError
	}

	var childNode = map[string]interface{}{}
	childNode["value"] = child

	var parentNode = map[string]interface{}{}
	parentNode["value"] = parent
	parentNode["child"] = childNode

	c.Fields[customFieldID] = parentNode
	return nil
}
