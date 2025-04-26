package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateScheme_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		t       DateScheme
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "random time",
			t:       DateScheme(time.Date(2021, 6, 22, 7, 34, 41, 123, time.UTC)),
			want:    []byte(`"2021-06-22"`),
			wantErr: assert.NoError,
		},
		{
			name:    "timezone",
			t:       DateScheme(time.Date(2021, 6, 22, 7, 34, 41, 123, time.FixedZone("UTC+2", 7200))),
			want:    []byte(`"2021-06-22"`),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.MarshalJSON()
			if !tt.wantErr(t, err, "MarshalJSON()") {
				return
			}
			assert.Equalf(t, string(tt.want), string(got), "MarshalJSON()")
		})
	}
}

func TestDateScheme_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    DateScheme
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "null",
			args: args{
				data: []byte(`null`),
			},
			want:    DateScheme{},
			wantErr: assert.NoError,
		},
		{
			name: "random date",
			args: args{
				data: []byte(`"2021-06-22"`),
			},
			want:    DateScheme(time.Date(2021, 6, 22, 0, 0, 0, 0, time.UTC)),
			wantErr: assert.NoError,
		},
		{
			name: "invalid date",
			args: args{
				data: []byte(`"asdf"`),
			},
			want:    DateScheme{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DateScheme{}
			tt.wantErr(t, d.UnmarshalJSON(tt.args.data), fmt.Sprintf("UnmarshalJSON(%v)", tt.args.data))
			assert.Equalf(t, tt.want, d, "MarshalJSON()")
		})
	}
}

func TestDateTimeScheme_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		t       DateTimeScheme
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "random time",
			t:       DateTimeScheme(time.Date(2021, 6, 22, 7, 34, 41, 123, time.FixedZone("UTC-9", -32400))),
			want:    []byte(`"2021-06-22T07:34:41-0900"`),
			wantErr: assert.NoError,
		},
		{
			name:    "timezone",
			t:       DateTimeScheme(time.Date(2021, 6, 22, 7, 34, 41, 123, time.FixedZone("UTC+2", 7200))),
			want:    []byte(`"2021-06-22T07:34:41+0200"`),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.MarshalJSON()
			if !tt.wantErr(t, err, "MarshalJSON()") {
				return
			}
			assert.Equalf(t, string(tt.want), string(got), "MarshalJSON()")
		})
	}
}

func TestDateTimeScheme_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    DateTimeScheme
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "null",
			args: args{
				data: []byte(`null`),
			},
			want:    DateTimeScheme{},
			wantErr: assert.NoError,
		},
		{
			name: "random date",
			args: args{
				data: []byte(`"2021-06-22T07:34:41+0200"`),
			},
			want:    DateTimeScheme(time.Date(2021, 6, 22, 7, 34, 41, 0, time.FixedZone("", 7200))),
			wantErr: assert.NoError,
		},
		{
			name: "invalid date",
			args: args{
				data: []byte(`"asdf"`),
			},
			want:    DateTimeScheme{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DateTimeScheme{}
			tt.wantErr(t, d.UnmarshalJSON(tt.args.data), fmt.Sprintf("UnmarshalJSON(%v)", tt.args.data))
			assert.Equalf(t, tt.want, d, "MarshalJSON()")
		})
	}
}
