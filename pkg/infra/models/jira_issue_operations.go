package models

import "fmt"

type UpdateOperations struct{ Fields []map[string]interface{} }

func (u *UpdateOperations) AddArrayOperation(customFieldID string, mapping map[string]string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
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
	return
}

func (u *UpdateOperations) AddStringOperation(customFieldID, operation, value string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(operation) == 0 {
		return fmt.Errorf("error, please provide a valid operation value")
	}

	if len(value) == 0 {
		return fmt.Errorf("error, please provide a valid value value")
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

	return
}
