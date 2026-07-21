package lc3499

import (
	"reflect"
	"testing"
)

func TestMaxActiveSectionsAfterTrade(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{

		{
			name: "example 1",
			s:    "01",
			want: 1,
		},

		{
			name: "example 2",
			s:    "0100",
			want: 4,
		},

		{
			name: "example 3",
			s:    "1000100",
			want: 7,
		},

		{
			name: "example 4",
			s:    "01010",
			want: 4,
		},

		{
			name: "example 5",
			s:    "100",
			want: 1,
		},

		{
			name: "example 6",
			s:    "100",
			want: 1,
		},

		{
			name: "example 7",
			s:    "0111",
			want: 3,
		},

		{
			name: "example 8",
			s:    "01101",
			want: 5,
		},

		{
			name: "example 9",
			s:    "100011100111110000100000",
			want: 19,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxActiveSectionsAfterTrade(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
