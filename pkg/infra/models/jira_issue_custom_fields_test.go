package models

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParseMultiSelectField(t *testing.T) {

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
			got, err := ParseMultiSelectCustomField(testCase.args.buffer, testCase.args.customField)
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

func TestParseMultiGroupPickerField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
	"fields": {
		"customfield_10052": [
      {
        "name": "jira-administrators",
        "groupId": "1da6f895-2b42-423b-8bfb-1e09ee8d7764",
        "self": "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=1da6f895-2b42-423b-8bfb-1e09ee8d7764"
      },
      {
        "name": "jira-administrators-system",
        "groupId": "be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
        "self": "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=be9ba0ab-ecdc-445b-9ce6-b95202026c1a"
      }
    ]
	}
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
	"field_no_mapped": {
		"customfield_10052": [
      {
        "name": "jira-administrators",
        "groupId": "1da6f895-2b42-423b-8bfb-1e09ee8d7764",
        "self": "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=1da6f895-2b42-423b-8bfb-1e09ee8d7764"
      },
      {
        "name": "jira-administrators-system",
        "groupId": "be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
        "self": "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=be9ba0ab-ecdc-445b-9ce6-b95202026c1a"
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
		"customfield_10052": null
	}
}`)

	bufferMockedWithInvalidType := bytes.Buffer{}
	bufferMockedWithInvalidType.WriteString(`
{
	"fields": {
		"customfield_10052": "Test field sample"
	}
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    []*GroupDetailScheme
		want1   bool
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10052",
			},
			want: []*GroupDetailScheme{
				{
					Self:    "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=1da6f895-2b42-423b-8bfb-1e09ee8d7764",
					Name:    "jira-administrators",
					GroupID: "1da6f895-2b42-423b-8bfb-1e09ee8d7764",
				},
				{
					Self:    "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
					Name:    "jira-administrators-system",
					GroupID: "be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
				},
			},
			want1:   true,
			wantErr: false,
		},

		{
			name: "when the buffer no contains information",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10052",
			},
			want:    nil,
			want1:   false,
			wantErr: false,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10052",
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
				customField: "customfield_10052",
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
				customField: "customfield_10052",
			},
			want:    nil,
			want1:   false,
			wantErr: true,
			Err:     ErrNoCustomFieldUnmarshalError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseMultiGroupPickerCustomField(testCase.args.buffer, testCase.args.customField)
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

func TestParseMultiUserPickerField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
	"fields": {
		"customfield_10055": [
     {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
        "accountId": "5e5f6a63157ed50cd2b9eaca",
        "avatarUrls": {
          "48x48": "https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
          "24x24": "https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
          "16x16": "https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
          "32x32": "https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png"
        },
        "displayName": "Carlos Treminio",
        "active": true,
        "timeZone": "Asia/Dhaka",
        "accountType": "atlassian"
      },
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
        "accountId": "5b86be50b8e3cb5895860d6d",
        "emailAddress": "ctreminiom079@gmail.com",
        "avatarUrls": {
          "48x48": "https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
          "24x24": "https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
          "16x16": "https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
          "32x32": "https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png"
        },
        "displayName": "Carlos Treminio",
        "active": true,
        "timeZone": "America/Guatemala",
        "accountType": "atlassian"
      },
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=557058%3Ad6b5955a-e193-41e1-b051-79cdb0755d68",
        "accountId": "557058:d6b5955a-e193-41e1-b051-79cdb0755d68",
        "avatarUrls": {
          "48x48": "https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
          "24x24": "https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
          "16x16": "https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
          "32x32": "https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png"
        },
        "displayName": "Trello",
        "active": true,
        "timeZone": "Europe/London",
        "accountType": "app"
      }
    ]
	}
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
	"field_no_mapped": {
		"customfield_10055": [
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
        "accountId": "5e5f6a63157ed50cd2b9eaca",
        "avatarUrls": {
          "48x48": "https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
          "24x24": "https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
          "16x16": "https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
          "32x32": "https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png"
        },
        "displayName": "Carlos Treminio",
        "active": true,
        "timeZone": "Asia/Dhaka",
        "accountType": "atlassian"
      },
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
        "accountId": "5b86be50b8e3cb5895860d6d",
        "emailAddress": "ctreminiom079@gmail.com",
        "avatarUrls": {
          "48x48": "https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
          "24x24": "https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
          "16x16": "https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
          "32x32": "https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png"
        },
        "displayName": "Carlos Treminio",
        "active": true,
        "timeZone": "America/Guatemala",
        "accountType": "atlassian"
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
		"customfield_10055": null
	}
}`)

	bufferMockedWithInvalidType := bytes.Buffer{}
	bufferMockedWithInvalidType.WriteString(`
{
	"fields": {
		"customfield_10055": "Test field sample"
	}
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    []*UserDetailScheme
		want1   bool
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10055",
			},
			want: []*UserDetailScheme{
				{
					Self:        "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
					AccountID:   "5e5f6a63157ed50cd2b9eaca",
					DisplayName: "Carlos Treminio",
					Active:      true,
					TimeZone:    "Asia/Dhaka",
					AccountType: "atlassian",
				},
				{
					Self:         "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
					AccountID:    "5b86be50b8e3cb5895860d6d",
					EmailAddress: "ctreminiom079@gmail.com",
					DisplayName:  "Carlos Treminio",
					Active:       true,
					TimeZone:     "America/Guatemala",
					AccountType:  "atlassian",
				},
				{
					Self:        "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=557058%3Ad6b5955a-e193-41e1-b051-79cdb0755d68",
					AccountID:   "557058:d6b5955a-e193-41e1-b051-79cdb0755d68",
					DisplayName: "Trello",
					Active:      true,
					TimeZone:    "Europe/London",
					AccountType: "app",
				},
			},
			want1:   true,
			wantErr: false,
		},

		{
			name: "when the buffer no contains information",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10055",
			},
			want:    nil,
			want1:   false,
			wantErr: false,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10055",
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
				customField: "customfield_10055",
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
				customField: "customfield_10055",
			},
			want:    nil,
			want1:   false,
			wantErr: true,
			Err:     ErrNoCustomFieldUnmarshalError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseMultiUserPickerCustomField(testCase.args.buffer, testCase.args.customField)
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

