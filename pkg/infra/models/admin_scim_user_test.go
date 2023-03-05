package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSCIMUserToPathScheme_AddStringOperation(t *testing.T) {
	type fields struct {
		Schemas    []string
		Operations []*SCIMUserToPathOperationScheme
	}
	type args struct {
		operation string
		path      string
		value     string
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
				operation: "operator_sample",
				path:      "path_sample",
				value:     "value_sample",
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the operation is not provided",
			fields: fields{},
			args: args{
				operation: "",
				path:      "path_sample",
				value:     "value_sample",
			},
			wantErr: true,
			Err:     ErrNoSCIMOperationError,
		},

		{
			name:   "when the path is not provided",
			fields: fields{},
			args: args{
				operation: "operator_sample",
				path:      "",
				value:     "value_sample",
			},
			wantErr: true,
			Err:     ErrNoSCIMPathError,
		},

		{
			name:   "when the value is not provided",
			fields: fields{},
			args: args{
				operation: "operator_sample",
				path:      "path_sample",
				value:     "",
			},
			wantErr: true,
			Err:     ErrNoSCIMValueError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &SCIMUserToPathScheme{
				Schemas:    testCase.fields.Schemas,
				Operations: testCase.fields.Operations,
			}

			err := s.AddStringOperation(testCase.args.operation, testCase.args.path, testCase.args.value)
			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.Err.Error())

			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSCIMUserToPathScheme_AddBoolOperation(t *testing.T) {
	type fields struct {
		Schemas    []string
		Operations []*SCIMUserToPathOperationScheme
	}
	type args struct {
		operation string
		path      string
		value     bool
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
				operation: "operator_sample",
				path:      "path_sample",
				value:     true,
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the operation is not provided",
			fields: fields{},
			args: args{
				operation: "",
				path:      "path_sample",
				value:     true,
			},
			wantErr: true,
			Err:     ErrNoSCIMOperationError,
		},

		{
			name:   "when the path is not provided",
			fields: fields{},
			args: args{
				operation: "operator_sample",
				path:      "",
				value:     true,
			},
			wantErr: true,
			Err:     ErrNoSCIMPathError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &SCIMUserToPathScheme{
				Schemas:    testCase.fields.Schemas,
				Operations: testCase.fields.Operations,
			}

			err := s.AddBoolOperation(testCase.args.operation, testCase.args.path, testCase.args.value)
			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.Err.Error())

			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSCIMUserToPathScheme_AddComplexOperation(t *testing.T) {
	type fields struct {
		Schemas    []string
		Operations []*SCIMUserToPathOperationScheme
	}
	type args struct {
		operation string
		path      string
		values    []*SCIMUserComplexOperationScheme
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
				operation: "operator_sample",
				path:      "path_sample",
				values: []*SCIMUserComplexOperationScheme{
					{
						Value:     "value_sample",
						ValueType: "value_type_sample",
						Primary:   false,
					},
				},
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the operation is not provided",
			fields: fields{},
			args: args{
				operation: "",
				path:      "path_sample",
			},
			wantErr: true,
			Err:     ErrNoSCIMOperationError,
		},

		{
			name:   "when the path is not provided",
			fields: fields{},
			args: args{
				operation: "operator_sample",
				path:      "",
			},
			wantErr: true,
			Err:     ErrNoSCIMPathError,
		},

		{
			name:   "when the value is not provided",
			fields: fields{},
			args: args{
				operation: "operator_sample",
				path:      "path_sample",
				values:    nil,
			},
			wantErr: true,
			Err:     ErrNoSCIMComplexValueError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &SCIMUserToPathScheme{
				Schemas:    testCase.fields.Schemas,
				Operations: testCase.fields.Operations,
			}

			err := s.AddComplexOperation(testCase.args.operation, testCase.args.path, testCase.args.values)
			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.Err.Error())

			} else {
				assert.NoError(t, err)
			}
		})
	}
}
