package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestAuditService_Get(t *testing.T) {

	var mockDateFormat = "2021-02-26T15:04:05.999-0700"

	mockDateTimeAsTime, err := time.Parse(DateFormatJira, mockDateFormat)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name               string
		options            *AuditRecordGetOptions
		offset, limit      int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "GetAuditRecordsWhenTheOptionsAreSet",
			options: &AuditRecordGetOptions{
				Filter: "",
				From:   mockDateTimeAsTime.AddDate(0, -1, 0).Format(DateFormatJira),
				To:     mockDateTimeAsTime.Format(DateFormatJira),
			},
			offset:             0,
			limit:              1000,
			mockFile:           "./mocks/get-audit-records.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/auditing/record?from=2021-01-26T15%3A04%3A05.999-0700&limit=1000&offset=0&to=2021-02-26T15%3A04%3A05.999-0700",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name: "GetAuditRecordsWhenTheOptionFilterIsSet",
			options: &AuditRecordGetOptions{
				Filter: "Workflow",
				From:   mockDateTimeAsTime.AddDate(0, -1, 0).Format(DateFormatJira),
				To:     mockDateTimeAsTime.Format(DateFormatJira),
			},
			offset:             0,
			limit:              1000,
			mockFile:           "./mocks/get-audit-records.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/auditing/record?filter=Workflow&from=2021-01-26T15%3A04%3A05.999-0700&limit=1000&offset=0&to=2021-02-26T15%3A04%3A05.999-0700",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetAuditRecordsWhenTheOptionsAreNil",
			options:            nil,
			offset:             0,
			limit:              1000,
			mockFile:           "./mocks/get-audit-records.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/auditing/record?limit=1000&offset=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name: "GetAuditRecordsWhenTheRequestMethodIsIncorrect",
			options: &AuditRecordGetOptions{
				Filter: "",
				From:   mockDateTimeAsTime.AddDate(0, -1, 0).Format(DateFormatJira),
				To:     mockDateTimeAsTime.Format(DateFormatJira),
			},
			offset:             0,
			limit:              1000,
			mockFile:           "./mocks/get-audit-records.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/auditing/record?from=2021-01-26T15%3A04%3A05.999-0700&limit=1000&offset=0&to=2021-02-26T15%3A04%3A05.999-0700",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetAuditRecordsWhenTheStatusCodeIsIncorrect",
			options: &AuditRecordGetOptions{
				Filter: "",
				From:   mockDateTimeAsTime.AddDate(0, -1, 0).Format(DateFormatJira),
				To:     mockDateTimeAsTime.Format(DateFormatJira),
			},
			offset:             0,
			limit:              1000,
			mockFile:           "./mocks/get-audit-records.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/auditing/record?from=2021-01-26T15%3A04%3A05.999-0700&limit=1000&offset=0&to=2021-02-26T15%3A04%3A05.999-0700",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "GetAuditRecordsWhenTheContextIsNil",
			options: &AuditRecordGetOptions{
				Filter: "",
				From:   mockDateTimeAsTime.AddDate(0, -1, 0).Format(DateFormatJira),
				To:     mockDateTimeAsTime.Format(DateFormatJira),
			},
			offset:             0,
			limit:              1000,
			mockFile:           "./mocks/get-audit-records.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/auditing/record?from=2021-01-26T15%3A04%3A05.999-0700&limit=1000&offset=0&to=2021-02-26T15%3A04%3A05.999-0700",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetAuditRecordsWhenTheEndpointIsIncorrect",
			options: &AuditRecordGetOptions{
				Filter: "",
				From:   mockDateTimeAsTime.AddDate(0, -1, 0).Format(DateFormatJira),
				To:     mockDateTimeAsTime.Format(DateFormatJira),
			},
			offset:             0,
			limit:              1000,
			mockFile:           "./mocks/get-audit-records.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           " /rest/api/3/auditing/record?from=2021-01-26T15%3A04%3A05.999-0700&limit=1000&offset=0&to=2021-02-26T15%3A04%3A05.999-0700",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetAuditRecordsWhenTheResponseBodyHasADifferentFormat",
			options: &AuditRecordGetOptions{
				Filter: "",
				From:   mockDateTimeAsTime.AddDate(0, -1, 0).Format(DateFormatJira),
				To:     mockDateTimeAsTime.Format(DateFormatJira),
			},
			offset:             0,
			limit:              1000,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/auditing/record?from=2021-01-26T15%3A04%3A05.999-0700&limit=1000&offset=0&to=2021-02-26T15%3A04%3A05.999-0700",
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

			i := &AuditService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.options, testCase.offset, testCase.limit)

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

				for _, record := range gotResult.Records {

					t.Log("-----------------------")
					t.Logf("Audit Record ID: %v", record.ID)
					t.Logf("Audit Record Summary: %v", record.Summary)
					t.Logf("Audit Record Category: %v", record.Category)
					t.Logf("Audit Record Created: %v", record.Created)
					t.Logf("Audit Record RemoteAddress: %v", record.RemoteAddress)
					t.Logf("Audit Record AuthorKey: %v", record.AuthorKey)
					t.Log("----------------------- \n")

				}

			}
		})

	}

}
