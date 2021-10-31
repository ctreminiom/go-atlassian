package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestCustomFields_Cascading(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		parent, child string
		wantErr       bool
	}{
		{
			name:          "CreateCascadingCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			parent:        "America",
			child:         "Costa Rica",
			wantErr:       false,
		},

		{
			name:          "CreateCascadingCustomFieldWhenTheCustomFieldIsNotSet",
			customFieldID: "",
			parent:        "America",
			child:         "Costa Rica",
			wantErr:       true,
		},

		{
			name:          "CreateCascadingCustomFieldWhenTheParentIsNotSet",
			customFieldID: "customfield_000",
			parent:        "",
			child:         "Costa Rica",
			wantErr:       true,
		},

		{
			name:          "CreateCascadingCustomFieldWhenTheChildIsNotSet",
			customFieldID: "customfield_000",
			parent:        "America",
			child:         "",
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.Cascading(testCase.customFieldID, testCase.parent, testCase.child)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestCustomFields_CheckBox(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		options       []string
		wantErr       bool
	}{
		{
			name:          "CreateCheckBoxCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			options:       []string{"Option 1", "Option 2"},
			wantErr:       false,
		},

		{
			name:          "CreateCheckBoxCustomFieldWhenCustomFieldIDIsNotSet",
			customFieldID: "",
			options:       []string{"Option 1", "Option 2"},
			wantErr:       true,
		},

		{
			name:          "CreateCheckBoxCustomFieldWhenTheOptionsAreNotSet",
			customFieldID: "customfield_000",
			options:       nil,
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.CheckBox(testCase.customFieldID, testCase.options)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestCustomFields_Date(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		dateValue     time.Time
		wantErr       bool
	}{
		{
			name:          "CreateDateCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			dateValue:     time.Now(),
			wantErr:       false,
		},

		{
			name:          "CreateDateCustomFieldWhenTheDateValueHasACustomDateValue",
			customFieldID: "customfield_000",
			dateValue:     time.Now().AddDate(0, -2, 0),
			wantErr:       false,
		},

		{
			name:          "CreateDateCustomFieldWhenTheCustomFieldIsNotSet",
			customFieldID: "",
			dateValue:     time.Now(),
			wantErr:       true,
		},

		{
			name:          "CreateDateCustomFieldWhenTheDataValueIsNotSet",
			customFieldID: "customfield_000",
			dateValue:     time.Time{},
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.Date(testCase.customFieldID, testCase.dateValue)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestCustomFields_DateTime(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		dateValue     time.Time
		wantErr       bool
	}{
		{
			name:          "CreateDateTimeCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			dateValue:     time.Now(),
			wantErr:       false,
		},

		{
			name:          "CreateDateTimeCustomFieldWhenTheDateValueHasACustomDateValue",
			customFieldID: "customfield_000",
			dateValue:     time.Now().AddDate(0, -2, 0),
			wantErr:       false,
		},

		{
			name:          "CreateDateTimeCustomFieldWhenTheCustomFieldIsNotSet",
			customFieldID: "",
			dateValue:     time.Now(),
			wantErr:       true,
		},

		{
			name:          "CreateDateTimeCustomFieldWhenTheDataValueIsNotSet",
			customFieldID: "customfield_000",
			dateValue:     time.Time{},
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.DateTime(testCase.customFieldID, testCase.dateValue)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestCustomFields_Group(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		group         string
		wantErr       bool
	}{
		{
			name:          "CreateGroupCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			group:         "jira-system-admins",
			wantErr:       false,
		},

		{
			name:          "CreateGroupCustomFieldWhenTheCustomFieldIDIsNotSet",
			customFieldID: "",
			group:         "jira-system-admins",
			wantErr:       true,
		},

		{
			name:          "CreateGroupCustomFieldWhenTheGroupIsNotSet",
			customFieldID: "customfield_000",
			group:         "",
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.Group(testCase.customFieldID, testCase.group)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestCustomFields_Groups(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		groups        []string
		wantErr       bool
	}{
		{
			name:          "CreateGroupsCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			groups:        []string{"jira-system-admin", "confluence-system-admin"},
			wantErr:       false,
		},

		{
			name:          "CreateGroupsCustomFieldWhenCustomFieldIDIsNotSet",
			customFieldID: "",
			groups:        []string{"jira-system-admin", "confluence-system-admin"},
			wantErr:       true,
		},

		{
			name:          "CreateGroupsCustomFieldWhenTheOptionsAreNotSet",
			customFieldID: "customfield_000",
			groups:        nil,
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.Groups(testCase.customFieldID, testCase.groups)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestCustomFields_MultiSelect(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		options       []string
		wantErr       bool
	}{
		{
			name:          "CreateMultiSelectCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			options:       []string{"Option 1", "Option 2"},
			wantErr:       false,
		},

		{
			name:          "CreateMultiSelectCustomFieldWhenCustomFieldIDIsNotSet",
			customFieldID: "",
			options:       []string{"Option 1", "Option 2"},
			wantErr:       true,
		},

		{
			name:          "CreateMultiSelectCustomFieldWhenTheOptionsAreNotSet",
			customFieldID: "customfield_000",
			options:       nil,
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.MultiSelect(testCase.customFieldID, testCase.options)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestCustomFields_Number(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		numberValue   float64
		wantErr       bool
	}{
		{
			name:          "CreateNumberCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			numberValue:   10000.4333,
			wantErr:       false,
		},

		{
			name:          "CreateNumberCustomFieldWhenTheCustomFieldIsNotSet",
			customFieldID: "",
			numberValue:   10000.4333,
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.Number(testCase.customFieldID, testCase.numberValue)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestCustomFields_RadioButton(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		button        string
		wantErr       bool
	}{
		{
			name:          "CreateRadioButtonCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			button:        "Option 1",
			wantErr:       false,
		},

		{
			name:          "CreateRadioButtonCustomFieldWhenTheCustomFieldIDIsNotSet",
			customFieldID: "",
			button:        "Option 1",
			wantErr:       true,
		},

		{
			name:          "CreateRadioButtonCustomFieldWhenTheOptionIsNotSet",
			customFieldID: "customfield_000",
			button:        "",
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.RadioButton(testCase.customFieldID, testCase.button)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestCustomFields_Select(t *testing.T) {
	testCases := []struct {
		name          string
		customFieldID string
		option        string
		wantErr       bool
	}{
		{
			name:          "CreateSelectCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			option:        "Option 1",
			wantErr:       false,
		},

		{
			name:          "CreateSelectCustomFieldWhenTheCustomFieldIDIsNotSet",
			customFieldID: "",
			option:        "Option 1",
			wantErr:       true,
		},

		{
			name:          "CreateSelectButtonCustomFieldWhenTheOptionIsNotSet",
			customFieldID: "customfield_000",
			option:        "",
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.Select(testCase.customFieldID, testCase.option)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}
}

func TestCustomFields_Text(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		textValue     string
		wantErr       bool
	}{
		{
			name:          "CreateTextCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			textValue:     "Option 1",
			wantErr:       false,
		},

		{
			name:          "CreateTextCustomFieldWhenTheCustomFieldIDIsNotSet",
			customFieldID: "",
			textValue:     "Option 1",
			wantErr:       true,
		},

		{
			name:          "CreateTextButtonCustomFieldWhenTheTextValueIsNotSet",
			customFieldID: "customfield_000",
			textValue:     "",
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.Text(testCase.customFieldID, testCase.textValue)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}
}

func TestCustomFields_URL(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		url           string
		wantErr       bool
	}{
		{
			name:          "CreateURLCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			url:           "https://www.google.com/",
			wantErr:       false,
		},

		{
			name:          "CreateURLCustomFieldWhenTheCustomFieldIDIsNotSet",
			customFieldID: "",
			url:           "https://www.google.com/",
			wantErr:       true,
		},

		{
			name:          "CreateURLButtonCustomFieldWhenTheTextValueIsNotSet",
			customFieldID: "customfield_000",
			url:           "",
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.URL(testCase.customFieldID, testCase.url)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}
}

func TestCustomFields_User(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		accountID     string
		wantErr       bool
	}{
		{
			name:          "CreateUserCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			accountID:     "1310743b-be3b-4281-b45c-869fb7b31c3e",
			wantErr:       false,
		},

		{
			name:          "CreateUserCustomFieldWhenTheCustomFieldIDIsNotSet",
			customFieldID: "",
			accountID:     "1310743b-be3b-4281-b45c-869fb7b31c3e",
			wantErr:       true,
		},

		{
			name:          "CreateUserButtonCustomFieldWhenTheTextValueIsNotSet",
			customFieldID: "customfield_000",
			accountID:     "",
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.User(testCase.customFieldID, testCase.accountID)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestCustomFields_Users(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		accountIDs    []string
		wantErr       bool
	}{
		{
			name:          "CreateUsersCustomFieldWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			accountIDs:    []string{"46e7ea94-9bc9-4c5e-b9b8-a3736eb8c23d", "b0d47b0e-8d42-45fb-a9f7-b80bac42342f"},
			wantErr:       false,
		},

		{
			name:          "CreateUsersCustomFieldWhenCustomFieldIDIsNotSet",
			customFieldID: "",
			accountIDs:    []string{"46e7ea94-9bc9-4c5e-b9b8-a3736eb8c23d", "b0d47b0e-8d42-45fb-a9f7-b80bac42342f"},
			wantErr:       true,
		},

		{
			name:          "CreateUsersCustomFieldWhenTheOptionsAreNotSet",
			customFieldID: "customfield_000",
			accountIDs:    nil,
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var customFields = CustomFields{}
			err := customFields.Users(testCase.customFieldID, testCase.accountIDs)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestIssueScheme_MergeCustomFields(t *testing.T) {

	var customFieldMockedWithFields = CustomFields{}

	// Add a new custom field
	err := customFieldMockedWithFields.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		t.Fatal(err)
	}

	err = customFieldMockedWithFields.Number("customfield_10043", 1000.3232)
	if err != nil {
		t.Fatal(err)
	}

	var customFieldMockedWithOutFields = CustomFields{
		nil,
	}

	testCases := []struct {
		name         string
		fields       *CustomFields
		issuePayload *IssueScheme
		wantErr      bool
	}{
		{
			name: "MergeCustomFieldsWhenTheCustomFieldsAreValue",
			issuePayload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			fields:  &customFieldMockedWithFields,
			wantErr: false,
		},

		{
			name: "MergeCustomFieldsWhenTheCustomFieldsAreEmpty",
			issuePayload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			fields:  &customFieldMockedWithOutFields,
			wantErr: true,
		},

		{
			name: "MergeCustomFieldsWhenTheCustomFieldsIsNil",
			issuePayload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			fields:  nil,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			issueSchemeWithCustomFields, err := testCase.issuePayload.MergeCustomFields(testCase.fields)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

			empJSON, err := json.MarshalIndent(issueSchemeWithCustomFields, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}

			t.Log(string(empJSON))
		})

	}

}

func TestIssueScheme_MergeOperations(t *testing.T) {

	var operations = &UpdateOperations{}

	err := operations.AddArrayOperation("labels", map[string]string{
		"triaged":   "remove",
		"triaged-2": "remove",
		"triaged-1": "remove",
		"blocker":   "remove",
	})

	if err != nil {
		t.Fatal(err)
	}

	var operationsWithOutFields = &UpdateOperations{}

	testCases := []struct {
		name         string
		operations   *UpdateOperations
		issuePayload *IssueScheme
		wantErr      bool
	}{
		{
			name: "MergeOperationsWhenTheOperationsAreCorrect",
			issuePayload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			operations: operations,
			wantErr:    false,
		},
		{
			name: "MergeOperationsWhenTheOperationsAreNil",
			issuePayload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			operations: nil,
			wantErr:    true,
		},
		{
			name: "MergeOperationsWhenTheOperationsDoNotHaveFields",
			issuePayload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			operations: operationsWithOutFields,
			wantErr:    true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			issueSchemeWithOperations, err := testCase.issuePayload.MergeOperations(testCase.operations)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

			empJSON, err := json.MarshalIndent(issueSchemeWithOperations, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}

			t.Log(string(empJSON))
		})

	}

}

func TestIssueService_Assign(t *testing.T) {

	testCases := []struct {
		name                    string
		issueKeyOrID, accountID string
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:               "AssignUserToIssueWhenTheParamsAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			accountID:          "8db2bf25-239d-43b6-8acf-b9769b4a5374",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3/assignee",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "AssignUserToIssueWhenTheIssueKeyOrIDIsNotProvided",
			issueKeyOrID:       "",
			accountID:          "8db2bf25-239d-43b6-8acf-b9769b4a5374",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3/assignee",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignUserToIssueWhenTheAccountIDIsSetAsNil",
			issueKeyOrID:       "DUMMY-3",
			accountID:          "",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3/assignee",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "AssignUserToIssueWhenTheAccountIDIsTheDefaultAssign",
			issueKeyOrID:       "DUMMY-3",
			accountID:          "-1",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3/assignee",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "AssignUserToIssueWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			accountID:          "8db2bf25-239d-43b6-8acf-b9769b4a5374",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/assignee",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignUserToIssueWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			accountID:          "8db2bf25-239d-43b6-8acf-b9769b4a5374",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3/assignee",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AssignUserToIssueWhenTheEndpointIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			accountID:          "8db2bf25-239d-43b6-8acf-b9769b4a5374",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issue/DUMMY-3/assignee",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignUserToIssueWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			accountID:          "8db2bf25-239d-43b6-8acf-b9769b4a5374",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3/assignee",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueService{client: mockClient}

			gotResponse, err := i.Assign(testCase.context, testCase.issueKeyOrID, testCase.accountID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}
		})

	}

}

func TestIssueService_Create(t *testing.T) {

	var customFieldMockedWithFields = CustomFields{}

	// Add a new custom field
	err := customFieldMockedWithFields.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		t.Fatal(err)
	}

	err = customFieldMockedWithFields.Number("customfield_10043", 1000.3232)
	if err != nil {
		t.Fatal(err)
	}

	var customFieldMockedWithOutFields = CustomFields{}

	testCases := []struct {
		name               string
		payload            *IssueScheme
		customFields       *CustomFields
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateIssueWhenTheCustomFieldsAreProvided",
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			customFields:       &customFieldMockedWithFields,
			mockFile:           "./mocks/create-issue.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateIssueWhenTheCustomFieldsAndPayloadAreNotProvided",
			payload:            nil,
			customFields:       nil,
			mockFile:           "./mocks/create-issue.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueWhenTheCustomFieldsAreProvidedButNotContainsFields",
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			customFields:       &customFieldMockedWithFields,
			mockFile:           "./mocks/create-issue.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name: "CreateIssueWhenTheCustomFieldsAreNotProvided",
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			customFields:       &customFieldMockedWithOutFields,
			mockFile:           "./mocks/create-issue.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueWhenTheCustomFieldsStructIsNil",
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			customFields:       nil,
			mockFile:           "./mocks/create-issue.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name: "CreateIssueWhenTheRequestMethodIsIncorrect",
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			customFields:       &customFieldMockedWithFields,
			mockFile:           "./mocks/create-issue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueWhenTheStatusCodeIsIncorrect",
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			customFields:       &customFieldMockedWithFields,
			mockFile:           "./mocks/create-issue.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name: "CreateIssueWhenTheEndpointIsIncorrect",
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			customFields:       &customFieldMockedWithFields,
			mockFile:           "./mocks/create-issue.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "CreateIssueWhenTheContextIsNil",
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			customFields:       &customFieldMockedWithFields,
			mockFile:           "./mocks/create-issue.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueWhenTheContextIsNilAndCustomFieldsAreNotProvided",
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			customFields:       nil,
			mockFile:           "./mocks/create-issue.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueWhenTheTheResponseBodyHasADifferentFormat",
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			customFields:       &customFieldMockedWithFields,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "CreateIssueWhenTheResponseBodyLengthIsZero",
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			customFields:       &customFieldMockedWithFields,
			mockFile:           "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context, testCase.payload, testCase.customFields)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

				t.Log("-----------------------")
				t.Logf("Create Issue Key: %v", gotResult.Key)
				t.Logf("Create Issue ID: %v", gotResult.ID)
				t.Logf("Create Issue Self: %v", gotResult.Self)
				t.Log("----------------------- \n")

			}
		})

	}

}

func TestIssueService_Creates(t *testing.T) {

	var customFieldMockedWithFields = CustomFields{}

	// Add a new custom field
	err := customFieldMockedWithFields.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		t.Fatal(err)
	}

	err = customFieldMockedWithFields.Number("customfield_10043", 1000.3232)
	if err != nil {
		t.Fatal(err)
	}

	var customFieldMockedWithOutFields = CustomFields{}

	var newIssuePayloadMockWithCustomFields00 = IssueBulkScheme{
		Payload: &IssueScheme{
			Fields: &IssueFieldsScheme{
				Summary:   "New summary test",
				Project:   &ProjectScheme{ID: "10000"},
				IssueType: &IssueTypeScheme{Name: "Story"},
			},
		},
		CustomFields: &customFieldMockedWithFields,
	}

	var newIssuePayloadMockWithCustomFields01 = IssueBulkScheme{
		Payload: &IssueScheme{
			Fields: &IssueFieldsScheme{
				Summary:   "New summary test",
				Project:   &ProjectScheme{ID: "10000"},
				IssueType: &IssueTypeScheme{Name: "Story"},
			},
		},
		CustomFields: &customFieldMockedWithFields,
	}

	var newIssuePayloadMockWithCustomFieldsValueAsNil = IssueBulkScheme{
		Payload: &IssueScheme{
			Fields: &IssueFieldsScheme{
				Summary:   "New summary test",
				Project:   &ProjectScheme{ID: "10000"},
				IssueType: &IssueTypeScheme{Name: "Story"},
			},
		},
		CustomFields: nil,
	}

	var newIssuePayloadMockWithOutCustomFields00 = IssueBulkScheme{
		Payload: &IssueScheme{
			Fields: &IssueFieldsScheme{
				Summary:   "New summary test",
				Project:   &ProjectScheme{ID: "10000"},
				IssueType: &IssueTypeScheme{Name: "Story"},
			},
		},
		CustomFields: &customFieldMockedWithOutFields,
	}

	var newIssuePayloadMockWithOutCustomFields01 = IssueBulkScheme{
		Payload: &IssueScheme{
			Fields: &IssueFieldsScheme{
				Summary:   "New summary test",
				Project:   &ProjectScheme{ID: "10000"},
				IssueType: &IssueTypeScheme{Name: "Story"},
			},
		},
		CustomFields: &customFieldMockedWithOutFields,
	}

	var newIssuePayloadMockWithPayloadIsNil = IssueBulkScheme{
		Payload:      nil,
		CustomFields: &customFieldMockedWithOutFields,
	}

	var payloadMockWithCorrectPayloads []*IssueBulkScheme
	payloadMockWithCorrectPayloads = append(
		payloadMockWithCorrectPayloads,
		&newIssuePayloadMockWithCustomFields00,
		&newIssuePayloadMockWithCustomFields01,
	)

	var payloadMockWithIncorrectPayloads []*IssueBulkScheme
	payloadMockWithIncorrectPayloads = append(
		payloadMockWithIncorrectPayloads,
		&newIssuePayloadMockWithOutCustomFields00,
		&newIssuePayloadMockWithOutCustomFields01,
		&newIssuePayloadMockWithPayloadIsNil,
		&newIssuePayloadMockWithCustomFieldsValueAsNil,
	)

	var payloadMockWithNilPayloads []*IssueBulkScheme
	payloadMockWithNilPayloads = append(payloadMockWithNilPayloads,
		&newIssuePayloadMockWithCustomFields00,
		&newIssuePayloadMockWithPayloadIsNil,
	)

	testCases := []struct {
		name               string
		payload            []*IssueBulkScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CreateIssuesWhenThePayloadsNodesAreCorrect",
			payload:            payloadMockWithCorrectPayloads,
			mockFile:           "./mocks/create-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/bulk",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateIssuesWhenThePayloadsNodesIsEmpty",
			payload:            nil,
			mockFile:           "./mocks/create-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/bulk",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateIssuesWhenThePayloadsContainsOneNodeWithOutCustomFieldsElements",
			payload:            payloadMockWithIncorrectPayloads,
			mockFile:           "./mocks/create-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/bulk",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateIssuesWhenThePayloadsContainsOneNilNode",
			payload:            payloadMockWithNilPayloads,
			mockFile:           "./mocks/create-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/bulk",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateIssuesWhenTheRequestMethodIsIncorrect",
			payload:            payloadMockWithCorrectPayloads,
			mockFile:           "./mocks/create-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/bulk",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateIssuesWhenTheStatusCodeIsIncorrect",
			payload:            payloadMockWithCorrectPayloads,
			mockFile:           "./mocks/create-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/bulk",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "CreateIssuesWhenTheContextIsNil",
			payload:            payloadMockWithCorrectPayloads,
			mockFile:           "./mocks/create-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/bulk",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateIssuesWhenTheEndpointIsIncorrect",
			payload:            payloadMockWithCorrectPayloads,
			mockFile:           "./mocks/create-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issue/bulk",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateIssuesWhenTheResponseBodyHasADifferentFormat",
			payload:            payloadMockWithCorrectPayloads,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/bulk",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueService{client: mockClient}

			gotResult, gotResponse, err := i.Creates(testCase.context, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

				for pos, issue := range gotResult.Issues {
					t.Log("-----------------------")
					t.Logf("Create Issue #%v Key: %v", pos+1, issue.Key)
					t.Logf("Create Issue #%v ID: %v", pos+1, issue.ID)
					t.Logf("Create Issue #%v Self: %v", pos+1, issue.Self)
					t.Log("----------------------- \n")
				}
			}
		})

	}

}

func TestIssueService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		deleteSubTasks     bool
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteIssueWhenTheIssueKeyIsSet",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteIssueWhenTheSubTasksIsDeleted",
			issueKeyOrID:       "DUMMY-3",
			deleteSubTasks:     true,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3?deleteSubtasks=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteIssueWhenTheIssueKeyIsNotSet",
			issueKeyOrID:       "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueWhenTheEndpointIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueWhenTheEndpointIsEmpty",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.issueKeyOrID, testCase.deleteSubTasks)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

			}
		})

	}

}

func TestIssueService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		fields             []string
		expands            []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueWhenTheFieldsAndExpandParametersAreSet",
			issueKeyOrID:       "DUMMY-3",
			fields:             []string{"summary", "description", "status", "customfield_00000"},
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			mockFile:           "./mocks/get-issue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&fields=summary%2Cdescription%2Cstatus%2Ccustomfield_00000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueWhenTheFieldsAndExpandParametersAreNotSet",
			issueKeyOrID:       "DUMMY-3",
			fields:             nil,
			expands:            nil,
			mockFile:           "./mocks/get-issue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueWhenTheFieldsParameterIsOnlySet",
			issueKeyOrID:       "DUMMY-3",
			fields:             []string{"summary", "description", "status", "customfield_00000"},
			expands:            nil,
			mockFile:           "./mocks/get-issue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3?fields=summary%2Cdescription%2Cstatus%2Ccustomfield_00000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueWhenTheExpandParameterIsSet",
			issueKeyOrID:       "DUMMY-3",
			fields:             nil,
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			mockFile:           "./mocks/get-issue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			fields:             []string{"summary", "description", "status", "customfield_00000"},
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			mockFile:           "./mocks/get-issue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&fields=summary%2Cdescription%2Cstatus%2Ccustomfield_00000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			fields:             []string{"summary", "description", "status", "customfield_00000"},
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			mockFile:           "./mocks/get-issue.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&fields=summary%2Cdescription%2Cstatus%2Ccustomfield_00000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			fields:             []string{"summary", "description", "status", "customfield_00000"},
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			mockFile:           "./mocks/get-issue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&fields=summary%2Cdescription%2Cstatus%2Ccustomfield_00000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetIssueWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			fields:             []string{"summary", "description", "status", "customfield_00000"},
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			mockFile:           "./mocks/get-issue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&fields=summary%2Cdescription%2Cstatus%2Ccustomfield_00000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWhenTheEndpointIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			fields:             []string{"summary", "description", "status", "customfield_00000"},
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			mockFile:           "./mocks/get-issue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/DUMMY-3?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&fields=summary%2Cdescription%2Cstatus%2Ccustomfield_00000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			fields:             []string{"summary", "description", "status", "customfield_00000"},
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&fields=summary%2Cdescription%2Cstatus%2Ccustomfield_00000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.issueKeyOrID, testCase.fields, testCase.expands)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

				t.Log("-----------------------")
				t.Logf("Create Returned Key: %v", gotResult.Key)
				t.Logf("Create Returned ID: %v", gotResult.ID)
				t.Logf("Create Returned Self: %v", gotResult.Self)
				t.Log("----------------------- \n")
			}
		})

	}

}

