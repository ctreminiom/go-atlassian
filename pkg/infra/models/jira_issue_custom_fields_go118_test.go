package models

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParseCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
	"fields": {
		"customfield_10046": {
			"MyFieldName": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10044"
		}
	}
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
	"field_no_mapped": {
		"customfield_10046": {
			"MyFieldName": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10044"
		}
	}
}`)

	bufferMockedWithNoJSON := bytes.Buffer{}
	bufferMockedWithNoJSON.WriteString(`{}{`)

	bufferMockedWithNoInfo := bytes.Buffer{}
	bufferMockedWithNoInfo.WriteString(`
{
	"fields": {
		"w": null
	}
}`)

	bufferMockedWithInvalidType := bytes.Buffer{}
	bufferMockedWithInvalidType.WriteString(`
{
	"fields": {
		"customfield_10046": "Test field sample"
	}
}`)

	type CustomType struct {
		MyFieldName string
	}

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    *T
		wantErr bool
		Err     error
	}
	tests := []testCase[CustomType]{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: &CustomType{
				MyFieldName: "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10044",
			},
			wantErr: false,
		},

		{
			name: "when the buffer no contains information",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10046",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoCustomTypeError,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10046",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},

		{
			name: "when the buffer does not contains a valid field type",
			args: args{
				buffer:      bufferMockedWithInvalidType,
				customField: "customfield_10046",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoCustomTypeError,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoJSON,
				customField: "customfield_10046",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCustomField[CustomType](tt.args.buffer, tt.args.customField)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCustomField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCustomField() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(err, tt.Err) {
				t.Errorf("ParseCustomField() got = (%v), want (%v)", err, tt.Err)
			}
		})
	}
}

func TestParseCustomFields(t *testing.T) {
	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
    "expand": "names,schema",
    "startAt": 0,
    "maxResults": 50,
    "total": 1,
    "issues": [
        {
            "expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
            "id": "10035",
            "self": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
            "key": "KP-22",
            "fields": {
                "customfield_10046": {
                    "MyFieldName": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10044"
                }
            }
        }
    ]
}`)

	bufferMockedWithNoIssues := bytes.Buffer{}
	bufferMockedWithNoIssues.WriteString(`
{
    "expand": "names,schema",
    "startAt": 0,
    "maxResults": 50,
    "total": 1,
    "no_issues": [
        {
            "expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
            "id": "10035",
            "self": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
            "key": "KP-22",
            "no_fields": {
                "customfield_10046": {
                    "MyFieldName": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10044"
                }
            }
        }
    ]
}`)

	bufferMockedWithNoInfo := bytes.Buffer{}
	bufferMockedWithNoInfo.WriteString(`
{
    "expand": "names,schema",
    "startAt": 0,
    "maxResults": 50,
    "total": 1,
    "data": [
        {
            "expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
            "id": "10035",
            "self": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
            "key": "KP-22",
            "fields": {
                "customfield_10046": null
            }
        }
    ]
}`)

	type CustomType struct {
		MyFieldName string
	}

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    map[string]*T
		wantErr bool
		Err     error
	}
	tests := []testCase[CustomType]{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string]*CustomType{
				"KP-22": {
					MyFieldName: "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10044",
				},
			},
			wantErr: false,
		},

		{
			name: "when the buffer does not contain the issues object",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10046",
			},
			wantErr: true,
			Err:     ErrNoIssuesSliceError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCustomFields[CustomType](tt.args.buffer, tt.args.customField)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCustomFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCustomFields() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(err, tt.Err) {
				t.Errorf("ParseCustomFields() got = (%v), want (%v)", err, tt.Err)
			}
		})
	}
}
