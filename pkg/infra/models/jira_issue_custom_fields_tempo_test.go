package models

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParseTempoAccountCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
  "fields":{
    "customfield_10036":{
      "id":22,
      "value":"SP Datacenter"
    }
  }
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
  "no_fields":{
    "customfield_10036":{
      "id":22,
      "value":"SP Datacenter"
    }
  }
}`)

	bufferMockedWithNoJSON := bytes.Buffer{}
	bufferMockedWithNoJSON.WriteString(`{}{`)

	bufferMockedWithNoInfo := bytes.Buffer{}
	bufferMockedWithNoInfo.WriteString(`
{
  "fields":{
    "customfield_10036":null
  }
}`)

	bufferMockedWithInvalidType := bytes.Buffer{}
	bufferMockedWithInvalidType.WriteString(`
{
  "fields":{
    "customfield_10036":""
  }
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    *CustomFieldTempoAccountScheme
		want1   bool
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10036",
			},
			want: &CustomFieldTempoAccountScheme{
				ID:    22,
				Value: "SP Datacenter",
			},
			wantErr: false,
		},

		{
			name: "when the buffer no contains information",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10036",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoTempoAccountType,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10046",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoFieldInformation,
		},

		{
			name: "when the buffer does not contains a valid field type",
			args: args{
				buffer:      bufferMockedWithInvalidType,
				customField: "customfield_10036",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoTempoAccountType,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoJSON,
				customField: "customfield_10046",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoFieldInformation,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseTempoAccountCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseTempoAccountCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseTempoAccountCustomField() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseTempoAccountCustomField() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseTempoAccountCustomFields(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
  "expand":"names,schema",
  "startAt":0,
  "maxResults":50,
  "total":1,
  "issues":[
    {
      "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
      "id":"10035",
      "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
      "key":"KP-22",
      "fields":{
        "customfield_10036":{
          "id":22,
          "value":"SP Datacenter"
        }
      }
    }
  ]
}`)

	bufferMockedWithNoIssues := bytes.Buffer{}
	bufferMockedWithNoIssues.WriteString(`
{
  "expand":"names,schema",
  "startAt":0,
  "maxResults":50,
  "total":1,
  "no_issues":[
    {
      "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
      "id":"10035",
      "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
      "key":"KP-22",
      "no_fields":{
        "customfield_10036":{
          "id":22,
          "value":"SP Datacenter"
        }
      }
    }
  ]
}`)

	bufferMockedWithNoInfo := bytes.Buffer{}
	bufferMockedWithNoInfo.WriteString(`
{
  "expand":"names,schema",
  "startAt":0,
  "maxResults":50,
  "total":1,
  "issues":[
    {
      "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
      "id":"10035",
      "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
      "key":"KP-22",
      "fields":{
        "customfield_10036":null
      }
    }
  ]
}`)

	bufferMockedWithInvalidType := bytes.Buffer{}
	bufferMockedWithInvalidType.WriteString(`
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
                "customfield_10046": "string"
            }
        }
    ]
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*CustomFieldTempoAccountScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10036",
			},
			want: map[string]*CustomFieldTempoAccountScheme{
				"KP-22": {
					ID:    22,
					Value: "SP Datacenter",
				},
			},
			wantErr: false,
		},

		{
			name: "when the buffer does not contain the issues object",
			args: args{
				buffer:      bufferMockedWithNoIssues,
				customField: "customfield_10036",
			},
			wantErr: true,
			Err:     ErrNoIssuesSlice,
		},

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10036",
			},
			wantErr: true,
			Err:     ErrNoMapValues,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseTempoAccountCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseMultiSelectCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseMultiSelectCustomField() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseMultiSelectCustomField() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}
