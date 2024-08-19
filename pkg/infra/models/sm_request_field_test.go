package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomerRequestPayloadScheme_DateTimeCustomField(t *testing.T) {

	type fields struct {
		Channel             string
		Form                *CreateCustomerRequestFormPayloadScheme
		IsAdfRequest        bool
		RaiseOnBehalfOf     string
		RequestFieldValues  map[string]interface{}
		RequestParticipants []string
		RequestTypeID       string
		ServiceDeskID       string
	}

	type args struct {
		id    string
		value time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
		want    *CreateCustomerRequestPayloadScheme
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				id:    "customfield_10021",
				value: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
			want: &CreateCustomerRequestPayloadScheme{
				RequestFieldValues: map[string]interface{}{"customfield_10021": "2019-01-01T00:00:00Z"},
			},
		},

		{
			name:   "when the customfield id is not provided",
			fields: fields{},
			args: args{
				id:    "",
				value: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: true,
			Err:     ErrNoCustomFieldID,
		},

		{
			name:   "when the date value is not provided",
			fields: fields{},
			args: args{
				id:    "customfield_10021",
				value: time.Time{},
			},
			wantErr: true,
			Err:     ErrNoDatePickerType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &CreateCustomerRequestPayloadScheme{
				Channel:             tt.fields.Channel,
				Form:                tt.fields.Form,
				IsAdfRequest:        tt.fields.IsAdfRequest,
				RaiseOnBehalfOf:     tt.fields.RaiseOnBehalfOf,
				RequestFieldValues:  tt.fields.RequestFieldValues,
				RequestParticipants: tt.fields.RequestParticipants,
				RequestTypeID:       tt.fields.RequestTypeID,
				ServiceDeskID:       tt.fields.ServiceDeskID,
			}

			err := c.DateTimeCustomField(tt.args.id, tt.args.value)

			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, tt.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, c)
			}
		})
	}
}

func TestCreateCustomerRequestPayloadScheme_DateCustomField(t *testing.T) {

	type fields struct {
		Channel             string
		Form                *CreateCustomerRequestFormPayloadScheme
		IsAdfRequest        bool
		RaiseOnBehalfOf     string
		RequestFieldValues  map[string]interface{}
		RequestParticipants []string
		RequestTypeID       string
		ServiceDeskID       string
	}

	type args struct {
		id    string
		value time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
		want    *CreateCustomerRequestPayloadScheme
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				id:    "customfield_10021",
				value: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
			want: &CreateCustomerRequestPayloadScheme{
				RequestFieldValues: map[string]interface{}{"customfield_10021": "2019-01-01"},
			},
		},

		{
			name:   "when the customfield id is not provided",
			fields: fields{},
			args: args{
				id:    "",
				value: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: true,
			Err:     ErrNoCustomFieldID,
		},

		{
			name:   "when the date value is not provided",
			fields: fields{},
			args: args{
				id:    "customfield_10021",
				value: time.Time{},
			},
			wantErr: true,
			Err:     ErrNoDatePickerType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &CreateCustomerRequestPayloadScheme{
				Channel:             tt.fields.Channel,
				Form:                tt.fields.Form,
				IsAdfRequest:        tt.fields.IsAdfRequest,
				RaiseOnBehalfOf:     tt.fields.RaiseOnBehalfOf,
				RequestFieldValues:  tt.fields.RequestFieldValues,
				RequestParticipants: tt.fields.RequestParticipants,
				RequestTypeID:       tt.fields.RequestTypeID,
				ServiceDeskID:       tt.fields.ServiceDeskID,
			}

			err := c.DateCustomField(tt.args.id, tt.args.value)

			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, tt.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, c)
			}
		})
	}
}

