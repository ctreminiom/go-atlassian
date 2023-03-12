package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCustomerRequestFields_Attachments(t *testing.T) {

	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		attachments []string
	}
	testCases := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				attachments: []string{"uuid-sample"},
			},
			wantErr: false,
		},

		{
			name:   "when the attachments are not provided",
			fields: fields{},
			args: args{
				attachments: nil,
			},
			wantErr: true,
			err:     ErrNoAttachmentIdsError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Attachments(testCase.args.attachments)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_Labels(t *testing.T) {

	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		labels []string
	}
	testCases := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				labels: []string{"label-sample"},
			},
			wantErr: false,
		},

		{
			name:   "when the labels are not provided",
			fields: fields{},
			args: args{
				labels: nil,
			},
			wantErr: true,
			err:     ErrNoLabelsError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Labels(testCase.args.labels)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_Components(t *testing.T) {

	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		components []string
	}
	testCases := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				components: []string{"component-sample"},
			},
			wantErr: false,
		},

		{
			name:   "when the components are not provided",
			fields: fields{},
			args: args{
				components: nil,
			},
			wantErr: true,
			err:     ErrNoComponentsError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Components(testCase.args.components)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_Groups(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				groups:        []string{"group-sample"},
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				groups:        []string{"group-sample"},
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the groups are not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				groups:        nil,
			},
			wantErr: true,
			err:     ErrNoGroupsNameError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Groups(testCase.args.customFieldID, testCase.args.groups)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_Group(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				group:         "group-sample",
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				group:         "group-sample",
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the group is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				group:         "",
			},
			wantErr: true,
			err:     ErrNoGroupNameError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Group(testCase.args.customFieldID, testCase.args.group)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_URL(t *testing.T) {

	type fields struct {
		Fields []map[string]interface{}
	}
	type args struct {
		customFieldID string
		url           string
	}
	testCases := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				url:           "url-sample",
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				url:           "url-sample",
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the url is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				url:           "",
			},
			wantErr: true,
			err:     ErrNoUrlTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.URL(testCase.args.customFieldID, testCase.args.url)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_Text(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				textValue:     "text-sample",
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				textValue:     "url-sample",
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the text is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				textValue:     "",
			},
			wantErr: true,
			err:     ErrNoTextTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Text(testCase.args.customFieldID, testCase.args.textValue)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_DateTime(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				dateValue:     time.Now().AddDate(0, -1, 0),
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				dateValue:     time.Now().AddDate(0, -1, 0),
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the date-time is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
			},
			wantErr: true,
			err:     ErrNoDateTimeTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.DateTime(testCase.args.customFieldID, testCase.args.dateValue)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_Date(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				dateValue:     time.Now().AddDate(0, -1, 0),
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				dateValue:     time.Now().AddDate(0, -1, 0),
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the date-time is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
			},
			wantErr: true,
			err:     ErrNoDateTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Date(testCase.args.customFieldID, testCase.args.dateValue)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_MultiSelect(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				options:       []string{"option-1"},
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				options:       []string{"option-1"},
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the groups are not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				options:       nil,
			},
			wantErr: true,
			err:     ErrNoMultiSelectTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.MultiSelect(testCase.args.customFieldID, testCase.args.options)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_Select(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				option:        "option-sample",
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				option:        "option-sample",
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the group is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				option:        "",
			},
			wantErr: true,
			err:     ErrNoSelectTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Select(testCase.args.customFieldID, testCase.args.option)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_RadioButton(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				option:        "option-sample",
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				option:        "option-sample",
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the group is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				option:        "",
			},
			wantErr: true,
			err:     ErrNoButtonTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.RadioButton(testCase.args.customFieldID, testCase.args.option)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_User(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				accountID:     "account-id-sample-uuid",
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				accountID:     "account-id-sample-uuid",
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the group is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				accountID:     "",
			},
			wantErr: true,
			err:     ErrNoUserTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.User(testCase.args.customFieldID, testCase.args.accountID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_Users(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				accountIDs:    []string{"account-id-sample-uuid"},
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				accountIDs:    []string{"account-id-sample-uuid"},
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the group is not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				accountIDs:    nil,
			},
			wantErr: true,
			err:     ErrNoMultiUserTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Users(testCase.args.customFieldID, testCase.args.accountIDs)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_Number(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				numberValue:   10000.232,
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				numberValue:   10000.232,
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.Number(testCase.args.customFieldID, testCase.args.numberValue)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_CheckBox(t *testing.T) {

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
		err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				options:       []string{"option-1"},
			},
			wantErr: false,
		},

		{
			name:   "when the customfield is not provided",
			fields: fields{},
			args: args{
				customFieldID: "",
				options:       []string{"option-1"},
			},
			wantErr: true,
			err:     ErrNoCustomFieldIDError,
		},

		{
			name:   "when the groups are not provided",
			fields: fields{},
			args: args{
				customFieldID: "customfield_10000",
				options:       nil,
			},
			wantErr: true,
			err:     ErrNoCheckBoxTypeError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &CustomerRequestFields{
				Fields: testCase.fields.Fields,
			}

			err := c.CheckBox(testCase.args.customFieldID, testCase.args.options)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.err.Error())

			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestCustomerRequestFields_Cascading(t *testing.T) {
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
			Err:     ErrNoCustomFieldIDError,
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

			c := &CustomerRequestFields{
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
