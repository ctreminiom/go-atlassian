package models

// UpdateOperations represents a collection of update operations.
// Fields is a slice of maps, each containing a string key and an interface{} value.
type UpdateOperations struct{ Fields []map[string]interface{} }

// AddArrayOperation adds an array operation to the collection.
// It takes a custom field ID and a mapping of string to string as parameters.
// If the custom field ID is not provided, it returns an ErrNoFieldID.
// It creates an operation node for each value-operation pair in the mapping,
// with the operation as the key and the value as the value,
// and appends the operation node to the operations.
// It then creates a field node with the custom field ID as the key and the operations as the value,
// creates an update node with the "update" key and the field node as the value,
// and appends the update node to the Fields of the UpdateOperations.
// It returns nil if the operation is successful.
func (u *UpdateOperations) AddArrayOperation(customFieldID string, mapping map[string]string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldID
	}

	var operations []map[string]interface{}
	for value, operation := range mapping {

		var operationNode = map[string]interface{}{}
		operationNode[operation] = value

		operations = append(operations, operationNode)
	}

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = operations

	var updateNode = map[string]interface{}{}
	updateNode["update"] = fieldNode

	u.Fields = append(u.Fields, updateNode)
	return nil
}

// AddStringOperation adds a string operation to the collection.
// It takes a custom field ID, an operation, and a value as parameters.
// If the custom field ID is not provided, it returns an ErrNoFieldID.
// If the operation is not provided, it returns an ErrNoEditOperator.
// If the value is not provided, it returns an ErrNoEditValue.
// It creates an operation node with the operation as the key and the value as the value,
// appends the operation node to the operations,
// creates a field node with the custom field ID as the key and the operations as the value,
// creates an update node with the "update" key and the field node as the value,
// and appends the update node to the Fields of the UpdateOperations.
// It returns nil if the operation is successful.
func (u *UpdateOperations) AddStringOperation(customFieldID, operation, value string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldID
	}

	if len(operation) == 0 {
		return ErrNoEditOperator
	}

	if len(value) == 0 {
		return ErrNoEditValue
	}

	var operations []map[string]interface{}
	var operationNode = map[string]interface{}{}
	operationNode[operation] = value

	operations = append(operations, operationNode)

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = operations

	var updateNode = map[string]interface{}{}
	updateNode["update"] = fieldNode

	u.Fields = append(u.Fields, updateNode)

	return nil
}

// AddMultiRawOperation adds a multi raw operation to the collection.
// It takes a custom field ID and a slice of mappings as parameters.
// Each mapping is a map with string keys and interface{} values.
// If the custom field ID is not provided, it returns an ErrNoFieldID.
// It appends the mappings to the operations, creates a field node with the custom field ID and the operations,
// creates an update node with the "update" key and the field node, and appends the update node to the Fields of the UpdateOperations.
// It returns nil if the operation is successful.
func (u *UpdateOperations) AddMultiRawOperation(customFieldID string, mappings []map[string]interface{}) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldID
	}

	var operations []map[string]interface{}
	operations = append(operations, mappings...)

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = operations

	var updateNode = map[string]interface{}{}
	updateNode["update"] = fieldNode

	u.Fields = append(u.Fields, updateNode)
	return nil
}
