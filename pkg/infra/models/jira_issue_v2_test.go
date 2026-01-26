package models

import (
	"reflect"
	"testing"
)

func TestIssueSchemeV2_MergeCustomFields(t *testing.T) {

	var customFields = &CustomFields{}
	err := customFields.Number("customfield_10043", 1000.3232)
	if err != nil {
		t.Error(err)
	}

	type fields struct {
		ID          string
		Key         string
		Self        string
		Transitions []*IssueTransitionScheme
		Changelog   *IssueChangelogScheme
		Fields      *IssueFieldsSchemeV2
	}

	type args struct {
		fields *CustomFields
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
		Err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				fields: customFields,
			},
			want: map[string]interface{}{
				"fields": map[string]interface{}{
					"customfield_10043": 1000.3232,
				},
			},
			wantErr: false,
			Err:     nil,
		},
		{
			name: "when custom fields is nil, should preserve issue scheme fields",
			fields: fields{
				ID:  "10001",
				Key: "TEST-1",
				Fields: &IssueFieldsSchemeV2{
					Summary: "Test Summary",
				},
			},
			args: args{
				fields: nil,
			},
			want: map[string]interface{}{
				"id":  "10001",
				"key": "TEST-1",
				"fields": map[string]interface{}{
					"summary": "Test Summary",
				},
			},
			wantErr: false,
			Err:     nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			i := &IssueSchemeV2{
				ID:          testCase.fields.ID,
				Key:         testCase.fields.Key,
				Self:        testCase.fields.Self,
				Transitions: testCase.fields.Transitions,
				Changelog:   testCase.fields.Changelog,
				Fields:      testCase.fields.Fields,
			}
			got, err := i.MergeCustomFields(testCase.args.fields)
			if (err != nil) != testCase.wantErr {
				t.Errorf("MergeCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("MergeCustomFields() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("AddArrayOperation() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestIssueSchemeV2_MergeOperations(t *testing.T) {

	var operations = &UpdateOperations{}
	err := operations.AddArrayOperation("labels", map[string]string{"triaged": "remove"})
	if err != nil {
		t.Error(err)
	}

	type fields struct {
		ID          string
		Key         string
		Self        string
		Transitions []*IssueTransitionScheme
		Changelog   *IssueChangelogScheme
		Fields      *IssueFieldsSchemeV2
	}
	type args struct {
		operations *UpdateOperations
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
		Err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				operations: operations,
			},
			want: map[string]interface{}{
				"update": map[string]interface{}{
					"labels": []map[string]interface{}{
						{
							"remove": "triaged",
						},
					},
				},
			},
			wantErr: false,
			Err:     nil,
		},
		{
			name: "when operations is nil, should preserve issue scheme fields",
			fields: fields{
				ID:  "10001",
				Key: "TEST-1",
				Fields: &IssueFieldsSchemeV2{
					Summary: "Test Summary",
				},
			},
			args: args{
				operations: nil,
			},
			want: map[string]interface{}{
				"id":  "10001",
				"key": "TEST-1",
				"fields": map[string]interface{}{
					"summary": "Test Summary",
				},
			},
			wantErr: false,
			Err:     nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			i := &IssueSchemeV2{
				ID:          testCase.fields.ID,
				Key:         testCase.fields.Key,
				Self:        testCase.fields.Self,
				Transitions: testCase.fields.Transitions,
				Changelog:   testCase.fields.Changelog,
				Fields:      testCase.fields.Fields,
			}
			got, err := i.MergeOperations(testCase.args.operations)
			if (err != nil) != testCase.wantErr {
				t.Errorf("MergeOperations() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("MergeOperations() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("MergeOperations() got = (%v), want (%v)", err, testCase.Err)
			}

		})
	}
}

func TestIssueSchemeV2_ToMap(t *testing.T) {
	type fields struct {
		ID          string
		Key         string
		Self        string
		Transitions []*IssueTransitionScheme
		Changelog   *IssueChangelogScheme
		Fields      *IssueFieldsSchemeV2
	}
	testCases := []struct {
		name    string
		fields  fields
		want    map[string]interface{}
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			fields: fields{
				Key: "DUMMY-1",
				Fields: &IssueFieldsSchemeV2{
					Summary: "Test",
				},
			},
			want: map[string]interface{}{
				"key": "DUMMY-1",
				"fields": map[string]interface{}{
					"summary": "Test",
				},
			},
			wantErr: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			i := &IssueSchemeV2{
				ID:          testCase.fields.ID,
				Key:         testCase.fields.Key,
				Self:        testCase.fields.Self,
				Transitions: testCase.fields.Transitions,
				Changelog:   testCase.fields.Changelog,
				Fields:      testCase.fields.Fields,
			}
			got, err := i.ToMap()

			if (err != nil) != testCase.wantErr {
				t.Errorf("ToMap() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ToMap() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ToMap() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}
