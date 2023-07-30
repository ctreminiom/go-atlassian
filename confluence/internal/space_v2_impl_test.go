package internal

import (
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_internalSpaceV2Impl_Bulk(t *testing.T) {

	optionsMocked := &model.GetSpacesOptionSchemeV2{
		IDs:               []int{10001, 10002},
		Keys:              []string{"DUMMY"},
		Type:              "global",
		Status:            "current",
		Labels:            []string{"test-label"},
		Sort:              "-name",
		DescriptionFormat: "view",
		SerializeIDs:      true,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx     context.Context
		options *model.GetSpacesOptionSchemeV2
		cursor  string
		limit   int
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:     context.TODO(),
				options: optionsMocked,
				cursor:  "cursor_sample_uuid",
				limit:   50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces?cursor=cursor_sample_uuid&description-format=view&ids=10001%2C10002&keys=DUMMY&labels=test-label&limit=50&serialize-ids-as-strings=true&sort=-name&status=current&type=global",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SpaceChunkV2Scheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.TODO(),
				options: optionsMocked,
				cursor:  "cursor_sample_uuid",
				limit:   50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces?cursor=cursor_sample_uuid&description-format=view&ids=10001%2C10002&keys=DUMMY&labels=test-label&limit=50&serialize-ids-as-strings=true&sort=-name&status=current&type=global",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSpaceV2Service(testCase.fields.c)

			gotResult, gotResponse, err := newService.Bulk(testCase.args.ctx, testCase.args.options, testCase.args.cursor,
				testCase.args.limit)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalSpaceV2Impl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx               context.Context
		spaceID           int
		descriptionFormat string
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:               context.TODO(),
				spaceID:           10001,
				descriptionFormat: "view",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces/10001?description-format=view",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SpaceSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the space id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoSpaceIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:               context.TODO(),
				spaceID:           10001,
				descriptionFormat: "view",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/spaces/10001?description-format=view",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSpaceV2Service(testCase.fields.c)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.spaceID, testCase.args.descriptionFormat)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}
