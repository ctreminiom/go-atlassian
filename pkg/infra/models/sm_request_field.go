package models

import (
	"time"
)

type CreateCustomerRequestPayloadScheme struct {
	Channel             string                                  `json:"channel,omitempty"`
	Form                *CreateCustomerRequestFormPayloadScheme `json:"form,omitempty"`
	IsAdfRequest        bool                                    `json:"isAdfRequest,omitempty"`
	RaiseOnBehalfOf     string                                  `json:"raiseOnBehalfOf,omitempty"`
	RequestFieldValues  map[string]interface{}                  `json:"requestFieldValues,omitempty"`
	RequestParticipants []string                                `json:"requestParticipants,omitempty"`
	RequestTypeID       string                                  `json:"requestTypeId,omitempty"`
	ServiceDeskID       string                                  `json:"serviceDeskId,omitempty"`
}

type CreateCustomerRequestFormPayloadScheme struct {
	Answers interface{} `json:"answers,omitempty"`
}

func (c *CreateCustomerRequestPayloadScheme) AddCustomField(key string, value interface{}) error {

	if c.RequestFieldValues == nil {
		c.RequestFieldValues = make(map[string]interface{})
	}

	c.RequestFieldValues[key] = value
	return nil
}

func (c *CreateCustomerRequestPayloadScheme) DateTimeCustomField(id string, value time.Time) error {

	if id == "" {
		return ErrNoCustomFieldIDError
	}

	if value.IsZero() {
		return ErrNoDatePickerTypeError
	}

	return c.AddCustomField(id, value.Format(time.RFC3339))
}

func (c *CreateCustomerRequestPayloadScheme) DateCustomField(id string, value time.Time) error {

	if id == "" {
		return ErrNoCustomFieldIDError
	}

	if value.IsZero() {
		return ErrNoDatePickerTypeError
	}

	return c.AddCustomField(id, value.Format("2006-01-02"))
}

func (c *CreateCustomerRequestPayloadScheme) MultiSelectOrCheckBoxCustomField(id string, values []string) error {

	if id == "" {
		return ErrNoCustomFieldIDError
	}

	if len(values) == 0 {
		return ErrNoMultiSelectTypeError
	}

	var options []map[string]interface{}
	for _, option := range values {
		optionNode := make(map[string]interface{})
		optionNode["value"] = option
		options = append(options, optionNode)
	}

	return c.AddCustomField(id, options)
}

func (c *CreateCustomerRequestPayloadScheme) UserCustomField(id, accountID string) error {

	if id == "" {
		return ErrNoCustomFieldIDError
	}

	if accountID == "" {
		return ErrNoUserTypeError
	}

	return c.AddCustomField(id, map[string]interface{}{"accountId": accountID})
}

func (c *CreateCustomerRequestPayloadScheme) UsersCustomField(id string, accountIDs []string) error {

	if id == "" {
		return ErrNoCustomFieldIDError
	}

	if len(accountIDs) == 0 {
		return ErrNoMultiUserTypeError
	}

	var accounts []map[string]interface{}
	for _, accountID := range accountIDs {

		if accountID == "" {
			return ErrNoUserTypeError
		}

		accounts = append(accounts, map[string]interface{}{"accountId": accountID})
	}

	return c.AddCustomField(id, accounts)
}

func (c *CreateCustomerRequestPayloadScheme) CascadingCustomField(id, parent, child string) error {

	if id == "" {
		return ErrNoCustomFieldIDError
	}

	if parent == "" {
		return ErrNoCascadingParentError
	}

	if child == "" {
		return ErrNoCascadingChildError
	}

	childNode := map[string]interface{}{"value": child}
	return c.AddCustomField(id, map[string]interface{}{"value": parent, "child": childNode})
}

func (c *CreateCustomerRequestPayloadScheme) GroupsCustomField(id string, names []string) error {

	if id == "" {
		return ErrNoCustomFieldIDError
	}

	if len(names) == 0 {
		return ErrNoGroupsNameError
	}

	var groups []map[string]interface{}
	for _, name := range names {

		if name == "" {
			return ErrNoGroupNameError
		}

		groups = append(groups, map[string]interface{}{"name": name})
	}

	return c.AddCustomField(id, groups)
}

func (c *CreateCustomerRequestPayloadScheme) GroupCustomField(id, name string) error {

	if id == "" {
		return ErrNoCustomFieldIDError
	}

	if name == "" {
		return ErrNoGroupNameError
	}

	return c.AddCustomField(id, map[string]interface{}{"name": name})
}

func (c *CreateCustomerRequestPayloadScheme) RadioButtonOrSelectCustomField(id string, option string) error {

	if id == "" {
		return ErrNoCustomFieldIDError
	}

	if option == "" {
		return ErrNoSelectTypeError
	}

	return c.AddCustomField(id, map[string]interface{}{"value": option})
}

func (c *CreateCustomerRequestPayloadScheme) Components(components []string) error {

	if len(components) == 0 {
		return ErrNoComponentsError
	}

	var values []map[string]interface{}
	for _, component := range components {

		if component == "" {
			return ErrNCoComponentError
		}

		values = append(values, map[string]interface{}{"name": component})
	}

	return c.AddCustomField("components", values)
}