func TestIssueService_Move(t *testing.T) {

	var payload = &IssueScheme{
		Fields: &IssueFieldsScheme{
			Summary: "New summary test test",
		},
	}

	var customFieldMockedWithFields = CustomFields{}

	// Add a new custom field
	err := customFieldMockedWithFields.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		t.Fatal(err)
	}

	err = customFieldMockedWithFields.Number("customfield_10043", 1000.3232)
	if err != nil {
		t.Fatal(err)
	}

	var customFieldMockedWithOutFields = CustomFields{}

	var operationMockedWithFields = UpdateOperations{}

	err = operationMockedWithFields.AddArrayOperation("labels", map[string]string{
		"triaged":   "remove",
		"triaged-2": "remove",
		"triaged-1": "remove",
		"blocker":   "remove",
	})

	if err != nil {
		t.Fatal(err)
	}

	operationMockedWithOutFields := UpdateOperations{}

	testCases := []struct {
		name               string
		issueKeyOrID       string
		transitionID       string
		options            *IssueMoveOptions
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:           "MoveIssueStatusWhenTheTheCustomFieldsIsProvidedButNotOperations",
			issueKeyOrID:   "DUMMY-3",
			transitionID:   "10001",
			wantHTTPMethod: http.MethodPost,
			options: &IssueMoveOptions{
				Fields:       payload,
				CustomFields: &customFieldMockedWithFields,
				Operations:   nil,
			},
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:           "MoveIssueStatusWhenTheTheOperationsIsProvidedButNotCustomFields",
			issueKeyOrID:   "DUMMY-3",
			transitionID:   "10001",
			wantHTTPMethod: http.MethodPost,
			options: &IssueMoveOptions{
				Fields:       payload,
				CustomFields: nil,
				Operations:   &operationMockedWithFields,
			},
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:           "MoveIssueStatusWhenTheTheOperationsIsProvidedButNotCustomFieldsAndContext",
			issueKeyOrID:   "DUMMY-3",
			transitionID:   "10001",
			wantHTTPMethod: http.MethodPost,
			options: &IssueMoveOptions{
				Fields:       payload,
				CustomFields: nil,
				Operations:   &operationMockedWithFields,
			},
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:           "MoveIssueStatusWhenTheTheOperationsIsProvidedButNotCustomFieldsAndNotFieldsOperations",
			issueKeyOrID:   "DUMMY-3",
			transitionID:   "10001",
			wantHTTPMethod: http.MethodPost,
			options: &IssueMoveOptions{
				Fields:       payload,
				CustomFields: nil,
				Operations:   &operationMockedWithOutFields,
			},
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:           "MoveIssueStatusWhenTheTheCustomFieldsIsProvidedButNotOperationsAndContext",
			issueKeyOrID:   "DUMMY-3",
			transitionID:   "10001",
			wantHTTPMethod: http.MethodPost,
			options: &IssueMoveOptions{
				Fields:       payload,
				CustomFields: &customFieldMockedWithFields,
				Operations:   nil,
			},
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:           "MoveIssueStatusWhenTheTheCustomFieldsIsProvidedButWithNotFields",
			issueKeyOrID:   "DUMMY-3",
			transitionID:   "10001",
			wantHTTPMethod: http.MethodPost,
			options: &IssueMoveOptions{
				Fields:       payload,
				CustomFields: &customFieldMockedWithOutFields,
				Operations:   nil,
			},
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:           "MoveIssueStatusWhenTheTheCustomFieldsAndOperationAreProvided",
			issueKeyOrID:   "DUMMY-3",
			transitionID:   "10001",
			wantHTTPMethod: http.MethodPost,
			options: &IssueMoveOptions{
				Fields:       payload,
				CustomFields: &customFieldMockedWithFields,
				Operations:   &operationMockedWithFields,
			},
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:           "MoveIssueStatusWhenTheTheCustomFieldsAndOperationAreProvidedButNotContext",
			issueKeyOrID:   "DUMMY-3",
			transitionID:   "10001",
			wantHTTPMethod: http.MethodPost,
			options: &IssueMoveOptions{
				Fields:       payload,
				CustomFields: &customFieldMockedWithFields,
				Operations:   &operationMockedWithFields,
			},
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:           "MoveIssueStatusWhenTheTheOperationAreProvidedButNotTheCustomFields",
			issueKeyOrID:   "DUMMY-3",
			transitionID:   "10001",
			wantHTTPMethod: http.MethodPost,
			options: &IssueMoveOptions{
				Fields:       payload,
				CustomFields: &customFieldMockedWithOutFields,
				Operations:   &operationMockedWithFields,
			},
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:           "MoveIssueStatusWhenTheTheCustomFieldsAreProvidedButNotTheOperations",
			issueKeyOrID:   "DUMMY-3",
			transitionID:   "10001",
			wantHTTPMethod: http.MethodPost,
			options: &IssueMoveOptions{
				Fields:       payload,
				CustomFields: &customFieldMockedWithFields,
				Operations:   &operationMockedWithOutFields,
			},
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "MoveIssueStatusWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			transitionID:       "10001",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "MoveIssueStatusWhenTheTransitionIDIsNotSet",
			issueKeyOrID:       "DUMMY-3",
			transitionID:       "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "MoveIssueStatusWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			transitionID:       "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "MoveIssueStatusWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			transitionID:       "10001",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "MoveIssueStatusWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			transitionID:       "10001",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "MoveIssueStatusWhenTheEndpointIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			transitionID:       "10001",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueService{client: mockClient}

			gotResponse, err := i.Move(testCase.context, testCase.issueKeyOrID, testCase.transitionID, testCase.options)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

			}
		})

	}

}

