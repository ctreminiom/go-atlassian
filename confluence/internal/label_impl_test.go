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

func Test_internalLabelImpl_Get(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                  context.Context
		labelName, labelType string
		start, limit         int
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
				ctx:       context.TODO(),
				labelName: "blogs",
				labelType: "blogpost",
				start:     200,
				limit:     50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/label?limit=50&name=blogs&start=200&type=blogpost",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.LabelDetailsScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.TODO(),
				labelName: "blogs",
				labelType: "blogpost",
				start:     200,
				limit:     50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/label?limit=50&name=blogs&start=200&type=blogpost",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the label name is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoLabelNameError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewLabelService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.labelName, testCase.args.labelType,
				testCase.args.start, testCase.args.limit)

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
