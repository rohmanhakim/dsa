package lc91

import (
	"reflect"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{

		{
			name: "example 1",
			s:    "12",
			want: 2,
		},

		{
			name: "example 2",
			s:    "226",
			want: 3,
		},

		{
			name: "example 3",
			s:    "06",
			want: 0,
		},

		{
			name: "example 4",
			s:    "10",
			want: 1,
		},

		{
			name: "example 5",
			s:    "1234",
			want: 3,
		},

		{
			name: "example 6",
			s:    "0123",
			want: 0,
		},

		{
			name: "example 7",
			s:    "120",
			want: 1,
		},

		{
			name: "example 8",
			s:    "1230",
			want: 0,
		},

		{
			name: "example 9",
			s:    "2101",
			want: 1,
		},

		{
			name: "example 10",
			s:    "1123",
			want: 5,
		},

		{
			name: "example 11",
			s:    "112211",
			want: 13,
		},

		{
			name: "example 12",
			s:    "1122112",
			want: 21,
		},

		{
			name: "example 13",
			s:    "27",
			want: 1,
		},

		{
			name: "example 14",
			s:    "2611055971756562",
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := numDecodings(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