func TestCreateCustomerRequestPayloadScheme_MultiSelectOrCheckBoxCustomField(t *testing.T) {

	type fields struct {
		Channel             string
		Form                *CreateCustomerRequestFormPayloadScheme
		IsAdfRequest        bool
		RaiseOnBehalfOf     string
		RequestFieldValues  map[string]interface{}
		RequestParticipants []string
		RequestTypeID       string
		ServiceDeskID       string
	}

	type args struct {
		id      string
		options []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
		want    *CreateCustomerRequestPayloadScheme
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				id:      "customfield_10021",
				options: []string{"option 01", "option 02"},
			},
			wantErr: false,
			want: &CreateCustomerRequestPayloadScheme{
				RequestFieldValues: map[string]interface{}{"customfield_10021": []map[string]interface{}{map[string]interface{}{"value": "option 01"}, map[string]interface{}{"value": "option 02"}}},
			},
		},

		{
			name:   "when the customfield id is not provided",
			fields: fields{},
			args: args{
				id: "",
			},
			wantErr: true,
			Err:     ErrNoCustomFieldID,
		},

		{
			name:   "when the option name is not provided",
			fields: fields{},
			args: args{
				id: "customfield_10021",
			},
			wantErr: true,
			Err:     ErrNoMultiSelectType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &CreateCustomerRequestPayloadScheme{
				Channel:             tt.fields.Channel,
				Form:                tt.fields.Form,
				IsAdfRequest:        tt.fields.IsAdfRequest,
				RaiseOnBehalfOf:     tt.fields.RaiseOnBehalfOf,
				RequestFieldValues:  tt.fields.RequestFieldValues,
				RequestParticipants: tt.fields.RequestParticipants,
				RequestTypeID:       tt.fields.RequestTypeID,
				ServiceDeskID:       tt.fields.ServiceDeskID,
			}

			err := c.MultiSelectOrCheckBoxCustomField(tt.args.id, tt.args.options)

			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, tt.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, c)
			}
		})
	}
}

func TestCreateCustomerRequestPayloadScheme_UserCustomField(t *testing.T) {

	type fields struct {
		Channel             string
		Form                *CreateCustomerRequestFormPayloadScheme
		IsAdfRequest        bool
		RaiseOnBehalfOf     string
		RequestFieldValues  map[string]interface{}
		RequestParticipants []string
		RequestTypeID       string
		ServiceDeskID       string
	}

	type args struct {
		id        string
		accountID string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
		want    *CreateCustomerRequestPayloadScheme
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				id:        "customfield_10021",
				accountID: "uuid-sample",
			},
			wantErr: false,
			want: &CreateCustomerRequestPayloadScheme{
				RequestFieldValues: map[string]interface{}{"customfield_10021": map[string]interface{}{"accountId": "uuid-sample"}},
			},
		},

		{
			name:   "when the customfield id is not provided",
			fields: fields{},
			args: args{
				id: "",
			},
			wantErr: true,
			Err:     ErrNoCustomFieldID,
		},

		{
			name:   "when the date value is not provided",
			fields: fields{},
			args: args{
				id: "customfield_10021",
			},
			wantErr: true,
			Err:     ErrNoUserType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &CreateCustomerRequestPayloadScheme{
				Channel:             tt.fields.Channel,
				Form:                tt.fields.Form,
				IsAdfRequest:        tt.fields.IsAdfRequest,
				RaiseOnBehalfOf:     tt.fields.RaiseOnBehalfOf,
				RequestFieldValues:  tt.fields.RequestFieldValues,
				RequestParticipants: tt.fields.RequestParticipants,
				RequestTypeID:       tt.fields.RequestTypeID,
				ServiceDeskID:       tt.fields.ServiceDeskID,
			}

			err := c.UserCustomField(tt.args.id, tt.args.accountID)

			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, tt.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, c)
			}
		})
	}
}

