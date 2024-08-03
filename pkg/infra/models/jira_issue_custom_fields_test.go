package models

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestParseMultiSelectCustomField(t *testing.T) {

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
			Err:     ErrNoMultiSelectTypeError,
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
			Err:     ErrNoMultiSelectTypeError,
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
			wantErr: false,
		},

		{
			name: "when the buffer no contains information",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10052",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoMultiSelectTypeError,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10052",
			},
			want:    nil,
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
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseMultiGroupPickerCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseMultiGroupPickerCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseMultiGroupPickerCustomField() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseMultiGroupPickerCustomField() got = (%v), want (%v)", err, testCase.Err)
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
			wantErr: true,
			Err:     ErrNoMultiSelectTypeError,
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
			Err:     ErrNoFieldInformationError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseMultiUserPickerCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseMultiUserPickerCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseMultiUserPickerCustomField() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseMultiUserPickerCustomField() got = (%v), want (%v)", err, testCase.Err)
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
				ID:    "10054",
				Child: &CascadingSelectChildScheme{
					Self:  "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10057",
					Value: "Costa Rica",
					ID:    "10057",
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
			wantErr: true,
			Err:     ErrNoCascadingParentError,
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
			Err:     ErrNoCascadingParentError,
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
			Err:     ErrNoFieldInformationError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseCascadingSelectCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseCascadingSelectCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseCascadingSelectCustomField() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseCascadingSelectCustomField() got = (%v), want (%v)", err, testCase.Err)
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
			wantErr: true,
			Err:     ErrNoMultiVersionTypeError,
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
			Err:     ErrNoFieldInformationError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseMultiVersionCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseUserPickerCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseUserPickerCustomField() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseUserPickerCustomField() got = (%v), want (%v)", err, testCase.Err)
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
	  "emailAddress": "example@example.com",
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
				EmailAddress: "example@example.com",
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
			wantErr: true,
			Err:     ErrNoUserTypeError,
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
			Err:     ErrNoFieldInformationError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseUserPickerCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseUserPickerCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseUserPickerCustomField() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseUserPickerCustomField() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseStringCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
  "fields": {
    "customfield_10045": "Lorem ipsum dolor sit amet, consetetur sadipscing elitr"
  }
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
  "fields_no_mapped": {
    "customfield_10045": "At vero eos et accusam et justo"
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
		"customfield_10045": 10023.345
	}
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10045",
			},
			want:    "Lorem ipsum dolor sit amet, consetetur sadipscing elitr",
			wantErr: false,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10045",
			},
			want:    "",
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},

		{
			name: "when the buffer field value is null",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10045",
			},
			want:    "",
			wantErr: true,
			Err:     ErrNoTextTypeError,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10045",
			},
			want:    "",
			wantErr: true,
			Err:     ErrNoTextTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseStringCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseStringCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseStringCustomField() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseStringCustomField() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseDatePickerCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
  "fields": {
    "customfield_10045": "2023-09-22"
  }
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
  "fields_no_mapped": {
    "customfield_10045": "2023-09-22"
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
		"customfield_10045": 10023.345
	}
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10045",
			},
			want:    time.Date(2023, time.September, 22, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10045",
			},
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},

		{
			name: "when the buffer field value is null",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10045",
			},
			wantErr: true,
			Err:     ErrNoDatePickerTypeError,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10045",
			},
			wantErr: true,
			Err:     ErrNoDatePickerTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseDatePickerCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseDatePickerCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseDatePickerCustomField() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseDatePickerCustomField() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseDateTimeCustomField(t *testing.T) {

	mockedTime, err := time.Parse("2006-01-02T15:04:05.000-0700", "2023-07-12T16:00:00.000+0100")
	if err != nil {
		t.Fatal(err)
	}

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
  "fields": {
    "customfield_10045": "2023-07-12T16:00:00.000+0100"
  }
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
  "fields_no_mapped": {
    "customfield_10045": "2023-07-12T16:00:00.000+0100"
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
		"customfield_10045": 10023.345
	}
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10045",
			},
			want:    mockedTime,
			wantErr: false,
		},

		{
			name: "when the buffer does not contains the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10045",
			},
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},

		{
			name: "when the buffer field value is null",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10045",
			},
			wantErr: true,
			Err:     ErrNoDateTimeTypeError,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoInfo,
				customField: "customfield_10045",
			},
			wantErr: true,
			Err:     ErrNoDateTimeTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseDateTimeCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseDateTimeCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseDateTimeCustomField() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseDateTimeCustomField() got = (%v), want (%v)", err, testCase.Err)
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
			wantErr: true,
			Err:     ErrNoFloatTypeError,
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
			Err:     ErrNoFloatTypeError,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoJSON,
				customField: "customfield_10045",
			},
			want:    0,
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseFloatCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseFloatCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseFloatCustomField() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseFloatCustomField() got = (%v), want (%v)", err, testCase.Err)
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
			wantErr: true,
			Err:     ErrNoLabelsTypeError,
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
			Err:     ErrNoLabelsTypeError,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoJSON,
				customField: "customfield_10045",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseLabelCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseLabelCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseLabelCustomField() got = %v, want %v", got, testCase.want)
			}

			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseLabelCustomField() got = (%v), want (%v)", err, testCase.Err)
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
        "id": 5,
        "name": "KP Sprint 3",
        "state": "active",
        "boardId": 4,
        "goal": "",
        "startDate": "2023-03-04T02:03:16.273Z",
        "endDate": "2023-03-17T02:03:00.000Z",
		"completeDate": "2023-03-04T02:03:16.273Z"
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
        "id": 5,
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
					ID:           5,
					State:        "active",
					Name:         "KP Sprint 3",
					StartDate:    "2023-03-04T02:03:16.273Z",
					CompleteDate: "2023-03-04T02:03:16.273Z",
					EndDate:      "2023-03-17T02:03:00.000Z",
					BoardID:      4,
					Goal:         "",
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
			wantErr: true,
			Err:     ErrNoSprintTypeError,
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
			Err:     ErrNoSprintTypeError,
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
			Err:     ErrNoFieldInformationError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseSprintCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseSprintCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseSprintCustomField() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseSprintCustomField() got = (%v), want (%v)", err, testCase.Err)
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
			wantErr: true,
			Err:     ErrNoSelectTypeError,
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
			Err:     ErrNoSelectTypeError,
		},

		{
			name: "when the buffer cannot be parsed",
			args: args{
				buffer:      bufferMockedWithNoJSON,
				customField: "customfield_10045",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoFieldInformationError,
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

func TestParseAssetCustomField(t *testing.T) {

	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
	"fields": {
		"customfield_10046": [
		 {
			"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f1",
			"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
			"objectId": "1"
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
			"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f1",
			"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
			"objectId": "1"
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

	type args struct {
		buffer      bytes.Buffer
		customField string
	}

	testCases := []struct {
		name    string
		args    args
		want    []*CustomFieldAssetScheme
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
			want: []*CustomFieldAssetScheme{
				{
					WorkspaceID: "5e037d73-1c0a-43ce-adca-f169a42557f1",
					ID:          "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
					ObjectID:    "1",
				},
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
			Err:     ErrNoAssetTypeError,
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
			Err:     ErrNoAssetTypeError,
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
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseAssetCustomField(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseAssetCustomField() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseAssetCustomField() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseAssetCustomField() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseRequestTypeCustomField(t *testing.T) {
	bufferMocked := bytes.Buffer{}
	bufferMocked.WriteString(`
{
  "fields": {
    "customfield_10010": {
      "_links": {
        "jiraRest": "https://mydomain.atlassian.net/rest/api/2/issue/144906",
        "web": "https://mydomain.atlassian.net/servicedesk/customer/portal/2/ESD-40928",
        "self": "https://mydomain.atlassian.net/rest/servicedeskapi/request/144906",
        "agent": "https://mydomain.atlassian.net/browse/ESD-40928"
      },
      "requestType": {
        "_expands": [
          "field"
        ],
        "id": "96",
        "_links": {
          "self": "https://mydomain.atlassian.net/rest/servicedeskapi/servicedesk/2/requesttype/96"
        },
        "name": "General Service Request",
        "description": "",
        "helpText": "",
        "issueTypeId": "10039",
        "serviceDeskId": "2",
        "portalId": "2",
        "groupIds": [],
        "icon": {
          "id": "10466",
          "_links": {
            "iconUrls": {
              "48x48": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=large",
              "24x24": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=small",
              "16x16": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=xsmall",
              "32x32": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=medium"
            }
          }
        }
      },
      "currentStatus": {
        "status": "Check billing account",
        "statusCategory": "NEW",
        "statusDate": {
          "jira": "2024-06-28T08:33:57.313+0200",
          "friendly": "Friday 15:33",
          "epochMillis": 1719556437313
        }
      }
    }
  }
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
{
  "no_fields": {
    "customfield_10010": {
      "_links": {
        "jiraRest": "https://mydomain.atlassian.net/rest/api/2/issue/144906",
        "web": "https://mydomain.atlassian.net/servicedesk/customer/portal/2/ESD-40928",
        "self": "https://mydomain.atlassian.net/rest/servicedeskapi/request/144906",
        "agent": "https://mydomain.atlassian.net/browse/ESD-40928"
      },
      "requestType": {
        "_expands": [
          "field"
        ],
        "id": "96",
        "_links": {
          "self": "https://mydomain.atlassian.net/rest/servicedeskapi/servicedesk/2/requesttype/96"
        },
        "name": "General Service Request",
        "description": "",
        "helpText": "",
        "issueTypeId": "10039",
        "serviceDeskId": "2",
        "portalId": "2",
        "groupIds": [],
        "icon": {
          "id": "10466",
          "_links": {
            "iconUrls": {
              "48x48": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=large",
              "24x24": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=small",
              "16x16": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=xsmall",
              "32x32": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=medium"
            }
          }
        }
      },
      "currentStatus": {
        "status": "Check billing account",
        "statusCategory": "NEW",
        "statusDate": {
          "iso8601": "2024-06-28T15:33:57+0900",
          "jira": "2024-06-28T08:33:57.313+0200",
          "friendly": "Friday 15:33",
          "epochMillis": 1719556437313
        }
      }
    }
  }
}`)

	bufferMockedWithWrongType := bytes.Buffer{}
	bufferMockedWithWrongType.WriteString(`
{
  "fields": {
    "customfield_10010": ""
  }
}`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    *CustomFieldRequestTypeScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10010",
			},
			want: &CustomFieldRequestTypeScheme{
				Links: &CustomFieldRequestTypeLinkScheme{
					JiraRest: "https://mydomain.atlassian.net/rest/api/2/issue/144906",
					Web:      "https://mydomain.atlassian.net/servicedesk/customer/portal/2/ESD-40928",
					Self:     "https://mydomain.atlassian.net/rest/servicedeskapi/request/144906",
					Agent:    "https://mydomain.atlassian.net/browse/ESD-40928",
				},
				RequestType: &CustomerRequestTypeScheme{
					ID:            "96",
					Name:          "General Service Request",
					Description:   "",
					HelpText:      "",
					IssueTypeID:   "10039",
					ServiceDeskID: "2",
					GroupIds:      []string{},
				},
				CurrentStatus: &CustomerRequestCurrentStatusScheme{
					Status:         "Check billing account",
					StatusCategory: "NEW",
					StatusDate: &CustomerRequestCurrentStatusDateScheme{
						Jira:        "2024-06-28T08:33:57.313+0200",
						Friendly:    "Friday 15:33",
						EpochMillis: 1719556437313,
					},
				},
			},
		},
		{
			name: "when the buffer contains no custom field",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10020",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoRequestTypeError,
		},
		{
			name: "when the buffer contains no fields",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10010",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoFieldInformationError,
		},
		{
			name: "when the buffer contains wrong type",
			args: args{
				buffer:      bufferMockedWithWrongType,
				customField: "customfield_10010",
			},
			want:    nil,
			wantErr: true,
			Err:     ErrNoRequestTypeError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRequestTypeCustomField(tt.args.buffer, tt.args.customField)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRequestTypeCustomField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRequestTypeCustomField() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(err, tt.Err) {
				t.Errorf("ParseRequestTypeCustomField() got = (%v), want (%v)", err, tt.Err)
			}
		})
	}
}

func TestParseMultiSelectCustomFields(t *testing.T) {

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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
		want    map[string][]*CustomFieldContextOptionScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string][]*CustomFieldContextOptionScheme{
				"KP-22": {
					{
						ID:    "10046",
						Value: "Option 3",
					},
					{
						ID:    "10047",
						Value: "Option 4",
					},
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string][]*CustomFieldContextOptionScheme{
				"KP-22": {
					{
						ID:    "10046",
						Value: "Option 3",
					},
					{
						ID:    "10047",
						Value: "Option 4",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseMultiSelectCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseMultiSelectCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseMultiSelectCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseMultiSelectCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseMultiGroupPickerCustomFields(t *testing.T) {

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
                "customfield_10046": [
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
        },
		{
            "expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
            "id": "10035",
            "self": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
            "key": "KP-23",
            "fields": {
                "customfield_10046": [
				  {
					"name": "jira-administrators-2",
					"groupId": "1da6f895-2b42-423b-8bfb-1e09ee8d7764",
					"self": "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=1da6f895-2b42-423b-8bfb-1e09ee8d7764"
				  },
				  {
					"name": "jira-administrators-system-2",
					"groupId": "be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
					"self": "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=be9ba0ab-ecdc-445b-9ce6-b95202026c1a"
				  }
				]
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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
                "customfield_10046": null,
            }
        },
		{
            "expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
            "id": "10035",
            "self": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
            "key": "KP-23",
            "fields": {
                "customfield_10046": [
				  {
					"name": "jira-administrators-2",
					"groupId": "1da6f895-2b42-423b-8bfb-1e09ee8d7764",
					"self": "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=1da6f895-2b42-423b-8bfb-1e09ee8d7764"
				  },
				  {
					"name": "jira-administrators-system-2",
					"groupId": "be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
					"self": "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=be9ba0ab-ecdc-445b-9ce6-b95202026c1a"
				  }
				]
            }
        }
    ]
}`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
                "customfield_10046": "self": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
            }
        },
		{
            "expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
            "id": "10035",
            "self": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
            "key": "KP-23",
            "fields": {
                "customfield_10046": [
				  {
					"name": "jira-administrators-2",
					"groupId": "1da6f895-2b42-423b-8bfb-1e09ee8d7764",
					"self": "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=1da6f895-2b42-423b-8bfb-1e09ee8d7764"
				  },
				  {
					"name": "jira-administrators-system-2",
					"groupId": "be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
					"self": "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=be9ba0ab-ecdc-445b-9ce6-b95202026c1a"
				  }
				]
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
		want    map[string][]*GroupDetailScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string][]*GroupDetailScheme{
				"KP-22": {
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
				"KP-23": {
					{
						Self:    "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=1da6f895-2b42-423b-8bfb-1e09ee8d7764",
						Name:    "jira-administrators-2",
						GroupID: "1da6f895-2b42-423b-8bfb-1e09ee8d7764",
					},
					{
						Self:    "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
						Name:    "jira-administrators-system-2",
						GroupID: "be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
					},
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string][]*GroupDetailScheme{
				"KP-23": {
					{
						Self:    "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=1da6f895-2b42-423b-8bfb-1e09ee8d7764",
						Name:    "jira-administrators-2",
						GroupID: "1da6f895-2b42-423b-8bfb-1e09ee8d7764",
					},
					{
						Self:    "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
						Name:    "jira-administrators-system-2",
						GroupID: "be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
					},
				},
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			want: map[string][]*GroupDetailScheme{
				"KP-23": {
					{
						Self:    "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=1da6f895-2b42-423b-8bfb-1e09ee8d7764",
						Name:    "jira-administrators-2",
						GroupID: "1da6f895-2b42-423b-8bfb-1e09ee8d7764",
					},
					{
						Self:    "https://ctreminiom.atlassian.net/rest/api/3/group?groupId=be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
						Name:    "jira-administrators-system-2",
						GroupID: "be9ba0ab-ecdc-445b-9ce6-b95202026c1a",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseMultiGroupPickerCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseMultiGroupPickerCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseMultiGroupPickerCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseMultiGroupPickerCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseMultiUserPickerCustomFields(t *testing.T) {

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
            "customfield_10046":[
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
                  "accountId":"5e5f6a63157ed50cd2b9eaca",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "24x24":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "16x16":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "32x32":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"Asia/Dhaka",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
                  "accountId":"5b86be50b8e3cb5895860d6d",
                  "emailAddress":"ctreminiom079@gmail.com",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "24x24":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "16x16":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "32x32":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"America/Guatemala",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=557058%3Ad6b5955a-e193-41e1-b051-79cdb0755d68",
                  "accountId":"557058:d6b5955a-e193-41e1-b051-79cdb0755d68",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "24x24":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "16x16":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "32x32":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png"
                  },
                  "displayName":"Trello",
                  "active":true,
                  "timeZone":"Europe/London",
                  "accountType":"app"
               }
            ]
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
                  "accountId":"5e5f6a63157ed50cd2b9eaca",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "24x24":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "16x16":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "32x32":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"Asia/Dhaka",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
                  "accountId":"5b86be50b8e3cb5895860d6d",
                  "emailAddress":"ctreminiom079@gmail.com",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "24x24":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "16x16":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "32x32":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"America/Guatemala",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=557058%3Ad6b5955a-e193-41e1-b051-79cdb0755d68",
                  "accountId":"557058:d6b5955a-e193-41e1-b051-79cdb0755d68",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "24x24":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "16x16":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "32x32":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png"
                  },
                  "displayName":"Trello",
                  "active":true,
                  "timeZone":"Europe/London",
                  "accountType":"app"
               }
            ]
         }
      }
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046": null,
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
                  "accountId":"5e5f6a63157ed50cd2b9eaca",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "24x24":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "16x16":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "32x32":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"Asia/Dhaka",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
                  "accountId":"5b86be50b8e3cb5895860d6d",
                  "emailAddress":"ctreminiom079@gmail.com",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "24x24":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "16x16":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "32x32":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"America/Guatemala",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=557058%3Ad6b5955a-e193-41e1-b051-79cdb0755d68",
                  "accountId":"557058:d6b5955a-e193-41e1-b051-79cdb0755d68",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "24x24":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "16x16":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "32x32":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png"
                  },
                  "displayName":"Trello",
                  "active":true,
                  "timeZone":"Europe/London",
                  "accountType":"app"
               }
            ]
         }
      }
   ]
}`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
                  "accountId":"5e5f6a63157ed50cd2b9eaca",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "24x24":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "16x16":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "32x32":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"Asia/Dhaka",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
                  "accountId":"5b86be50b8e3cb5895860d6d",
                  "emailAddress":"ctreminiom079@gmail.com",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "24x24":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "16x16":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "32x32":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"America/Guatemala",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=557058%3Ad6b5955a-e193-41e1-b051-79cdb0755d68",
                  "accountId":"557058:d6b5955a-e193-41e1-b051-79cdb0755d68",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "24x24":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "16x16":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "32x32":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png"
                  },
                  "displayName":"Trello",
                  "active":true,
                  "timeZone":"Europe/London",
                  "accountType":"app"
               }
            ]
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]*UserDetailScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string][]*UserDetailScheme{
				"KP-22": {
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
						DisplayName:  "Carlos Treminio",
						EmailAddress: "ctreminiom079@gmail.com",
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
				"KP-23": {
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
						DisplayName:  "Carlos Treminio",
						EmailAddress: "ctreminiom079@gmail.com",
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string][]*UserDetailScheme{
				"KP-23": {
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
						DisplayName:  "Carlos Treminio",
						EmailAddress: "ctreminiom079@gmail.com",
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
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			want: map[string][]*UserDetailScheme{
				"KP-23": {
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
						DisplayName:  "Carlos Treminio",
						EmailAddress: "ctreminiom079@gmail.com",
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
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseMultiUserPickerCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseMultiUserPickerCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseMultiUserPickerCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseMultiUserPickerCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseCascadingCustomFields(t *testing.T) {

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
            "customfield_10046":{
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
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":{
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
      },
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046":{
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
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": null
         }
      },
   ]
}
`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
                  "accountId":"5e5f6a63157ed50cd2b9eaca",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "24x24":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "16x16":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "32x32":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"Asia/Dhaka",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
                  "accountId":"5b86be50b8e3cb5895860d6d",
                  "emailAddress":"ctreminiom079@gmail.com",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "24x24":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "16x16":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "32x32":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"America/Guatemala",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=557058%3Ad6b5955a-e193-41e1-b051-79cdb0755d68",
                  "accountId":"557058:d6b5955a-e193-41e1-b051-79cdb0755d68",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "24x24":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "16x16":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "32x32":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png"
                  },
                  "displayName":"Trello",
                  "active":true,
                  "timeZone":"Europe/London",
                  "accountType":"app"
               }
            ]
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*CascadingSelectScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string]*CascadingSelectScheme{
				"KP-22": {
					Self:  "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10054",
					Value: "America",
					ID:    "10054",
					Child: &CascadingSelectChildScheme{
						Self:  "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10057",
						Value: "Costa Rica",
						ID:    "10057",
					},
				},
				"KP-23": {
					Self:  "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10054",
					Value: "America",
					ID:    "10054",
					Child: &CascadingSelectChildScheme{
						Self:  "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10057",
						Value: "Costa Rica",
						ID:    "10057",
					},
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string]*CascadingSelectScheme{
				"KP-22": {
					Self:  "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10054",
					Value: "America",
					ID:    "10054",
					Child: &CascadingSelectChildScheme{
						Self:  "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10057",
						Value: "Costa Rica",
						ID:    "10057",
					},
				},
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseCascadingCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseCascadingCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseCascadingCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseCascadingCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseMultiVersionCustomFields(t *testing.T) {

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
            "customfield_10046":[
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
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
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
      },
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046":[
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
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": null
         }
      },
   ]
}
`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": "https://ctreminiom.atlassian.net/rest/api/3/version/10000"
         }
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": "https://ctreminiom.atlassian.net/rest/api/3/version/10000"
         }
      },
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]*VersionDetailScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string][]*VersionDetailScheme{

				"KP-22": {
					{
						Self:        "https://ctreminiom.atlassian.net/rest/api/3/version/10000",
						ID:          "10000",
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
				"KP-23": {
					{
						Self:        "https://ctreminiom.atlassian.net/rest/api/3/version/10000",
						ID:          "10000",
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string][]*VersionDetailScheme{

				"KP-22": {
					{
						Self:        "https://ctreminiom.atlassian.net/rest/api/3/version/10000",
						ID:          "10000",
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
			},

			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseMultiVersionCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseMultiVersionCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseMultiVersionCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseMultiVersionCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseUserPickerCustomFields(t *testing.T) {

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
            "customfield_10046":{
			  "self": "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6acefc1fca0af44135f8",
			  "accountId": "5e5f6acefc1fca0af44135f8",
			  "emailAddress": "example@example.com",
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
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":{
			  "self": "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6acefc1fca0af44135f8",
			  "accountId": "5e5f6acefc1fca0af44135f8",
			  "emailAddress": "example@example.com",
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
      },
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046":{
			  "self": "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6acefc1fca0af44135f8",
			  "accountId": "5e5f6acefc1fca0af44135f8",
			  "emailAddress": "example@example.com",
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
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": null
         }
      },
   ]
}
`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
                  "accountId":"5e5f6a63157ed50cd2b9eaca",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "24x24":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "16x16":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "32x32":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"Asia/Dhaka",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
                  "accountId":"5b86be50b8e3cb5895860d6d",
                  "emailAddress":"ctreminiom079@gmail.com",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "24x24":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "16x16":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "32x32":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"America/Guatemala",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=557058%3Ad6b5955a-e193-41e1-b051-79cdb0755d68",
                  "accountId":"557058:d6b5955a-e193-41e1-b051-79cdb0755d68",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "24x24":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "16x16":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "32x32":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png"
                  },
                  "displayName":"Trello",
                  "active":true,
                  "timeZone":"Europe/London",
                  "accountType":"app"
               }
            ]
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*UserDetailScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string]*UserDetailScheme{
				"KP-22": {
					Self:         "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6acefc1fca0af44135f8",
					AccountID:    "5e5f6acefc1fca0af44135f8",
					EmailAddress: "example@example.com",
					DisplayName:  "Eduardo Navarro",
					Active:       true,
					TimeZone:     "Europe/London",
					AccountType:  "atlassian",
				},
				"KP-23": {
					Self:         "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6acefc1fca0af44135f8",
					AccountID:    "5e5f6acefc1fca0af44135f8",
					EmailAddress: "example@example.com",
					DisplayName:  "Eduardo Navarro",
					Active:       true,
					TimeZone:     "Europe/London",
					AccountType:  "atlassian",
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string]*UserDetailScheme{
				"KP-22": {
					Self:         "https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6acefc1fca0af44135f8",
					AccountID:    "5e5f6acefc1fca0af44135f8",
					EmailAddress: "example@example.com",
					DisplayName:  "Eduardo Navarro",
					Active:       true,
					TimeZone:     "Europe/London",
					AccountType:  "atlassian",
				},
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseUserPickerCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseUserPickerCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseUserPickerCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseUserPickerCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseStringCustomFields(t *testing.T) {

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
            "customfield_10046":  "Lorem ipsum dolor sit amet, consetetur sadipscing elitr"
         }
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":  "Lorem ipsum dolor sit amet, consetetur sadipscing elitr"
         }
      },
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046":  "Lorem ipsum dolor sit amet, consetetur sadipscing elitr"
         }
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": null
         }
      },
   ]
}
`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": ["self": "asd"],
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
                  "accountId":"5e5f6a63157ed50cd2b9eaca",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "24x24":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "16x16":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "32x32":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"Asia/Dhaka",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
                  "accountId":"5b86be50b8e3cb5895860d6d",
                  "emailAddress":"ctreminiom079@gmail.com",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "24x24":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "16x16":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "32x32":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"America/Guatemala",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=557058%3Ad6b5955a-e193-41e1-b051-79cdb0755d68",
                  "accountId":"557058:d6b5955a-e193-41e1-b051-79cdb0755d68",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "24x24":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "16x16":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "32x32":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png"
                  },
                  "displayName":"Trello",
                  "active":true,
                  "timeZone":"Europe/London",
                  "accountType":"app"
               }
            ]
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string]string{
				"KP-22": "Lorem ipsum dolor sit amet, consetetur sadipscing elitr",
				"KP-23": "Lorem ipsum dolor sit amet, consetetur sadipscing elitr",
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string]string{
				"KP-22": "Lorem ipsum dolor sit amet, consetetur sadipscing elitr",
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseStringCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseStringCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseStringCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseStringCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseFloatCustomFields(t *testing.T) {

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
            "customfield_10046":  4003939
         }
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": 4003939
         }
      },
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046": 4003939
         }
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": null
         }
      },
   ]
}
`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": ["self": "asd"],
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
                  "accountId":"5e5f6a63157ed50cd2b9eaca",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "24x24":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "16x16":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "32x32":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"Asia/Dhaka",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
                  "accountId":"5b86be50b8e3cb5895860d6d",
                  "emailAddress":"ctreminiom079@gmail.com",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "24x24":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "16x16":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "32x32":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"America/Guatemala",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=557058%3Ad6b5955a-e193-41e1-b051-79cdb0755d68",
                  "accountId":"557058:d6b5955a-e193-41e1-b051-79cdb0755d68",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "24x24":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "16x16":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "32x32":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png"
                  },
                  "displayName":"Trello",
                  "active":true,
                  "timeZone":"Europe/London",
                  "accountType":"app"
               }
            ]
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]float64
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string]float64{
				"KP-22": 4003939,
				"KP-23": 4003939,
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string]float64{
				"KP-22": 4003939,
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseFloatCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseFloatCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseFloatCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseFloatCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseLabelCustomFields(t *testing.T) {

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
            "customfield_10046":  ["label-1", "label-2"]
         }
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":  ["label-1", "label-2"]
         }
      },
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046":  ["label-1", "label-2"]
         }
      },
     {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": null
         }
      },
   ]
}
`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": ["self": "asd"],
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5e5f6a63157ed50cd2b9eaca",
                  "accountId":"5e5f6a63157ed50cd2b9eaca",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "24x24":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "16x16":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png",
                     "32x32":"https://secure.gravatar.com/avatar/2e6d2ee8550c63137e196a2890bc25a9?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-4.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"Asia/Dhaka",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=5b86be50b8e3cb5895860d6d",
                  "accountId":"5b86be50b8e3cb5895860d6d",
                  "emailAddress":"ctreminiom079@gmail.com",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "24x24":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "16x16":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png",
                     "32x32":"https://secure.gravatar.com/avatar/b830f79c6cc32dcbcb9842f98cd3d3cd?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FCT-6.png"
                  },
                  "displayName":"Carlos Treminio",
                  "active":true,
                  "timeZone":"America/Guatemala",
                  "accountType":"atlassian"
               },
               {
                  "self":"https://ctreminiom.atlassian.net/rest/api/3/user?accountId=557058%3Ad6b5955a-e193-41e1-b051-79cdb0755d68",
                  "accountId":"557058:d6b5955a-e193-41e1-b051-79cdb0755d68",
                  "avatarUrls":{
                     "48x48":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "24x24":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "16x16":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png",
                     "32x32":"https://secure.gravatar.com/avatar/53e3e37950768a905d53cebdfcbd63e3?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FT-1.png"
                  },
                  "displayName":"Trello",
                  "active":true,
                  "timeZone":"Europe/London",
                  "accountType":"app"
               }
            ]
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]string
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string][]string{
				"KP-22": {"label-1", "label-2"},
				"KP-23": {"label-1", "label-2"},
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string][]string{
				"KP-22": {"label-1", "label-2"},
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseLabelCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseLabelCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseLabelCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseLabelCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseSprintCustomFields(t *testing.T) {

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
            "customfield_10046":[
			  {
				"id": 5,
				"name": "KP Sprint 3",
				"state": "active",
				"boardId": 4,
				"goal": "",
				"startDate": "2023-03-04T02:03:16.273Z",
				"endDate": "2023-03-17T02:03:00.000Z",
				"completeDate": "2023-03-04T02:03:16.273Z"
			  }
			]
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
			  {
				"id": 5,
				"name": "KP Sprint 3",
				"state": "active",
				"boardId": 4,
				"goal": "",
				"startDate": "2023-03-04T02:03:16.273Z",
				"endDate": "2023-03-17T02:03:00.000Z",
				"completeDate": "2023-03-04T02:03:16.273Z"
			  }
			]
         }
      }
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046": null,
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
			  {
				"id": 5,
				"name": "KP Sprint 3",
				"state": "active",
				"boardId": 4,
				"goal": "",
				"startDate": "2023-03-04T02:03:16.273Z",
				"endDate": "2023-03-17T02:03:00.000Z",
				"completeDate": "2023-03-04T02:03:16.273Z"
			  }
			]
         }
      }
   ]
}`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":"5e5f6a63157ed50cd2b9eaca"
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]*SprintDetailScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string][]*SprintDetailScheme{
				"KP-22": {
					{
						ID:           5,
						State:        "active",
						Name:         "KP Sprint 3",
						StartDate:    "2023-03-04T02:03:16.273Z",
						EndDate:      "2023-03-17T02:03:00.000Z",
						CompleteDate: "2023-03-04T02:03:16.273Z",
						BoardID:      4,
					},
				},
				"KP-23": {
					{
						ID:           5,
						State:        "active",
						Name:         "KP Sprint 3",
						StartDate:    "2023-03-04T02:03:16.273Z",
						EndDate:      "2023-03-17T02:03:00.000Z",
						CompleteDate: "2023-03-04T02:03:16.273Z",
						BoardID:      4,
					},
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string][]*SprintDetailScheme{
				"KP-23": {
					{
						ID:           5,
						State:        "active",
						Name:         "KP Sprint 3",
						StartDate:    "2023-03-04T02:03:16.273Z",
						EndDate:      "2023-03-17T02:03:00.000Z",
						CompleteDate: "2023-03-04T02:03:16.273Z",
						BoardID:      4,
					},
				},
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseSprintCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseSprintCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseSprintCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseSprintCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseSelectCustomFields(t *testing.T) {

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
            "customfield_10046":{
			  "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10058",
			  "value": "Scranton 1",
			  "id": "10058"
			}
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":{
			  "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10058",
			  "value": "Scranton 1",
			  "id": "10058"
			}
         }
      }
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046": null,
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":{
			  "self": "https://ctreminiom.atlassian.net/rest/api/3/customFieldOption/10058",
			  "value": "Scranton 1",
			  "id": "10058"
			}
         }
      }
   ]
}`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":"5e5f6a63157ed50cd2b9eaca"
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*CustomFieldContextOptionScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string]*CustomFieldContextOptionScheme{
				"KP-22": {
					ID:    "10058",
					Value: "Scranton 1",
				},
				"KP-23": {
					ID:    "10058",
					Value: "Scranton 1",
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string]*CustomFieldContextOptionScheme{
				"KP-23": {
					ID:    "10058",
					Value: "Scranton 1",
				},
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseSelectCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseSelectCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseSelectCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseSelectCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseAssetCustomFields(t *testing.T) {

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
            "customfield_10046":[
			 {
				"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f1",
				"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
				"objectId": "1"
			  },
				{
				"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f2",
				"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:2",
				"objectId": "1"
			  }
			]
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
			 {
				"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f1",
				"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
				"objectId": "1"
			  },
				{
				"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f2",
				"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:2",
				"objectId": "1"
			  }
			]
         }
      }
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046": null
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
			 {
				"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f1",
				"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
				"objectId": "1"
			  },
				{
				"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f2",
				"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:2",
				"objectId": "1"
			  }
			]
         }
      }
   ]
}`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": "test"
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046":[
			 {
				"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f1",
				"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
				"objectId": "1"
			  },
				{
				"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f2",
				"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:2",
				"objectId": "1"
			  }
			]
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]*CustomFieldAssetScheme
		wantErr bool
		Err     error
	}{

		/*
			"customfield_10046":[
					 {
						"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f1",
						"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
						"objectId": "1"
					  },
						{
						"workspaceId": "5e037d73-1c0a-43ce-adca-f169a42557f2",
						"id": "5e037d73-1c0a-43ce-adca-f169a42557f1:2",
						"objectId": "1"
					  }
					]
		*/

		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string][]*CustomFieldAssetScheme{
				"KP-22": {
					{
						WorkspaceID: "5e037d73-1c0a-43ce-adca-f169a42557f1",
						ID:          "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
						ObjectID:    "1",
					},
					{
						WorkspaceID: "5e037d73-1c0a-43ce-adca-f169a42557f2",
						ID:          "5e037d73-1c0a-43ce-adca-f169a42557f1:2",
						ObjectID:    "1",
					},
				},
				"KP-23": {
					{
						WorkspaceID: "5e037d73-1c0a-43ce-adca-f169a42557f1",
						ID:          "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
						ObjectID:    "1",
					},
					{
						WorkspaceID: "5e037d73-1c0a-43ce-adca-f169a42557f2",
						ID:          "5e037d73-1c0a-43ce-adca-f169a42557f1:2",
						ObjectID:    "1",
					},
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string][]*CustomFieldAssetScheme{
				"KP-23": {
					{
						WorkspaceID: "5e037d73-1c0a-43ce-adca-f169a42557f1",
						ID:          "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
						ObjectID:    "1",
					},
					{
						WorkspaceID: "5e037d73-1c0a-43ce-adca-f169a42557f2",
						ID:          "5e037d73-1c0a-43ce-adca-f169a42557f1:2",
						ObjectID:    "1",
					},
				},
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			want: map[string][]*CustomFieldAssetScheme{
				"KP-23": {
					{
						WorkspaceID: "5e037d73-1c0a-43ce-adca-f169a42557f1",
						ID:          "5e037d73-1c0a-43ce-adca-f169a42557f1:1",
						ObjectID:    "1",
					},
					{
						WorkspaceID: "5e037d73-1c0a-43ce-adca-f169a42557f2",
						ID:          "5e037d73-1c0a-43ce-adca-f169a42557f1:2",
						ObjectID:    "1",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseAssetCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseAssetCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseAssetCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseAssetCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseDatePickerCustomFields(t *testing.T) {

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
            "customfield_10046": "2023-09-22"
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
             "customfield_10046": "2023-09-23"
         }
      }
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046": null,
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": "2023-09-23"
         }
      }
   ]
}`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": 33030303,
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": true
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]time.Time
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string]time.Time{
				"KP-22": time.Date(2023, time.September, 22, 0, 0, 0, 0, time.UTC),
				"KP-23": time.Date(2023, time.September, 23, 0, 0, 0, 0, time.UTC),
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string]time.Time{
				"KP-23": time.Date(2023, time.September, 23, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseDatePickerCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseSelectCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseSelectCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseSelectCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseDateTimeCustomFields(t *testing.T) {

	mockedTime, err := time.Parse("2006-01-02T15:04:05.000-0700", "2023-07-12T16:00:00.000+0100")
	if err != nil {
		t.Fatal(err)
	}

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
            "customfield_10046": "2023-07-12T16:00:00.000+0100"
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
             "customfield_10046": "2023-07-12T16:00:00.000+0100"
         }
      }
   ]
}
`)

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
                "customfield_10046": [
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10046",
                        "value": "Option 3",
                        "id": "10046"
                    },
                    {
                        "self": "https://ctreminiom.atlassian.net/rest/api/2/customFieldOption/10047",
                        "value": "Option 4",
                        "id": "10047"
                    }
                ]
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

	bufferMockedWithNullValues := bytes.Buffer{}
	bufferMockedWithNullValues.WriteString(`
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
            "customfield_10046": null,
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": "2023-07-12T16:00:00.000+0100"
         }
      }
   ]
}`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10046": 33030303,
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10046": true
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]time.Time
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10046",
			},
			want: map[string]time.Time{
				"KP-22": mockedTime,
				"KP-23": mockedTime,
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

		{
			name: "when the buffer contains null customfields",
			args: args{
				buffer:      bufferMockedWithNullValues,
				customField: "customfield_10046",
			},
			want: map[string]time.Time{
				"KP-23": mockedTime,
			},
			wantErr: false,
		},

		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10046",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseDateTimeCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseSelectCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseSelectCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseSelectCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}

func TestParseRequestTypeCustomFields(t *testing.T) {
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
        "customfield_10010": {
          "_links": {
            "jiraRest": "https://mydomain.atlassian.net/rest/api/2/issue/144906",
            "web": "https://mydomain.atlassian.net/servicedesk/customer/portal/2/ESD-40928",
            "self": "https://mydomain.atlassian.net/rest/servicedeskapi/request/144906",
            "agent": "https://mydomain.atlassian.net/browse/ESD-40928"
          },
          "requestType": {
            "_expands": [
              "field"
            ],
            "id": "96",
            "_links": {
              "self": "https://mydomain.atlassian.net/rest/servicedeskapi/servicedesk/2/requesttype/96"
            },
            "name": "General Service Request",
            "description": "",
            "helpText": "",
            "issueTypeId": "10039",
            "serviceDeskId": "2",
            "portalId": "2",
            "groupIds": [],
            "icon": {
              "id": "10466",
              "_links": {
                "iconUrls": {
                  "48x48": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=large",
                  "24x24": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=small",
                  "16x16": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=xsmall",
                  "32x32": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=medium"
                }
              }
            }
          },
          "currentStatus": {
            "status": "Check billing account",
            "statusCategory": "NEW",
            "statusDate": {
              "jira": "2024-06-28T08:33:57.313+0200",
              "friendly": "Friday 15:33",
              "epochMillis": 1719556437313
            }
          }
        }
      }
    },
    {
      "expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
      "id": "10035",
      "self": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
      "key": "KP-23",
      "fields": {
        "customfield_10010": {
          "_links": {
            "jiraRest": "https://mydomain.atlassian.net/rest/api/2/issue/144906",
            "web": "https://mydomain.atlassian.net/servicedesk/customer/portal/2/ESD-40928",
            "self": "https://mydomain.atlassian.net/rest/servicedeskapi/request/144906",
            "agent": "https://mydomain.atlassian.net/browse/ESD-40928"
          },
          "requestType": {
            "_expands": [
              "field"
            ],
            "id": "96",
            "_links": {
              "self": "https://mydomain.atlassian.net/rest/servicedeskapi/servicedesk/2/requesttype/96"
            },
            "name": "General Service Request",
            "description": "",
            "helpText": "",
            "issueTypeId": "10039",
            "serviceDeskId": "2",
            "portalId": "2",
            "groupIds": [],
            "icon": {
              "id": "10466",
              "_links": {
                "iconUrls": {
                  "48x48": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=large",
                  "24x24": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=small",
                  "16x16": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=xsmall",
                  "32x32": "https://mydomain.atlassian.net/rest/api/2/universal_avatar/view/type/SD_REQTYPE/avatar/10466?size=medium"
                }
              }
            }
          },
          "currentStatus": {
            "status": "Check billing account",
            "statusCategory": "NEW",
            "statusDate": {
              "jira": "2024-06-28T08:33:57.313+0200",
              "friendly": "Friday 15:33",
              "epochMillis": 1719556437313
            }
          }
        }
      }
    }
  ]
}
`)

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
      "fields": {
        "customfield_10010": {}
      }
    },
    {
      "expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
      "id": "10035",
      "self": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
      "key": "KP-23",
      "fields": {
        "customfield_10010": {}
      }
    }
  ]
}`)

	bufferMockedWithNoFields := bytes.Buffer{}
	bufferMockedWithNoFields.WriteString(`
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
      "no_fields": {
        "customfield_10010": {}
      }
    },
    {
      "expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
      "id": "10035",
      "self": "https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
      "key": "KP-23",
      "no_fields": {
        "customfield_10010": {}
      }
    }
  ]
}`)

	bufferMockedWithInvalidTypes := bytes.Buffer{}
	bufferMockedWithInvalidTypes.WriteString(`
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
            "customfield_10010": 33030303,
         }
      },
      {
         "expand":"operations,versionedRepresentations,editmeta,changelog,renderedFields",
         "id":"10035",
         "self":"https://ctreminiom.atlassian.net/rest/api/2/issue/10035",
         "key":"KP-23",
         "fields":{
            "customfield_10010": true
         }
      }
   ]
}
`)

	type args struct {
		buffer      bytes.Buffer
		customField string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*CustomFieldRequestTypeScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the buffer contains information",
			args: args{
				buffer:      bufferMocked,
				customField: "customfield_10010",
			},
			want: map[string]*CustomFieldRequestTypeScheme{
				"KP-22": {
					Links: &CustomFieldRequestTypeLinkScheme{
						JiraRest: "https://mydomain.atlassian.net/rest/api/2/issue/144906",
						Web:      "https://mydomain.atlassian.net/servicedesk/customer/portal/2/ESD-40928",
						Self:     "https://mydomain.atlassian.net/rest/servicedeskapi/request/144906",
						Agent:    "https://mydomain.atlassian.net/browse/ESD-40928",
					},
					RequestType: &CustomerRequestTypeScheme{
						ID:            "96",
						Name:          "General Service Request",
						Description:   "",
						HelpText:      "",
						IssueTypeID:   "10039",
						ServiceDeskID: "2",
						GroupIds:      []string{},
					},
					CurrentStatus: &CustomerRequestCurrentStatusScheme{
						Status:         "Check billing account",
						StatusCategory: "NEW",
						StatusDate: &CustomerRequestCurrentStatusDateScheme{
							Jira:        "2024-06-28T08:33:57.313+0200",
							Friendly:    "Friday 15:33",
							EpochMillis: 1719556437313,
						},
					},
				},
				"KP-23": {
					Links: &CustomFieldRequestTypeLinkScheme{
						JiraRest: "https://mydomain.atlassian.net/rest/api/2/issue/144906",
						Web:      "https://mydomain.atlassian.net/servicedesk/customer/portal/2/ESD-40928",
						Self:     "https://mydomain.atlassian.net/rest/servicedeskapi/request/144906",
						Agent:    "https://mydomain.atlassian.net/browse/ESD-40928",
					},
					RequestType: &CustomerRequestTypeScheme{
						ID:            "96",
						Name:          "General Service Request",
						Description:   "",
						HelpText:      "",
						IssueTypeID:   "10039",
						ServiceDeskID: "2",
						GroupIds:      []string{},
					},
					CurrentStatus: &CustomerRequestCurrentStatusScheme{
						Status:         "Check billing account",
						StatusCategory: "NEW",
						StatusDate: &CustomerRequestCurrentStatusDateScheme{
							Jira:        "2024-06-28T08:33:57.313+0200",
							Friendly:    "Friday 15:33",
							EpochMillis: 1719556437313,
						},
					},
				},
			},

			wantErr: false,
		},
		{
			name: "when the buffer does not contain the issues object",
			args: args{
				buffer:      bufferMockedWithNoIssues,
				customField: "customfield_10010",
			},
			wantErr: true,
			Err:     ErrNoIssuesSliceError,
		},
		{
			name: "when the buffer does not contain the fields object",
			args: args{
				buffer:      bufferMockedWithNoFields,
				customField: "customfield_10010",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
		{
			name: "when the buffer contains invalid types",
			args: args{
				buffer:      bufferMockedWithInvalidTypes,
				customField: "customfield_10010",
			},
			wantErr: true,
			Err:     ErrNoMapValuesError,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseRequestTypeCustomFields(testCase.args.buffer, testCase.args.customField)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseRequestTypeCustomFields() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ParseRequestTypeCustomFields() got = %v, want %v", got, testCase.want)
			}
			if !reflect.DeepEqual(err, testCase.Err) {
				t.Errorf("ParseRequestTypeCustomFields() got = (%v), want (%v)", err, testCase.Err)
			}
		})
	}
}