func TestParseCascadingSelectField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
  "fields": {
    "customfield_10045": {
      "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10054",
      "value": "America",
      "id": "10054",
      "child": {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10057",
        "value": "Costa Rica",
        "id": "10057"
      }
    }
  }
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
  "fields_no_mapped": {
    "customfield_10045": {
      "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10054",
      "value": "America",
      "id": "10054",
      "child": {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10057",
        "value": "Costa Rica",
        "id": "10057"
      }
    }
  }
}`)

	bufferMockedWithNoJSON := bytes.Buffer{}
	bufferMockedWithNoJSON.WriteString(`{}{`)

	bufferMockedWithNoInfo := bytes.Buffer{}
	bufferMockedWithNoInfo.WriteString(`
{
	"fields": {
		"customfield_10045": null
	}
}`)

	bufferMockedWithInvalidType := bytes.Buffer{}
	bufferMockedWithInvalidType.WriteString(`
{
	"fields": {
		"customfield_10045": "Test field sample"
	}
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    *CascadingSelectScheme
		want1   bool
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10045",
			},
			want: &CascadingSelectScheme{
				Self:  "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10054",
				Value: "America",
				Id:    "10054",
				Child: &CascadingSelectChildScheme{
					Self:  "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10057",
					Value: "Costa Rica",
					Id:    "10057",
				},
			},
			want1:   false,
			wantErr: false,
		},

		{
			name: "when the buffer no contains information",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10045",
			},
			want:    nil,
			want1:   false,
			wantErr: false,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10045",
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
				customField: "customfield_10045",
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
				customField: "customfield_10045",
			},
			want:    nil,
			want1:   false,
			wantErr: true,
			Err:     ErrNoCustomFieldUnmarshalError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseCascadingSelectCustomField(testCase.args.buffer, testCase.args.customField)
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

func TestParseMultiCheckboxesCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
	"fields": {
		"customfield_10046": [
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10037",
        "value": "Option 2",
        "id": "10037"
      },
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10039",
        "value": "Options 4",
        "id": "10039"
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
        "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10037",
        "value": "Option 2",
        "id": "10037"
      },
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10039",
        "value": "Options 4",
        "id": "10039"
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
					ID:    "10037",
					Value: "Option 2",
				},
				{
					ID:    "10039",
					Value: "Options 4",
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
			got, err := ParseMultiCheckboxesCustomField(testCase.args.buffer, testCase.args.customField)
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

func TestParseMultiVersionCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
	"fields": {
		"customfield_10046": [
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/version/10000",
        "id": "10000",
        "description": "",
        "name": "Version 00",
        "archived": false,
        "released": false,
        "releaseDate": "2021-02-23"
      },
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/version/10002",
        "id": "10002",
        "description": "Version Sandbox description - UPDATED",
        "name": "Version Sandbox - UPDATED",
        "archived": false,
        "released": true,
        "releaseDate": "2021-03-06"
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
        "self": "https://ctreminiom.atlassian.net/rest/api/3/version/10000",
        "id": "10000",
        "description": "",
        "name": "Version 00",
        "archived": false,
        "released": false,
        "releaseDate": "2021-02-23"
      },
      {
        "self": "https://ctreminiom.atlassian.net/rest/api/3/version/10002",
        "id": "10002",
        "description": "Version Sandbox description - UPDATED",
        "name": "Version Sandbox - UPDATED",
        "archived": false,
        "released": true,
        "releaseDate": "2021-03-06"
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
		want    []*VersionDetailScheme
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
			want: []*VersionDetailScheme{
				{
					Self:        "https://ctreminiom.atlassian.net/rest/api/3/version/10000",
					ID:          "10000",
					Description: "",
					Name:        "Version 00",
					Archived:    false,
					Released:    false,
					ReleaseDate: "2021-02-23",
				},
				{
					Self:        "https://ctreminiom.atlassian.net/rest/api/3/version/10002",
					ID:          "10002",
					Description: "Version Sandbox description - UPDATED",
					Name:        "Version Sandbox - UPDATED",
					Archived:    false,
					Released:    true,
					ReleaseDate: "2021-03-06",
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
			got, err := ParseMultiVersionCustomField(testCase.args.buffer, testCase.args.customField)
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

func TestParseUserPickerCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
  "fields": {
    "customfield_10045": {
      "self": "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6acefc1fca0af44135f8",
      "accountId": "5e5f6acefc1fca0af44135f8",
      "avatarUrls": {
        "48x48": "https://secure.gravatar.com/avatar/6c20a29c5ab36b3cbc121782edaadfc9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FEN-4.png",
        "24x24": "https://secure.gravatar.com/avatar/6c20a29c5ab36b3cbc121782edaadfc9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FEN-4.png",
        "16x16": "https://secure.gravatar.com/avatar/6c20a29c5ab36b3cbc121782edaadfc9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FEN-4.png",
        "32x32": "https://secure.gravatar.com/avatar/6c20a29c5ab36b3cbc121782edaadfc9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FEN-4.png"
      },
      "displayName": "Eduardo Navarro",
      "active": true,
      "timeZone": "Europe/London",
      "accountType": "atlassian"
    }
  }
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
  "fields_no_mapped": {
    "customfield_10045": {
      "self": "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6acefc1fca0af44135f8",
      "accountId": "5e5f6acefc1fca0af44135f8",
      "avatarUrls": {
        "48x48": "https://secure.gravatar.com/avatar/6c20a29c5ab36b3cbc121782edaadfc9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FEN-4.png",
        "24x24": "https://secure.gravatar.com/avatar/6c20a29c5ab36b3cbc121782edaadfc9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FEN-4.png",
        "16x16": "https://secure.gravatar.com/avatar/6c20a29c5ab36b3cbc121782edaadfc9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FEN-4.png",
        "32x32": "https://secure.gravatar.com/avatar/6c20a29c5ab36b3cbc121782edaadfc9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FEN-4.png"
      },
      "displayName": "Eduardo Navarro",
      "active": true,
      "timeZone": "Europe/London",
      "accountType": "atlassian"
    }
  }
}`)

	bufferMockedWithNoJSON := bytes.Buffer{}
	bufferMockedWithNoJSON.WriteString(`{}{`)

	bufferMockedWithNoInfo := bytes.Buffer{}
	bufferMockedWithNoInfo.WriteString(`
{
	"fields": {
		"customfield_10045": null
	}
}`)

	bufferMockedWithInvalidType := bytes.Buffer{}
	bufferMockedWithInvalidType.WriteString(`
{
	"fields": {
		"customfield_10045": "Test field sample"
	}
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    *UserDetailScheme
		want1   bool
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10045",
			},
			want: &UserDetailScheme{
				Self:         "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6acefc1fca0af44135f8",
				AccountID:    "5e5f6acefc1fca0af44135f8",
				EmailAddress: "",
				DisplayName:  "Eduardo Navarro",
				Active:       true,
				TimeZone:     "Europe/London",
				AccountType:  "atlassian",
			},
			want1:   false,
			wantErr: false,
		},

		{
			name: "when the buffer no contains information",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10045",
			},
			want:    nil,
			want1:   false,
			wantErr: false,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10045",
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
				customField: "customfield_10045",
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
				customField: "customfield_10045",
			},
			want:    nil,
			want1:   false,
			wantErr: true,
			Err:     ErrNoCustomFieldUnmarshalError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseUserPickerCustomField(testCase.args.buffer, testCase.args.customField)
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

func TestParseFloatCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
  "fields": {
    "customfield_10045": 1000.323
  }
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
  "fields_no_mapped": {
    "customfield_10045": 1000.323
  }
}`)

	bufferMockedWithNoJSON := bytes.Buffer{}
	bufferMockedWithNoJSON.WriteString(`{}{`)

	bufferMockedWithNoInfo := bytes.Buffer{}
	bufferMockedWithNoInfo.WriteString(`
{
	"fields": {
		"customfield_10045": null
	}
}`)

	bufferMockedWithInvalidType := bytes.Buffer{}
	bufferMockedWithInvalidType.WriteString(`
{
	"fields": {
		"customfield_10045": "Test field sample"
	}
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10045",
			},
			want:    1000.323,
			wantErr: false,
		},

		{
			name: "when the buffer no contains information",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10045",
			},
			want:    0,
			wantErr: false,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10045",
			},
			want:    0,
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},

		{
			name: "when the buffer does not contains a valid field type",
			args: args{
				buffer:      bufferMockedWithInvalidType,
				customField: "customfield_10045",
			},
			want:    0,
			wantErr: true,
			Err:     ErrNoMultiSelectTypeError,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoJSON,
				customField: "customfield_10045",
			},
			want:    0,
			wantErr: true,
			Err:     ErrNoCustomFieldUnmarshalError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseFloatCustomField(testCase.args.buffer, testCase.args.customField)
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

func TestParseLabelCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
  "fields": {
    "customfield_10045": [
      "asd",
      "asds"
    ]
  }
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
  "fields_no_mapped": {
   "customfield_10045": [
      "asd",
      "asds"
    ]
  }
}`)

	bufferMockedWithNoJSON := bytes.Buffer{}
	bufferMockedWithNoJSON.WriteString(`{}{`)

	bufferMockedWithNoInfo := bytes.Buffer{}
	bufferMockedWithNoInfo.WriteString(`
{
	"fields": {
		"customfield_10045": null
	}
}`)

	bufferMockedWithInvalidType := bytes.Buffer{}
	bufferMockedWithInvalidType.WriteString(`
{
	"fields": {
		"customfield_10045": "Test field sample"
	}
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10045",
			},
			want:    []string{"asd", "asds"},
			wantErr: false,
		},

		{
			name: "when the buffer no contains information",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10045",
			},
			want:    nil,
			wantErr: false,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10045",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},

		{
			name: "when the buffer does not contains a valid field type",
			args: args{
				buffer:      bufferMockedWithInvalidType,
				customField: "customfield_10045",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoMultiSelectTypeError,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoJSON,
				customField: "customfield_10045",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoCustomFieldUnmarshalError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseLabelCustomField(testCase.args.buffer, testCase.args.customField)
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

func TestParseSprintCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
	"fields": {
		"customfield_10046": [
      {
        "id": 4,
        "name": "KP Sprint 3",
        "state": "active",
        "boardId": 4,
        "goal": "",
        "startDate": "2023-03-04T02:03:16.273Z",
        "endDate": "2023-03-17T02:03:00.000Z"
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
        "id": 4,
        "name": "KP Sprint 3",
        "state": "active",
        "boardId": 4,
        "goal": "",
        "startDate": "2023-03-04T02:03:16.273Z",
        "endDate": "2023-03-17T02:03:00.000Z"
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
		want    []*SprintDetailScheme
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
			want: []*SprintDetailScheme{
				{
					ID:            4,
					State:         "active",
					Name:          "KP Sprint 3",
					StartDate:     "2023-03-04T02:03:16.273Z",
					EndDate:       "2023-03-17T02:03:00.000Z",
					OriginBoardID: 4,
					Goal:          "",
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
			got, err := ParseSprintCustomField(testCase.args.buffer, testCase.args.customField)
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

func TestParseSelectCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
  "fields": {
    "customfield_10045": {
      "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10058",
      "value": "Scranton 1",
      "id": "10058"
    }
  }
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
  "fields_no_mapped": {
   "customfield_10045": {
      "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10058",
      "value": "Scranton 1",
      "id": "10058"
    }
  }
}`)

	bufferMockedWithNoJSON := bytes.Buffer{}
	bufferMockedWithNoJSON.WriteString(`{}{`)

	bufferMockedWithNoInfo := bytes.Buffer{}
	bufferMockedWithNoInfo.WriteString(`
{
	"fields": {
		"customfield_10045": null
	}
}`)

	bufferMockedWithInvalidType := bytes.Buffer{}
	bufferMockedWithInvalidType.WriteString(`
{
	"fields": {
		"customfield_10045": "Test field sample"
	}
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    *CustomFieldContextOptionScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10045",
			},
			want: &CustomFieldContextOptionScheme{
				ID:    "10058",
				Value: "Scranton 1",
			},
			wantErr: false,
		},

		{
			name: "when the buffer no contains information",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10045",
			},
			want:    nil,
			wantErr: false,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10045",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},

		{
			name: "when the buffer does not contains a valid field type",
			args: args{
				buffer:      bufferMockedWithInvalidType,
				customField: "customfield_10045",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoMultiSelectTypeError,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoJSON,
				customField: "customfield_10045",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoCustomFieldUnmarshalError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseSelectCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseSelectCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseSelectCustomField() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseSelectCustomField() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}
