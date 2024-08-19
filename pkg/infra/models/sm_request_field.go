package models

import (
	"time"
)

// CreateCustomerRequestPayloadScheme represents the payload for creating a customer request.
type CreateCustomerRequestPayloadScheme struct {
	Channel             string                                  `json:"channel,omitempty"`             // The channel through which the request is made.
	Form                *CreateCustomerRequestFormPayloadScheme `json:"form,omitempty"`                // The form associated with the request.
	IsAdfRequest        bool                                    `json:"isAdfRequest,omitempty"`        // Indicates if the request is an Atlassian Document Format (ADF) request.
	RaiseOnBehalfOf     string                                  `json:"raiseOnBehalfOf,omitempty"`     // The account ID of the user on whose behalf the request is raised.
	RequestFieldValues  map[string]interface{}                  `json:"requestFieldValues,omitempty"`  // The custom field values for the request.
	RequestParticipants []string                                `json:"requestParticipants,omitempty"` // The account IDs of the participants of the request.
	RequestTypeID       string                                  `json:"requestTypeId,omitempty"`       // The ID of the request type.
	ServiceDeskID       string                                  `json:"serviceDeskId,omitempty"`       // The ID of the service desk.
}

// CreateCustomerRequestFormPayloadScheme represents the form payload for creating a customer request.
type CreateCustomerRequestFormPayloadScheme struct {
	Answers interface{} `json:"answers,omitempty"` // The answers provided in the form.
}

// AddCustomField adds a custom field to the request payload.
// It takes a key which is the name of the custom field and a value which is the value of the custom field.
// If the RequestFieldValues map is not initialized, it initializes it.
// It returns an error if any occurs during the process.
func (c *CreateCustomerRequestPayloadScheme) AddCustomField(key string, value interface{}) error {

	if c.RequestFieldValues == nil {
		c.RequestFieldValues = make(map[string]interface{})
	}

	c.RequestFieldValues[key] = value
	return nil
}

// DateTimeCustomField adds a custom field of type DateTime to the request payload.
// It takes an id which is the name of the custom field and a value which is the value of the custom field.
// The value is formatted as RFC3339.
// It returns an error if the id is empty or the value is zero.
func (c *CreateCustomerRequestPayloadScheme) DateTimeCustomField(id string, value time.Time) error {

	if id == "" {
		return ErrNoCustomFieldID
	}

	if value.IsZero() {
		return ErrNoDatePickerType
	}

	return c.AddCustomField(id, value.Format(time.RFC3339))
}

// DateCustomField adds a custom field of type Date to the request payload.
// It takes an id which is the name of the custom field and a value which is the value of the custom field.
// The value is formatted as "2006-01-02".
// It returns an error if the id is empty or the value is zero.
func (c *CreateCustomerRequestPayloadScheme) DateCustomField(id string, value time.Time) error {

	if id == "" {
		return ErrNoCustomFieldID
	}

	if value.IsZero() {
		return ErrNoDatePickerType
	}

	return c.AddCustomField(id, value.Format("2006-01-02"))
}

// MultiSelectOrCheckBoxCustomField adds a custom field of type MultiSelect or CheckBox to the request payload.
// It takes an id which is the name of the custom field and values which are the values of the custom field.
// It returns an error if the id is empty or the values slice is empty.
func (c *CreateCustomerRequestPayloadScheme) MultiSelectOrCheckBoxCustomField(id string, values []string) error {

	if id == "" {
		return ErrNoCustomFieldID
	}

	if len(values) == 0 {
		return ErrNoMultiSelectType
	}

	var options []map[string]interface{}
	for _, option := range values {
		optionNode := make(map[string]interface{})
		optionNode["value"] = option
		options = append(options, optionNode)
	}

	return c.AddCustomField(id, options)
}