func TestCreateCustomerRequestPayloadScheme_UsersCustomField(t *testing.T) {

	type fields struct {
		Channel             string
		Form                *CreateCustomerRequestFormPayloadScheme
		IsAdfRequest        bool
		RaiseOnBehalfOf     string
		RequestFieldValues  map[string]interface{}
		RequestParticipants []string
		RequestTypeID       string
		ServiceDeskID       string
	}

	type args struct {
		id      string
		options []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
		want    *CreateCustomerRequestPayloadScheme
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				id:      "customfield_10021",
				options: []string{"uuid-sample", "uuid-sample-1"},
			},
			wantErr: false,
			want: &CreateCustomerRequestPayloadScheme{
				RequestFieldValues: map[string]interface{}{"customfield_10021": []map[string]interface{}{map[string]interface{}{"accountId": "uuid-sample"}, map[string]interface{}{"accountId": "uuid-sample-1"}}},
			},
		},

		{
			name:   "when the customfield id is not provided",
			fields: fields{},
			args: args{
				id: "",
			},
			wantErr: true,
			Err:     ErrNoCustomFieldID,
		},

		{
			name:   "when the date value is not provided",
			fields: fields{},
			args: args{
				id: "customfield_10021",
			},
			wantErr: true,
			Err:     ErrNoMultiUserType,
		},

		{
			name:   "when the users slice contains an empty value",
			fields: fields{},
			args: args{
				id:      "customfield_10021",
				options: []string{"uuid-sample", ""},
			},
			wantErr: true,
			Err:     ErrNoUserType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &CreateCustomerRequestPayloadScheme{
				Channel:             tt.fields.Channel,
				Form:                tt.fields.Form,
				IsAdfRequest:        tt.fields.IsAdfRequest,
				RaiseOnBehalfOf:     tt.fields.RaiseOnBehalfOf,
				RequestFieldValues:  tt.fields.RequestFieldValues,
				RequestParticipants: tt.fields.RequestParticipants,
				RequestTypeID:       tt.fields.RequestTypeID,
				ServiceDeskID:       tt.fields.ServiceDeskID,
			}

			err := c.UsersCustomField(tt.args.id, tt.args.options)

			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, tt.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, c)
			}
		})
	}
}

func TestCreateCustomerRequestPayloadScheme_CascadingCustomField(t *testing.T) {

	type fields struct {
		Channel             string
		Form                *CreateCustomerRequestFormPayloadScheme
		IsAdfRequest        bool
		RaiseOnBehalfOf     string
		RequestFieldValues  map[string]interface{}
		RequestParticipants []string
		RequestTypeID       string
		ServiceDeskID       string
	}

	type args struct {
		id, parent, child string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
		want    *CreateCustomerRequestPayloadScheme
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				id:     "customfield_10021",
				parent: "America",
				child:  "Costa Rica",
			},
			wantErr: false,
			want: &CreateCustomerRequestPayloadScheme{
				RequestFieldValues: map[string]interface{}{"customfield_10021": map[string]interface{}{"child": map[string]interface{}{"value": "Costa Rica"}, "value": "America"}},
			},
		},

		{
			name:   "when the customfield id is not provided",
			fields: fields{},
			args: args{
				id: "",
			},
			wantErr: true,
			Err:     ErrNoCustomFieldID,
		},

		{
			name:   "when the parent value is not provided",
			fields: fields{},
			args: args{
				id: "customfield_10021",
			},
			wantErr: true,
			Err:     ErrNoCascadingParent,
		},

		{
			name:   "when the child value is not provided",
			fields: fields{},
			args: args{
				id:     "customfield_10021",
				parent: "America",
			},
			wantErr: true,
			Err:     ErrNoCascadingChild,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &CreateCustomerRequestPayloadScheme{
				Channel:             tt.fields.Channel,
				Form:                tt.fields.Form,
				IsAdfRequest:        tt.fields.IsAdfRequest,
				RaiseOnBehalfOf:     tt.fields.RaiseOnBehalfOf,
				RequestFieldValues:  tt.fields.RequestFieldValues,
				RequestParticipants: tt.fields.RequestParticipants,
				RequestTypeID:       tt.fields.RequestTypeID,
				ServiceDeskID:       tt.fields.ServiceDeskID,
			}

			err := c.CascadingCustomField(tt.args.id, tt.args.parent, tt.args.child)

			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, tt.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, c)
			}
		})
	}
}

