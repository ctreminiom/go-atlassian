// Package models provides the data structures used in the admin package.
package models

import (
	"testing" // Standard Go testing package

	"github.com/stretchr/testify/assert" // Assert package for testing
)

// TestSCIMUserToPathScheme_AddStringOperation tests the AddStringOperation method of the SCIMUserToPathScheme struct.
func TestSCIMUserToPathScheme_AddStringOperation(t *testing.T) {
	// fields struct for holding the fields of SCIMUserToPathScheme
	type fields struct {
		Schemas    []string
		Operations []*SCIMUserToPathOperationScheme
	}
	// args struct for holding the arguments to the AddStringOperation method
	type args struct {
		operation string
		path      string
		value     string
	}
	// testCases struct for holding each test case
	testCases := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
	}{
		// Test case when the parameters are correct
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
		// Test case when the operation is not provided
		{
			name:   "when the operation is not provided",
			fields: fields{},
			args: args{
				operation: "",
				path:      "path_sample",
				value:     "value_sample",
			},
			wantErr: true,
			Err:     ErrNoSCIMOperation,
		},
		// Test case when the path is not provided
		{
			name:   "when the path is not provided",
			fields: fields{},
			args: args{
				operation: "operator_sample",
				path:      "",
				value:     "value_sample",
			},
			wantErr: true,
			Err:     ErrNoSCIMPath,
		},
		// Test case when the value is not provided
		{
			name:   "when the value is not provided",
			fields: fields{},
			args: args{
				operation: "operator_sample",
				path:      "path_sample",
				value:     "",
			},
			wantErr: true,
			Err:     ErrNoSCIMValue,
		},
	}
	// Loop over each test case
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Create a new SCIMUserToPathScheme
			s := &SCIMUserToPathScheme{
				Schemas:    testCase.fields.Schemas,
				Operations: testCase.fields.Operations,
			}

			// Call the AddStringOperation method
			err := s.AddStringOperation(testCase.args.operation, testCase.args.path, testCase.args.value)
			// If an error is expected
			if testCase.wantErr {
				// Log the error if it exists
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				// Assert that the error is as expected
				assert.EqualError(t, err, testCase.Err.Error())

			} else {
				// Assert that no error occurred
				assert.NoError(t, err)
			}
		})
	}
}

// TestSCIMUserToPathScheme_AddBoolOperation tests the AddBoolOperation method of the SCIMUserToPathScheme struct.
func TestSCIMUserToPathScheme_AddBoolOperation(t *testing.T) {
	// fields struct for holding the fields of SCIMUserToPathScheme
	type fields struct {
		Schemas    []string
		Operations []*SCIMUserToPathOperationScheme
	}
	// args struct for holding the arguments to the AddBoolOperation method
	type args struct {
		operation string
		path      string
		value     bool
	}
	// testCases struct for holding each test case
	testCases := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
	}{
		// Test case when the parameters are correct
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
		// Test case when the operation is not provided
		{
			name:   "when the operation is not provided",
			fields: fields{},
			args: args{
				operation: "",
				path:      "path_sample",
				value:     true,
			},
			wantErr: true,
			Err:     ErrNoSCIMOperation,
		},
		// Test case when the path is not provided
		{
			name:   "when the path is not provided",
			fields: fields{},
			args: args{
				operation: "operator_sample",
				path:      "",
				value:     true,
			},
			wantErr: true,
			Err:     ErrNoSCIMPath,
		},
	}
	// Loop over each test case
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Create a new SCIMUserToPathScheme
			s := &SCIMUserToPathScheme{
				Schemas:    testCase.fields.Schemas,
				Operations: testCase.fields.Operations,
			}

			// Call the AddBoolOperation method
			err := s.AddBoolOperation(testCase.args.operation, testCase.args.path, testCase.args.value)
			// If an error is expected
			if testCase.wantErr {
				// Log the error if it exists
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				// Assert that the error is as expected
				assert.EqualError(t, err, testCase.Err.Error())

			} else {
				// Assert that no error occurred
				assert.NoError(t, err)
			}
		})
	}
}

// TestSCIMUserToPathScheme_AddComplexOperation tests the AddComplexOperation method of the SCIMUserToPathScheme struct.
func TestSCIMUserToPathScheme_AddComplexOperation(t *testing.T) {
	// fields struct for holding the fields of SCIMUserToPathScheme
	type fields struct {
		Schemas    []string
		Operations []*SCIMUserToPathOperationScheme
	}
	// args struct for holding the arguments to the AddComplexOperation method
	type args struct {
		operation string
		path      string
		values    []*SCIMUserComplexOperationScheme
	}
	// testCases struct for holding each test case
	testCases := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
	}{
		// Test case when the parameters are correct
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
		// Test case when the operation is not provided
		{
			name:   "when the operation is not provided",
			fields: fields{},
			args: args{
				operation: "",
				path:      "path_sample",
			},
			wantErr: true,
			Err:     ErrNoSCIMOperation,
		},
		// Test case when the path is not provided
		{
			name:   "when the path is not provided",
			fields: fields{},
			args: args{
				operation: "operator_sample",
				path:      "",
			},
			wantErr: true,
			Err:     ErrNoSCIMPath,
		},
		// Test case when the value is not provided
		{
			name:   "when the value is not provided",
			fields: fields{},
			args: args{
				operation: "operator_sample",
				path:      "path_sample",
				values:    nil,
			},
			wantErr: true,
			Err:     ErrNoSCIMComplexValue,
		},
	}
	// Loop over each test case
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Create a new SCIMUserToPathScheme
			s := &SCIMUserToPathScheme{
				Schemas:    testCase.fields.Schemas,
				Operations: testCase.fields.Operations,
			}

			// Call the AddComplexOperation method
			err := s.AddComplexOperation(testCase.args.operation, testCase.args.path, testCase.args.values)
			// If an error is expected
			if testCase.wantErr {
				// Log the error if it exists
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				// Assert that the error is as expected
				assert.EqualError(t, err, testCase.Err.Error())

			} else {
				// Assert that no error occurred
				assert.NoError(t, err)
			}
		})
	}
}