func TestIssueService_Notify(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		options            *IssueNotifyOptionsScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:         "SendMailNotificationWhenTheOptionsAreCorrect",
			issueKeyOrID: "DUMMY-3",
			options: &IssueNotifyOptionsScheme{
				HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",
				Subject:  "SUBJECT EMAIL EXAMPLE",
				To: &IssueNotifyToScheme{
					Reporter: true,
					Assignee: true,
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/notify",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:         "SendMailNotificationWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID: "",
			options: &IssueNotifyOptionsScheme{
				HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",
				Subject:  "SUBJECT EMAIL EXAMPLE",
				To: &IssueNotifyToScheme{
					Reporter: true,
					Assignee: true,
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/notify",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "SendMailNotificationWhenTheOptionsAreNil",
			issueKeyOrID:       "DUMMY-3",
			options:            nil,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/notify",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "SendMailNotificationWhenTheRequestMethodIsIncorrect",
			issueKeyOrID: "DUMMY-3",
			options: &IssueNotifyOptionsScheme{
				HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",
				Subject:  "SUBJECT EMAIL EXAMPLE",
				To: &IssueNotifyToScheme{
					Reporter: true,
					Assignee: true,
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3/notify",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "SendMailNotificationWhenTheStatusCodeIsIncorrect",
			issueKeyOrID: "DUMMY-3",
			options: &IssueNotifyOptionsScheme{
				HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",
				Subject:  "SUBJECT EMAIL EXAMPLE",
				To: &IssueNotifyToScheme{
					Reporter: true,
					Assignee: true,
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/notify",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:         "SendMailNotificationWhenTheContextIsNil",
			issueKeyOrID: "DUMMY-3",
			options: &IssueNotifyOptionsScheme{
				HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",
				Subject:  "SUBJECT EMAIL EXAMPLE",
				To: &IssueNotifyToScheme{
					Reporter: true,
					Assignee: true,
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/notify",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "SendMailNotificationWhenTheEndpointIsIncorrect",
			issueKeyOrID: "DUMMY-3",
			options: &IssueNotifyOptionsScheme{
				HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",
				Subject:  "SUBJECT EMAIL EXAMPLE",
				To: &IssueNotifyToScheme{
					Reporter: true,
					Assignee: true,
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issue/DUMMY-3/notify",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueService{client: mockClient}

			gotResponse, err := i.Notify(testCase.context, testCase.issueKeyOrID, testCase.options)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

			}
		})

	}

}

func TestIssueService_Transitions(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueTransitionsWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			mockFile:           "./mocks/get-transitions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueTransitionsWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			mockFile:           "./mocks/get-transitions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTransitionsWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			mockFile:           "./mocks/get-transitions.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTransitionsWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			mockFile:           "./mocks/get-transitions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetIssueTransitionsWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			mockFile:           "./mocks/get-transitions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssueTransitionsWhenTheEndpointsIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			mockFile:           "./mocks/get-transitions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssueTransitionsWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/transitions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueService{client: mockClient}

			gotResult, gotResponse, err := i.Transitions(testCase.context, testCase.issueKeyOrID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

				for _, transition := range gotResult.Transitions {

					t.Log("------------------------------")
					t.Logf("Transition Name: %v", transition.Name)
					t.Logf("Transition ID: %v", transition.ID)
					t.Logf("Transition HasScreen: %v", transition.HasScreen)

					t.Logf("Transition To Name: %v", transition.To.Name)
					t.Logf("Transition To ID: %v", transition.To.ID)
					t.Logf("Transition To Self: %v", transition.To.Self)
					t.Log("------------------------------")
				}
			}
		})

	}

}

func TestIssueService_Update(t *testing.T) {

	var customFieldMockedWithFields = CustomFields{}

	// Add a new custom field
	err := customFieldMockedWithFields.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		t.Fatal(err)
	}

	err = customFieldMockedWithFields.Number("customfield_10043", 1000.3232)
	if err != nil {
		t.Fatal(err)
	}

	var customFieldMockedWithOutFields = CustomFields{}

	var operationMockedWithFields = UpdateOperations{}

	err = operationMockedWithFields.AddArrayOperation("labels", map[string]string{
		"triaged":   "remove",
		"triaged-2": "remove",
		"triaged-1": "remove",
		"blocker":   "remove",
	})

	if err != nil {
		t.Fatal(err)
	}

	var operationMockedWithOutFields = UpdateOperations{}

	testCases := []struct {
		name               string
		issueKeyOrID       string
		notify             bool
		payload            *IssueScheme
		customFields       *CustomFields
		operations         *UpdateOperations
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:         "EditIssueWhenTheCustomFieldsAndOperationAreProvided",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithFields,
			customFields:       &customFieldMockedWithFields,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:         "EditIssueWhenTheCustomFieldsAreProvided",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         nil,
			customFields:       &customFieldMockedWithFields,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:         "EditIssueWhenTheOperationsAreProvided",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithFields,
			customFields:       nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:         "EditIssueWhenTheNotificationIsDisabled",
			issueKeyOrID: "DUMMY-3",
			notify:       false,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithFields,
			customFields:       &customFieldMockedWithFields,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3?notifyUsers=false",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:         "EditIssueWhenTheIssueKeyOrIDIsNotProvided",
			issueKeyOrID: "",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithFields,
			customFields:       &customFieldMockedWithFields,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheCustomFieldsAndOperationAreNotProvided",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         nil,
			customFields:       nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "EditIssueWhenTheCustomFields,OperationAndPayloadAreNotProvided",
			issueKeyOrID:       "DUMMY-3",
			notify:             true,
			payload:            nil,
			operations:         nil,
			customFields:       nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheCustomFields,OperationAndContextAreNotProvided",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         nil,
			customFields:       nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheCustomFieldsAndOperationAndProvideButNoTheContext",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithFields,
			customFields:       &customFieldMockedWithFields,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheCustomFieldsAndOperationAreProvidedButCustomfieldDoNotHaveFields",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithFields,
			customFields:       &customFieldMockedWithOutFields,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheCustomFieldsAndOperationAreProvidedButOperationsDoNotHaveFields",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithOutFields,
			customFields:       &customFieldMockedWithFields,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheCustomFieldsAreProvidedButCustomfieldDoNotHaveFields",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         nil,
			customFields:       &customFieldMockedWithOutFields,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheOperationsAreProvidedButOperationsDoNotHaveFields",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithOutFields,
			customFields:       nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheCustomFieldsAreProvidedButContextIsNil",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         nil,
			customFields:       &customFieldMockedWithFields,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheOperationsAreProvidedButContextIsNil",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithFields,
			customFields:       nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheRequestMethodIsIncorrect",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithFields,
			customFields:       &customFieldMockedWithFields,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheStatusCodeIsIncorrect",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithFields,
			customFields:       &customFieldMockedWithFields,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:         "EditIssueWhenTheEndpointIsEmpty",
			issueKeyOrID: "DUMMY-3",
			notify:       true,
			payload: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary: "New summary test",
				},
			},
			operations:         &operationMockedWithFields,
			customFields:       &customFieldMockedWithFields,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueService{client: mockClient}

			gotResponse, err := i.Update(
				testCase.context,
				testCase.issueKeyOrID,
				testCase.notify,
				testCase.payload,
				testCase.customFields,
				testCase.operations,
			)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

			}
		})

	}

}

