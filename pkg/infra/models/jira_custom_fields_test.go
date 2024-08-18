package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomFields_Cascading(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		parent        string
		child         string
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
				customFieldID: "customfield_10001",
				parent:        "America",
				child:         "US",
			},
			wantErr: false,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				parent:        "America",
				child:         "US",
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the parent value is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10001",
				parent:        "",
				child:         "US",
			},
			wantErr: true,
			Err:     ErrNoCascadingParentError,
		},

		{
			name:   "when the child value is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10001",
				parent:        "America",
				child:         "",
			},
			wantErr: true,
			Err:     ErrNoCascadingChildError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Cascading(testCase.args.customFieldID, testCase.args.parent, testCase.args.child)
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

func TestCustomFields_CheckBox(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		options       []string
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
				customFieldID: "customfield_10001",
				options:       []string{"Value"},
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				options:       []string{"Value"},
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the options are not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10001",
				options:       nil,
			},
			wantErr: true,
			Err:     ErrNoCheckBoxTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.CheckBox(testCase.args.customFieldID, testCase.args.options)
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

func TestCustomFields_Date(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		dateTimeValue time.Time
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
				customFieldID: "customfield_10001",
				dateTimeValue: time.Now().AddDate(0, -1, 0),
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				dateTimeValue: time.Now().AddDate(0, -1, 0),
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the date is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10001",
				dateTimeValue: time.Time{},
			},
			wantErr: true,
			Err:     ErrNoDateTimeTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Date(testCase.args.customFieldID, testCase.args.dateTimeValue)

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

func TestCustomFields_DateTime(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		dateValue     time.Time
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
				customFieldID: "customfield_10001",
				dateValue:     time.Now().AddDate(0, -1, 0),
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				dateValue:     time.Now().AddDate(0, -1, 0),
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the date-time is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10001",
				dateValue:     time.Time{},
			},
			wantErr: true,
			Err:     ErrNoDatePickerTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.DateTime(testCase.args.customFieldID, testCase.args.dateValue)

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

func TestCustomFields_Group(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		group         string
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
				customFieldID: "customfield_10001",
				group:         "jira-users",
			},
			wantErr: false,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				group:         "jira-users",
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the group name is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10001",
				group:         "",
			},
			wantErr: true,
			Err:     ErrNoGroupNameError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Group(testCase.args.customFieldID, testCase.args.group)
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

func TestCustomFields_Groups(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		groups        []string
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
				customFieldID: "customfield_1000",
				groups:        []string{"jira-users", "jira-admins"},
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				groups:        []string{"jira-users", "jira-admins"},
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the groups names are not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_1000",
				groups:        nil,
			},
			wantErr: true,
			Err:     ErrNoGroupsNameError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Groups(testCase.args.customFieldID, testCase.args.groups)
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

func TestCustomFields_MultiSelect(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		options       []string
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
				customFieldID: "customfield_1000",
				options:       []string{"options"},
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				options:       []string{"options"},
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the options are not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_1000",
				options:       nil,
			},
			wantErr: true,
			Err:     ErrNoMultiSelectTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.MultiSelect(testCase.args.customFieldID, testCase.args.options)
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

func TestCustomFields_Number(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		numberValue   float64
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
				customFieldID: "customfield_1000",
				numberValue:   0,
			},
			wantErr: false,
			Err:     nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Number(testCase.args.customFieldID, testCase.args.numberValue)
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

func TestCustomFields_RadioButton(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		button        string
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
				customFieldID: "customfield_1000",
				button:        "Button 1 ",
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				button:        "Button 1 ",
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the option is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_1000",
				button:        "",
			},
			wantErr: true,
			Err:     ErrNoButtonTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.RadioButton(testCase.args.customFieldID, testCase.args.button)
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

func TestCustomFields_Select(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		option        string
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
				customFieldID: "customfield_1000",
				option:        "Option 1",
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				option:        "Option 1",
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the option is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_1000",
				option:        "",
			},
			wantErr: true,
			Err:     ErrNoSelectTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Select(testCase.args.customFieldID, testCase.args.option)
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

func TestCustomFields_Text(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		textValue     string
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
				customFieldID: "customfield_1000",
				textValue:     "Application",
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				textValue:     "Application",
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the value is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_1000",
				textValue:     "",
			},
			wantErr: true,
			Err:     ErrNoTextTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Text(testCase.args.customFieldID, testCase.args.textValue)
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

func TestCustomFields_URL(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		URL           string
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
				customFieldID: "customfield_1000",
				URL:           "https://www.google.com/",
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				URL:           "https://www.google.com/",
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the url is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_1000",
				URL:           "",
			},
			wantErr: true,
			Err:     ErrNoURLTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.URL(testCase.args.customFieldID, testCase.args.URL)
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

func TestCustomFields_User(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		accountID     string
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
				customFieldID: "customfield_1000",
				accountID:     "uuid-sample",
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				accountID:     "uuid-sample",
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},

		{
			name:   "when the user is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_1000",
				accountID:     "",
			},
			wantErr: true,
			Err:     ErrNoUserTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.User(testCase.args.customFieldID, testCase.args.accountID)
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

func TestCustomFields_Users(t *testing.T) {
	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		accountIDs    []string
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
				customFieldID: "customfield_1000",
				accountIDs:    []string{"user-1", "user-2"},
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the custom-field is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				accountIDs:    []string{"user-1", "user-2"},
			},
			wantErr: true,
			Err:     ErrNoFieldIDError,
		},
		{
			name:   "when the users are not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_1000",
				accountIDs:    nil,
			},
			wantErr: true,
			Err:     ErrNoMultiUserTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Users(testCase.args.customFieldID, testCase.args.accountIDs)
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
