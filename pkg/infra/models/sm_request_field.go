package models

import (
	"encoding/json"
	"fmt"
	"github.com/imdario/mergo"
	"time"
)

type CreateCustomerRequestPayloadScheme struct {
	RequestParticipants []string `json:"requestParticipants,omitempty"`
	ServiceDeskID       string   `json:"serviceDeskId,omitempty"`
	RequestTypeID       string   `json:"requestTypeId,omitempty"`
}

func (c *CreateCustomerRequestPayloadScheme) MergeFields(fields *CustomerRequestFields) (map[string]interface{}, error) {

	if fields == nil {
		return nil, fmt.Errorf("error, please provide a value *CustomFields pointer")
	}

	if len(fields.Fields) == 0 {
		return nil, fmt.Errorf("error!, the Fields tag does not contains custom fields")
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

	//For each customField created, merge it into the eAsMap
	for _, field := range fields.Fields {

		if err := mergo.Merge(&issueSchemeAsMap, field, mergo.WithOverride); err != nil {
			return nil, err
		}
	}

	return issueSchemeAsMap, nil
}

type CustomerRequestFields struct{ Fields []map[string]interface{} }

func (c *CustomerRequestFields) Groups(customFieldID string, groups []string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(groups) == 0 {
		return fmt.Errorf("error, please provide a valid groups value")
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
	fieldsNode["requestFieldValues"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) Group(customFieldID, group string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(group) == 0 {
		return fmt.Errorf("error, please provide a valid group value")
	}

	var groupNode = map[string]interface{}{}
	groupNode["name"] = group

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = groupNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["requestFieldValues"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) URL(customFieldID, URL string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(URL) == 0 {
		return fmt.Errorf("error, please provide a valid URL value")
	}

	var urlNode = map[string]interface{}{}
	urlNode[customFieldID] = URL

	var fieldsNode = map[string]interface{}{}
	fieldsNode["requestFieldValues"] = urlNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) Text(customFieldID, textValue string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(textValue) == 0 {
		return fmt.Errorf("error, please provide a valid textValue value")
	}

	var urlNode = map[string]interface{}{}
	urlNode[customFieldID] = textValue

	var fieldsNode = map[string]interface{}{}
	fieldsNode["requestFieldValues"] = urlNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) DateTime(customFieldID string, dateValue time.Time) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if dateValue.IsZero() {
		return fmt.Errorf("error, please provide a valid dateValue value")
	}

	var dateNode = map[string]interface{}{}
	dateNode[customFieldID] = dateValue.Format(time.RFC3339)

	var fieldsNode = map[string]interface{}{}
	fieldsNode["requestFieldValues"] = dateNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) Date(customFieldID string, dateTimeValue time.Time) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if dateTimeValue.IsZero() {
		return fmt.Errorf("error, please provide a valid dateValue value")
	}

	var dateTimeNode = map[string]interface{}{}
	dateTimeNode[customFieldID] = dateTimeValue.Format("2006-01-02")

	var fieldsNode = map[string]interface{}{}
	fieldsNode["requestFieldValues"] = dateTimeNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) MultiSelect(customFieldID string, options []string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(options) == 0 {
		return fmt.Errorf("error, please provide a valid options value")
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
	fieldsNode["requestFieldValues"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) Select(customFieldID string, option string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(option) == 0 {
		return fmt.Errorf("error, please provide a valid option value")
	}

	var selectNode = map[string]interface{}{}
	selectNode["value"] = option

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = selectNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["requestFieldValues"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) RadioButton(customFieldID, button string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(button) == 0 {
		return fmt.Errorf("error, please provide a button option value")
	}

	var selectNode = map[string]interface{}{}
	selectNode["value"] = button

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = selectNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["requestFieldValues"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) User(customFieldID string, accountID string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(accountID) == 0 {
		return fmt.Errorf("error, please provide a accountID option value")
	}

	var userNode = map[string]interface{}{}
	userNode["accountId"] = accountID

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = userNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["requestFieldValues"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) Users(customFieldID string, accountIDs []string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(accountIDs) == 0 {
		return fmt.Errorf("error, please provide a accountIDs value")
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
	fieldsNode["requestFieldValues"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) Number(customFieldID string, numberValue float64) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	var urlNode = map[string]interface{}{}
	urlNode[customFieldID] = numberValue

	var fieldsNode = map[string]interface{}{}
	fieldsNode["requestFieldValues"] = urlNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) CheckBox(customFieldID string, options []string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(options) == 0 {
		return fmt.Errorf("error, please provide a valid options value")
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
	fieldsNode["requestFieldValues"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomerRequestFields) Cascading(customFieldID, parent, child string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(parent) == 0 {
		return fmt.Errorf("error, please provide a parent value")
	}

	if len(child) == 0 {
		return fmt.Errorf("error, please provide a child value")
	}

	var childNode = map[string]interface{}{}
	childNode["value"] = child

	var parentNode = map[string]interface{}{}
	parentNode["value"] = parent
	parentNode["child"] = childNode

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = parentNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["requestFieldValues"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}
