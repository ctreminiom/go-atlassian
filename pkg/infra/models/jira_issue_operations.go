package models

type UpdateOperations struct{ Fields []map[string]interface{} }

func (u *UpdateOperations) AddArrayOperation(customFieldID string, mapping map[string]string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
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

func (u *UpdateOperations) AddStringOperation(customFieldID, operation, value string) error {

	if len(customFieldID) == 0 {
		return ErrNoFieldIDError
	}

	if len(operation) == 0 {
		return ErrNoEditOperatorError
	}

	if len(value) == 0 {
		return ErrNoEditValueError
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