func TestUpdateOperations_AddArrayOperation(t *testing.T) {

	testCases := []struct {
		name          string
		customFieldID string
		mapping       map[string]string
		wantErr       bool
	}{
		{
			name:          "AddArrayOperationWhenTheParametersAreSet",
			customFieldID: "customfield_000",
			mapping: map[string]string{
				"triaged":   "remove",
				"triaged-2": "remove",
				"triaged-1": "remove",
				"blocker":   "remove",
			},
			wantErr: false,
		},

		{
			name:          "AddArrayOperationWhenTheCustomFieldIDIsNotSet",
			customFieldID: "",
			mapping: map[string]string{
				"triaged":   "remove",
				"triaged-2": "remove",
				"triaged-1": "remove",
				"blocker":   "remove",
			},
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var operations = UpdateOperations{}
			err := operations.AddArrayOperation(testCase.customFieldID, testCase.mapping)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestUpdateOperations_AddStringOperation(t *testing.T) {

	testCases := []struct {
		name                            string
		customFieldID, operation, value string
		wantErr                         bool
	}{
		{
			name:          "AddStringOperationWhenTheParametersAreSet",
			customFieldID: "summary",
			operation:     "set",
			value:         "new summary using operation",
			wantErr:       false,
		},

		{
			name:          "AddStringOperationWhenTheCustomFieldIDIsNotSet",
			customFieldID: "",
			operation:     "set",
			value:         "new summary using operation",
			wantErr:       true,
		},

		{
			name:          "AddStringOperationWhenTheOperationIsNotSet",
			customFieldID: "summary",
			operation:     "",
			value:         "new summary using operation",
			wantErr:       true,
		},

		{
			name:          "AddStringOperationWhenTheValueIsNotSet",
			customFieldID: "summary",
			operation:     "set",
			value:         "",
			wantErr:       true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			var operations = UpdateOperations{}
			err := operations.AddStringOperation(testCase.customFieldID, testCase.operation, testCase.value)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			}

		})

	}

}

func TestIssueScheme_ToMap(t *testing.T) {

	testCases := []struct {
		name    string
		issue   *IssueScheme
		wantErr bool
	}{
		{
			name: "ConvertIssueStructToMapWhenTheParametersAreCorrect",
			issue: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			wantErr: false,
		},

		{
			name:    "ConvertIssueStructToMapWhenTheIssueStructIsNil",
			issue:   nil,
			wantErr: false,
		},

		{
			name: "ConvertIssueStructToMapWhenTheParametersAreCorrect",
			issue: &IssueScheme{
				Fields: &IssueFieldsScheme{
					Summary:   "New summary test",
					Project:   &ProjectScheme{ID: "10000"},
					IssueType: &IssueTypeScheme{Name: "Story"},
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			result, err := testCase.issue.ToMap()

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
			} else {
				t.Log(result)
			}

		})

	}

}