func TestCreateCustomerRequestPayloadScheme_GroupsCustomField(t *testing.T) {

	type fields struct {
		Channel             string
		Form                *CreateCustomerRequestFormPayloadScheme
		IsAdfRequest        bool
		RaiseOnBehalfOf     string
		RequestFieldValues  map[string]interface{}
		RequestParticipants []string
		RequestTypeID       string
		ServiceDeskID       string
	}

	type args struct {
		id      string
		options []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
		want    *CreateCustomerRequestPayloadScheme
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				id:      "customfield_10021",
				options: []string{"group-name-01", "group-name-02"},
			},
			wantErr: false,
			want: &CreateCustomerRequestPayloadScheme{
				RequestFieldValues: map[string]interface{}{"customfield_10021": []map[string]interface{}{map[string]interface{}{"name": "group-name-01"}, map[string]interface{}{"name": "group-name-02"}}},
			},
		},

		{
			name:   "when the customfield id is not provided",
			fields: fields{},
			args: args{
				id: "",
			},
			wantErr: true,
			Err:     ErrNoCustomFieldID,
		},

		{
			name:   "when the groups value is not provided",
			fields: fields{},
			args: args{
				id: "customfield_10021",
			},
			wantErr: true,
			Err:     ErrNoGroupsName,
		},

		{
			name:   "when the groups slice contains an empty element",
			fields: fields{},
			args: args{
				id:      "customfield_10021",
				options: []string{"group-name-01", ""},
			},
			wantErr: true,
			Err:     ErrNoGroupName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &CreateCustomerRequestPayloadScheme{
				Channel:             tt.fields.Channel,
				Form:                tt.fields.Form,
				IsAdfRequest:        tt.fields.IsAdfRequest,
				RaiseOnBehalfOf:     tt.fields.RaiseOnBehalfOf,
				RequestFieldValues:  tt.fields.RequestFieldValues,
				RequestParticipants: tt.fields.RequestParticipants,
				RequestTypeID:       tt.fields.RequestTypeID,
				ServiceDeskID:       tt.fields.ServiceDeskID,
			}

			err := c.GroupsCustomField(tt.args.id, tt.args.options)

			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, tt.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, c)
			}
		})
	}
}

func TestCreateCustomerRequestPayloadScheme_GroupCustomField(t *testing.T) {

	type fields struct {
		Channel             string
		Form                *CreateCustomerRequestFormPayloadScheme
		IsAdfRequest        bool
		RaiseOnBehalfOf     string
		RequestFieldValues  map[string]interface{}
		RequestParticipants []string
		RequestTypeID       string
		ServiceDeskID       string
	}

	type args struct {
		id   string
		name string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
		want    *CreateCustomerRequestPayloadScheme
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				id:   "customfield_10021",
				name: "jira-admins",
			},
			wantErr: false,
			want: &CreateCustomerRequestPayloadScheme{
				RequestFieldValues: map[string]interface{}{"customfield_10021": map[string]interface{}{"name": "jira-admins"}},
			},
		},

		{
			name:   "when the customfield id is not provided",
			fields: fields{},
			args: args{
				id: "",
			},
			wantErr: true,
			Err:     ErrNoCustomFieldID,
		},

		{
			name:   "when the group name is not provided",
			fields: fields{},
			args: args{
				id: "customfield_10021",
			},
			wantErr: true,
			Err:     ErrNoGroupName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &CreateCustomerRequestPayloadScheme{
				Channel:             tt.fields.Channel,
				Form:                tt.fields.Form,
				IsAdfRequest:        tt.fields.IsAdfRequest,
				RaiseOnBehalfOf:     tt.fields.RaiseOnBehalfOf,
				RequestFieldValues:  tt.fields.RequestFieldValues,
				RequestParticipants: tt.fields.RequestParticipants,
				RequestTypeID:       tt.fields.RequestTypeID,
				ServiceDeskID:       tt.fields.ServiceDeskID,
			}

			err := c.GroupCustomField(tt.args.id, tt.args.name)

			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, tt.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, c)
			}
		})
	}
}

