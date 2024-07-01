package models

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestUpdateOperations_AddArrayOperation(t *testing.T) {

	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		mapping       map[string]string
	}
	testCases := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
	}{
		{
			name:   "when the values are correct",
			fields: fields{},
			args: args{
				customFieldID: "custom_field_id",
				mapping: map[string]string{
					"value1": "verb"},
			},
			wantErr: false,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				mapping: map[string]string{
					"value1": "verb"},
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			u := &UpdateOperations{
				Fields: testCase.fields.Fields,
			}
			if err := u.AddArrayOperation(testCase.args.customFieldID, testCase.args.mapping); (err != nil) != testCase.wantErr {

				if !reflect.DeepEqual(err, testCase.Err) {
					t.Errorf("AddArrayOperation() got = (%v), want (%v)", err, testCase.Err)
				}

				t.Errorf("AddArrayOperation() error = %v, wantErr %v", err, testCase.wantErr)
			}
		})
	}
}

func TestUpdateOperations_AddStringOperation(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		operation     string
		value         string
	}
	testCases := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "custom_field_id",
				operation:     "operation_sample",
				value:         "value_sample",
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				operation:     "operation_sample",
				value:         "value_sample",
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the operation is not provided",
			fields: fields{},
			args: args{
				customFieldID: "custom_field_id",
				operation:     "",
				value:         "value_sample",
			},
			wantErr: true,
			Err:     ErrNoEditOperatorError,
		},

		{
			name:   "when the operator value is not provided",
			fields: fields{},
			args: args{
				customFieldID: "custom_field_id",
				operation:     "operation_sample",
				value:         "",
			},
			wantErr: true,
			Err:     ErrNoEditValueError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			u := &UpdateOperations{
				Fields: testCase.fields.Fields,
			}
			if err := u.AddStringOperation(testCase.args.customFieldID, testCase.args.operation, testCase.args.value); (err != nil) != testCase.wantErr {

				if !reflect.DeepEqual(err, testCase.Err) {
					t.Errorf("AddStringOperation() got = (%v), want (%v)", err, testCase.Err)
				}

				t.Errorf("AddStringOperation() error = %v, wantErr %v", err, testCase.wantErr)
			}
		})
	}
}

func TestUpdateOperations_AddMultiRawOperation(t *testing.T) {

	expectedJson := `[{"update":{"custom_field_id":[{"add":{"id":"10001"}},{"remove":{"name":"Version 00"}},{"add":{"id":"1010"}}]}}]`

	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		mappings      []map[string]interface{}
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     assert.ErrorAssertionFunc
		expectedErr bool
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "custom_field_id",
				mappings: []map[string]interface{}{
					{
						"add": map[string]interface{}{
							"id": "10001",
						},
					},
					{
						"remove": map[string]interface{}{
							"name": "Version 00",
						},
					},
					{
						"add": map[string]interface{}{
							"id": "1010",
						},
					},
				},
			},
			wantErr: assert.NoError,
		},

		{
			name:   "when the customfieldID is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
			},
			wantErr:     assert.Error,
			expectedErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UpdateOperations{
				Fields: tt.fields.Fields,
			}

			tt.wantErr(t, u.AddMultiRawOperation(tt.args.customFieldID, tt.args.mappings),
				fmt.Sprintf("AddMultiRawOperation(%v, %v)", tt.args.customFieldID, tt.args.mappings))

			if !tt.expectedErr {
				actualJSON, err := json.Marshal(u.Fields)
				if err != nil {
					t.Fatal(err)
				}

				assert.Equal(t, expectedJson, string(actualJSON))
			}
		})
	}
}
