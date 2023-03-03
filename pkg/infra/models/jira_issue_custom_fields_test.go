package models

import (
	"bytes"
	"reflect"
	"testing"
)

func TestExtractMultiSelectField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
	"fields": {
		"customfield_10046": [
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10044",
        "value": "Option 1",
        "id": "10044",
		"optionId": "12222",
		"disabled": true
      },
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10046",
        "value": "Option 3",
        "id": "10046"
      },
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10047",
        "value": "Option 4",
        "id": "10047"
      }
    ]
	}
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
	"field_no_mapped": {
		"customfield_10046": [
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10044",
        "value": "Option 1",
        "id": "10044"
      },
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10046",
        "value": "Option 3",
        "id": "10046"
      }
    ]
	}
}`)

	bufferMockedWithNoJSON := bytes.Buffer{}
	bufferMockedWithNoJSON.WriteString(`{}{`)

	bufferMockedWithNoInfo := bytes.Buffer{}
	bufferMockedWithNoInfo.WriteString(`
{
	"fields": {
		"customfield_10046": null
	}
}`)

	bufferMockedWithInvalidType := bytes.Buffer{}
	bufferMockedWithInvalidType.WriteString(`
{
	"fields": {
		"customfield_10046": "Test field sample"
	}
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    []*CustomFieldContextOptionScheme
		want1   bool
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: []*CustomFieldContextOptionScheme{
				{
					ID:       "10044",
					Value:    "Option 1",
					OptionID: "12222",
					Disabled: true,
				},
				{
					ID:    "10046",
					Value: "Option 3",
				},
				{
					ID:    "10047",
					Value: "Option 4",
				},
			},
			want1:   true,
			wantErr: false,
		},

		{
			name: "when the buffer no contains information",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10046",
			},
			want:    nil,
			want1:   false,
			wantErr: false,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10046",
			},
			want:    nil,
			want1:   false,
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
			want1:   false,
			wantErr: true,
			Err:     ErrNoMultiSelectTypeError,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoJSON,
				customField: "customfield_10046",
			},
			want:    nil,
			want1:   false,
			wantErr: true,
			Err:     ErrNoCustomFieldUnmarshalError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, got1, err := ExtractMultiSelectField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ExtractMultiSelectField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ExtractMultiSelectField() got = %v, want %v", got, testCase.want)
			}
			if got1 != testCase.want1 {
				t.Errorf("ExtractMultiSelectField() got1 = %v, want %v", got1, testCase.want1)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ExtractMultiSelectField() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}