func TestCreateCustomerRequestPayloadScheme_RadioButtonOrSelectCustomField(t *testing.T) {

	type fields struct {
		Channel             string
		Form                *CreateCustomerRequestFormPayloadScheme
		IsAdfRequest        bool
		RaiseOnBehalfOf     string
		RequestFieldValues  map[string]interface{}
		RequestParticipants []string
		RequestTypeID       string
		ServiceDeskID       string
	}

	type args struct {
		id     string
		option string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
		want    *CreateCustomerRequestPayloadScheme
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				id:     "customfield_10021",
				option: "Option 01",
			},
			wantErr: false,
			want: &CreateCustomerRequestPayloadScheme{
				RequestFieldValues: map[string]interface{}{"customfield_10021": map[string]interface{}{"value": "Option 01"}},
			},
		},

		{
			name:   "when the customfield id is not provided",
			fields: fields{},
			args: args{
				id: "",
			},
			wantErr: true,
			Err:     ErrNoCustomFieldID,
		},

		{
			name:   "when the option name is not provided",
			fields: fields{},
			args: args{
				id: "customfield_10021",
			},
			wantErr: true,
			Err:     ErrNoSelectType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &CreateCustomerRequestPayloadScheme{
				Channel:             tt.fields.Channel,
				Form:                tt.fields.Form,
				IsAdfRequest:        tt.fields.IsAdfRequest,
				RaiseOnBehalfOf:     tt.fields.RaiseOnBehalfOf,
				RequestFieldValues:  tt.fields.RequestFieldValues,
				RequestParticipants: tt.fields.RequestParticipants,
				RequestTypeID:       tt.fields.RequestTypeID,
				ServiceDeskID:       tt.fields.ServiceDeskID,
			}

			err := c.RadioButtonOrSelectCustomField(tt.args.id, tt.args.option)

			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, tt.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, c)
			}
		})
	}
}

func TestCreateCustomerRequestPayloadScheme_Components(t *testing.T) {

	type fields struct {
		Channel             string
		Form                *CreateCustomerRequestFormPayloadScheme
		IsAdfRequest        bool
		RaiseOnBehalfOf     string
		RequestFieldValues  map[string]interface{}
		RequestParticipants []string
		RequestTypeID       string
		ServiceDeskID       string
	}

	type args struct {
		components []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		Err     error
		want    *CreateCustomerRequestPayloadScheme
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				components: []string{"Jira Cloud", "Confluence Cloud"},
			},
			wantErr: false,
			want: &CreateCustomerRequestPayloadScheme{
				RequestFieldValues: map[string]interface{}{"components": []map[string]interface{}{map[string]interface{}{"name": "Jira Cloud"}, map[string]interface{}{"name": "Confluence Cloud"}}},
			},
		},

		{
			name:   "when the components are not provided",
			fields: fields{},
			args: args{
				components: nil,
			},
			wantErr: true,
			Err:     ErrNoComponents,
		},

		{
			name:   "when the one component on the slice is empty",
			fields: fields{},
			args: args{
				components: []string{"Jira Cloud", ""},
			},
			wantErr: true,
			Err:     ErrNCoComponent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &CreateCustomerRequestPayloadScheme{
				Channel:             tt.fields.Channel,
				Form:                tt.fields.Form,
				IsAdfRequest:        tt.fields.IsAdfRequest,
				RaiseOnBehalfOf:     tt.fields.RaiseOnBehalfOf,
				RequestFieldValues:  tt.fields.RequestFieldValues,
				RequestParticipants: tt.fields.RequestParticipants,
				RequestTypeID:       tt.fields.RequestTypeID,
				ServiceDeskID:       tt.fields.ServiceDeskID,
			}

			err := c.Components(tt.args.components)

			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, tt.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, c)
			}
		})
	}
}