// UserCustomField adds a custom field of type User to the request payload.
// It takes an id which is the name of the custom field and an accountID which is the account ID of the user.
// It returns an error if the id or accountID is empty.
func (c *CreateCustomerRequestPayloadScheme) UserCustomField(id, accountID string) error {

	if id == "" {
		return ErrNoCustomFieldID
	}

	if accountID == "" {
		return ErrNoUserType
	}

	return c.AddCustomField(id, map[string]interface{}{"accountId": accountID})
}

// UsersCustomField adds a custom field of type Users to the request payload.
// It takes an id which is the name of the custom field and accountIDs which are the account IDs of the users.
// It returns an error if the id is empty or the accountIDs slice is empty or contains an empty string.
func (c *CreateCustomerRequestPayloadScheme) UsersCustomField(id string, accountIDs []string) error {

	if id == "" {
		return ErrNoCustomFieldID
	}

	if len(accountIDs) == 0 {
		return ErrNoMultiUserType
	}

	var accounts []map[string]interface{}
	for _, accountID := range accountIDs {

		if accountID == "" {
			return ErrNoUserType
		}

		accounts = append(accounts, map[string]interface{}{"accountId": accountID})
	}

	return c.AddCustomField(id, accounts)
}

// CascadingCustomField adds a custom field of type Cascading to the request payload.
// It takes an id which is the name of the custom field, a parent which is the parent value of the cascading field,
// and a child which is the child value of the cascading field.
// It returns an error if the id, parent, or child is empty.
func (c *CreateCustomerRequestPayloadScheme) CascadingCustomField(id, parent, child string) error {

	if id == "" {
		return ErrNoCustomFieldID
	}

	if parent == "" {
		return ErrNoCascadingParent
	}

	if child == "" {
		return ErrNoCascadingChild
	}

	childNode := map[string]interface{}{"value": child}
	return c.AddCustomField(id, map[string]interface{}{"value": parent, "child": childNode})
}

// GroupsCustomField adds a custom field of type Groups to the request payload.
// It takes an id which is the name of the custom field and names which are the names of the groups.
// It returns an error if the id is empty or the names slice is empty or contains an empty string.
func (c *CreateCustomerRequestPayloadScheme) GroupsCustomField(id string, names []string) error {

	if id == "" {
		return ErrNoCustomFieldID
	}

	if len(names) == 0 {
		return ErrNoGroupsName
	}

	var groups []map[string]interface{}
	for _, name := range names {

		if name == "" {
			return ErrNoGroupName
		}

		groups = append(groups, map[string]interface{}{"name": name})
	}

	return c.AddCustomField(id, groups)
}

// GroupCustomField adds a custom field of type Group to the request payload.
// It takes an id which is the name of the custom field and a name which is the name of the group.
// It returns an error if the id or name is empty.
func (c *CreateCustomerRequestPayloadScheme) GroupCustomField(id, name string) error {

	if id == "" {
		return ErrNoCustomFieldID
	}

	if name == "" {
		return ErrNoGroupName
	}

	return c.AddCustomField(id, map[string]interface{}{"name": name})
}

// RadioButtonOrSelectCustomField adds a custom field of type RadioButton or Select to the request payload.
// It takes an id which is the name of the custom field and an option which is the selected option.
// It returns an error if the id or option is empty.
func (c *CreateCustomerRequestPayloadScheme) RadioButtonOrSelectCustomField(id string, option string) error {

	if id == "" {
		return ErrNoCustomFieldID
	}

	if option == "" {
		return ErrNoSelectType
	}

	return c.AddCustomField(id, map[string]interface{}{"value": option})
}

// Components adds a custom field of type Components to the request payload.
// It takes components which are the names of the components.
// It returns an error if the components slice is empty or contains an empty string.
func (c *CreateCustomerRequestPayloadScheme) Components(components []string) error {

	if len(components) == 0 {
		return ErrNoComponents
	}

	var values []map[string]interface{}
	for _, component := range components {

		if component == "" {
			return ErrNCoComponent
		}

		values = append(values, map[string]interface{}{"name": component})
	}

	return c.AddCustomField("components", values)
}
