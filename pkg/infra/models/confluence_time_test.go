package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfluenceDateTimeScheme_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    ConfluenceDateTimeScheme
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "null",
			args:    args{data: []byte(`null`)},
			want:    ConfluenceDateTimeScheme{},
			wantErr: assert.NoError,
		},
		{
			name:    "ISO 8601 with milliseconds",
			args:    args{data: []byte(`"2024-09-23T20:17:35.607Z"`)},
			want:    ConfluenceDateTimeScheme(time.Date(2024, 9, 23, 20, 17, 35, 607000000, time.UTC)),
			wantErr: assert.NoError,
		},
		{
			name:    "RFC 3339 without fractional seconds",
			args:    args{data: []byte(`"2024-09-23T20:17:35Z"`)},
			want:    ConfluenceDateTimeScheme(time.Date(2024, 9, 23, 20, 17, 35, 0, time.UTC)),
			wantErr: assert.NoError,
		},
		{
			name:    "RFC 3339 with timezone offset",
			args:    args{data: []byte(`"2024-09-23T22:17:35+02:00"`)},
			want:    ConfluenceDateTimeScheme(time.Date(2024, 9, 23, 22, 17, 35, 0, time.FixedZone("", 7200))),
			wantErr: assert.NoError,
		},
		{
			name:    "Unix epoch milliseconds",
			args:    args{data: []byte(`1727122655607`)},
			want:    ConfluenceDateTimeScheme(time.Date(2024, 9, 23, 20, 17, 35, 607000000, time.UTC)),
			wantErr: assert.NoError,
		},
		{
			name:    "invalid string",
			args:    args{data: []byte(`"not-a-date"`)},
			want:    ConfluenceDateTimeScheme{},
			wantErr: assert.Error,
		},
		{
			name:    "invalid token",
			args:    args{data: []byte(`true`)},
			want:    ConfluenceDateTimeScheme{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := ConfluenceDateTimeScheme{}
			tt.wantErr(t, d.UnmarshalJSON(tt.args.data), fmt.Sprintf("UnmarshalJSON(%s)", tt.args.data))
			assert.Truef(t, time.Time(tt.want).Equal(time.Time(d)), "expected %v, got %v", time.Time(tt.want), time.Time(d))
		})
	}
}

func TestConfluenceDateTimeScheme_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		t       ConfluenceDateTimeScheme
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "UTC time",
			t:       ConfluenceDateTimeScheme(time.Date(2024, 9, 23, 20, 17, 35, 0, time.UTC)),
			want:    `"2024-09-23T20:17:35Z"`,
			wantErr: assert.NoError,
		},
		{
			name:    "time with offset",
			t:       ConfluenceDateTimeScheme(time.Date(2024, 9, 23, 22, 17, 35, 0, time.FixedZone("CET", 7200))),
			want:    `"2024-09-23T22:17:35+02:00"`,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.MarshalJSON()
			if !tt.wantErr(t, err, "MarshalJSON()") {
				return
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}
