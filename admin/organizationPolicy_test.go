package admin

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestOrganizationPolicyService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     string
		policyType         string
		cursor             string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetOrganizationPoliciesWhenTheParametersAreCorrect",
			mockFile:           "./mocks/get-organization-policies.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyType:         "policy-type-sample",
			cursor:             "eu-realm-policy",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies?cursor=eu-realm-policy&type=policy-type-sample",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetOrganizationPoliciesWhenTheCursorIsNotSet",
			mockFile:           "./mocks/get-organization-policies.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyType:         "policy-type-sample",
			cursor:             "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies?type=policy-type-sample",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetOrganizationPoliciesWhenTheQueryParamsAreNotSet",
			mockFile:           "./mocks/get-organization-policies.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyType:         "",
			cursor:             "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetOrganizationPoliciesWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-organization-policies.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyType:         "policy-type-sample",
			cursor:             "eu-realm-policy",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies?cursor=eu-realm-policy&type=policy-type-sample",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPoliciesWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-organization-policies.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyType:         "policy-type-sample",
			cursor:             "eu-realm-policy",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies?cursor=eu-realm-policy&type=policy-type-sample",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPoliciesWhenTheContextIsNil",
			mockFile:           "./mocks/get-organization-policies.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyType:         "policy-type-sample",
			cursor:             "eu-realm-policy",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies?cursor=eu-realm-policy&type=policy-type-sample",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPoliciesWhenTheEndpointIsEmpty",
			mockFile:           "./mocks/get-organization-policies.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyType:         "policy-type-sample",
			cursor:             "eu-realm-policy",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPoliciesWhenTheResponseBodyIsEmpty",
			mockFile:           "./mocks/empty.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyType:         "policy-type-sample",
			cursor:             "eu-realm-policy",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies?cursor=eu-realm-policy&type=policy-type-sample",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPoliciesWhenTheOrganizationIDIsNotSet",
			mockFile:           "./mocks/get-organization-policies.json",
			organizationID:     "",
			policyType:         "policy-type-sample",
			cursor:             "eu-realm-policy",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies?cursor=eu-realm-policy&type=policy-type-sample",
			context:            context.Background(),
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &OrganizationPolicyService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.organizationID, testCase.policyType, testCase.cursor)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)

				for _, organization := range gotResult.Data {
					t.Log(organization.ID, organization.Attributes.Name)
				}
			}

		})
	}

}

func TestOrganizationPolicyService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     string
		policyID           string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetOrganizationPolicyWhenTheParametersAreCorrect",
			mockFile:           "./mocks/get-organization-policy.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetOrganizationPolicyWhenWhenTheOrganizationIDIsNotSet",
			mockFile:           "./mocks/get-organization-policy.json",
			organizationID:     "",
			policyID:           "policy-type-sample",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPolicyWhenThePolicyIDIsNotSet",
			mockFile:           "./mocks/get-organization-policy.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPolicyWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-organization-policy.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "policy-type-sample",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPolicyWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-organization-policy.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "policy-type-sample",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPolicyWhenTheContextIsNil",
			mockFile:           "./mocks/get-organization-policy.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "policy-type-sample",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPolicyWhenTheEndpointIsEmpty",
			mockFile:           "./mocks/get-organization-policy.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "policy-type-sample",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPolicyWhenTheResponseBodyIsEmpty",
			mockFile:           "./mocks/empty.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "policy-type-sample",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationPolicyWhenTheResponseBodyIsIncorrect",
			mockFile:           "./mocks/get-organization-policies.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &OrganizationPolicyService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.organizationID, testCase.policyID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)

				t.Log(gotResult)
			}

		})
	}

}

func TestOrganizationPolicyService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     string
		payload            *OrganizationPolicyData
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:           "CreateOrganizationPolicyWhenTheParametersAreCorrect",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:           "CreateOrganizationPolicyWhenTheOrganizationIDIsNotSet",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "CreateOrganizationPolicyWhenThePayloadIsNil",
			mockFile:           "./mocks/get-organization-policy.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			payload:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "CreateOrganizationPolicyWhenTheRequestMethodIsIncorrect",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "CreateOrganizationPolicyWhenTheStatusCodeIsIncorrect",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "CreateOrganizationPolicyWhenTheContextIsNil",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies",
			context:            nil,
			wantErr:            true,
		},

		{
			name:           "CreateOrganizationPolicyWhenTheEndpointIsEmpty",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "CreateOrganizationPolicyWhenTheResponseBodyIsEmpty",
			mockFile:       "./mocks/empty.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "CreateOrganizationPolicyWhenTheResponseBodyIsIncorrect",
			mockFile:       "./mocks/get-organization-policies.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies",
			context:            context.Background(),
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &OrganizationPolicyService{client: mockClient}
			gotResult, gotResponse, err := service.Create(testCase.context, testCase.organizationID, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)

				t.Log(gotResult)
			}

		})
	}

}

func TestOrganizationPolicyService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     string
		policyID           string
		payload            *OrganizationPolicyData
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:           "UpdateOrganizationPolicyWhenTheParametersAreCorrect",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:       "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:           "UpdateOrganizationPolicyWhenTheOrganizationIDIsNotSet",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "",
			policyID:       "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "UpdateOrganizationPolicyWhenThePolicyIDIsNotSet",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:       "",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "UpdateOrganizationPolicyWhenThePayloadIsNil",
			mockFile:           "./mocks/get-organization-policy.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			payload:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "UpdateOrganizationPolicyWhenTheRequestMethodIsIncorrect",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:       "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "UpdateOrganizationPolicyWhenTheStatusCodeIsIncorrect",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:       "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "UpdateOrganizationPolicyWhenTheContextIsNil",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:       "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            nil,
			wantErr:            true,
		},

		{
			name:           "UpdateOrganizationPolicyWhenTheEndpointIsEmpty",
			mockFile:       "./mocks/get-organization-policy.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:       "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "UpdateOrganizationPolicyWhenTheResponseBodyIsEmpty",
			mockFile:       "./mocks/empty.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:       "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "UpdateOrganizationPolicyWhenTheResponseBodyIsIncorrect",
			mockFile:       "./mocks/get-organization-policies.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:       "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			payload: &OrganizationPolicyData{
				Type: "policy",
				Attributes: &OrganizationPolicyAttributes{
					Type:   "data-residency", //ip-allowlist
					Name:   "SCIMUserNameScheme of this Policy",
					Status: "enabled", //disabled
				},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &OrganizationPolicyService{client: mockClient}
			gotResult, gotResponse, err := service.Update(testCase.context, testCase.organizationID, testCase.policyID, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)

				t.Log(gotResult)
			}

		})
	}

}

func TestOrganizationPolicyService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     string
		policyID           string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteOrganizationPolicyWhenTheParametersAreCorrect",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "DeleteOrganizationPolicyWhenTheOrganizationIDIsNotSet",
			organizationID:     "",
			policyID:           "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "DeleteOrganizationPolicyWhenThePolicyIDIsNotSet",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "DeleteOrganizationPolicyWhenTheRequestMethodIsIncorrect",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "DeleteOrganizationPolicyWhenTheStatusCodeIsIncorrect",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "DeleteOrganizationPolicyWhenTheContextIsNil",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "DeleteOrganizationPolicyWhenTheEndpointIsEmpty",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			policyID:           "60f0f660-be3e-4d70-bd34-9c2858ec040f",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &OrganizationPolicyService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.organizationID, testCase.policyID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)

			}

		})
	}

}
